package types

type Workspace struct {
	Tag string `json:"tag"`
	WorkspaceId string `json:"workspaceId"`
	ImageId *string `json:"imageId"`
	Name string `json:"name"`
}
