package types

type Warning struct {
	Code *string `json:"code,omitempty"`
	Message string `json:"message"`
}
