package types

type AssetPriceSeriesQuery struct {
	BaseSymbolID string `json:"baseSymbolId"`
	Period ChartPeriod `json:"period"`
}
