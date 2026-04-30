package sdk

import (
	"log/slog"
	nethttp "net/http"
	"time"

	"github.com/bronlabs/bron-sdk-go/sdk/api"
	"github.com/bronlabs/bron-sdk-go/sdk/auth"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/realtime"
)

type BronClientConfig struct {
	APIKey      string
	WorkspaceID string
	BaseURL     string
	// Proxy is an optional HTTP/HTTPS proxy URL applied to both REST and
	// WebSocket transport. Format: scheme://[user:pass@]host:port. Leave
	// empty to fall back to HTTP_PROXY / HTTPS_PROXY env vars.
	Proxy string
}

type BronClient struct {
	http        *http.Client
	workspaceID string
	baseURL     string
	apiKey      string

	// API clients
	Accounts          *api.AccountsAPI
	Balances          *api.BalancesAPI
	Transactions      *api.TransactionsAPI
	Addresses         *api.AddressesAPI
	Assets            *api.AssetsAPI
	Workspaces        *api.WorkspacesAPI
	TransactionLimits *api.TransactionLimitsAPI
	AddressBook       *api.AddressBookAPI
	Stake             *api.StakeAPI
}

// Backwards-compatible constructor
func NewBronClient(config BronClientConfig) *BronClient {
	return NewBronClientWithOptions(config)
}

// Functional options for DI
type ClientOption func(*clientOptions)

type clientOptions struct {
	stdHTTP     *nethttp.Client
	signer      func(auth.BronJwtOptions) (string, error)
	clock       func() time.Time
	retry       *http.RetryPolicy
	rtLifecycle func(realtime.LifecycleEvent)
	rtLogger    *slog.Logger
}

func WithNetHTTPClient(c *nethttp.Client) ClientOption {
	return func(o *clientOptions) { o.stdHTTP = c }
}

func WithSigner(s func(auth.BronJwtOptions) (string, error)) ClientOption {
	return func(o *clientOptions) { o.signer = s }
}

func WithClock(cl func() time.Time) ClientOption {
	return func(o *clientOptions) { o.clock = cl }
}

func WithRetryPolicy(p http.RetryPolicy) ClientOption {
	return func(o *clientOptions) { o.retry = &p }
}

// WithRealtimeLifecycleHandler registers a callback for WebSocket connection
// state changes (disconnect, reconnect attempts, reconnect success). Useful
// for stderr logging in CLIs and metrics in long-running services.
func WithRealtimeLifecycleHandler(fn func(realtime.LifecycleEvent)) ClientOption {
	return func(o *clientOptions) { o.rtLifecycle = fn }
}

// WithRealtimeLogger plugs a structured logger into the WebSocket transport.
// Set the handler to LevelDebug for frame-level tracing (each ping, each
// envelope, dial attempts). Default is silent.
func WithRealtimeLogger(l *slog.Logger) ClientOption {
	return func(o *clientOptions) { o.rtLogger = l }
}

func NewBronClientWithOptions(config BronClientConfig, opts ...ClientOption) *BronClient {
	co := &clientOptions{}
	for _, opt := range opts {
		opt(co)
	}

	var httpClient *http.Client
	if co.stdHTTP != nil {
		httpClient = http.NewClientWithHTTP(config.BaseURL, config.APIKey, co.stdHTTP)
	} else {
		httpClient = http.NewClient(config.BaseURL, config.APIKey)
	}
	if co.signer != nil {
		httpClient.SetSigner(co.signer)
	}
	if co.clock != nil {
		httpClient.SetClock(co.clock)
	}
	if co.retry != nil {
		httpClient.SetRetryPolicy(*co.retry)
	}

	client := &BronClient{
		http:        httpClient,
		workspaceID: config.WorkspaceID,
		baseURL:     config.BaseURL,
		apiKey:      config.APIKey,
	}

	rtOpts := []realtime.Option{realtime.WithProxy(config.Proxy)}
	if co.rtLifecycle != nil {
		rtOpts = append(rtOpts, realtime.WithLifecycleHandler(co.rtLifecycle))
	}
	if co.rtLogger != nil {
		rtOpts = append(rtOpts, realtime.WithLogger(co.rtLogger))
	}
	rt := realtime.NewClient(config.BaseURL, config.APIKey, rtOpts...)

	client.Accounts = api.NewAccountsAPI(httpClient, config.WorkspaceID)
	client.Balances = api.NewBalancesAPI(httpClient, config.WorkspaceID)
	client.Transactions = api.NewTransactionsAPI(httpClient, config.WorkspaceID, rt)
	client.Addresses = api.NewAddressesAPI(httpClient, config.WorkspaceID)
	client.Assets = api.NewAssetsAPI(httpClient, config.WorkspaceID)
	client.Workspaces = api.NewWorkspacesAPI(httpClient, config.WorkspaceID)
	client.TransactionLimits = api.NewTransactionLimitsAPI(httpClient, config.WorkspaceID)
	client.AddressBook = api.NewAddressBookAPI(httpClient, config.WorkspaceID)
	client.Stake = api.NewStakeAPI(httpClient, config.WorkspaceID)

	return client
}
