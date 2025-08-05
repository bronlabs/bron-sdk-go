package types

type WorkspaceMemberEmbedded struct {
	Identities *[]Identity `json:"identities"`
	PermissionGroups *[]string `json:"permissionGroups"`
	Profile *UserProfile `json:"profile"`
}
