// Package realtime is the WebSocket subscription transport for the Bron API.
//
// Conceptually a subscription is "GET extended": you send the same query DTO
// you would on the matching list endpoint, the server replays the historical
// match as the first frame, then keeps the connection open and pushes live
// updates as additional frames of the same response shape (a list with one
// element per change in steady state).
//
// This package handles the wire envelope (SUBSCRIBE / UNSUBSCRIBE, JWT
// signing, status decoding, body string-or-object handling, ping keepalive,
// and transparent auto-reconnect). Typed convenience wrappers live in
// `sdk/api/<resource>.go`.
//
// Reconnect behavior matches the bron-web frontend (websocket-context.tsx):
//   - keep-alive ping every 15s (server idles out at ~60s)
//   - on transient close (1006 abnormal, network error, server restart)
//     transparently re-dial and re-SUBSCRIBE with the same Correlation-Id
//   - linear backoff 1s → 2s → … → 10s, reset to 0 on a successful
//     re-subscribe
//   - close code 4000 (logout) ends the stream permanently
//   - close code 4001 (token refresh) reconnects immediately with no backoff
//
// On every reconnect the server replays the snapshot frame again (same as
// frontend) — callers should be prepared to see duplicate items. The CLI
// `--no-history` flag (limit=0) makes this a no-op.
package realtime

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	nethttp "net/http"
	"net/url"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
)

// Frame is one envelope received from the server.
type Frame struct {
	Status  int
	Headers map[string]string
	Body    json.RawMessage // already unwrapped if the server sent a string-encoded JSON body
}

// Subscription describes what to subscribe to. URI is the same path you'd hit
// on the matching list endpoint; Filter is the same query DTO. CorrelationID
// is optional (auto-generated if empty) and is reused verbatim across
// reconnects so log correlation stays stable for the whole subscription.
type Subscription struct {
	URI           string
	Filter        interface{}
	CorrelationID string
}

// Stream is an open subscription. Frames arrive on the returned channel; when
// the stream ends (peer logout 4000, ctx cancel, Close called, or a
// non-recoverable error), the channel is closed and Err() returns the reason
// (nil for clean shutdown).
//
// Always defer Close() — it cancels reconnects, sends UNSUBSCRIBE on the
// current connection (best-effort), and tears down the WebSocket. Calling
// Close() multiple times is safe.
type Stream struct {
	frames chan Frame
	sub    Subscription
	uri    string // filter-encoded URI (sent on every (re)subscribe)
	client *Client

	closeOnce sync.Once
	cancelCtx context.CancelFunc
	done      chan struct{} // closed when runLoop exits
	finalErr  atomicError

	// connMu protects conn / writeMu / closed. conn+writeMu are replaced on
	// every successful reconnect; `closed` flips once Close is called and
	// blocks any further setConn from installing a fresh connection (avoids
	// a race where a dial in flight returns after Close and lands a live
	// conn that nothing will ever close).
	connMu  sync.Mutex
	conn    *websocket.Conn
	writeMu *sync.Mutex
	closed  bool
}

func (s *Stream) Frames() <-chan Frame { return s.frames }

// Err returns the reason the stream ended; nil for clean shutdown
// (Close called, ctx cancel, server logout 4000).
func (s *Stream) Err() error { return s.finalErr.load() }

// Close stops reconnects, best-effort UNSUBSCRIBE if the writer is idle,
// and waits for the run loop to exit. Safe to call concurrently and
// multiple times.
//
// Order matters:
//  1. flip `closed` so any in-flight reconnect drops its dial.
//  2. send UNSUBSCRIBE while the conn is still alive (TryLock — if
//     pingLoop is mid-write we skip the goodbye and rely on close).
//  3. cancelCtx — wakes the serveOnce ctx-watcher which force-closes the
//     conn so ReadMessage unblocks.
//  4. wait for runLoop to exit.
//
// Sending UNSUBSCRIBE BEFORE cancelCtx avoids the ctx-watcher racing us
// to a conn.Close (which would make the WriteJSON in step 2 fail).
func (s *Stream) Close() error {
	s.closeOnce.Do(func() {
		s.connMu.Lock()
		s.closed = true
		conn, mu := s.conn, s.writeMu
		s.connMu.Unlock()

		if conn != nil && mu.TryLock() {
			_ = conn.SetWriteDeadline(time.Now().Add(time.Second))
			_ = conn.WriteJSON(map[string]interface{}{
				"method":  "UNSUBSCRIBE",
				"uri":     s.uri,
				"headers": map[string]string{"Correlation-Id": s.sub.CorrelationID},
			})
			mu.Unlock()
		}

		s.cancelCtx()
		// Belt-and-suspenders: explicitly close the conn we just sent
		// UNSUBSCRIBE on. The serveOnce ctx-watcher will see it already
		// closed and skip its own close (no double-close on Conn — gorilla
		// returns net.ErrClosed which we ignore).
		if conn != nil {
			_ = conn.Close()
		}
		<-s.done
	})
	return nil
}

