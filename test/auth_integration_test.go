package test

import (
	"context"
	"os"
	"testing"

	"github.com/bronlabs/bron-sdk-go/sdk"
)

func TestAuthIntegration(t *testing.T) {
	if os.Getenv("BRON_API_KEY") == "" {
		t.Skip("set env vars to run")
	}
	baseURL := os.Getenv("BRON_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.bron.org"
	}
	c := sdk.NewBronClient(sdk.BronClientConfig{APIKey: os.Getenv("BRON_API_KEY"), WorkspaceID: os.Getenv("BRON_WORKSPACE_ID"), BaseURL: baseURL})
	if _, err := c.Workspaces.GetWorkspaceById(context.Background(), nil); err != nil {
		t.Fatalf("api error %v", err)
	}
}
