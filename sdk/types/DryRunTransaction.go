package types

type DryRunTransaction struct {
	AccountID string `json:"accountId"`
	Estimations *[]TransactionEstimation `json:"estimations,omitempty"`
	ExternalID *string `json:"externalId,omitempty"`
	Extra *TransactionExtra `json:"extra,omitempty"`
	Params *map[string]interface{} `json:"params,omitempty"`
	TransactionType TransactionType `json:"transactionType"`
	Warning *Warning `json:"warning,omitempty"`
}
