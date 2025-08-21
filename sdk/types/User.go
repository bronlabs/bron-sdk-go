package types

type User struct {
	AllowedIPs *[]string `json:"allowedIps,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	LastSignInAt *string `json:"lastSignInAt,omitempty"`
	UserID string `json:"userId"`
}