// closeConn closes the currently-installed connection (if any). Used by
// runLoop between reconnect cycles. Close() itself bypasses this and closes
// directly under connMu so it can prevent a racing setConn from reinstalling.
func (s *Stream) closeConn() {
	s.connMu.Lock()
	conn := s.conn
	s.connMu.Unlock()
	if conn != nil {
		_ = conn.Close()
	}
}

// Client is a thin WS client. Reuse one across many subscriptions.
type Client struct {
	baseURL   string
	apiKey    string
	dialer    *websocket.Dialer
	signer    func(auth.BronJwtOptions) (string, error)
	clock     func() time.Time
	onEvent   func(LifecycleEvent)
	stableMin time.Duration // 0 → use defaultStableThreshold
	logger    *slog.Logger
}

// log returns the configured logger or a discard handler. Always non-nil.
func (c *Client) log() *slog.Logger {
	if c.logger != nil {
		return c.logger
	}
	return discardLogger
}

var discardLogger = slog.New(slog.NewTextHandler(io.Discard, nil))

// stableThreshold returns the configured "connection considered stable"
// duration (override via WithStableThreshold for tests, otherwise default).
func (c *Client) stableThreshold() time.Duration {
	if c.stableMin > 0 {
		return c.stableMin
	}
	return defaultStableThreshold
}

// LifecycleEvent surfaces reconnect-related state changes to the caller.
// Useful for stderr logging in CLIs and metrics/alerts in long-running
// services. Default handler is a no-op.
type LifecycleEvent struct {
	Kind          LifecycleKind
	CorrelationID string
	URI           string
	Attempt       int           // 1-based; meaningful for Reconnecting / Reconnected
	Backoff       time.Duration // delay before this Reconnecting attempt
	Err           error         // last error that triggered the reconnect
}

type LifecycleKind int

const (
	// EventDisconnected fires once when the active connection ends with an
	// error that the client will try to recover from.
	EventDisconnected LifecycleKind = iota
	// EventReconnecting fires before each re-dial attempt (after the
	// backoff has elapsed).
	EventReconnecting
	// EventReconnected fires once a re-dial + re-SUBSCRIBE succeeds.
	EventReconnected
)

func (k LifecycleKind) String() string {
	switch k {
	case EventDisconnected:
		return "disconnected"
	case EventReconnecting:
		return "reconnecting"
	case EventReconnected:
		return "reconnected"
	default:
		return "unknown"
	}
}

type Option func(*Client)

// WithProxy makes the dialer route through proxyURL. Empty string falls back
// to HTTP_PROXY / HTTPS_PROXY env vars; pass an explicit URL to override.
func WithProxy(proxyURL string) Option {
	return func(c *Client) {
		if proxyURL == "" {
			c.dialer.Proxy = nethttp.ProxyFromEnvironment
			return
		}
		u, err := url.Parse(proxyURL)
		if err != nil || u.Scheme == "" || u.Host == "" {
			c.dialer.Proxy = nethttp.ProxyFromEnvironment
			return
		}
		c.dialer.Proxy = nethttp.ProxyURL(u)
	}
}

// WithSigner injects a custom JWT signer (default: auth.GenerateBronJwt).
func WithSigner(s func(auth.BronJwtOptions) (string, error)) Option {
	return func(c *Client) { c.signer = s }
}

// WithClock injects a clock for iat (default: time.Now).
func WithClock(cl func() time.Time) Option {
	return func(c *Client) { c.clock = cl }
}

