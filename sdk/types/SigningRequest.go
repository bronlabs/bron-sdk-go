package types

type SigningRequest struct {
	AccountId *string `json:"accountId,omitempty"`
	BlockchainNonce *string `json:"blockchainNonce,omitempty"`
	MessagesForSigning *MessagesForSigning `json:"messagesForSigning,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	RequestParameters *map[string]interface{} `json:"requestParameters,omitempty"`
	ShouldBeBroadcasted *bool `json:"shouldBeBroadcasted,omitempty"`
	Signed *Signed `json:"signed,omitempty"`
	SigningData *BlockchainSigningRequest `json:"signingData,omitempty"`
	SigningRequestId *string `json:"signingRequestId,omitempty"`
	Status *SigningRequestStatus `json:"status,omitempty"`
	TransactionId *string `json:"transactionId,omitempty"`
	TransactionType *TransactionType `json:"transactionType,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}
