package types

type TransactionEmbedded struct {
	AttestationSignature *AttestationSignature `json:"attestationSignature,omitempty"`
	CurrentSigningRequest *SigningRequest `json:"currentSigningRequest,omitempty"`
	Events *[]TransactionEvent `json:"events,omitempty"`
	IntentsQuote *IntentsQuoteEmbedded `json:"intentsQuote,omitempty"`
	QuoteDetails *QuoteDetails `json:"quoteDetails,omitempty"`
}
