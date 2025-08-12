package api

import (
	"fmt"

	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type WorkspacesAPI struct {
	http        *http.Client
	workspaceID string
}

func NewWorkspacesAPI(http *http.Client, workspaceID string) *WorkspacesAPI {
	return &WorkspacesAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *WorkspacesAPI) GetWorkspaceById(query ...*types.WorkspaceByIdQuery) (*types.Workspace, error) {
	path := fmt.Sprintf("/workspaces/%s", api.workspaceID)
	var result types.Workspace
	var queryParam *types.WorkspaceByIdQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *WorkspacesAPI) GetActivities(query ...*types.ActivitiesQuery) (*types.Activities, error) {
	path := fmt.Sprintf("/workspaces/%s/activities", api.workspaceID)
	var result types.Activities
	var queryParam *types.ActivitiesQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

func (api *WorkspacesAPI) GetWorkspaceMembers(query ...*types.WorkspaceMembersQuery) (*types.WorkspaceMembers, error) {
	path := fmt.Sprintf("/workspaces/%s/members", api.workspaceID)
	var result types.WorkspaceMembers
	var queryParam *types.WorkspaceMembersQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.Request(&result, options)
	return &result, err
}

