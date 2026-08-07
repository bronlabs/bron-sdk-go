package types

type TransactionEmbedded struct {
	AttestationSignature *AttestationSignature `json:"attestationSignature,omitempty"`
	CurrentSigningRequest *SigningRequest `json:"currentSigningRequest,omitempty"`
	Events *[]TransactionEvent `json:"events,omitempty"`
}
