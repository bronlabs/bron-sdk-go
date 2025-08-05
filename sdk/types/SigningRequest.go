package types

type SigningRequest struct {
	SigningData *BlockchainSigningRequest `json:"signingData"`
	SigningRequestId *string `json:"signingRequestId"`
	ShouldBeBroadcasted *bool `json:"shouldBeBroadcasted"`
	Status *SigningRequestStatus `json:"status"`
	TransactionId *string `json:"transactionId"`
	TransactionType *TransactionType `json:"transactionType"`
	WorkspaceId *string `json:"workspaceId"`
	AccountId *string `json:"accountId"`
	BlockchainNonce *string `json:"blockchainNonce"`
	MessagesForSigning *MessagesForSigning `json:"messagesForSigning"`
	NetworkId *string `json:"networkId"`
	RequestParameters *map[string]interface{} `json:"requestParameters"`
	Signed *Signed `json:"signed"`
}