// WithDialer replaces the underlying *websocket.Dialer wholesale (e.g. when
// the caller has custom TLS config, NetDial, or HandshakeTimeout to set).
func WithDialer(d *websocket.Dialer) Option {
	return func(c *Client) {
		if d != nil {
			c.dialer = d
		}
	}
}

// WithLifecycleHandler registers a callback for connection state changes
// (disconnect, reconnect attempts, reconnect success). The callback is
// invoked from the run loop goroutine — keep it short and non-blocking.
func WithLifecycleHandler(fn func(LifecycleEvent)) Option {
	return func(c *Client) { c.onEvent = fn }
}

// WithStableThreshold overrides the duration a reconnected connection has
// to stay alive before the backoff escalation resets to zero. Useful for
// deterministic unit tests; production code should leave the default.
func WithStableThreshold(d time.Duration) Option {
	return func(c *Client) {
		if d > 0 {
			c.stableMin = d
		}
	}
}

// WithLogger plugs a structured logger in. Default is a discard handler.
//
// Levels used:
//   - DEBUG: each frame received, each ping sent, dial attempts, envelopes
//   - INFO:  connect, disconnect, reconnect events
//   - WARN:  reconnect-with-backoff (transient flap)
//   - ERROR: terminal failures (e.g. server logout, non-recoverable error)
//
// Authorization tokens never appear in logs — only `correlationId`,
// `uri`, `attempt`, `backoff`, etc. Set the handler level to DEBUG via the
// caller's `slog.HandlerOptions` to enable frame-level tracing.
func WithLogger(l *slog.Logger) Option {
	return func(c *Client) {
		if l != nil {
			c.logger = l
		}
	}
}

