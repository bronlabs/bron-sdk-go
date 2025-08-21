# Bron SDK Go

Go SDK for the Bron API. This is a complete port of the TypeScript SDK to Go, maintaining the same structure and functionality.

## Features

- **Complete API Coverage**: All Bron API endpoints are supported
- **JWT Authentication**: Automatic JWT generation for API requests
- **Key Generation**: Built-in JWK key pair generation
- **Type Safety**: Strongly typed Go structs for all API responses
- **Code Generation**: Automatic code generation from OpenAPI spec
- **Optional Query Parameters**: No need to pass nil for empty queries
- **Proper Return Types**: All methods return strongly typed responses
- **Testing**: Comprehensive test suite

## Installation

```bash
go get github.com/bronlabs/bron-sdk-go
```

### Example

```go
package main

import (
	"log"
	"os"
	"context"
  
	"github.com/bronlabs/bron-sdk-go/sdk"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client := sdk.NewBronClient(sdk.BronClientConfig{
		APIKey:      os.Getenv("BRON_API_KEY"),
		WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"),
	})

	ctx := context.Background()

	// Just change these values:
	accountID := "your_account_id" // Your account ID
	toAddress := "0x..."           // Where to send
	amount := "0.001"              // How much to send
	symbol := "ETH"                // What to send (ETH, BRON, etc.)
	networkID := "testETH"         // Network (ETH=mainnet, testETH=testnet)

	// Create transaction - returns the created transaction
	txBody := types.NewWithdrawalTx(accountID, uuid.New().String(), types.WithdrawalParams{
		Amount:    amount,
		NetworkID: &networkID,
		Symbol:    &symbol,
		ToAddress: &toAddress,
	})
	tx, err := client.Transactions.CreateTransaction(ctx, txBody)

	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Printf("✅ Transaction created: %s", tx.TransactionID)
}
```

**Get Accounts & Balances:**

```go
ctx := context.Background()

// Get all accounts - no query parameters needed
accounts, err := client.Accounts.GetAccounts(ctx)
if err != nil {
  log.Fatal(err)
}

// Get all balances - no query parameters needed
balances, err := client.Balances.GetBalances(ctx)
if err != nil {
  log.Fatal(err)
}

// Get balances for first account with specific query
if len(accounts.Accounts) > 0 {
	account := accounts.Accounts[0]
	accountIDs := []string{account.AccountID}
	
	balances, err := client.Balances.GetBalances(ctx, &types.BalancesQuery{
		AccountIDs: &accountIDs,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, balance := range balances.Balances {
		log.Printf("Balance %s (%s): %s", balance.AssetID, balance.Symbol, balance.TotalBalance)
	}

	// Create transaction - returns the created transaction
	txBody := types.NewWithdrawalTx(account.AccountID, uuid.New().String(), types.WithdrawalParams{
		Amount:    "73.042",
		AssetID:   func() *string { v := "2"; return &v }(),
		ToAddress: func() *string { v := "0x428CdE5631142916F295d7bb2DA9d1b5f49F0eF9"; return &v }(),
	})
	tx, err := client.Transactions.CreateTransaction(ctx, txBody)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created transaction '%s': send %s", tx.TransactionID, tx.Params["amount"])
}
```

**More Examples:**

```go
ctx := context.Background()

// Get all transactions - no query parameters needed
transactions, err := client.Transactions.GetTransactions(ctx)
if err != nil {
  log.Fatal(err)
}

// Get filtered transactions with query parameters
limit := "10"
filteredTransactions, err := client.Transactions.GetTransactions(ctx, &types.TransactionsQuery{
	Limit: &limit,
})
if err != nil {
  log.Fatal(err)
}

// Get all assets - no query parameters needed
assets, err := client.Assets.GetAssets(ctx)
if err != nil {
  log.Fatal(err)
}

// Create address book record - returns the created record
record, err := client.AddressBook.CreateAddressBookRecord(ctx, types.CreateAddressBookRecord{
	Name:      "My Address",
	Address:   "0x428CdE5631142916F295d7bb2DA9d1b5f49F0eF9",
	NetworkID: "testETH",
})
if err != nil {
  log.Fatal(err)
}

log.Printf("Created address book record: %s", record.RecordID)
```

