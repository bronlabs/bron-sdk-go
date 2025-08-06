package types

type StakesQuery struct {
	AccountId *string `json:"accountId,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	RewardPeriod *string `json:"rewardPeriod,omitempty"`
}
