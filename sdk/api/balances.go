package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type BalancesAPI struct {
	http *http.Client
	workspaceID string
}

func NewBalancesAPI(http *http.Client, workspaceID string) *BalancesAPI {
	return &BalancesAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *BalancesAPI) GetBalances() (*types.Balances, error) {
		path := fmt.Sprintf("/workspaces/%s/balances", api.workspaceID)
		var result types.Balances
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *BalancesAPI) GetBalanceById(balanceId string) (*types.Balance, error) {
		path := fmt.Sprintf("/workspaces/%s/balances/%s", api.workspaceID, balanceId)
		var result types.Balance
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


