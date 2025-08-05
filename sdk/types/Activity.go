package types

type Activity struct {
	AccountId *string `json:"accountId"`
	ActivityId string `json:"activityId"`
	ActivityType ActivityType `json:"activityType"`
	CreatedAt string `json:"createdAt"`
	Description *string `json:"description"`
	Title string `json:"title"`
	UserId *string `json:"userId"`
	WorkspaceId *string `json:"workspaceId"`
}
