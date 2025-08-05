package types

type SymbolMarketPrice struct {
	Price string `json:"price"`
	QuoteSymbolId string `json:"quoteSymbolId"`
	BaseSymbolId string `json:"baseSymbolId"`
}
