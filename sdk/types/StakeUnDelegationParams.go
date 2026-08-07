package types

type StakeUnDelegationParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	Resource *StakeResource `json:"resource,omitempty"`
	StakeID *string `json:"stakeId,omitempty"`
}
