package types

type TransactionExtra struct {
	BlockchainDetails *[]BlockchainTxDetails `json:"blockchainDetails"`
	BlockchainRequest *BlockchainRequest `json:"blockchainRequest"`
	DepositTransactionId *string `json:"depositTransactionId"`
	Description *string `json:"description"`
	FromAccountId *string `json:"fromAccountId"`
	Memo *string `json:"memo"`
	ToAccountId *string `json:"toAccountId"`
	Confirmations *string `json:"confirmations"`
	ExternalBroadcast *bool `json:"externalBroadcast"`
	FromAddress *string `json:"fromAddress"`
	SigningRequestId *string `json:"signingRequestId"`
	ToAddress *string `json:"toAddress"`
	WithdrawTransactionId *string `json:"withdrawTransactionId"`
	Approvers *TransactionApprovers `json:"approvers"`
}
