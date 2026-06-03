package types

type TransactionLimitsQuery struct {
	LimitIDs *[]string `json:"limitIds,omitempty"`
	Statuses *[]TransactionLimitStatus `json:"statuses,omitempty"`
	FromAccountIDs *[]string `json:"fromAccountIds,omitempty"`
	ToAddressBookRecordIDs *[]string `json:"toAddressBookRecordIds,omitempty"`
	ToAccountIDs *[]string `json:"toAccountIds,omitempty"`
	AppliesToUserIDs *[]string `json:"appliesToUserIds,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
