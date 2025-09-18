package types

type TransactionEmbedded struct {
	CurrentSigningRequest *SigningRequest `json:"currentSigningRequest,omitempty"`
	Events *[]TransactionEvent `json:"events,omitempty"`
}
