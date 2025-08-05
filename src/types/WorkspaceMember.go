package types

type WorkspaceMember struct {
	WorkspaceId   string                   `json:"workspaceId"`
	_embedded     *WorkspaceMemberEmbedded `json:"_embedded"`
	CreatedAt     string                   `json:"createdAt"`
	DeactivatedAt *string                  `json:"deactivatedAt"`
	Status        MemberStatus             `json:"status"`
	UpdatedAt     *string                  `json:"updatedAt"`
	UserId        string                   `json:"userId"`
}
