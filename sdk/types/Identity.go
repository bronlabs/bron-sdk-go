package types

type Identity struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	IdentityId string `json:"identityId"`
	IdentityType IdentityType `json:"identityType"`
	IdentityValue string `json:"identityValue"`
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserId string `json:"userId"`
}
