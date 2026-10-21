package types

type SlippageReason string

const (
	SlippageReason_STOCK_OFF_SESSION SlippageReason = "stock-off-session"
	SlippageReason_STOCK_SESSION_UNKNOWN SlippageReason = "stock-session-unknown"
)
