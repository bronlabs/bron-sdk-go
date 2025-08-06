package types

type SymbolByIdQuery struct {
	SymbolIds *[]string `json:"symbolIds,omitempty"`
	AssetIds *[]string `json:"assetIds,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
