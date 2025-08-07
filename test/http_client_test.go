package test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	sdkhttp "github.com/bronlabs/bron-sdk-go/sdk/http"
)

func TestHttpClient(t *testing.T) {
	ts := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if h := r.Header.Get("Authorization"); !strings.HasPrefix(h, "ApiKey ") {
			t.Fatalf("no auth header")
		}
		if r.URL.RawQuery != "limit=10&offset=0" {
			t.Fatalf("bad query")
		}
		w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	jwk := genMockJwk()
	client := sdkhttp.NewClient(ts.URL, jwk)
	_ = client.Request(nil, sdkhttp.RequestOptions{
		Method: "GET",
		Path:   "/foo",
		Query: struct {
			Limit  int `json:"limit"`
			Offset int `json:"offset"`
		}{10, 0},
	})
}
