package types

type WorkspaceMemberEmbedded struct {
	Profile *UserProfile `json:"profile"`
	Identities *[]Identity `json:"identities"`
	PermissionGroups *[]string `json:"permissionGroups"`
}
