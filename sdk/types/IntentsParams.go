package types

type IntentsParams struct {
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	IntentID string `json:"intentId"`
}