func NewClient(baseURL, apiKey string, opts ...Option) *Client {
	dialer := *websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second
	dialer.Proxy = nethttp.ProxyFromEnvironment

	c := &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		dialer:  &dialer,
		signer:  auth.GenerateBronJwt,
		clock:   time.Now,
		onEvent: func(LifecycleEvent) {},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Reconnect tunables, sized to match the bron-web frontend
// (front/mono/shared/helpers/websocket-context.tsx).
const (
	pingInterval        = 15 * time.Second
	reconnectStep       = 1 * time.Second
	reconnectMax        = 10 * time.Second
	tokenRefreshBackoff = 1 * time.Second // close code 4001
	closeCodeLogout     = 4000
	closeCodeReauth     = 4001
	// defaultStableThreshold is how long a reconnected stream must keep
	// ticking before the backoff/attempt counters reset. Without this a
	// flapping server (accepts SUBSCRIBE → drops immediately) would loop
	// with 0-delay forever. Two ping intervals — if the conn lasted long
	// enough for at least one ping round-trip, the disconnect was "real"
	// rather than the server slamming the door right after handshake.
	defaultStableThreshold = 2 * pingInterval
)

// errLogout is the sentinel returned from a connection cycle when the server
// closed with code 4000 — runLoop treats it as terminal (no reconnect).
var errLogout = errors.New("realtime: server closed (logout)")

// Subscribe opens a subscription. Returns synchronously after the first
// successful dial + SUBSCRIBE. After that, network drops and idle timeouts
// trigger transparent re-dial + re-SUBSCRIBE; the caller keeps reading from
// the same Frames() channel.
func (c *Client) Subscribe(ctx context.Context, sub Subscription) (*Stream, error) {
	if sub.URI == "" {
		return nil, fmt.Errorf("realtime: Subscription.URI is required")
	}
	if sub.CorrelationID == "" {
		sub.CorrelationID = newCorrelationID()
	}

	uri, err := encodeFilterToURI(sub.URI, sub.Filter)
	if err != nil {
		return nil, fmt.Errorf("realtime: encode filter: %w", err)
	}

	streamCtx, cancel := context.WithCancel(ctx)

	stream := &Stream{
		frames:    make(chan Frame, 16),
		sub:       sub,
		uri:       uri,
		client:    c,
		cancelCtx: cancel,
		done:      make(chan struct{}),
	}

	// First dial is synchronous: surface the error to the caller before we
	// hand back a Stream. The caller's Subscribe() should fail fast on bad
	// credentials, unreachable host, etc.
	conn, mu, err := c.dialAndSubscribe(streamCtx, uri, sub.CorrelationID)
	if err != nil {
		cancel()
		return nil, err
	}
	// Initial setConn cannot race Close — the caller hasn't seen the
	// stream yet. Subsequent reconnects use the same call site but go
	// through runLoop, which handles the ok=false case.
	stream.setConn(conn, mu)

	go stream.runLoop(streamCtx)

	return stream, nil
}

// dialAndSubscribe opens a fresh WebSocket and sends one SUBSCRIBE envelope.
// Returns the connection plus its write mutex; the caller owns lifecycle.
func (c *Client) dialAndSubscribe(ctx context.Context, uri, correlationID string) (*websocket.Conn, *sync.Mutex, error) {
	var jwk map[string]interface{}
	if err := json.Unmarshal([]byte(c.apiKey), &jwk); err != nil {
		return nil, nil, fmt.Errorf("realtime: parse api key as JWK: %w", err)
	}
	kid, _ := jwk["kid"].(string)
	if kid == "" {
		return nil, nil, fmt.Errorf("realtime: api key JWK is missing 'kid'")
	}

	// JWT signs (iat, method, uri, body). Wire body is `{}` since filters
	// travel on the URI query (parametersEntity ignores the body).
	const wireBody = "{}"
	iat := c.clock().Unix()
	token, err := c.signer(auth.BronJwtOptions{
		Method:     "SUBSCRIBE",
		Path:       uri,
		Body:       wireBody,
		Kid:        kid,
		PrivateKey: c.apiKey,
		Iat:        &iat,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("realtime: sign jwt: %w", err)
	}

	wsURL := httpToWs(c.baseURL) + "/ws"
	c.log().Debug("realtime: dial", "url", wsURL, "uri", uri, "correlationId", correlationID)
	conn, _, err := c.dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		c.log().Debug("realtime: dial failed", "url", wsURL, "err", err)
		return nil, nil, fmt.Errorf("realtime: dial %s: %w", wsURL, err)
	}

	subEnvelope := map[string]interface{}{
		"method": "SUBSCRIBE",
		"uri":    uri,
		"headers": map[string]string{
			"Correlation-Id": correlationID,
			"Authorization":  "ApiKey " + token,
			"Content-Type":   "application/json",
		},
		"body": map[string]interface{}{},
	}

	mu := &sync.Mutex{}
	mu.Lock()
	_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	writeErr := conn.WriteJSON(subEnvelope)
	mu.Unlock()
	if writeErr != nil {
		_ = conn.Close()
		return nil, nil, fmt.Errorf("realtime: send subscribe: %w", writeErr)
	}

	c.log().Debug("realtime: subscribed", "uri", uri, "correlationId", correlationID)
	return conn, mu, nil
}

// setConn installs a freshly-dialed connection. Returns false (and the
// caller MUST close `conn` itself) if Close already fired — this is the
// race where dialAndSubscribe completed concurrently with Close. Without
// this check the connection would never be observed by Close's `closeConn`
// and ReadMessage in serveOnce would block forever, deadlocking the
// `<-done` wait inside Close.
func (s *Stream) setConn(conn *websocket.Conn, mu *sync.Mutex) bool {
	s.connMu.Lock()
	defer s.connMu.Unlock()
	if s.closed {
		return false
	}
	s.conn = conn
	s.writeMu = mu
	return true
}

// runLoop owns the connection lifecycle:
//   - serve frames from the current connection
//   - on disconnect, decide whether to reconnect, and with what backoff
//   - on (re)connect, hand off to serveOnce again
//
// Backoff strategy mirrors the bron-web frontend:
//   - "stable" disconnect (conn lived ≥ stableThreshold) → reset to 0,
//     reconnect immediately
//   - "flap" disconnect (conn lived < stableThreshold) → escalate backoff
//     by one step per cycle so we don't pummel a bouncing server
//   - dial failures escalate backoff too (each failed attempt waits longer)
//   - close code 4001 (token refresh) overrides to a fixed 1s
//
// Done — closes both `frames` and `done`. After this returns, callers see
// the channel close and `Err()` returns the terminal error (or nil for a
// clean shutdown).
func (s *Stream) runLoop(ctx context.Context) {
	defer close(s.done)
	defer close(s.frames)

	backoff := time.Duration(0)
	attempt := 0
	connectedAt := s.client.clock()

	for {
		err := s.serveOnce(ctx)

		// Terminal cases — don't reconnect.
		if ctx.Err() != nil {
			return
		}
		if errors.Is(err, errLogout) {
			s.finalErr.store(err)
			return
		}
		if err != nil && !isReconnectable(err) && !errors.Is(err, errReauth) {
			s.finalErr.store(err)
			return
		}

		s.emit(LifecycleEvent{Kind: EventDisconnected, Err: err})
		s.closeConn()

		// Decide initial reconnect delay based on how the conn died.
		switch {
		case errors.Is(err, errReauth):
			// Server-initiated token refresh — frontend does instant 1s.
			backoff = tokenRefreshBackoff
			attempt = 0
		case s.client.clock().Sub(connectedAt) >= s.client.stableThreshold():
			// Healthy conn that just dropped — first reconnect is immediate.
			backoff = 0
			attempt = 0
		default:
			// Flap (or initial dial-fail loop) — escalate.
			backoff = nextBackoff(backoff)
		}

		// Inner loop: keep dialing until success, ctx cancel, or Close.
		// Each failed dial escalates backoff one more step.
		newConn, newMu, dialErr := s.reconnect(ctx, &attempt, &backoff, err)
		if dialErr != nil {
			// ctx cancel during reconnect — exit cleanly.
			return
		}
		if newConn == nil {
			// Close() landed while reconnecting — exit cleanly.
			return
		}

		if !s.setConn(newConn, newMu) {
			_ = newConn.Close()
			return
		}
		s.emit(LifecycleEvent{Kind: EventReconnected, Attempt: attempt})
		connectedAt = s.client.clock()
	}
}

// reconnect runs the inner re-dial loop with backoff escalation. Returns
// (conn, mu, nil) on success, (nil, nil, ctx.Err) on caller cancel.
// `attempt` and `backoff` are bumped in place so caller observability can
// see the latest values via the lifecycle handler.
func (s *Stream) reconnect(ctx context.Context, attempt *int, backoff *time.Duration, lastErr error) (*websocket.Conn, *sync.Mutex, error) {
	for {
		if *backoff > 0 {
			select {
			case <-ctx.Done():
				return nil, nil, ctx.Err()
			case <-time.After(*backoff):
			}
		}

		*attempt++
		s.emit(LifecycleEvent{Kind: EventReconnecting, Attempt: *attempt, Backoff: *backoff, Err: lastErr})

		conn, mu, derr := s.client.dialAndSubscribe(ctx, s.uri, s.sub.CorrelationID)
		if derr == nil {
			// Parent ctx may have been canceled DURING the dial — if so,
			// drop the brand-new conn and bail. Without this check, the
			// caller could install a fresh conn whose subsequent
			// ReadMessage would block until the server idle-times-out
			// (~60s) since the gorilla socket doesn't honour ctx.
			if ctx.Err() != nil {
				_ = conn.Close()
				return nil, nil, ctx.Err()
			}
			return conn, mu, nil
		}

		if ctx.Err() != nil {
			return nil, nil, ctx.Err()
		}
		// Dial failed — escalate and retry. Don't surface the dial error;
		// the lifecycle handler already sees Reconnecting attempts and the
		// next dial's outcome will be visible via Reconnected or another
		// Reconnecting event.
		*backoff = nextBackoff(*backoff)
	}
}

// emit fans a lifecycle event out to the registered handler AND mirrors it
// to the structured logger. CorrelationID and URI are filled in here so
// callers don't have to remember.
func (s *Stream) emit(evt LifecycleEvent) {
	evt.CorrelationID = s.sub.CorrelationID
	evt.URI = s.uri

	logger := s.client.log()
	switch evt.Kind {
	case EventDisconnected:
		logger.Info("realtime: disconnected",
			"correlationId", evt.CorrelationID,
			"uri", evt.URI,
			"err", evt.Err,
		)
	case EventReconnecting:
		level := slog.LevelInfo
		if evt.Backoff > 0 || evt.Attempt > 1 {
			level = slog.LevelWarn
		}
		logger.Log(context.Background(), level, "realtime: reconnecting",
			"correlationId", evt.CorrelationID,
			"uri", evt.URI,
			"attempt", evt.Attempt,
			"backoff", evt.Backoff,
			"err", evt.Err,
		)
	case EventReconnected:
		logger.Info("realtime: reconnected",
			"correlationId", evt.CorrelationID,
			"uri", evt.URI,
			"attempt", evt.Attempt,
		)
	}

	if s.client.onEvent != nil {
		s.client.onEvent(evt)
	}
}

// errReauth signals a server-initiated token refresh (close code 4001) —
// runLoop should reconnect immediately with no backoff.
var errReauth = errors.New("realtime: server requested token refresh")

// serveOnce runs read+ping for a single connection. Captures the conn and
// writeMu at start so concurrent reconnects don't make pingLoop ping the
// fresh socket. Spawns a ctx-watcher that force-closes the captured conn
// when the parent ctx is canceled — without it, ReadMessage would block on
// a healthy socket even after Close() / parent-ctx cancel.
//
// Returns:
//   - nil on clean ctx cancel (Close() or parent ctx canceled)
//   - errLogout on close code 4000
//   - errReauth on close code 4001 (caller should reconnect immediately)
//   - any other error if the connection ended unexpectedly
func (s *Stream) serveOnce(ctx context.Context) error {
	s.connMu.Lock()
	conn, mu := s.conn, s.writeMu
	s.connMu.Unlock()
	if conn == nil {
		return errors.New("realtime: no connection")
	}

	connCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go pingLoop(connCtx, conn, mu)
	// ctx-watcher: gorilla/websocket's ReadMessage doesn't honour ctx, so a
	// caller-side Close()/parent-ctx cancel wouldn't unblock the read on a
	// healthy socket without forcibly closing the underlying conn. The
	// watcher does that. Distinguishes parent-ctx cancel (close the conn)
	// from serveOnce's own defer cancel (do nothing — conn is being torn
	// down by runLoop or already errored out).
	go func() {
		<-connCtx.Done()
		if ctx.Err() != nil {
			_ = conn.Close()
		}
	}()

	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			if websocket.IsCloseError(err, closeCodeLogout) {
				return errLogout
			}
			if websocket.IsCloseError(err, closeCodeReauth) {
				return errReauth
			}
			return err
		}

		frame, ok := parseEnvelope(raw)
		if !ok {
			frame = Frame{Body: append(json.RawMessage(nil), raw...)}
		}

		s.client.log().Debug("realtime: frame",
			"correlationId", s.sub.CorrelationID,
			"status", frame.Status,
			"bodyBytes", len(frame.Body),
		)

		select {
		case <-ctx.Done():
			return nil
		case s.frames <- frame:
		}
	}
}

