package types

type Asset struct {
	AssetID string `json:"assetId"`
	AssetType *AssetType `json:"assetType,omitempty"`
	ContractInformation *SmartContractInformation `json:"contractInformation,omitempty"`
	Decimals *string `json:"decimals,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	SymbolID *string `json:"symbolId,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}
