package types

type Stake struct {
	AccountID *string `json:"accountId,omitempty"`
	Amount *string `json:"amount,omitempty"`
	AssetID *string `json:"assetId,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	PoolID *string `json:"poolId,omitempty"`
	StakeID *string `json:"stakeId,omitempty"`
	StakeInfo *StakeInfo `json:"stakeInfo,omitempty"`
	Status *StakeStatus `json:"status,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Warning *Warning `json:"warning,omitempty"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}
