package types

type TransactionApprovers struct {
	ApprovedBy *[]string `json:"approvedBy,omitempty"`
	AvailableApprovers *[]string `json:"availableApprovers,omitempty"`
	LimitID *string `json:"limitId,omitempty"`
	Number *string `json:"number,omitempty"`
	SecurityDelayDuration *string `json:"securityDelayDuration,omitempty"`
	SecurityDelayExpiresAt *string `json:"securityDelayExpiresAt,omitempty"`
	SkipApproval *bool `json:"skipApproval,omitempty"`
}
