package types

type WithdrawalParams struct {
	Amount string `json:"amount"`
	AssetId *string `json:"assetId,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	IncludeFee *bool `json:"includeFee,omitempty"`
	Memo *string `json:"memo,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	ToAccountId *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	ToAddressBookRecordId *string `json:"toAddressBookRecordId,omitempty"`
}
