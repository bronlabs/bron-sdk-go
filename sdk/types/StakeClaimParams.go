package types

type StakeClaimParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	StakeID *string `json:"stakeId,omitempty"`
}
