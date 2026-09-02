package types

type EventStakeInfo struct {
	Amount *string `json:"amount,omitempty"`
	PoolID *string `json:"poolId,omitempty"`
	Resource *string `json:"resource,omitempty"`
	StakeID *string `json:"stakeId,omitempty"`
}
