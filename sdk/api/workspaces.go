package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type WorkspacesAPI struct {
	http *http.Client
	workspaceID string
}

func NewWorkspacesAPI(http *http.Client, workspaceID string) *WorkspacesAPI {
	return &WorkspacesAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *WorkspacesAPI) GetWorkspaceMembers() (*types.WorkspaceMembers, error) {
		path := fmt.Sprintf("/workspaces/%s/members", api.workspaceID)
		var result types.WorkspaceMembers
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *WorkspacesAPI) GetActivities() (*types.Activities, error) {
		path := fmt.Sprintf("/workspaces/%s/activities", api.workspaceID)
		var result types.Activities
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


	func (api *WorkspacesAPI) GetWorkspace() (*types.Workspace, error) {
		path := fmt.Sprintf("/workspaces/%s", api.workspaceID)
		var result types.Workspace
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


