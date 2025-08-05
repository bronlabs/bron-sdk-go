package types

type WithdrawalParams struct {
	ToAddressBookRecordId *string `json:"toAddressBookRecordId"`
	Amount string `json:"amount"`
	AssetId *string `json:"assetId"`
	ToAddress *string `json:"toAddress"`
	FeeLevel *FeeLevel `json:"feeLevel"`
	IncludeFee *bool `json:"includeFee"`
	Memo *string `json:"memo"`
	NetworkFees *RequestedNetworkFees `json:"networkFees"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	ToAccountId *string `json:"toAccountId"`
}
