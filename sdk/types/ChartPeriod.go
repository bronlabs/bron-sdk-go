package types

type ChartPeriod string

const (
	ChartPeriod_TYPE_1H ChartPeriod = "1h"
	ChartPeriod_TYPE_1D ChartPeriod = "1d"
	ChartPeriod_TYPE_1W ChartPeriod = "1w"
	ChartPeriod_TYPE_1M ChartPeriod = "1m"
	ChartPeriod_TYPE_6M ChartPeriod = "6m"
	ChartPeriod_TYPE_1Y ChartPeriod = "1y"
	ChartPeriod_YTD ChartPeriod = "ytd"
	ChartPeriod_ALL ChartPeriod = "all"
)
