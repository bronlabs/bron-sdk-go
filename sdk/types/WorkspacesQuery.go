package types

type WorkspacesQuery struct {
	WorkspaceIds *[]string `json:"workspaceIds,omitempty"`
	IncludeSettings *bool `json:"includeSettings,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
	AccountIds *[]string `json:"accountIds,omitempty"`
	Search *string `json:"search,omitempty"`
	UserIds *[]string `json:"userIds,omitempty"`
	ActivityTypes *[]ActivityType `json:"activityTypes,omitempty"`
	ExcludedActivityTypes *[]ActivityType `json:"excludedActivityTypes,omitempty"`
	IncludePermissionGroups *bool `json:"includePermissionGroups,omitempty"`
	IncludeUsersProfiles *bool `json:"includeUsersProfiles,omitempty"`
	IncludeEmails *bool `json:"includeEmails,omitempty"`
}
