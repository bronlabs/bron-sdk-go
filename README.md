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

To validate a JWK:

```bash
go run cmd/keygen/main.go --validate '{"kty":"EC",...}'
```

### 2. Basic Usage

**Send any token:**

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

	// Just change these values:
	accountID := "your_account_id" // Your account ID
	toAddress := "0x..."           // Where to send
	amount := "0.001"              // How much to send
	symbol := "ETH"                // What to send (ETH, BRON, etc.)
	networkId := "testETH"         // Network (ETH=mainnet, testETH=testnet)

	err := client.Transactions.CreateTransaction(types.CreateTransaction{
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
}
```

**Get Accounts & Balances:**

```go
// Get all accounts
accounts, err := client.Accounts.GetAccounts()
if err != nil {
log.Fatal(err)
}

// Get balances
balances, err := client.Balances.GetBalances()
if err != nil {
log.Fatal(err)
}

// Get specific account balance
balance, err := client.GetAccountBalance("account_id")
if err != nil {
log.Fatal(err)
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