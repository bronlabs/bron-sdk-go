package types

type Asset struct {
	AssetId string `json:"assetId"`
	AssetType *AssetType `json:"assetType"`
	ContractInformation *SmartContractInformation `json:"contractInformation"`
	Decimals *string `json:"decimals"`
	NetworkId *string `json:"networkId"`
	SymbolId *string `json:"symbolId"`
	Verified *bool `json:"verified"`
}
