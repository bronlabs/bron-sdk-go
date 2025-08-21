package types

type SymbolsQuery struct {
	SymbolIDs *[]string `json:"symbolIds,omitempty"`
	AssetIDs *[]string `json:"assetIds,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
