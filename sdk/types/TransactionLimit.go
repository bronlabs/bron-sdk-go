package types

type TransactionLimit struct {
	UpdatedAt *string `json:"updatedAt"`
	WorkspaceId string `json:"workspaceId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	Destinations LimitDestinations `json:"destinations"`
	LimitRule LimitRule `json:"limitRule"`
	Sources LimitSources `json:"sources"`
	Status TransactionLimitStatus `json:"status"`
	UpdatedBy *string `json:"updatedBy"`
	AppliesTo LimitAppliesTo `json:"appliesTo"`
	ExternalId string `json:"externalId"`
	LimitId string `json:"limitId"`
	LimitType TransactionLimitType `json:"limitType"`
	TransactionParams LimitTransactionParams `json:"transactionParams"`
}
