package types

type RawTransactionParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	Data *string `json:"data,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	RawTransactions *[]string `json:"rawTransactions,omitempty"`
	ToAddress string `json:"toAddress"`
}
