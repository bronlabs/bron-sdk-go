package types

type StakesQuery struct {
	AccountID *string `json:"accountId,omitempty"`
	AssetID *string `json:"assetId,omitempty"`
	RewardPeriod *string `json:"rewardPeriod,omitempty"`
}
