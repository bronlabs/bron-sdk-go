package types

type BronLockBurnExtra struct {
	RedeemRequestID *string `json:"redeemRequestId,omitempty"`
	ReleaseBlockchainTxID *string `json:"releaseBlockchainTxId,omitempty"`
	ReleaseSigningRequestID *string `json:"releaseSigningRequestId,omitempty"`
	ReleaseTransactionID *string `json:"releaseTransactionId,omitempty"`
}
