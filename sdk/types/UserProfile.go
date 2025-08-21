package types

type UserProfile struct {
	ImageID *string `json:"imageId,omitempty"`
	Name *string `json:"name,omitempty"`
	UserID string `json:"userId"`
}
