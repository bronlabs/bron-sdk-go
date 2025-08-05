package types

type SigningRequest struct {
	NetworkId *string `json:"networkId"`
	SigningRequestId *string `json:"signingRequestId"`
	TransactionId *string `json:"transactionId"`
	TransactionType *TransactionType `json:"transactionType"`
	WorkspaceId *string `json:"workspaceId"`
	AccountId *string `json:"accountId"`
	RequestParameters *map[string]interface{} `json:"requestParameters"`
	ShouldBeBroadcasted *bool `json:"shouldBeBroadcasted"`
	Signed *Signed `json:"signed"`
	SigningData *BlockchainSigningRequest `json:"signingData"`
	Status *SigningRequestStatus `json:"status"`
	BlockchainNonce *string `json:"blockchainNonce"`
	MessagesForSigning *MessagesForSigning `json:"messagesForSigning"`
}
