package api

import (
	"fmt"

	"context"
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

func (api *AddressBookAPI) GetAddressBookRecords(ctx context.Context, query ...*types.AddressBookRecordsQuery) (*types.AddressBookRecords, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
	var result types.AddressBookRecords
	var queryParam *types.AddressBookRecordsQuery
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

func (api *AddressBookAPI) CreateAddressBookRecord(ctx context.Context, body types.CreateAddressBookRecord) (*types.AddressBookRecord, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records", api.workspaceID)
	var result types.AddressBookRecord
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *AddressBookAPI) DeactivateAddressBookRecord(ctx context.Context, recordId string) (*types.Unit, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
	var result types.Unit
	options := http.RequestOptions{
		Method: "DELETE",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *AddressBookAPI) GetAddressBookRecordByID(ctx context.Context, recordId string) (*types.AddressBookRecord, error) {
	path := fmt.Sprintf("/workspaces/%s/address-book-records/%s", api.workspaceID, recordId)
	var result types.AddressBookRecord
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

