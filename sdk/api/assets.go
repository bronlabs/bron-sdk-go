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

func (api *AssetsAPI) GetNetworks() (*types.Networks, error) {
	path := fmt.Sprintf("/dictionary/networks", api.workspaceID)
	var result types.Networks
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetPrices() (*types.SymbolMarketPrices, error) {
	path := fmt.Sprintf("/dictionary/symbol-market-prices", api.workspaceID)
	var result types.SymbolMarketPrices
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetAssets() (*types.Assets, error) {
	path := fmt.Sprintf("/dictionary/assets", api.workspaceID)
	var result types.Assets
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *AssetsAPI) GetSymbols() (*types.Symbols, error) {
	path := fmt.Sprintf("/dictionary/symbols", api.workspaceID)
	var result types.Symbols
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.Request(&result, options)
	return &result, err
}
