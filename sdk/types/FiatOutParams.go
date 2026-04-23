package types

type FiatOutParams struct {
	Amount string `json:"amount"`
	AssetID string `json:"assetId"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	FiatAssetID string `json:"fiatAssetId"`
	NetworkID string `json:"networkId"`
	ToAddressBookRecordID string `json:"toAddressBookRecordId"`
}
