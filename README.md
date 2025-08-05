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

## Project Structure

```
bron-sdk-go/
├── src/
│   ├── client.go          # Main client
│   ├── api/               # API implementations
│   │   └── accounts.go
│   ├── types/             # Type definitions
│   │   ├── Account.go
│   │   ├── Transaction.go
│   │   └── Workspace.go
│   ├── http/              # HTTP client
│   │   └── client.go
│   └── utils/             # Utilities
│       ├── auth.go        # JWT generation
│       ├── keyGenerator.go # Key generation
│       └── generator.go   # Code generator
├── cmd/
│   ├── keygen/            # Key generation CLI
│   └── generator/         # Code generation CLI
├── test/                  # Tests
├── my-bron-app-go/        # Demo application
└── bron-open-api-public.json # OpenAPI spec
```

## Transaction API

Use the transaction API to send any token:

```go
// Send any token
err := client.Transactions.CreateTransaction(types.CreateTransaction{
ExternalId:      uuid.New().String(),
AccountId:       accountID,
TransactionType: "withdrawal",
Params: map[string]interface{}{
"amount":    amount,
"networkId": networkId, // "ETH" for mainnet, "testETH" for testnet
"symbol":    symbol, // "ETH", "BRON", etc.
"toAddress": toAddress,
},
})
```

## API Examples

### Accounts

```go
// Get all accounts
accounts, err := client.Accounts.GetAccounts()

// Get specific account
account, err := client.Accounts().RetrieveAccountById("account-id")
```

### Transactions

```go
// Create transaction
transaction, err := client.Transactions().CreateTransaction(bron.CreateTransaction{
ExternalID:      "unique-id",
AccountID:       "account-id",
TransactionType: "withdrawal",
Params: map[string]interface{}{
"amount":    "0.001",
"networkId": "testETH",
"symbol":    "ETH",
"toAddress": "0x...",
},
})
```

## Development

### Running Tests

```bash
go test ./test/...
```

### Code Generation

```bash
make generate
```

This will generate Go types and API methods from the OpenAPI specification.

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

## Demo Application

A complete demo application is included in `my-bron-app-go/`:

```bash
cd my-bron-app-go
go run demo.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

MIT License - see LICENSE file for details. 
