package api

import (
	"fmt"
	"net/url"

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
	path := fmt.Sprintf("/workspaces/%s/transactions", url.PathEscape(api.workspaceID))
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
	path := fmt.Sprintf("/workspaces/%s/transactions", url.PathEscape(api.workspaceID))
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
	path := fmt.Sprintf("/workspaces/%s/transactions/bulk-create", url.PathEscape(api.workspaceID))
	var result types.Transactions
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) DryRunTransaction(ctx context.Context, body types.CreateTransaction) (*types.DryRunTransaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/dry-run", url.PathEscape(api.workspaceID))
	var result types.DryRunTransaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionByID(ctx context.Context, transactionId string) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) AcceptDepositOffer(ctx context.Context, transactionId string, body types.OfferActions) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/accept-deposit-offer", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) ApproveTransaction(ctx context.Context, transactionId string, body types.ApproveTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/approve", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) CancelTransaction(ctx context.Context, transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/cancel", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
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
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/create-signing-request", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) DeclineTransaction(ctx context.Context, transactionId string, body types.CancelTransaction) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/decline", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) GetTransactionEvents(ctx context.Context, transactionId string) (*types.TransactionEvents, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/events", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.TransactionEvents
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *TransactionsAPI) RejectOutgoingOffer(ctx context.Context, transactionId string, body types.OfferActions) (*types.Transaction, error) {
	path := fmt.Sprintf("/workspaces/%s/transactions/%s/reject-outgoing-offer", url.PathEscape(api.workspaceID), url.PathEscape(transactionId))
	var result types.Transaction
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

