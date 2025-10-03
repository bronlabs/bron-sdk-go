package api

import (
	"fmt"

	"context"
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

func (api *TransactionsAPI) GetTransactions(ctx context.Context, query ...*types.TransactionsQuery) (*types.Transactions, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	var result types.Transactions
	var queryParam *types.TransactionsQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateTransaction(ctx context.Context, body types.CreateTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions", api.workspaceID)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateMultipleTransactions(ctx context.Context, body types.CreateTransactions) (*types.Transactions, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/bulk-create", api.workspaceID)
	var result types.Transactions
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) DryRunTransaction(ctx context.Context, body types.CreateTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/dry-run", api.workspaceID)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionByID(ctx context.Context, transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CancelTransaction(ctx context.Context, transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/cancel", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CreateSigningRequest(ctx context.Context, transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/create-signing-request", api.workspaceID, transactionId)
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionEvents(ctx context.Context, transactionId string) (*types.TransactionEvents, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/events", api.workspaceID, transactionId)
	var result types.TransactionEvents
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

