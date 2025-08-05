package types

type Asset struct {
	SymbolId *string `json:"symbolId"`
	Verified *bool `json:"verified"`
	AssetId string `json:"assetId"`
	AssetType *AssetType `json:"assetType"`
	ContractInformation *SmartContractInformation `json:"contractInformation"`
	Decimals *string `json:"decimals"`
	NetworkId *string `json:"networkId"`
}
