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

	// Just change these values:
	accountID := "your_account_id" // Your account ID
	toAddress := "0x..."           // Where to send
	amount := "0.001"              // How much to send
	symbol := "ETH"                // What to send (ETH, BRON, etc.)
	networkId := "testETH"         // Network (ETH=mainnet, testETH=testnet)

	// Create transaction - returns the created transaction
	tx, err := client.Transactions.CreateTransaction(types.CreateTransaction{
		ExternalId:      uuid.New().String(),
		AccountId:       accountID,
		TransactionType: "withdrawal",
		Params: map[string]interface{}{
			"amount":    amount,
			"networkId": networkId,
			"symbol":    symbol,
			"toAddress": toAddress,
		},
	})

	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Printf("✅ Transaction created: %s", tx.TransactionId)
}
```

**Get Accounts & Balances:**

```go
// Get all accounts - no query parameters needed
accounts, err := client.Accounts.GetAccounts()
if err != nil {
  log.Fatal(err)
}

// Get all balances - no query parameters needed
balances, err := client.Balances.GetBalances()
if err != nil {
  log.Fatal(err)
}

// Get balances for first account with specific query
if len(accounts.Accounts) > 0 {
	account := accounts.Accounts[0]
	accountIds := []string{account.AccountId}
	
	balances, err := client.Balances.GetBalances(&types.BalancesQuery{
		AccountIds: &accountIds,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, balance := range balances.Balances {
		log.Printf("Balance %s (%s): %s", balance.AssetId, balance.Symbol, balance.TotalBalance)
	}

	// Create transaction - returns the created transaction
	tx, err := client.Transactions.CreateTransaction(types.CreateTransaction{
		AccountId:       account.AccountId,
		ExternalId:      uuid.New().String(),
		TransactionType: "withdrawal",
		Params: map[string]interface{}{
			"amount":    "73.042",
			"assetId":   "2",
			"toAddress": "0x428CdE5631142916F295d7bb2DA9d1b5f49F0eF9",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created transaction '%s': send %s", tx.TransactionId, tx.Params["amount"])
}
```

**More Examples:**

```go
// Get all transactions - no query parameters needed
transactions, err := client.Transactions.GetTransactions()
if err != nil {
  log.Fatal(err)
}

// Get filtered transactions with query parameters
limit := "10"
filteredTransactions, err := client.Transactions.GetTransactions(&types.TransactionsQuery{
	Limit: &limit,
})
if err != nil {
  log.Fatal(err)
}

// Get all assets - no query parameters needed
assets, err := client.Assets.GetAssets()
if err != nil {
  log.Fatal(err)
}

// Create address book record - returns the created record
record, err := client.AddressBook.CreateAddressBookRecord(types.CreateAddressBookRecord{
	Name:      "My Address",
	Address:   "0x428CdE5631142916F295d7bb2DA9d1b5f49F0eF9",
	NetworkId: "testETH",
})
if err != nil {
  log.Fatal(err)
}

log.Printf("Created address book record: %s", record.RecordId)
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

## License

MIT License - see LICENSE file for details. 