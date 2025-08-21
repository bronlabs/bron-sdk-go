package types

type BlockchainTxDetails struct {
	BlockchainTxID *string `json:"blockchainTxId,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
}
