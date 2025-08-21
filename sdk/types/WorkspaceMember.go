package types

type WorkspaceMember struct {
	Embedded *WorkspaceMemberEmbedded `json:"_embedded,omitempty"`
	CreatedAt string `json:"createdAt"`
	DeactivatedAt *string `json:"deactivatedAt,omitempty"`
	Status MemberStatus `json:"status"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserID string `json:"userId"`
	WorkspaceID string `json:"workspaceId"`
}
