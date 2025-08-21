package types

type ActivitiesQuery struct {
	AccountIDs *[]string `json:"accountIds,omitempty"`
	Offset *string `json:"offset,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Search *string `json:"search,omitempty"`
	UserIDs *[]string `json:"userIds,omitempty"`
	ActivityTypes *[]ActivityType `json:"activityTypes,omitempty"`
	ExcludedActivityTypes *[]ActivityType `json:"excludedActivityTypes,omitempty"`
}
