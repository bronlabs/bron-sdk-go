package types

type Asset struct {
	AssetType *AssetType `json:"assetType"`
	ContractInformation *SmartContractInformation `json:"contractInformation"`
	Decimals *string `json:"decimals"`
	NetworkId *string `json:"networkId"`
	SymbolId *string `json:"symbolId"`
	Verified *bool `json:"verified"`
	AssetId string `json:"assetId"`
}
