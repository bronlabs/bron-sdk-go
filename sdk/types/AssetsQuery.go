package types

type AssetsQuery struct {
	AssetIDs *[]string `json:"assetIds,omitempty"`
	NetworkIDs *[]string `json:"networkIds,omitempty"`
	SymbolIDs *[]string `json:"symbolIds,omitempty"`
	ContractAddress *string `json:"contractAddress,omitempty"`
	AssetType *AssetType `json:"assetType,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
