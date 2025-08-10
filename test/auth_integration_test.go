package test

import (
	"os"
	"testing"

	"github.com/bronlabs/bron-sdk-go/sdk"
)

func TestAuthIntegration(t *testing.T) {
	if os.Getenv("BRON_API_KEY") == "" {
		t.Skip("set env vars to run")
	}
	c := sdk.NewBronClient(sdk.BronClientConfig{APIKey: os.Getenv("BRON_API_KEY"), WorkspaceID: os.Getenv("BRON_WORKSPACE_ID")})
	if _, err := c.Workspaces.GetWorkspaceById(nil); err != nil {
		t.Fatalf("api error %v", err)
	}
}
