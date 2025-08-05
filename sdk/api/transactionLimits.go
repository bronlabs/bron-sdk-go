package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type TransactionLimitsAPI struct {
	http *http.Client
	workspaceID string
}

func NewTransactionLimitsAPI(http *http.Client, workspaceID string) *TransactionLimitsAPI {
	return &TransactionLimitsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *TransactionLimitsAPI) GetTransactionLimits() (*types.TransactionLimits, error) {
		path := fmt.Sprintf("/workspaces/%s/transaction-limits", api.workspaceID)
		var result types.TransactionLimits
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *TransactionLimitsAPI) GetTransactionLimitById(limitId string) (*types.TransactionLimit, error) {
		path := fmt.Sprintf("/workspaces/%s/transaction-limits/%s", api.workspaceID, limitId)
		var result types.TransactionLimit
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


