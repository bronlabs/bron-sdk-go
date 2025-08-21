package types

type Workspace struct {
	ImageID *string `json:"imageId,omitempty"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	WorkspaceID string `json:"workspaceId"`
}
