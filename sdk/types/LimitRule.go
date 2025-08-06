package types

type LimitRule struct {
	Approve *LimitRuleApprove `json:"approve,omitempty"`
	SecurityDelay *LimitRuleSecurityDelay `json:"securityDelay,omitempty"`
	SkipApproval *bool `json:"skipApproval,omitempty"`
}
