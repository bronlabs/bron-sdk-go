package types

type TransactionLimit struct {
	AppliesTo LimitAppliesTo `json:"appliesTo"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	Destinations LimitDestinations `json:"destinations"`
	ExternalId string `json:"externalId"`
	LimitId string `json:"limitId"`
	LimitRule LimitRule `json:"limitRule"`
	LimitType TransactionLimitType `json:"limitType"`
	Sources LimitSources `json:"sources"`
	Status TransactionLimitStatus `json:"status"`
	TransactionParams LimitTransactionParams `json:"transactionParams"`
	UpdatedAt *string `json:"updatedAt"`
	UpdatedBy *string `json:"updatedBy"`
	WorkspaceId string `json:"workspaceId"`
}