// pingLoop pings the captured conn until ctx is done or a write fails.
// Bound to one connection cycle — runLoop spawns a fresh pingLoop after
// every reconnect, so an old loop that's still ticking when we replace the
// socket can't accidentally write to the new one.
func pingLoop(ctx context.Context, conn *websocket.Conn, mu *sync.Mutex) {
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			mu.Lock()
			_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			err := conn.WriteMessage(websocket.PingMessage, nil)
			mu.Unlock()
			if err != nil {
				return
			}
		}
	}
}

// nextBackoff bumps the linear delay; capped at reconnectMax.
func nextBackoff(prev time.Duration) time.Duration {
	next := prev + reconnectStep
	if next > reconnectMax {
		return reconnectMax
	}
	if next < reconnectStep {
		return reconnectStep
	}
	return next
}

// reconnectableCloseCodes is the explicit allowlist of WebSocket close
// codes that warrant a retry. Server-side bug codes (protocol violation,
// invalid payload, message-too-big, mandatory-extension, policy-violation,
// TLS-handshake) are deliberately omitted — retrying produces the same
// failure forever.
var reconnectableCloseCodes = []int{
	websocket.CloseNormalClosure,     // 1000 — server bounced gracefully
	websocket.CloseGoingAway,         // 1001 — server going away
	websocket.CloseAbnormalClosure,   // 1006 — TCP drop / idle timeout
	websocket.CloseInternalServerErr, // 1011
	websocket.CloseServiceRestart,    // 1012
	websocket.CloseTryAgainLater,     // 1013
}

