# bron-sdk-go

Go SDK for the [Bron](https://bron.org) API. Auto-generated from the OpenAPI spec; thin layer on top — typed models, JWT signing, retries, structured errors. Looking for a CLI instead? See [`bron-cli`](https://github.com/bronlabs/bron-cli).

## Install

```bash
go get github.com/bronlabs/bron-sdk-go
```

## Authenticate

Generate a P-256 keypair, paste the **public** half into the Bron UI (API keys), pass the **private** JWK to the client.

```bash
go run github.com/bronlabs/bron-sdk-go/cmd/keygen
```

```go
client := sdk.NewBronClient(sdk.BronClientConfig{
    APIKey:      os.Getenv("BRON_API_KEY"),       // private JWK
    WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"),
    BaseURL:     os.Getenv("BRON_BASE_URL"),      // optional, default https://api.bron.org
})
```

The SDK signs every request with a short-lived JWT (ES256). No token caching, no revocation flow.

## Use

```go
ctx := context.Background()

accounts, _ := client.Accounts.GetAccounts(ctx)                                            // no filters
accounts, _  = client.Accounts.GetAccounts(ctx, &types.AccountsQuery{Limit: ptr("50")})    // with filters

netID, asset, addr := "ETH", "5000", "0x428C..."
tx, err := client.Transactions.CreateTransaction(ctx, types.CreateTransaction{
    AccountID:       accounts.Accounts[0].AccountID,
    ExternalID:      uuid.New().String(),
    TransactionType: types.TransactionType_WITHDRAWAL,
    Params: types.WithdrawalParams{
        Amount: "0.1", AssetID: &asset, NetworkID: &netID, ToAddress: &addr,
    },
})
```

- **Resources:** `client.Accounts`, `Balances`, `Transactions`, `AddressBook`, `Assets`, `Networks`, `Symbols`, `Stakes`, `Workspaces`, …
- **Method shape:** `ctx` first; query struct is optional.
- **Returns:** `Get*`/`Create*` → `(*types.Resource, error)`; void ops → `error`.
- **Typed params per `transactionType`:** `WithdrawalParams`, `AllowanceParams`, `StakeDelegationParams`, `FiatOutParams`, … `Params` is `interface{}` — pass any struct (or `json.RawMessage`) for unsupported shapes.

## Subscribe (WebSocket)

A subscription is "GET extended": same query as the matching list endpoint,
the server replays the historical match as the first frame, then keeps the
connection open and pushes live updates as additional frames of the same
response shape (a list with one element per change in steady state).

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

stream, err := client.Transactions.Subscribe(ctx, &types.TransactionsQuery{
    TransactionStatuses: &[]types.TransactionStatus{
        types.TransactionStatus_SIGNING_REQUIRED,
        types.TransactionStatus_WAITING_APPROVAL,
    },
})
if err != nil { /* ... */ }
defer stream.Close()

for batch := range stream.Updates() {
    for _, tx := range batch.Transactions {
        log.Printf("tx %s → %s", tx.TransactionID, tx.Status)
    }
}
if err := stream.Err(); err != nil {
    log.Printf("stream ended: %v", err)
}
```

- **First frame** = historical match (same shape as `GetTransactions`).
- **Subsequent frames** = live updates, each typically a single-element list. Filters apply to both phases.
- **Skip the initial replay**: pass `Limit: ptr("0")`, or use `SubscribeWithFilter(ctx, map[string]interface{}{"limit": 0, ...})` when the typed `*string` field doesn't carry through (backend wants an integer).
- **Always `defer stream.Close()`** — sends `UNSUBSCRIBE` and tears down the WebSocket. Channel closes when ctx is cancelled, `Close` is called, or the connection drops; check `stream.Err()` after for the cause.
- **Proxy** is honored automatically via `BronClientConfig.Proxy` (or `HTTP_PROXY` / `HTTPS_PROXY` env vars). Inject a custom `*websocket.Dialer` via `realtime.WithDialer` if you need TLS or NetDial overrides.

### Auto-reconnect

The transport keeps the subscription alive across server-initiated disconnects:

- **Ping every 15s** so the server's ~60s idle timeout never fires.
- **Transparent re-dial + re-SUBSCRIBE** on close 1006 / network drops / server restart, reusing the same `Correlation-Id` so log correlation stays stable.
- **Linear backoff** 1s → 2s → … → 10s on flapping or dial failures; reset to 0 once a connection has been stable for 30s.
- **Close code 4000** (logout) ends the stream permanently — `stream.Err()` returns the sentinel.
- **Close code 4001** (token refresh) reconnects with a fixed 1s delay.

On every reconnect the server replays the snapshot frame again, so callers
should be prepared to see duplicates. Use `Limit: ptr("0")` (or
`SubscribeWithFilter` with `"limit": 0`) when only live updates matter.

### Observability

```go
import "log/slog"

client := sdk.NewBronClientWithOptions(cfg,
    // Structured logs to stderr, frame-level if you set Level: LevelDebug.
    sdk.WithRealtimeLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }))),

    // Or a programmatic hook for metrics/Sentry/whatever.
    sdk.WithRealtimeLifecycleHandler(func(evt realtime.LifecycleEvent) {
        switch evt.Kind {
        case realtime.EventReconnecting:
            metrics.Inc("ws.reconnecting", "uri", evt.URI)
        case realtime.EventReconnected:
            metrics.Inc("ws.reconnected", "attempts", evt.Attempt)
        }
    }),
)
```

`LifecycleEvent` carries `CorrelationID`, `URI`, `Attempt`, `Backoff`, and
the triggering `Err`. Authorization tokens never appear in logs.

## Errors

Non-2xx responses come back as `*http.APIError` (`Status`, `Code`, `Message`, `RequestID`):

```go
var apiErr *http.APIError
if errors.As(err, &apiErr) { /* ... */ }
```

## Retries

Exponential backoff + jitter for idempotent verbs, honors `Retry-After`. Off by default.

```go
client := sdk.NewBronClientWithOptions(cfg,
    sdk.WithRetryPolicy(http.RetryPolicy{Max: 3, Base: 200 * time.Millisecond}),
)
```

## Advanced

Inject your own transport, signer, or clock — useful for proxies, tracing, deterministic tests:

```go
client := sdk.NewBronClientWithOptions(cfg,
    sdk.WithNetHTTPClient(&http.Client{Timeout: 10 * time.Second}),
    sdk.WithSigner(func(o auth.BronJwtOptions) (string, error) { return auth.GenerateBronJwt(o) }),
    sdk.WithClock(func() time.Time { return time.Unix(1_700_000_000, 0) }),
)
```

## Contributors

```bash
make generate    # regenerate sdk/types + sdk/api from bron-open-api-public.json
make build
make test
make publish     # bumps patch, tags, pushes (release)
```

Keep `bron-open-api-public.json` in sync with the upstream public-api before regenerating.

## License

MIT License - see LICENSE file for details.
