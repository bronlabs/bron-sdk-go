package types

type AssetsQuery struct {
	AssetIds *[]string `json:"assetIds,omitempty"`
	NetworkIds *[]string `json:"networkIds,omitempty"`
	SymbolIds *[]string `json:"symbolIds,omitempty"`
	ContractAddress *string `json:"contractAddress,omitempty"`
	AssetType *AssetType `json:"assetType,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	BaseSymbolIds *[]string `json:"baseSymbolIds,omitempty"`
}
