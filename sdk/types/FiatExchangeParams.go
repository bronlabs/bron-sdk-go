package types

type FiatExchangeParams struct {
	Amount string `json:"amount"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	FromAccountID *string `json:"fromAccountId,omitempty"`
	FromAssetID string `json:"fromAssetId"`
	FromBankAccountID *string `json:"fromBankAccountId,omitempty"`
	FromNetworkID string `json:"fromNetworkId"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAssetID string `json:"toAssetId"`
	ToBankAccountID *string `json:"toBankAccountId,omitempty"`
	ToNetworkID string `json:"toNetworkId"`
}
