package types

type LimitTransactionParams struct {
	AboveAmount *LimitAmount `json:"aboveAmount,omitempty"`
	DurationHours *string `json:"durationHours,omitempty"`
}
