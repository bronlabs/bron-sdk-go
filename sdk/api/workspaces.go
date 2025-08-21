package api

import (
	"fmt"

	"context"
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

func (api *WorkspacesAPI) GetWorkspaceByID(ctx context.Context, query ...*types.WorkspaceByIDQuery) (*types.Workspace, error) {
	path := fmt.Sprintf("/workspaces/%s", api.workspaceID)
	var result types.Workspace
	var queryParam *types.WorkspaceByIDQuery
	if len(query) > 0 && query[0] != nil {
		queryParam = query[0]
	}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  queryParam,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *WorkspacesAPI) GetActivities(ctx context.Context, query ...*types.ActivitiesQuery) (*types.Activities, error) {
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
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *WorkspacesAPI) GetWorkspaceMembers(ctx context.Context, query ...*types.WorkspaceMembersQuery) (*types.WorkspaceMembers, error) {
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
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

