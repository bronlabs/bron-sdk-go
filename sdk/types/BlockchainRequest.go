package types

type BlockchainRequest struct {
	ExternalBroadcast *bool `json:"externalBroadcast,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
}
