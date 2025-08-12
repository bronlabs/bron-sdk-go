package api

import (
	"fmt"

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

func (api *TransactionLimitsAPI) GetTransactionLimits(query *types.TransactionLimitsQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/transaction-limits", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *TransactionLimitsAPI) GetTransactionLimitById(limitId string) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/transaction-limits/%s", api.workspaceID, limitId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return result, err
}

