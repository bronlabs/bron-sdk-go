package types

type SmartContractInformation struct {
	ContractAddress *string `json:"contractAddress,omitempty"`
	Standard *TokenStandard `json:"standard,omitempty"`
}
