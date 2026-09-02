package types

type PriceSeriesPoint struct {
	BaseSymbolID string `json:"baseSymbolId"`
	Close string `json:"close"`
	High string `json:"high"`
	Low string `json:"low"`
	Open string `json:"open"`
	Timestamp string `json:"timestamp"`
	Volume24h string `json:"volume24h"`
}
