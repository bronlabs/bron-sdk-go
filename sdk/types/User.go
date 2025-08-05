package types

type User struct {
	CreatedAt *string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	LastSignInAt *string `json:"lastSignInAt"`
	UserId string `json:"userId"`
	AllowedIps *[]string `json:"allowedIps"`
}
