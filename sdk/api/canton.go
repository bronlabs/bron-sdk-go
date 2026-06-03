package api

import (
	"fmt"

	"context"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type CantonAPI struct {
	http        *http.Client
	workspaceID string
}

func NewCantonAPI(http *http.Client, workspaceID string) *CantonAPI {
	return &CantonAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *CantonAPI) CantonLedgerAPIPassthrough(ctx context.Context, body types.CantonLedgerQuery) (*types.CantonLedgerQueryResult, error) {
	path := fmt.Sprintf("/workspaces/%s/canton/ledger-query", api.workspaceID)
	var result types.CantonLedgerQueryResult
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

