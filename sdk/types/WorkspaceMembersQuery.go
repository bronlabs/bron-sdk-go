package types

type WorkspaceMembersQuery struct {
	IncludePermissionGroups *bool `json:"includePermissionGroups,omitempty"`
	IncludeUsersProfiles *bool `json:"includeUsersProfiles,omitempty"`
	IncludeEmails *bool `json:"includeEmails,omitempty"`
}
