package types

type Signature struct {
	E *string `json:"e,omitempty"`
	R *string `json:"r,omitempty"`
	Rx *string `json:"rx,omitempty"`
	S *string `json:"s,omitempty"`
	Signature *string `json:"signature,omitempty"`
	V *string `json:"v,omitempty"`
}
