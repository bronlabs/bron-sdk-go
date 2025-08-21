package types

type Transaction struct {
	AccountID string `json:"accountId"`
	AccountType AccountType `json:"accountType"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	Embedded *TransactionEmbedded `json:"embedded,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	ExternalID string `json:"externalId"`
	Extra *TransactionExtra `json:"extra,omitempty"`
	Params interface{} `json:"params,omitempty"`
	Status TransactionStatus `json:"status"`
	TerminatedAt *string `json:"terminatedAt,omitempty"`
	TransactionID string `json:"transactionId"`
	TransactionType TransactionType `json:"transactionType"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	WorkspaceID string `json:"workspaceId"`
}
