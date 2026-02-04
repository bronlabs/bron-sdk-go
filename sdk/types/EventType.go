package types

type EventType string

const (
	EventType_IN EventType = "in"
	EventType_OUT EventType = "out"
	EventType_FEE EventType = "fee"
	EventType_NEGATIVE_DEPOSIT EventType = "negative-deposit"
	EventType_STAKE_DELEGATION EventType = "stake-delegation"
	EventType_STAKE_UNDELEGATION EventType = "stake-undelegation"
	EventType_STAKE_WITHDRAWN EventType = "stake-withdrawn"
	EventType_STAKE_EARN_REWARD EventType = "stake-earn-reward"
	EventType_STAKE_REWARD_ACCRUED EventType = "stake-reward-accrued"
	EventType_STAKE_POOL_CREATED EventType = "stake-pool-created"
	EventType_STAKE_POOL_UNJAILED EventType = "stake-pool-unjailed"
	EventType_STAKE_TRANSFERRED EventType = "stake-transferred"
	EventType_ALLOWANCE EventType = "allowance"
	EventType_NFT_IN EventType = "nft-in"
	EventType_NFT_OUT EventType = "nft-out"
	EventType_NFT_ALLOWANCE EventType = "nft-allowance"
	EventType_LOYALTY_LOCK EventType = "loyalty-lock"
	EventType_LOYALTY_UNLOCK EventType = "loyalty-unlock"
	EventType_LOYALTY_REWARD EventType = "loyalty-reward"
	EventType_CANTON_REWARD EventType = "canton-reward"
	EventType_MESSAGE_SIGNED EventType = "message-signed"
	EventType_DEPOSIT_OFFER EventType = "deposit-offer"
)
