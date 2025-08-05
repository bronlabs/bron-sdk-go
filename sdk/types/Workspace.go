package types

type Workspace struct {
	WorkspaceId string `json:"workspaceId"`
	ImageId *string `json:"imageId"`
	Name string `json:"name"`
	Tag string `json:"tag"`
}
