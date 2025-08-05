package api

import (
	"fmt"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
)

type StakeAPI struct {
	http *http.Client
	workspaceID string
}

func NewStakeAPI(http *http.Client, workspaceID string) *StakeAPI {
	return &StakeAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

	func (api *StakeAPI) GetStakes() (*types.Stakes, error) {
		path := fmt.Sprintf("/stakes/", api.workspaceID)
		var result types.Stakes
		options := http.RequestOptions{
			Method: "GET",
			Path:   path,
		}
		err := api.http.Request(&result, options)
		return &result, err
	}


