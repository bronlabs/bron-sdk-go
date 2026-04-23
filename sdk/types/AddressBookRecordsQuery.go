package types

type AddressBookRecordsQuery struct {
	RecordIDs *[]string `json:"recordIds,omitempty"`
	NetworkIDs *[]string `json:"networkIds,omitempty"`
	Addresses *[]string `json:"addresses,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	RecordType *RecordType `json:"recordType,omitempty"`
	RecordTypes *[]RecordType `json:"recordTypes,omitempty"`
	Statuses *[]RecordStatus `json:"statuses,omitempty"`
}
