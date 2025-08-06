package types

type User struct {
	AllowedIps *[]string `json:"allowedIps,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	LastSignInAt *string `json:"lastSignInAt,omitempty"`
	UserId string `json:"userId"`
}
