package api

import (
	"fmt"

	"context"
	"github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

type IntentsAPI struct {
	http        *http.Client
	workspaceID string
}

func NewIntentsAPI(http *http.Client, workspaceID string) *IntentsAPI {
	return &IntentsAPI{
		http:        http,
		workspaceID: workspaceID,
	}
}

func (api *IntentsAPI) CreateIntentRequest(ctx context.Context, body types.CreateIntent) (*types.Intent, error) {
	path := fmt.Sprintf("/workspaces/%s/intents", api.workspaceID)
	var result types.Intent
	options := http.RequestOptions{
		Method: "POST",
		Path:   path,
		Body:   body,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

func (api *IntentsAPI) GetIntentRequestByID(ctx context.Context, intentId string) (*types.Intent, error) {
	path := fmt.Sprintf("/workspaces/%s/intents/%s", api.workspaceID, intentId)
	var result types.Intent
	options := http.RequestOptions{
		Method: "GET",
		Path:   path,
	}
	err := api.http.RequestWithContext(ctx, &result, options)
	return &result, err
}

