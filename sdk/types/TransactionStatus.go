package types

type TransactionStatus string

const (
	TransactionStatus_NEW TransactionStatus = "new"
	TransactionStatus_WAITING_CONFIRMATIONS TransactionStatus = "waiting-confirmations"
	TransactionStatus_WAITING_APPROVAL TransactionStatus = "waiting-approval"
	TransactionStatus_APPROVED TransactionStatus = "approved"
	TransactionStatus_AWAITING_SECURITY_POLICY TransactionStatus = "awaiting-security-policy"
	TransactionStatus_COMPLETED TransactionStatus = "completed"
	TransactionStatus_CANCELED TransactionStatus = "canceled"
	TransactionStatus_EXPIRED TransactionStatus = "expired"
	TransactionStatus_SIGNING_REQUIRED TransactionStatus = "signing-required"
	TransactionStatus_SIGNING TransactionStatus = "signing"
	TransactionStatus_SIGNED TransactionStatus = "signed"
	TransactionStatus_BROADCASTED TransactionStatus = "broadcasted"
	TransactionStatus_MANUAL_RESOLVING TransactionStatus = "manual-resolving"
	TransactionStatus_FAILED_ON_BLOCKCHAIN TransactionStatus = "failed-on-blockchain"
	TransactionStatus_REMOVED_FROM_BLOCKCHAIN TransactionStatus = "removed-from-blockchain"
	TransactionStatus_ERROR TransactionStatus = "error"
	TransactionStatus_AWAITING_DEPOSIT TransactionStatus = "awaiting-deposit"
)
