package types

type WithdrawalParams struct {
	AssetId *string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	ToAddress *string `json:"toAddress"`
	Amount string `json:"amount"`
	IncludeFee *bool `json:"includeFee"`
	Memo *string `json:"memo"`
	NetworkFees *RequestedNetworkFees `json:"networkFees"`
	ToAccountId *string `json:"toAccountId"`
	ToAddressBookRecordId *string `json:"toAddressBookRecordId"`
}
