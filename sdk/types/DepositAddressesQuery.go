package types

type DepositAddressesQuery struct {
	AccountId *string `json:"accountId,omitempty"`
	AddressIds *[]string `json:"addressIds,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	AccountTypes *[]AccountType `json:"accountTypes,omitempty"`
	NetworkId *string `json:"networkId,omitempty"`
	Address *string `json:"address,omitempty"`
	Statuses *[]AddressStatus `json:"statuses,omitempty"`
	SortDirection *SortingDirection `json:"sortDirection,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
