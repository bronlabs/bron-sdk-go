package types

type WithdrawalParams struct {
	Memo *string `json:"memo"`
	NetworkFees *RequestedNetworkFees `json:"networkFees"`
	ToAccountId *string `json:"toAccountId"`
	ToAddressBookRecordId *string `json:"toAddressBookRecordId"`
	NetworkId *string `json:"networkId"`
	Symbol *string `json:"symbol"`
	ToAddress *string `json:"toAddress"`
	Amount string `json:"amount"`
	AssetId *string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel"`
	IncludeFee *bool `json:"includeFee"`
}
