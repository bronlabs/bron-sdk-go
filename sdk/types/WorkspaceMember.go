package types

type WorkspaceMember struct {
	_embedded *WorkspaceMemberEmbedded `json:"_embedded,omitempty"`
	CreatedAt string `json:"createdAt"`
	DeactivatedAt *string `json:"deactivatedAt,omitempty"`
	Status MemberStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserId string `json:"userId"`
	WorkspaceId string `json:"workspaceId"`
}
