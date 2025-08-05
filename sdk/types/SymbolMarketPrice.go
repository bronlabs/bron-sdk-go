package types

type SymbolMarketPrice struct {
	BaseSymbolId string `json:"baseSymbolId"`
	Price string `json:"price"`
	QuoteSymbolId string `json:"quoteSymbolId"`
}
