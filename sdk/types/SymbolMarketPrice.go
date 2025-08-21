package types

type SymbolMarketPrice struct {
	BaseSymbolID string `json:"baseSymbolId"`
	Price string `json:"price"`
	QuoteSymbolID string `json:"quoteSymbolId"`
}
