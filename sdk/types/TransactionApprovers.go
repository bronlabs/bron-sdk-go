package types

type TransactionApprovers struct {
	ApprovedBy *[]string `json:"approvedBy"`
	AvailableApprovers *[]string `json:"availableApprovers"`
	LimitId *string `json:"limitId"`
	Number *string `json:"number"`
	SecurityDelayDuration *string `json:"securityDelayDuration"`
	SecurityDelayExpiresAt *string `json:"securityDelayExpiresAt"`
	SkipApproval *bool `json:"skipApproval"`
}
