package types

type SigningRequest struct {
	AccountId *string `json:"accountId"`
	BlockchainNonce *string `json:"blockchainNonce"`
	ShouldBeBroadcasted *bool `json:"shouldBeBroadcasted"`
	Signed *Signed `json:"signed"`
	SigningData *BlockchainSigningRequest `json:"signingData"`
	SigningRequestId *string `json:"signingRequestId"`
	Status *SigningRequestStatus `json:"status"`
	TransactionId *string `json:"transactionId"`
	MessagesForSigning *MessagesForSigning `json:"messagesForSigning"`
	NetworkId *string `json:"networkId"`
	RequestParameters *map[string]interface{} `json:"requestParameters"`
	TransactionType *TransactionType `json:"transactionType"`
	WorkspaceId *string `json:"workspaceId"`
}
