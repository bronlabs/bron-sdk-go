package types

type Pool struct {
	PoolID *string `json:"poolId,omitempty"`
	PoolName *string `json:"poolName,omitempty"`
	TotalBonded *string `json:"totalBonded,omitempty"`
}
