package types

type StakeStatus string

const (
	StakeStatus_BONDING StakeStatus = "bonding"
	StakeStatus_ACTIVE StakeStatus = "active"
	StakeStatus_UNBONDING StakeStatus = "unbonding"
	StakeStatus_REQUIRE_CLAIM StakeStatus = "require-claim"
	StakeStatus_EXITED StakeStatus = "exited"
)
