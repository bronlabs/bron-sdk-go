package bron

import (
	"github.com/bronlabs/bron-sdk-go/src/api"
	"github.com/bronlabs/bron-sdk-go/src/http"
)

type BronClientConfig struct {
	APIKey      string
	WorkspaceID string
	BaseURL     string
}

type BronClient struct {
	WorkspaceID string
	http        *http.Client

	// API instances
	Accounts          *api.AccountsAPI
	Transactions      *api.TransactionsAPI
	Balances          *api.BalancesAPI
	Workspaces        *api.WorkspacesAPI
	AddressBook       *api.AddressBookAPI
	Assets            *api.AssetsAPI
	Addresses         *api.AddressesAPI
	TransactionLimits *api.TransactionLimitsAPI
	Stake             *api.StakeAPI
}

func NewBronClient(config BronClientConfig) *BronClient {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.bron.org"
	}

	httpClient := http.NewClient(baseURL, config.APIKey)

	client := &BronClient{
		WorkspaceID: config.WorkspaceID,
		http:        httpClient,
	}

	// Initialize API instances
	client.Accounts = api.NewAccountsAPI(httpClient, config.WorkspaceID)
	client.Transactions = api.NewTransactionsAPI(httpClient, config.WorkspaceID)
	client.Balances = api.NewBalancesAPI(httpClient, config.WorkspaceID)
	client.Workspaces = api.NewWorkspacesAPI(httpClient, config.WorkspaceID)
	client.AddressBook = api.NewAddressBookAPI(httpClient, config.WorkspaceID)
	client.Assets = api.NewAssetsAPI(httpClient, config.WorkspaceID)
	client.Addresses = api.NewAddressesAPI(httpClient, config.WorkspaceID)
	client.TransactionLimits = api.NewTransactionLimitsAPI(httpClient, config.WorkspaceID)
	client.Stake = api.NewStakeAPI(httpClient, config.WorkspaceID)

	return client
}
