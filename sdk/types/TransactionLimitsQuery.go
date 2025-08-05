package types

type TransactionLimitsQuery struct {
	Statuses *[]TransactionLimitStatus `json:"statuses,omitempty"`
	FromAccountIds *[]string `json:"fromAccountIds,omitempty"`
	ToAddressBookRecordIds *[]string `json:"toAddressBookRecordIds,omitempty"`
	ToAccountIds *[]string `json:"toAccountIds,omitempty"`
	AppliesToUserIds *[]string `json:"appliesToUserIds,omitempty"`
	AppliesToGroupIds *[]string `json:"appliesToGroupIds,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
