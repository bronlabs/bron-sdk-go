package types

type CancelTransaction struct {
	Reason *string `json:"reason,omitempty"`
}
