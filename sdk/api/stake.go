package api

import (
	"fmt"

	"context"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type StakeAPI struct {
	http        *http.Client
	workspaceID string
}

func NewStakeAPI(http *http.Client, workspaceID string) *StakeAPI {
	return &StakeAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *StakeAPI) GetStakes(ctx context.Context, query ...*types.StakesQuery) (*types.Stakes, error) {
	path := fmt.Sprintf("/workspaces/%s/stakes", api.workspaceID)
	var result types.Stakes
	var queryParam *types.StakesQuery
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

