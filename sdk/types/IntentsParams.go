package types

type IntentsParams struct {
	FeeAssetID *string `json:"feeAssetId,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	IntentID string `json:"intentId"`
}
