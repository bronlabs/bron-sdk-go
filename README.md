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

## Quick Start

### 1. Generate API Keys

```bash
go run cmd/keygen/main.go
```

This will output:

- **Public JWK** (send to Bron)
- **Private JWK** (keep safe)



### 2. Basic Usage

```sh
export BRON_API_KEY='{"kty":"EC","x":"VqW0Rzw4At***ADF2iFCzxc","y":"9AylQ7HHI0vRT0C***PqWuf2yT8","crv":"P-256","d":"DCQ0jrmYw8***9i64igNKuP0","kid":"cmdos3lj50000sayo6pl45zly"}'
export BRON_WORKSPACE_ID='htotobpkg7xqjfxenjid3n1o'
```

```go
package main

import (
	"log"
	"os"
	bron "github.com/bronlabs/bron-sdk-go/sdk"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client := bron.NewBronClient(bron.BronClientConfig{
		APIKey:      os.Getenv("BRON_API_KEY"),
		WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"),
	})

	// Get workspace
	workspace, err := client.Workspaces.GetWorkspaceById()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Workspace: %s", workspace.Name)

	// Get accounts
	accounts, err := client.Accounts.GetAccounts(nil)
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
		tx, err := client.Transactions.CreateTransaction(types.CreateTransaction{
			AccountId:       account.AccountId,
			ExternalId:      uuid.New().String(),
			TransactionType: "withdrawal",
			Params: map[string]interface{}{
				"amount":    "73.042",
				"assetId":   "2",
				"symbol":    "ETH",
				"networkId": "ETH",
				"toAddress": "0x428CdE5631142916F295d7bb2DA9d1b5f49F0eF9",
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created transaction '%s': send %s", tx.TransactionId, tx.Params["amount"])
	}
}
```

### Building

```bash
make build
```

### Key Generation

```bash
make generate-keys
```

## Configuration

The SDK supports the following configuration options:

- `APIKey`: Your private JWK (required)
- `WorkspaceID`: Your workspace ID (required)
- `BaseURL`: API base URL (defaults to https://api.bron.org)

## Authentication

The SDK automatically handles JWT generation for API requests. You only need to provide your private JWK as the API key.

## Error Handling

All API methods return errors that should be checked:

```go
accounts, err := client.Accounts().GetAccounts(nil)
if err != nil {
log.Printf("API error: %v", err)
return
}
```

## License

MIT License - see LICENSE file for details. 
