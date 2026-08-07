package types

type WorkspacePricesQuery struct {
	BaseSymbolIDs *[]string `json:"baseSymbolIds,omitempty"`
	BaseAssetIDs *[]string `json:"baseAssetIds,omitempty"`
	Used *bool `json:"used,omitempty"`
}
