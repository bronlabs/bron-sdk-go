package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	sdk "github.com/bronlabs/bron-sdk-go/sdk"
	"github.com/bronlabs/bron-sdk-go/sdk/auth"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

func newTestClient(srv *httptest.Server, signer func(auth.BronJwtOptions) (string, error)) *sdk.BronClient {
	return sdk.NewBronClientWithOptions(
		sdk.BronClientConfig{
			BaseURL:     srv.URL,
			WorkspaceID: "ws-test",
			APIKey:      `{"kid":"dummy"}`,
		},
		sdk.WithSigner(signer),
		sdk.WithClock(func() time.Time { return time.Unix(1700000000, 0) }),
	)
}

func dummySigner() func(auth.BronJwtOptions) (string, error) {
	return func(o auth.BronJwtOptions) (string, error) {
		return "dummy.jwt.token", nil
	}
}

func TestPathEscape_SlashInTransactionId(t *testing.T) {
	var capturedURI string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURI = r.URL.RequestURI()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := newTestClient(srv, dummySigner())

	_, _ = c.Transactions.CancelTransaction(
		context.Background(),
		"tx-target/approve?ignored=",
		types.CancelTransaction{},
	)

	if !strings.Contains(capturedURI, "%2F") {
		t.Fatalf("slash in transactionId must be percent-encoded, got URI: %s", capturedURI)
	}

	if !strings.HasSuffix(strings.SplitN(capturedURI, "?", 2)[0], "/cancel") {
		t.Fatalf("the /cancel suffix must remain as the final path segment, got URI: %s", capturedURI)
	}

	if strings.Contains(capturedURI, "/approve") {
		t.Fatalf("slash must not create a separate /approve path segment, got URI: %s", capturedURI)
	}
}

func TestPathEscape_SignerReceivesEscapedPath(t *testing.T) {
	var signedPath string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := newTestClient(srv, func(o auth.BronJwtOptions) (string, error) {
		signedPath = o.Path
		return "dummy.jwt.token", nil
	})

	_, _ = c.Transactions.GetTransactionByID(
		context.Background(),
		"id/with?special",
	)

	if !strings.Contains(signedPath, "%2F") {
		t.Fatalf("slash must be escaped in signed path, got: %s", signedPath)
	}

	if !strings.Contains(signedPath, "%3F") {
		t.Fatalf("question mark must be escaped in signed path, got: %s", signedPath)
	}
}

func TestPathEscape_NormalIdUnchanged(t *testing.T) {
	var capturedURI string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURI = r.URL.RequestURI()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := newTestClient(srv, dummySigner())

	_, _ = c.Transactions.CancelTransaction(
		context.Background(),
		"tx-normal-uuid-1234",
		types.CancelTransaction{},
	)

	want := "/workspaces/ws-test/transactions/tx-normal-uuid-1234/cancel"
	if capturedURI != want {
		t.Fatalf("normal IDs must pass through unchanged\ngot:  %s\nwant: %s", capturedURI, want)
	}
}
