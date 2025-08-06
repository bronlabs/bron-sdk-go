package types

type WorkspaceMemberEmbedded struct {
	Identities *[]Identity `json:"identities,omitempty"`
	PermissionGroups *[]string `json:"permissionGroups,omitempty"`
	Profile *UserProfile `json:"profile,omitempty"`
}
