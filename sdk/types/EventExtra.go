package types

type EventExtra struct {
	Allowance *[]EventAllowance `json:"allowance,omitempty"`
	In *[]EventInput `json:"in,omitempty"`
	Out *[]EventOutput `json:"out,omitempty"`
	RewardInfo *StakeRewardInfo `json:"rewardInfo,omitempty"`
	SigningMessage *SigningMessage `json:"signingMessage,omitempty"`
	StakeInfo *[]EventStakeInfo `json:"stakeInfo,omitempty"`
	TransactionFailed *bool `json:"transactionFailed,omitempty"`
}
