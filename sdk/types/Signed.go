package types

type Signed struct {
	Signature *string `json:"signature,omitempty"`
	Signatures *[]Signature `json:"signatures,omitempty"`
	TxRaw *string `json:"txRaw,omitempty"`
}
