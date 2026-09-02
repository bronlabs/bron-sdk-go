package types

type BlockchainSigningRequest struct {
	AssetID *string `json:"assetId,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	Sponsored *bool `json:"sponsored,omitempty"`
	TransactionType *TransactionType `json:"transactionType,omitempty"`
}