// isReconnectable decides whether an error from ReadMessage/WriteMessage
// warrants a transparent reconnect. Recognised classes:
//   - websocket close frames in the allowlist above
//   - io.ErrUnexpectedEOF / net.ErrClosed / ctx deadline → TCP drops
//   - net.Error with Timeout()=true → read or write deadline expired
//   - ECONNRESET / EPIPE / ETIMEDOUT → TCP-layer errno from the syscall
//
// Anything else (terminal close codes, JSON parse errors, application
// errors with words like "timeout" in their text) returns false so we
// don't loop on a broken state.
func isReconnectable(err error) bool {
	if err == nil {
		return false
	}
	if websocket.IsCloseError(err, reconnectableCloseCodes...) {
		return true
	}
	if errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, net.ErrClosed) ||
		errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	if errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.EPIPE) ||
		errors.Is(err, syscall.ETIMEDOUT) {
		return true
	}
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return true
	}
	return false
}

func parseEnvelope(raw []byte) (Frame, bool) {
	var env map[string]json.RawMessage
	if err := json.Unmarshal(raw, &env); err != nil {
		return Frame{}, false
	}

	frame := Frame{Headers: map[string]string{}}

	if statusRaw, ok := env["status"]; ok {
		var status int
		if err := json.Unmarshal(statusRaw, &status); err != nil {
			var s string
			if err := json.Unmarshal(statusRaw, &s); err == nil {
				_, _ = fmt.Sscanf(s, "%d", &status)
			}
		}
		frame.Status = status
	}

	if headersRaw, ok := env["headers"]; ok {
		var hdrs map[string]string
		if err := json.Unmarshal(headersRaw, &hdrs); err == nil {
			frame.Headers = hdrs
		}
	}

	body := env["body"]
	if len(body) > 0 && body[0] == '"' {
		// Server double-encoded the body as a JSON string — unwrap it.
		var s string
		if err := json.Unmarshal(body, &s); err == nil {
			body = json.RawMessage(s)
		}
	}
	frame.Body = body

	return frame, true
}