## Context requirement

All public API methods now take `context.Context` as the first argument. For simple usage, pass `context.Background()`. Use `context.WithTimeout`/`context.WithCancel` to control deadlines or cancellation across requests.

## Typed transaction params

Use typed builders to avoid `map[string]interface{}` for `params`:

```go
tx := types.NewWithdrawalTx(accountID, uuid.New().String(), types.WithdrawalParams{
    Amount:    "0.1",
    NetworkID: func() *string { v := "testETH"; return &v }(),
    Symbol:    func() *string { v := "ETH"; return &v }(),
    ToAddress: func() *string { v := "0x..."; return &v }(),
})
_, _ = client.Transactions.CreateTransaction(ctx, tx)
```

For advanced/custom cases not covered by typed params, use `json.RawMessage` via `types.NewCustomTx(...)`.

## Advanced configuration (functional options and DI)

You can inject your own transport, signer, and clock without changing call sites.

```go
// net/http client injection (timeouts, retries, proxy, TLS, tracing, etc.)
std := &http.Client{ Timeout: 10 * time.Second }
client := sdk.NewBronClientWithOptions(sdk.BronClientConfig{
  APIKey: os.Getenv("BRON_API_KEY"),
  WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"),
  BaseURL: os.Getenv("BRON_BASE_URL"),
}, sdk.WithNetHTTPClient(std))

// Custom signer and deterministic clock (useful for tests)
client = sdk.NewBronClientWithOptions(sdk.BronClientConfig{ /* ... */ },
  sdk.WithSigner(func(o auth.BronJwtOptions) (string, error) { return auth.GenerateBronJwt(o) }),
  sdk.WithClock(func() time.Time { return time.Unix(1_700_000_000, 0) }),
)

// Backwards compatible constructor still available:
// client := sdk.NewBronClient(config)
```

## Configuration

The SDK supports the following configuration options:

- `APIKey`: Your private JWK (required)
- `WorkspaceID`: Your workspace ID (required)
- `BaseURL`: API base URL (defaults to https://api.bron.org)

## Authentication

The SDK automatically handles JWT generation for API requests. You only need to provide your private JWK as the API key.

## Query Parameters

Query parameters are now optional! You can:

- **Call methods without parameters**: `client.Accounts.GetAccounts()`
- **Call methods with query parameters**: `client.Accounts.GetAccounts(&types.AccountsQuery{Limit: &limit})`

No need to pass `nil` when you don't want query parameters.

## Return Types

All API methods return strongly typed responses:

- **GET methods**: Return `(*types.ResponseType, error)`
- **POST methods**: Return `(*types.CreatedType, error)` (e.g., `CreateTransaction` returns `*types.Transaction`)
- **Methods without response body**: Return `error`

## Error Handling

All API methods return errors that should be checked:

```go
accounts, err := client.Accounts.GetAccounts()
if err != nil {
  log.Printf("API error: %v", err)
  return
}
```

## Retries

Enable exponential backoff + jitter retries (idempotent GET; honors Retry-After):

```go
// High-level
client := sdk.NewBronClientWithOptions(cfg,
  sdk.WithRetryPolicy(http.RetryPolicy{Max: 3, Base: 200 * time.Millisecond}),
)

// Low-level
hc := http.NewClient(cfg.BaseURL, cfg.APIKey)
hc.SetRetryPolicy(http.RetryPolicy{Max: 3, Base: 200 * time.Millisecond})
```

## Errors

Structured errors are returned as `*http.APIError` on non-2xx responses.

```go
resp, err := client.Workspaces.GetWorkspaceByID(ctx)
if err != nil {
  var apiErr *http.APIError
  if errors.As(err, &apiErr) {
    log.Printf("API error: status=%d code=%s requestID=%s msg=%s", apiErr.Status, apiErr.Code, apiErr.RequestID, apiErr.Message)
  } else {
    log.Printf("Unexpected error: %v", err)
  }
  return
}
```

## License

MIT License - see LICENSE file for details. 