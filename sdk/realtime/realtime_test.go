package realtime

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
)

// fakeJWK is the minimal JWK shape Subscribe needs to extract `kid`.
const fakeJWK = `{"kty":"EC","crv":"P-256","x":"x","y":"y","d":"d","kid":"test-kid"}`

func nopSigner(opts auth.BronJwtOptions) (string, error) { return "fake-jwt", nil }

// startServer spins a test WS server and returns its handler control + URL.
// onSubscribe is called for every SUBSCRIBE envelope received; the conn
// passed in is then handed back to the caller via ctrl so the test can drive
// the connection (write frames, close, etc.).
type serverCtrl struct {
	url     string
	t       *testing.T
	mu      sync.Mutex
	conns   []*websocket.Conn
	subEnvs []map[string]interface{}
	closeFn func()
}

func startServer(t *testing.T) *serverCtrl {
	t.Helper()
	c := &serverCtrl{t: t}
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("upgrade: %v", err)
			return
		}
		// Wait for SUBSCRIBE
		_, raw, err := ws.ReadMessage()
		if err != nil {
			_ = ws.Close()
			return
		}
		var env map[string]interface{}
		_ = json.Unmarshal(raw, &env)
		c.mu.Lock()
		c.conns = append(c.conns, ws)
		c.subEnvs = append(c.subEnvs, env)
		c.mu.Unlock()
		// Block forever (test will close it).
		for {
			if _, _, err := ws.ReadMessage(); err != nil {
				return
			}
		}
	}))
	c.url = "http://" + strings.TrimPrefix(srv.URL, "http://")
	c.closeFn = srv.Close
	return c
}

func (c *serverCtrl) waitForConns(n int, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		c.mu.Lock()
		if len(c.conns) >= n {
			c.mu.Unlock()
			return true
		}
		c.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
	return false
}

func (c *serverCtrl) latestConn() *websocket.Conn {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.conns) == 0 {
		return nil
	}
	return c.conns[len(c.conns)-1]
}

// sendFrame mirrors the server-side frame envelope shape (status/headers/body).
func (c *serverCtrl) sendFrame(body string) error {
	conn := c.latestConn()
	if conn == nil {
		return nil
	}
	return conn.WriteJSON(map[string]interface{}{
		"status":  200,
		"headers": map[string]string{},
		"body":    json.RawMessage(body),
	})
}

func (c *serverCtrl) dropLatest() {
	conn := c.latestConn()
	if conn != nil {
		_ = conn.Close()
	}
}

func TestSubscribe_ReceivesFrame(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/api/v1/transactions"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatalf("server didn't see subscribe")
	}

	if err := srv.sendFrame(`{"hello":"world"}`); err != nil {
		t.Fatalf("server send: %v", err)
	}

	select {
	case f := <-stream.Frames():
		if string(f.Body) != `{"hello":"world"}` {
			t.Fatalf("unexpected body: %s", f.Body)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("no frame received")
	}
}

