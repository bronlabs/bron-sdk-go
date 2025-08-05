package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type TransactionsAPI struct {
	http        *http.Client
	workspaceID string
}

func NewTransactionsAPI(http *http.Client, workspaceID string) *TransactionsAPI {
	return &TransactionsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *TransactionsAPI) GetTransactions() (*types.Transactions, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	var result types.Transactions
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateTransaction(body types.CreateTransaction) error {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	options := http.RequestOptions{
			Method: "POST",
		Path:   path,
		Body:   body,
	}
	return api.http.Request(nil, options)
}

func (api *TransactionsAPI) CreateMultipleTransactions(body types.CreateTransactions) error {
	path := fmt.Sprintf("/workspaces/%s/transactions/bulk-create", api.workspaceID)
	options := http.RequestOptions{
			Method: "POST",
		Path:   path,
		Body:   body,
	}
	return api.http.Request(nil, options)
}

func (api *TransactionsAPI) DryRunTransaction(body types.CreateTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/dry-run", api.workspaceID)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionById(transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *TransactionsAPI) CancelTransaction(transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/cancel", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateSigningRequest(transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/create-signing-request", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

