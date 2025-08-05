package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/src/types"
	"github.com/bronlabs/bron-sdk-go/src/http"
)

type AccountsAPI struct {
	http *http.Client
	workspaceID string
}

func NewAccountsAPI(http *http.Client, workspaceID string) *AccountsAPI {
	return &AccountsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *AccountsAPI) RetrieveAccountById(accountId string) (*types.Account, error) {
		path := fmt.Sprintf("/workspaces/%s/accounts/%s", api.workspaceID, accountId)
		var result types.Account
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *AccountsAPI) GetAccounts() (*types.Accounts, error) {
		path := fmt.Sprintf("/workspaces/%s/accounts", api.workspaceID)
		var result types.Accounts
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


