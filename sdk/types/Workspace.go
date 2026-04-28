package types

type Workspace struct {
	Icon *string `json:"icon,omitempty"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	WorkspaceID string `json:"workspaceId"`
}
