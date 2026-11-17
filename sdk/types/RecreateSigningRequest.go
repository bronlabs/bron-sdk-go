package types

type RecreateSigningRequest struct {
	CancelTransaction *bool `json:"cancelTransaction,omitempty"`
	FeeLevel *FeeLevel `json:"feeLevel,omitempty"`
	NetworkFees *RequestedNetworkFees `json:"networkFees,omitempty"`
	PrimitivesVersion *string `json:"primitivesVersion,omitempty"`
	SigningRequestID string `json:"signingRequestId"`
}
