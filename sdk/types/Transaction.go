package types

type Transaction struct {
	AccountId string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	Embedded *TransactionEmbedded `json:"embedded"`
	ExpiresAt *string `json:"expiresAt"`
	ExternalId string `json:"externalId"`
	Extra *TransactionExtra `json:"extra"`
	Params interface{} `json:"params"`
	Status TransactionStatus `json:"status"`
	TerminatedAt *string `json:"terminatedAt"`
	TransactionId string `json:"transactionId"`
	TransactionType TransactionType `json:"transactionType"`
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
}
