package types

type AssetMarketPrice struct {
	BaseAssetID string `json:"baseAssetId"`
	BaseSymbolID string `json:"baseSymbolId"`
	Price string `json:"price"`
	QuoteAssetID string `json:"quoteAssetId"`
	QuoteSymbolID string `json:"quoteSymbolId"`
}
