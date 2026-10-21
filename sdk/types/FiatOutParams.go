package types

type FiatOutParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	FiatAmount *string `json:"fiatAmount,omitempty"`
	FiatAssetID string `json:"fiatAssetId"`
	PaymentPurpose *string `json:"paymentPurpose,omitempty"`
	ProcessingTier *string `json:"processingTier,omitempty"`
	Reference *string `json:"reference,omitempty"`
	ToAddressBookRecordID string `json:"toAddressBookRecordId"`
}
