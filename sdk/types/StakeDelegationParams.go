package types

type StakeDelegationParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	PoolID *string `json:"poolId,omitempty"`
}
