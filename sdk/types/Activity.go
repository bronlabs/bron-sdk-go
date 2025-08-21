package types

type Activity struct {
	AccountID *string `json:"accountId,omitempty"`
	ActivityID string `json:"activityId"`
	ActivityType ActivityType `json:"activityType"`
	CreatedAt string `json:"createdAt"`
	Description *string `json:"description,omitempty"`
	Title string `json:"title"`
	UserID *string `json:"userId,omitempty"`
	WorkspaceID *string `json:"workspaceId,omitempty"`
}
