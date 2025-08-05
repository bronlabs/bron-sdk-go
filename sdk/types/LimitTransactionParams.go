package types

type LimitTransactionParams struct {
	AboveAmount *LimitAmount `json:"aboveAmount"`
	DurationHours *string `json:"durationHours"`
}
