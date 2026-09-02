package types

type StakesEmbedded struct {
	ResourceRentalRewardsAmounts *[]RewardsAmount `json:"resourceRentalRewardsAmounts,omitempty"`
	RewardsAmounts *[]RewardsAmount `json:"rewardsAmounts,omitempty"`
}
