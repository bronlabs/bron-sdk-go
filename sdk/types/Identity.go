package types

type Identity struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy *string `json:"createdBy,omitempty"`
	IdentityID string `json:"identityId"`
	IdentityType IdentityType `json:"identityType"`
	IdentityValue string `json:"identityValue"`
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserID string `json:"userId"`
}
