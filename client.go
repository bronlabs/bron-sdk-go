package bron

import (
	"github.com/bronlabs/bron-sdk-go/sdk/api"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type BronClientConfig struct {
	APIKey      string
	WorkspaceID string
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

func NewBronClient(config BronClientConfig) *BronClient {
	baseURL := "https://api.bron.org"
	httpClient := http.NewClient(baseURL, config.APIKey)

	client := &BronClient{
		http:        httpClient,
		workspaceID: config.WorkspaceID,
		baseURL:     baseURL,
		apiKey:      config.APIKey,
	}

	// Initialize API clients
	client.Accounts = api.NewAccountsAPI(httpClient, config.WorkspaceID)
	client.Balances = api.NewBalancesAPI(httpClient, config.WorkspaceID)
	client.Transactions = api.NewTransactionsAPI(httpClient, config.WorkspaceID)
	client.Addresses = api.NewAddressesAPI(httpClient, config.WorkspaceID)
	client.Assets = api.NewAssetsAPI(httpClient, config.WorkspaceID)
	client.Workspaces = api.NewWorkspacesAPI(httpClient, config.WorkspaceID)
	client.TransactionLimits = api.NewTransactionLimitsAPI(httpClient, config.WorkspaceID)
	client.AddressBook = api.NewAddressBookAPI(httpClient, config.WorkspaceID)
	client.Stake = api.NewStakeAPI(httpClient, config.WorkspaceID)

	return client
}
