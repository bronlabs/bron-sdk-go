package types

type TransactionLimitsQuery struct {
	Statuses *[]TransactionLimitStatus `json:"statuses,omitempty"`
	FromAccountIDs *[]string `json:"fromAccountIds,omitempty"`
	ToAddressBookRecordIDs *[]string `json:"toAddressBookRecordIds,omitempty"`
	ToAccountIDs *[]string `json:"toAccountIds,omitempty"`
	AppliesToUserIDs *[]string `json:"appliesToUserIds,omitempty"`
	AppliesToGroupIDs *[]string `json:"appliesToGroupIds,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
