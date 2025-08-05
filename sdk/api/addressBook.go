package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type AddressBookAPI struct {
	http *http.Client
	workspaceID string
}

func NewAddressBookAPI(http *http.Client, workspaceID string) *AddressBookAPI {
	return &AddressBookAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *AddressBookAPI) GetAddressBookRecords() (*types.AddressBookRecords, error) {
		path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
		var result types.AddressBookRecords
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *AddressBookAPI) CreateAddressBookRecord(body types.CreateAddressBookRecord) error {
		path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
		options := http.RequestOptions{
			Method: "POST",
			Path:   path,
			Body:   body,
		}
		return api.http.Request(nil, options)
	}


	func (api *AddressBookAPI) DeactivateAddressBookRecord(recordId string) (*types.Unit, error) {
		path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
		var result types.Unit
		options := http.RequestOptions{
			Method: "DELETE",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *AddressBookAPI) GetAddressBookRecordById(recordId string) (*types.AddressBookRecord, error) {
		path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
		var result types.AddressBookRecord
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


