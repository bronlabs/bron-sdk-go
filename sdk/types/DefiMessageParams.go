package types

type DefiMessageParams struct {
	Message string `json:"message"`
	NetworkID string `json:"networkId"`
	Origin string `json:"origin"`
	Version *string `json:"version,omitempty"`
}
