package types

type WithdrawalParams struct {
	Amount string `json:"amount"`
	AssetId *string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel"`
	IncludeFee *bool `json:"includeFee"`
	Memo *string `json:"memo"`
	NetworkFees *RequestedNetworkFees `json:"networkFees"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	ToAccountId *string `json:"toAccountId"`
	ToAddress *string `json:"toAddress"`
	ToAddressBookRecordId *string `json:"toAddressBookRecordId"`
}
