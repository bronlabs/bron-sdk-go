package types

type StakeRewardInfo struct {
	IncreaseOperableBalance *bool `json:"increaseOperableBalance,omitempty"`
	PoolIDs *[]string `json:"poolIds,omitempty"`
	RewardSource *RewardSource `json:"rewardSource,omitempty"`
	RewardWithoutTransaction *bool `json:"rewardWithoutTransaction,omitempty"`
	StakeID *string `json:"stakeId,omitempty"`
}
