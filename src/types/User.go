package types

type User struct {
	AllowedIps *[]string `json:"allowedIps"`
	CreatedAt *string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	LastSignInAt *string `json:"lastSignInAt"`
	UserId string `json:"userId"`
}
