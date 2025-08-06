package types

type UserProfile struct {
	ImageId *string `json:"imageId,omitempty"`
	Name *string `json:"name,omitempty"`
	UserId string `json:"userId"`
}
