package types

type BridgeParams struct {
	Amount string `json:"amount"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	SourceAssetID string `json:"sourceAssetId"`
}
