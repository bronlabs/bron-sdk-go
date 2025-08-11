package test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	sdkhttp "github.com/bronlabs/bron-sdk-go/sdk/http"
	"github.com/bronlabs/bron-sdk-go/sdk/version"
)

func TestHttpClient(t *testing.T) {
	client := sdkhttp.NewClient("https://api.bron.org", "test-key")
	if client == nil {
		t.Fatal("Failed to create client")
	}
}

func TestHttpClientUserAgent(t *testing.T) {
	// Create a test server to capture headers
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		userAgent := r.Header.Get("User-Agent")
		expectedUserAgent := "Bron SDK Go/" + version.SDK_VERSION

		if userAgent != expectedUserAgent {
			t.Errorf("Expected User-Agent: %s, got: %s", expectedUserAgent, userAgent)
		}

		w.WriteHeader(stdhttp.StatusOK)
		w.Write([]byte(`{"test": "response"}`))
	}))
	defer server.Close()

	client := sdkhttp.NewClient(server.URL, genMockJwk())

	var result map[string]interface{}
	err := client.Request(&result, sdkhttp.RequestOptions{
		Method: "GET",
		Path:   "/test",
	})

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}
