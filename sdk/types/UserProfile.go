package types

type UserProfile struct {
	Icon *string `json:"icon,omitempty"`
	Name *string `json:"name,omitempty"`
	UserID string `json:"userId"`
}
