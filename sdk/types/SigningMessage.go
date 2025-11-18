package types

type SigningMessage struct {
	Message string `json:"message"`
	Version *string `json:"version,omitempty"`
}
