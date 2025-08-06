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

func (api *AssetsAPI) GetAssets(query *types.AssetsQuery) (*types.Assets, error) {
	path := fmt.Sprintf("/dictionary/assets")
	var result types.Assets
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetAssetById(assetId string, query *types.AssetByIdQuery) (*types.Asset, error) {
	path := fmt.Sprintf("/dictionary/assets/%s", assetId)
	var result types.Asset
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetNetworks(query *types.NetworksQuery) (*types.Networks, error) {
	path := fmt.Sprintf("/dictionary/networks")
	var result types.Networks
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
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

func (api *AssetsAPI) GetPrices(query *types.PricesQuery) (*types.SymbolMarketPrices, error) {
	path := fmt.Sprintf("/dictionary/symbol-market-prices")
	var result types.SymbolMarketPrices
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetSymbols(query *types.SymbolsQuery) (*types.Symbols, error) {
	path := fmt.Sprintf("/dictionary/symbols")
	var result types.Symbols
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetSymbolById(symbolId string, query *types.SymbolByIdQuery) (*types.Symbol, error) {
	path := fmt.Sprintf("/dictionary/symbols/%s", symbolId)
	var result types.Symbol
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

