package types

type IntentsQuote struct {
	FromAmount string `json:"fromAmount"`
	FromAssetID string `json:"fromAssetId"`
	MinPrice string `json:"minPrice"`
	MinToAmount string `json:"minToAmount"`
	OracleFeePercent string `json:"oracleFeePercent"`
	SolverFeePercent string `json:"solverFeePercent"`
	ToAmount string `json:"toAmount"`
	ToAssetID string `json:"toAssetId"`
}
