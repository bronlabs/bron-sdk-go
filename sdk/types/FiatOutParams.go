package types

type FiatOutParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	FiatAssetID string `json:"fiatAssetId"`
	ToAddressBookRecordID string `json:"toAddressBookRecordId"`
}
