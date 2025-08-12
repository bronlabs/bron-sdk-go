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

func (api *WorkspacesAPI) GetWorkspaceById(query *types.WorkspaceByIdQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *WorkspacesAPI) GetActivities(query *types.ActivitiesQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/activities", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

func (api *WorkspacesAPI) GetWorkspaceMembers(query *types.WorkspaceMembersQuery) (interface{}, error) {
	path := fmt.Sprintf("/workspaces/%s/members", api.workspaceID)
	var result interface{}
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
		Query:  query,
	}
	err := api.http.Request(&result, options)
	return result, err
}

