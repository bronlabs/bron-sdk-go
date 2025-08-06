package types

type Asset struct {
	AssetId string `json:"assetId"`
	AssetType *AssetType `json:"assetType,omitempty"`
	ContractInformation *SmartContractInformation `json:"contractInformation,omitempty"`
	Decimals *string `json:"decimals,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	SymbolId *string `json:"symbolId,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}
