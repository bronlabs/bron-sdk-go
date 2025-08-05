package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type AddressesAPI struct {
	http        *http.Client
	workspaceID string
}

func NewAddressesAPI(http *http.Client, workspaceID string) *AddressesAPI {
	return &AddressesAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *AddressesAPI) GetDepositAddresses(query *types.AddressesQuery) (*types.Addresses, error) {
	path := fmt.Sprintf("/workspaces/%s/addresses", api.workspaceID)
	var result types.Addresses
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

