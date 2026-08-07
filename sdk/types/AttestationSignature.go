package types

type AttestationSignature struct {
	PublicKey *string `json:"publicKey,omitempty"`
	Signature *string `json:"signature,omitempty"`
}