func TestSubscribe_AutoReconnect_SameCorrelationID(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	var disconnects, reconnects int32
	client := NewClient(srv.url, fakeJWK,
		WithSigner(nopSigner),
		WithLifecycleHandler(func(evt LifecycleEvent) {
			switch evt.Kind {
			case EventDisconnected:
				atomic.AddInt32(&disconnects, 1)
			case EventReconnected:
				atomic.AddInt32(&reconnects, 1)
			}
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{
		URI:           "/api/v1/transactions",
		CorrelationID: "fixed-id",
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("server didn't see initial subscribe")
	}

	srv.dropLatest()

	if !srv.waitForConns(2, 3*time.Second) {
		t.Fatal("client didn't reconnect")
	}

	// Both SUBSCRIBE envelopes must carry the same Correlation-Id.
	srv.mu.Lock()
	envs := append([]map[string]interface{}{}, srv.subEnvs...)
	srv.mu.Unlock()
	if len(envs) < 2 {
		t.Fatalf("expected >=2 SUBSCRIBE envelopes, got %d", len(envs))
	}
	for i, env := range envs[:2] {
		hdrs, _ := env["headers"].(map[string]interface{})
		corr, _ := hdrs["Correlation-Id"].(string)
		if corr != "fixed-id" {
			t.Fatalf("env %d: correlation-id=%q want %q", i, corr, "fixed-id")
		}
	}

	// And we should still be able to receive a frame on the new conn.
	if err := srv.sendFrame(`{"after":"reconnect"}`); err != nil {
		t.Fatalf("server send post-reconnect: %v", err)
	}

	select {
	case f := <-stream.Frames():
		if !strings.Contains(string(f.Body), "after") {
			t.Fatalf("unexpected body: %s", f.Body)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("no frame after reconnect")
	}

	if atomic.LoadInt32(&disconnects) < 1 {
		t.Errorf("expected at least 1 disconnect event, got %d", disconnects)
	}
	if atomic.LoadInt32(&reconnects) < 1 {
		t.Errorf("expected at least 1 reconnect event, got %d", reconnects)
	}
}

func TestSubscribe_BackoffEscalates_WhenServerDown(t *testing.T) {
	// Start a server, dial through it once so the initial Subscribe()
	// succeeds, then close the server entirely so re-dial fails repeatedly.
	srv := startServer(t)

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("initial subscribe didn't land")
	}

	// Tear down the whole server — client will read EOF, attempt re-dial,
	// fail, escalate backoff each time.
	srv.closeFn()

	// Just verify the stream stays alive (no terminal error) for ~3s while
	// reconnect retries fail.
	select {
	case _, ok := <-stream.Frames():
		if !ok {
			// Channel closed — stream gave up. That's a regression: we
			// shouldn't surface dial failures as a terminal error.
			if err := stream.Err(); err != nil {
				t.Errorf("stream gave up with: %v (expected to keep retrying)", err)
			}
		}
	case <-time.After(3 * time.Second):
		// Good — still trying.
	}
}

func TestNextBackoff(t *testing.T) {
	cases := []struct {
		in, out time.Duration
	}{
		{0, reconnectStep},
		{reconnectStep, 2 * reconnectStep},
		{5 * time.Second, 6 * time.Second},
		{10 * time.Second, reconnectMax},
		{reconnectMax, reconnectMax},
		{2 * reconnectMax, reconnectMax}, // never exceed cap
	}
	for _, c := range cases {
		got := nextBackoff(c.in)
		if got != c.out {
			t.Errorf("nextBackoff(%s) = %s; want %s", c.in, got, c.out)
		}
	}
}

// closeWith writes a Close control frame with the given code on the latest
// conn and shuts down the read side, so the client's ReadMessage returns a
// CloseError with that exact code.
func (c *serverCtrl) closeWith(code int, reason string) {
	conn := c.latestConn()
	if conn == nil {
		return
	}
	_ = conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, reason),
		time.Now().Add(time.Second),
	)
	_ = conn.Close()
}

// TestSubscribe_Logout_SetsFinalErr — close code 4000 must surface through
// Stream.Err() so callers can distinguish a server kick from a clean
// caller-initiated Close.
func TestSubscribe_Logout_SetsFinalErr(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial subscribe")
	}

	srv.closeWith(closeCodeLogout, "logout")

	// Frames channel should close.
	select {
	case _, ok := <-stream.Frames():
		if ok {
			t.Fatal("expected channel close after logout")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("frames channel didn't close after logout")
	}
	// Err must surface errLogout.
	if !errors.Is(stream.Err(), errLogout) {
		t.Fatalf("Err() = %v, want errLogout", stream.Err())
	}
}

// TestSubscribe_TerminalCloseCode_NoReconnect — close codes that indicate
// a non-recoverable client/server disagreement (1008 policy violation,
// 1003 unsupported data, etc.) must NOT trigger a reconnect storm.
func TestSubscribe_TerminalCloseCode_NoReconnect(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial subscribe")
	}

	srv.closeWith(websocket.ClosePolicyViolation, "nope")

	select {
	case _, ok := <-stream.Frames():
		if ok {
			t.Fatal("expected channel close after policy violation")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("frames didn't close after terminal code")
	}

	// Should be exactly 1 conn — no reconnect attempted.
	srv.mu.Lock()
	n := len(srv.conns)
	srv.mu.Unlock()
	if n != 1 {
		t.Errorf("expected 1 connection (no reconnect on terminal code), got %d", n)
	}
	if stream.Err() == nil {
		t.Error("expected non-nil Err() after terminal close")
	}
}

// TestSubscribe_FlapEscalatesBackoff — when the server drops a connection
// quickly (well under stableThreshold), the next reconnect must NOT reset
// backoff to zero. Otherwise a bouncing server creates a tight 0-delay
// reconnect storm.
func TestSubscribe_FlapEscalatesBackoff(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	var attempts []int
	var backoffs []time.Duration
	var muLog sync.Mutex
	client := NewClient(srv.url, fakeJWK,
		WithSigner(nopSigner),
		// Keep stableThreshold at default so any quick reconnect counts
		// as a flap.
		WithLifecycleHandler(func(evt LifecycleEvent) {
			if evt.Kind == EventReconnecting {
				muLog.Lock()
				attempts = append(attempts, evt.Attempt)
				backoffs = append(backoffs, evt.Backoff)
				muLog.Unlock()
			}
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial conn")
	}

	// Flap: drop, wait for reconnect, drop again. Each cycle takes ms,
	// well below stableThreshold, so backoff must escalate. With linear
	// 1s/2s steps, total worst-case is ~3s for 3 flaps.
	for i := 0; i < 3; i++ {
		srv.dropLatest()
		if !srv.waitForConns(2+i, 5*time.Second) {
			t.Fatalf("reconnect %d didn't land", i+2)
		}
	}

	muLog.Lock()
	defer muLog.Unlock()
	if len(backoffs) < 2 {
		t.Fatalf("expected >=2 reconnect attempts logged, got %d", len(backoffs))
	}
	// All flap reconnects must be non-zero. The frontend resets to 0
	// only on a stable connection — see TestSubscribe_StableDisconnectResetsBackoff.
	for i, b := range backoffs {
		if b == 0 {
			t.Errorf("attempt %d backoff = 0 — flapping should escalate, not reset", i+1)
		}
	}
	// And the sequence must be non-decreasing — escalation, not regression.
	for i := 1; i < len(backoffs); i++ {
		if backoffs[i] < backoffs[i-1] {
			t.Errorf("attempt %d backoff = %s < attempt %d backoff %s — escalation went backwards", i+1, backoffs[i], i, backoffs[i-1])
		}
	}
}

// TestSubscribe_StableDisconnectResetsBackoff — a connection that lasted
// past stableThreshold and then drops must reconnect immediately (delay=0)
// even if a previous flap had escalated the counter. Use a tiny threshold
// (50ms) for deterministic test timing.
func TestSubscribe_StableDisconnectResetsBackoff(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	var backoffs []time.Duration
	var muLog sync.Mutex
	client := NewClient(srv.url, fakeJWK,
		WithSigner(nopSigner),
		WithStableThreshold(50*time.Millisecond),
		WithLifecycleHandler(func(evt LifecycleEvent) {
			if evt.Kind == EventReconnecting {
				muLog.Lock()
				backoffs = append(backoffs, evt.Backoff)
				muLog.Unlock()
			}
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial conn")
	}
	// Wait > stableThreshold so the next disconnect counts as stable.
	time.Sleep(100 * time.Millisecond)
	srv.dropLatest()
	if !srv.waitForConns(2, 3*time.Second) {
		t.Fatal("no reconnect")
	}

	muLog.Lock()
	defer muLog.Unlock()
	if len(backoffs) == 0 {
		t.Fatal("no reconnect attempts logged")
	}
	if backoffs[0] != 0 {
		t.Errorf("stable-disconnect first reconnect backoff = %s, want 0", backoffs[0])
	}
}

// TestClose_DuringActiveReconnect — Close while runLoop is mid-reconnect
// must return promptly, not deadlock waiting on the run loop.
func TestClose_DuringActiveReconnect(t *testing.T) {
	srv := startServer(t)

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial subscribe")
	}

	// Tear down the server entirely so the next reconnect attempt blocks
	// in dialAndSubscribe.
	srv.closeFn()
	srv.dropLatest()

	// Give the runLoop a moment to enter the dial-and-fail loop.
	time.Sleep(50 * time.Millisecond)

	done := make(chan struct{})
	go func() {
		_ = stream.Close()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("Close() didn't return — deadlock")
	}
}

// TestSubscribe_ParentCtxCancel — canceling the parent ctx (not via
// Close) must terminate the run loop. With the ctx-watcher in serveOnce,
// ReadMessage returns net.ErrClosed and runLoop exits cleanly.
func TestSubscribe_ParentCtxCancel(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	client := NewClient(srv.url, fakeJWK, WithSigner(nopSigner))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial conn")
	}

	cancel()

	// Frames must close within a reasonable window — without the
	// ctx-watcher, ReadMessage on the healthy socket would hang until
	// server idle timeout.
	select {
	case _, ok := <-stream.Frames():
		if ok {
			t.Fatal("expected frames channel to close after ctx cancel")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("frames didn't close after parent ctx cancel — runLoop may be stuck in ReadMessage")
	}
}

// TestSubscribe_TokenRefresh_4001 — close code 4001 must trigger a
// reconnect with exactly tokenRefreshBackoff (1s) delay, mirroring the
// frontend's special case.
func TestSubscribe_TokenRefresh_4001(t *testing.T) {
	srv := startServer(t)
	defer srv.closeFn()

	var firstBackoff time.Duration
	var firstAttempt int
	gotEvent := make(chan struct{}, 1)
	client := NewClient(srv.url, fakeJWK,
		WithSigner(nopSigner),
		WithLifecycleHandler(func(evt LifecycleEvent) {
			if evt.Kind == EventReconnecting && firstAttempt == 0 {
				firstBackoff = evt.Backoff
				firstAttempt = evt.Attempt
				select {
				case gotEvent <- struct{}{}:
				default:
				}
			}
		}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, Subscription{URI: "/x"})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer stream.Close()

	if !srv.waitForConns(1, time.Second) {
		t.Fatal("no initial conn")
	}

	srv.closeWith(closeCodeReauth, "refresh")

	select {
	case <-gotEvent:
	case <-time.After(3 * time.Second):
		t.Fatal("no reconnecting event after 4001")
	}

	if firstBackoff != tokenRefreshBackoff {
		t.Errorf("first reconnect backoff after 4001 = %s, want %s (tokenRefreshBackoff)", firstBackoff, tokenRefreshBackoff)
	}
}

func TestEncodeFilterToURI(t *testing.T) {
	cases := []struct {
		name   string
		base   string
		filter interface{}
		want   string
	}{
		{"nil filter", "/x", nil, "/x"},
		{"empty map", "/x", map[string]interface{}{}, "/x"},
		{"single string", "/x", map[string]interface{}{"a": "b"}, "/x?a=b"},
		{"array joined", "/x", map[string]interface{}{"k": []interface{}{"a", "b"}}, "/x?k=a%2Cb"},
		{"existing query", "/x?z=1", map[string]interface{}{"a": "b"}, "/x?z=1&a=b"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := encodeFilterToURI(c.base, c.filter)
			if err != nil {
				t.Fatal(err)
			}
			if got != c.want {
				t.Errorf("encodeFilterToURI(%q, %v) = %q; want %q", c.base, c.filter, got, c.want)
			}
		})
	}
}
