package types

type WorkspaceMember struct {
	CreatedAt string `json:"createdAt"`
	DeactivatedAt *string `json:"deactivatedAt"`
	Status MemberStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	UserId string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
	_embedded *WorkspaceMemberEmbedded `json:"_embedded"`
}
