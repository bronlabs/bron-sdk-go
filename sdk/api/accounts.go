package api

import (
	"fmt"

	"context"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type AccountsAPI struct {
	http        *http.Client
	workspaceID string
}

func NewAccountsAPI(http *http.Client, workspaceID string) *AccountsAPI {
	return &AccountsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *AccountsAPI) GetAccounts(ctx context.Context, query ...*types.AccountsQuery) (*types.Accounts, error) {
	path := fmt.Sprintf("/workspaces/%s/accounts", api.workspaceID)
	var result types.Accounts
	var queryParam *types.AccountsQuery
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

func (api *AccountsAPI) GetAccountByID(ctx context.Context, accountId string) (*types.Account, error) {
	path := fmt.Sprintf("/workspaces/%s/accounts/%s", api.workspaceID, accountId)
	var result types.Account
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

