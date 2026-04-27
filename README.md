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

MIT.
