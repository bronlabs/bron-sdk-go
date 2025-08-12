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

func (api *BalancesAPI) GetBalances(query ...*types.BalancesQuery) (*types.Balances, error) {
	path := fmt.Sprintf("/workspaces/%s/balances", api.workspaceID)
	var result types.Balances
	var queryParam *types.BalancesQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
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

