package types

type DepositAddressesQuery struct {
	AccountID *string `json:"accountId,omitempty"`
	AddressIDs *[]string `json:"addressIds,omitempty"`
	ExternalID *string `json:"externalId,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	NetworkID *string `json:"networkId,omitempty"`
	Address *string `json:"address,omitempty"`
	Statuses *[]AddressStatus `json:"statuses,omitempty"`
	SortDirection *SortingDirection `json:"sortDirection,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
