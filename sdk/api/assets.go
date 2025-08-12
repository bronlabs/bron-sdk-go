package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type AssetsAPI struct {
	http        *http.Client
	workspaceID string
}

func NewAssetsAPI(http *http.Client, workspaceID string) *AssetsAPI {
	return &AssetsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *AssetsAPI) GetAssets(query *types.AssetsQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/assets")
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetAssetById(assetId string, query *types.AssetByIdQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/assets/%s", assetId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetNetworks(query *types.NetworksQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/networks")
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetNetworkById(networkId string) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/networks/%s", networkId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetPrices(query *types.PricesQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/symbol-market-prices")
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetSymbols(query *types.SymbolsQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/symbols")
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *AssetsAPI) GetSymbolById(symbolId string, query *types.SymbolByIdQuery) (interface{}, error) {
	path := fmt.Sprintf("/dictionary/symbols/%s", symbolId)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

