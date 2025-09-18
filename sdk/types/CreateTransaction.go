package types

type CreateTransaction struct {
	AccountID string `json:"accountId"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	ExternalID string `json:"externalId"`
	Params interface{} `json:"params,omitempty"`
	TransactionType TransactionType `json:"transactionType"`
}
