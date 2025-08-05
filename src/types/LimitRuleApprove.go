package types

type LimitRuleApprove struct {
	AuthorisedApproversUserIds *[]string `json:"authorisedApproversUserIds"`
	NumberOfApprovals string `json:"numberOfApprovals"`
}
