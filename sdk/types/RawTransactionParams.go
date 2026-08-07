package types

type RawTransactionParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID *string `json:"assetId,omitempty"`
	Data *string `json:"data,omitempty"`
	ExternalBroadcast *bool `json:"externalBroadcast,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	RawTransactions *[]string `json:"rawTransactions,omitempty"`
	SkipSimulation *bool `json:"skipSimulation,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
}
