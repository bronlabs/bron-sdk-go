package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type AddressBookAPI struct {
	http        *http.Client
	workspaceID string
}

func NewAddressBookAPI(http *http.Client, workspaceID string) *AddressBookAPI {
	return &AddressBookAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *AddressBookAPI) GetAddressBookRecords(query *types.AddressBookRecordsQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AddressBookAPI) CreateAddressBookRecord(body types.CreateAddressBookRecord) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AddressBookAPI) DeactivateAddressBookRecord(recordId string) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
	var result interface{}
	options := http.RequestOptions{
		Method: "DELETE",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AddressBookAPI) GetAddressBookRecordById(recordId string) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return result, err
}

