package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type BalancesAPI struct {
	http        *http.Client
	workspaceID string
}

func NewBalancesAPI(http *http.Client, workspaceID string) *BalancesAPI {
	return &BalancesAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *BalancesAPI) GetBalances(query *types.BalancesQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/balances", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *BalancesAPI) GetBalanceById(balanceId string) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/balances/%s", api.workspaceID, balanceId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return result, err
}

