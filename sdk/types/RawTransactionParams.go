package types

type RawTransactionParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	Data *string `json:"data,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	RawTransactions *[]string `json:"rawTransactions,omitempty"`
	SkipSimulation *bool `json:"skipSimulation,omitempty"`
	ToAddress string `json:"toAddress"`
}
