package types

type StockSession string

const (
	StockSession_REGULAR StockSession = "regular"
	StockSession_PREMARKET StockSession = "premarket"
	StockSession_POSTMARKET StockSession = "postmarket"
	StockSession_OVERNIGHT StockSession = "overnight"
	StockSession_OFFHOURS StockSession = "offhours"
	StockSession_PAUSED StockSession = "paused"
	StockSession_CLOSED StockSession = "closed"
)
