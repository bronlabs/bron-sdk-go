package types

type AddressBookRecordsQuery struct {
	RecordIds *[]string `json:"recordIds,omitempty"`
	NetworkIds *[]string `json:"networkIds,omitempty"`
	Addresses *[]string `json:"addresses,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	Statuses *[]RecordStatus `json:"statuses,omitempty"`
}
