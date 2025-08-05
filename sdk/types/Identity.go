package types

type Identity struct {
	UpdatedAt *string `json:"updatedAt"`
	UserId string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	IdentityId string `json:"identityId"`
	IdentityType IdentityType `json:"identityType"`
	IdentityValue string `json:"identityValue"`
	LastUsedAt *string `json:"lastUsedAt"`
}
