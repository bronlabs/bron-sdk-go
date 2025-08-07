package test

import (
	"testing"

	sdk "github.com/bronlabs/bron-sdk-go/sdk"
)

func TestBronClientInit(t *testing.T) {
	client := sdk.NewBronClient(sdk.BronClientConfig{APIKey: "test", WorkspaceID: "ws"})
	if client == nil {
		t.Fatalf("client is nil")
	}
	// validate sub-APIs wired
	if client.Accounts == nil || client.Balances == nil || client.Transactions == nil || client.Assets == nil {
		t.Fatalf("sub-APIs not initialised")
	}
}
