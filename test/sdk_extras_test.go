package test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
	sdkhttp "github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/types"
)

func TestContextCancellation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer ts.Close()
	c := sdkhttp.NewClient(ts.URL, genMockJwk())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	var out map[string]any
	err := c.RequestWithContext(ctx, &out, sdkhttp.RequestOptions{Method: "GET", Path: "/workspaces/ws"})
	if !errors.Is(err, context.DeadlineExceeded) && !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestRetryGet(t *testing.T) {
	attempts := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts <= 2 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code":"ERR","message":"fail"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()
	c := sdkhttp.NewClient(ts.URL, genMockJwk())
	c.SetRetryPolicy(sdkhttp.RetryPolicy{Max: 2, Base: 1 * time.Millisecond})
	var out map[string]any
	if err := c.Request(&out, sdkhttp.RequestOptions{Method: "GET", Path: "/test"}); err != nil {
		t.Fatal(err)
	}
	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestNoRetryPostAndAPIError(t *testing.T) {
	attempts := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":"SERVER","message":"boom","requestId":"req-1"}`))
	}))
	defer ts.Close()
	c := sdkhttp.NewClient(ts.URL, genMockJwk())
	c.SetRetryPolicy(sdkhttp.RetryPolicy{Max: 5, Base: 1 * time.Millisecond})
	var out map[string]any
	err := c.Request(&out, sdkhttp.RequestOptions{Method: "POST", Path: "/test", Body: map[string]any{"a": 1}})
	if err == nil {
		t.Fatal("expected error")
	}
	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}
	var apiErr *sdkhttp.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.Status != 500 || apiErr.Code == "" || apiErr.Message == "" {
		t.Fatalf("bad api error: %+v", apiErr)
	}
}

func TestSignerDI(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")
		if a != "ApiKey TESTTOKEN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"code":"BAD","message":"bad auth"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer ts.Close()
	c := sdkhttp.NewClient(ts.URL, genMockJwk())
	c.SetSigner(func(o auth.BronJwtOptions) (string, error) { return "TESTTOKEN", nil })
	var out map[string]any
	if err := c.Request(&out, sdkhttp.RequestOptions{Method: "GET", Path: "/ok"}); err != nil {
		t.Fatal(err)
	}
}

func TestBuildersAndJSON(t *testing.T) {
	nid := "testETH"
	tx := types.NewWithdrawalTx("acc", "ext", types.WithdrawalParams{Amount: "1", NetworkID: &nid})
	b, err := json.Marshal(tx)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	for _, want := range []string{"\"accountId\"", "\"externalId\"", "\"params\"", "\"amount\"", "\"networkId\""} {
		if !strings.Contains(s, want) {
			t.Fatalf("missing %s in %s", want, s)
		}
	}
}

func TestInitialismsPresent(t *testing.T) {
	if _, ok := reflect.TypeOf(types.Transaction{}).FieldByName("TransactionID"); !ok {
		t.Fatal("TransactionID missing")
	}
	if _, ok := reflect.TypeOf(types.WithdrawalParams{}).FieldByName("NetworkID"); !ok {
		t.Fatal("NetworkID missing")
	}
	if _, ok := reflect.TypeOf(types.BalancesQuery{}).FieldByName("AccountIDs"); !ok {
		t.Fatal("AccountIDs missing")
	}
}
