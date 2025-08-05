package types

type WorkspaceMember struct {
	_embedded *WorkspaceMemberEmbedded `json:"_embedded"`
	CreatedAt string `json:"createdAt"`
	DeactivatedAt *string `json:"deactivatedAt"`
	Status MemberStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt"`
	UserId string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
}
