package types

type TransactionExtra struct {
	Approvers *TransactionApprovers `json:"approvers,omitempty"`
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails,omitempty"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest,omitempty"`
	Confirmations *string `json:"confirmations,omitempty"`
	DepositTransactionID *string `json:"depositTransactionId,omitempty"`
	Description *string `json:"description,omitempty"`
	FromAccountID *string `json:"fromAccountId,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	FromWorkspaceImageID *string `json:"fromWorkspaceImageId,omitempty"`
	FromWorkspaceName *string `json:"fromWorkspaceName,omitempty"`
	FromWorkspaceTag *string `json:"fromWorkspaceTag,omitempty"`
	Memo *string `json:"memo,omitempty"`
	SigningRequestID *string `json:"signingRequestId,omitempty"`
	ToAccountID *string `json:"toAccountId,omitempty"`
	ToAddress *string `json:"toAddress,omitempty"`
	ToWorkspaceImageID *string `json:"toWorkspaceImageId,omitempty"`
	ToWorkspaceName *string `json:"toWorkspaceName,omitempty"`
	ToWorkspaceTag *string `json:"toWorkspaceTag,omitempty"`
	WithdrawTransactionID *string `json:"withdrawTransactionId,omitempty"`
}
