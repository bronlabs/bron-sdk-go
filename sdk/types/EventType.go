package types

type EventType string

const (
	EventType_IN EventType = "in"
	EventType_OUT EventType = "out"
	EventType_FEE EventType = "fee"
	EventType_NEGATIVE_DEPOSIT EventType = "negative-deposit"
	EventType_STAKE_DELEGATION EventType = "stake-delegation"
	EventType_STAKE_UNDELEGATION EventType = "stake-undelegation"
	EventType_STAKE_CLAIM EventType = "stake-claim"
	EventType_STAKE_EARN_REWARD EventType = "stake-earn-reward"
	EventType_STAKE_REWARD_ACCRUED EventType = "stake-reward-accrued"
	EventType_ALLOWANCE EventType = "allowance"
	EventType_NFT_IN EventType = "nft-in"
	EventType_NFT_OUT EventType = "nft-out"
)
