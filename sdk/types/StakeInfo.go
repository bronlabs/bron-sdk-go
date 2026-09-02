package types

type StakeInfo struct {
	BondingAmount *string `json:"bondingAmount,omitempty"`
	BondingEndAt *string `json:"bondingEndAt,omitempty"`
	ClaimableRewardsAmount *string `json:"claimableRewardsAmount,omitempty"`
	OriginPool *Pool `json:"originPool,omitempty"`
	ReadyToClaimStakeAt *string `json:"readyToClaimStakeAt,omitempty"`
	ReadyToIncrement *bool `json:"readyToIncrement,omitempty"`
	ReadyToRedelegate *bool `json:"readyToRedelegate,omitempty"`
	ReadyToUnstake *bool `json:"readyToUnstake,omitempty"`
	ReadyToUnstakeAt *string `json:"readyToUnstakeAt,omitempty"`
	RewardsRequireClaim *bool `json:"rewardsRequireClaim,omitempty"`
	UnbondingAmount *string `json:"unbondingAmount,omitempty"`
	UnbondingEndAt *string `json:"unbondingEndAt,omitempty"`
	Warning *Warning `json:"warning,omitempty"`
}
