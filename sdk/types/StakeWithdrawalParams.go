package types

type StakeWithdrawalParams struct {
	Amount *string `json:"amount,omitempty"`
	AssetID string `json:"assetId"`
	PoolID *string `json:"poolId,omitempty"`
}
