package types

type SigningRequest struct {
	AccountID *string `json:"accountId,omitempty"`
	BlockchainNonce *string `json:"blockchainNonce,omitempty"`
	MessagesForSigning *MessagesForSigning `json:"messagesForSigning,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	RequestParameters *map[string]interface{} `json:"requestParameters,omitempty"`
	ShouldBeBroadcasted *bool `json:"shouldBeBroadcasted,omitempty"`
	Signed *Signed `json:"signed,omitempty"`
	SigningData *BlockchainSigningRequest `json:"signingData,omitempty"`
	SigningRequestID *string `json:"signingRequestId,omitempty"`
	Status *SigningRequestStatus `json:"status,omitempty"`
	TransactionID *string `json:"transactionId,omitempty"`
	TransactionType *TransactionType `json:"transactionType,omitempty"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}
