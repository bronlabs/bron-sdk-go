package types

type WorkspaceByIdQuery struct {
	WorkspaceIds *[]string `json:"workspaceIds,omitempty"`
	IncludeSettings *bool `json:"includeSettings,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}
