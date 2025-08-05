package types

type Transaction struct {
	ExternalId string `json:"externalId"`
	AccountType AccountType `json:"accountType"`
	Extra *TransactionExtra `json:"extra"`
	TransactionId string `json:"transactionId"`
	WorkspaceId string `json:"workspaceId"`
	AccountId string `json:"accountId"`
	Embedded *TransactionEmbedded `json:"embedded"`
	TerminatedAt *string `json:"terminatedAt"`
	TransactionType TransactionType `json:"transactionType"`
	UpdatedAt *string `json:"updatedAt"`
	CreatedBy *string `json:"createdBy"`
	ExpiresAt *string `json:"expiresAt"`
	Params interface{} `json:"params"`
	Status TransactionStatus `json:"status"`
	CreatedAt string `json:"createdAt"`
}
