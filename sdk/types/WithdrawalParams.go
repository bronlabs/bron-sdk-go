package types

type WithdrawalParams struct {
	Amount string `json:"amount"`
	AssetID *string `json:"assetId,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	IncludeFee *bool `json:"includeFee,omitempty"`
	Memo *string `json:"memo,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	ToAddressBookRecordID *string `json:"toAddressBookRecordId,omitempty"`
}
