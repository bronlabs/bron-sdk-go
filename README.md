# Bron SDK Go

Go SDK for the Bron API. This is a complete port of the TypeScript SDK to Go, maintaining the same structure and functionality.

## Features

- **Complete API Coverage**: All Bron API endpoints are supported
- **JWT Authentication**: Automatic JWT generation for API requests
- **Key Generation**: Built-in JWK key pair generation
- **Type Safety**: Strongly typed Go structs for all API responses
- **Code Generation**: Automatic code generation from OpenAPI spec
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

	client := sdk.NewBronClient(bron.BronClientConfig{
		APIKey:      os.Getenv("BRON_API_KEY"),
		WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"),
	})

	// Just change these values:
	accountID := "your_account_id" // Your account ID
	toAddress := "0x..."           // Where to send
	amount := "0.001"              // How much to send
	symbol := "ETH"                // What to send (ETH, BRON, etc.)
	networkId := "testETH"         // Network (ETH=mainnet, testETH=testnet)

	result, err := client.Transactions.CreateTransaction(types.CreateTransaction{
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

	log.Println("✅ Transaction sent!")
	log.Printf("Response: %+v", result)
}
```

**Get Accounts & Balances:**

```go
// Get all accounts
accounts, err := client.Accounts.GetAccounts(nil)
if err != nil {
  log.Fatal(err)
}

// Get balances
balances, err := client.Balances.GetBalances(nil)
if err != nil {
  log.Fatal(err)
}

	// Get balances for first account
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

		// Create transaction
		result, err := client.Transactions.CreateTransaction(types.CreateTransaction{
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

		log.Printf("Created transaction response: %+v", result)
	}
}
```

## Configuration

The SDK supports the following configuration options:

- `APIKey`: Your private JWK (required)
- `WorkspaceID`: Your workspace ID (required)
- `BaseURL`: API base URL (defaults to https://api.bron.org)

## Authentication

The SDK automatically handles JWT generation for API requests. You only need to provide your private JWK as the API key.

## Error Handling

All API methods return `(interface{}, error)` where the first value is the raw API response and the second is any error. Errors should always be checked:

```go
result, err := client.Accounts.GetAccounts(nil)
if err != nil {
  log.Printf("API error: %v", err)
  return
}
log.Printf("Response: %+v", result)
```

## License

MIT License - see LICENSE file for details. 