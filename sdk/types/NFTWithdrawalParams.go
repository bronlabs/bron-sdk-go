package types

type NFTWithdrawalParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	IncludeFee *bool `json:"includeFee,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	ToAddressBookRecordID *string `json:"toAddressBookRecordId,omitempty"`
	TokenID string `json:"tokenId"`
}
