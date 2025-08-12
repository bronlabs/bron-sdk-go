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

func (api *AssetsAPI) GetAssets(query ...*types.AssetsQuery) (*types.Assets, error) {
	path := fmt.Sprintf("/dictionary/assets")
	var result types.Assets
	var queryParam *types.AssetsQuery
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

func (api *AssetsAPI) GetAssetById(assetId string, query ...*types.AssetByIdQuery) (*types.Asset, error) {
	path := fmt.Sprintf("/dictionary/assets/%s", assetId)
	var result types.Asset
	var queryParam *types.AssetByIdQuery
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

func (api *AssetsAPI) GetNetworks(query ...*types.NetworksQuery) (*types.Networks, error) {
	path := fmt.Sprintf("/dictionary/networks")
	var result types.Networks
	var queryParam *types.NetworksQuery
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

func (api *AssetsAPI) GetNetworkById(networkId string) (*types.Network, error) {
	path := fmt.Sprintf("/dictionary/networks/%s", networkId)
	var result types.Network
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetPrices(query ...*types.PricesQuery) (*types.SymbolMarketPrices, error) {
	path := fmt.Sprintf("/dictionary/symbol-market-prices")
	var result types.SymbolMarketPrices
	var queryParam *types.PricesQuery
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

func (api *AssetsAPI) GetSymbols(query ...*types.SymbolsQuery) (*types.Symbols, error) {
	path := fmt.Sprintf("/dictionary/symbols")
	var result types.Symbols
	var queryParam *types.SymbolsQuery
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

func (api *AssetsAPI) GetSymbolById(symbolId string, query ...*types.SymbolByIdQuery) (*types.Symbol, error) {
	path := fmt.Sprintf("/dictionary/symbols/%s", symbolId)
	var result types.Symbol
	var queryParam *types.SymbolByIdQuery
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

