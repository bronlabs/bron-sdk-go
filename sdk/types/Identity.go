package types

type Identity struct {
	UserId string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	IdentityId string `json:"identityId"`
	IdentityType IdentityType `json:"identityType"`
	IdentityValue string `json:"identityValue"`
	LastUsedAt *string `json:"lastUsedAt"`
	UpdatedAt *string `json:"updatedAt"`
}
