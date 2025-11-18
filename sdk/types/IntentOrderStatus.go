package types

type IntentOrderStatus string

const (
	IntentOrderStatus_NOT_EXIST IntentOrderStatus = "not-exist"
	IntentOrderStatus_USER_INITIATED IntentOrderStatus = "user-initiated"
	IntentOrderStatus_AUCTION_IN_PROGRESS IntentOrderStatus = "auction-in-progress"
	IntentOrderStatus_WAIT_FOR_USER_TX IntentOrderStatus = "wait-for-user-tx"
	IntentOrderStatus_WAIT_FOR_ORACLE_CONFIRM_USER_TX IntentOrderStatus = "wait-for-oracle-confirm-user-tx"
	IntentOrderStatus_WAIT_FOR_SOLVER_TX IntentOrderStatus = "wait-for-solver-tx"
	IntentOrderStatus_WAIT_FOR_ORACLE_CONFIRM_SOLVER_TX IntentOrderStatus = "wait-for-oracle-confirm-solver-tx"
	IntentOrderStatus_COMPLETED IntentOrderStatus = "completed"
	IntentOrderStatus_LIQUIDATED IntentOrderStatus = "liquidated"
	IntentOrderStatus_CANCELLED IntentOrderStatus = "cancelled"
)
