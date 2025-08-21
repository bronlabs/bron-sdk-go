package api

import (
	"fmt"

	"context"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type TransactionLimitsAPI struct {
	http        *http.Client
	workspaceID string
}

func NewTransactionLimitsAPI(http *http.Client, workspaceID string) *TransactionLimitsAPI {
	return &TransactionLimitsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *TransactionLimitsAPI) GetTransactionLimits(ctx context.Context, query ...*types.TransactionLimitsQuery) (*types.TransactionLimits, error) {
	path := fmt.Sprintf("/workspaces/%s/transaction-limits", api.workspaceID)
	var result types.TransactionLimits
	var queryParam *types.TransactionLimitsQuery
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

func (api *TransactionLimitsAPI) GetTransactionLimitByID(ctx context.Context, limitId string) (*types.TransactionLimit, error) {
	path := fmt.Sprintf("/workspaces/%s/transaction-limits/%s", api.workspaceID, limitId)
	var result types.TransactionLimit
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