// encodeFilterToURI marshals filter as JSON, walks the resulting map, and
// appends each (key, value) pair as a URL query parameter on baseURI.
// Arrays are comma-joined (matches the REST client's encoding). Returns the
// baseURI verbatim if filter is nil or marshals to an empty object.
func encodeFilterToURI(baseURI string, filter interface{}) (string, error) {
	if filter == nil {
		return baseURI, nil
	}
	raw, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err != nil {
		return baseURI, nil // not a JSON object — nothing to encode
	}
	if len(m) == 0 {
		return baseURI, nil
	}
	params := url.Values{}
	for k, v := range m {
		if v == nil {
			continue
		}
		switch val := v.(type) {
		case []interface{}:
			parts := make([]string, len(val))
			for i, e := range val {
				parts[i] = fmt.Sprintf("%v", e)
			}
			params.Set(k, strings.Join(parts, ","))
		default:
			params.Set(k, fmt.Sprintf("%v", val))
		}
	}
	qs := params.Encode()
	if qs == "" {
		return baseURI, nil
	}
	if strings.Contains(baseURI, "?") {
		return baseURI + "&" + qs, nil
	}
	return baseURI + "?" + qs, nil
}

func httpToWs(base string) string {
	switch {
	case strings.HasPrefix(base, "https://"):
		return "wss://" + strings.TrimPrefix(base, "https://")
	case strings.HasPrefix(base, "http://"):
		return "ws://" + strings.TrimPrefix(base, "http://")
	}
	return base
}

// newCorrelationID returns a 32-char hex token (128 bits of entropy) prefixed
// with `sdk-`. Falls back to a timestamp-suffixed string only if crypto/rand
// is unavailable (effectively never).
func newCorrelationID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return fmt.Sprintf("sdk-%d", time.Now().UnixNano())
	}
	return "sdk-" + hex.EncodeToString(b[:])
}

// atomicError is a tiny wrapper around sync.Mutex + error for the run loop
// to publish a final error before closing the channel.
type atomicError struct {
	mu  sync.Mutex
	err error
}

func (a *atomicError) store(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.err = err
}

func (a *atomicError) load() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.err
}
