package types

type Workspace struct {
	ImageId *string `json:"imageId,omitempty"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	WorkspaceId string `json:"workspaceId"`
}
