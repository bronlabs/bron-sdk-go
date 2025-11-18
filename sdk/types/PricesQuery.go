package types

type PricesQuery struct {
	BaseSymbolIDs *[]string `json:"baseSymbolIds,omitempty"`
	BaseAssetIDs *[]string `json:"baseAssetIds,omitempty"`
}
