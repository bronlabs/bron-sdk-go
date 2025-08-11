package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strings"
	"time"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
	"github.com/bronlabs/bron-sdk-go/sdk/version"
)

type RequestOptions struct {
	Method string
	Path   string
	Body   interface{}
	Query  interface{}
}

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c *Client) Request(result interface{}, options RequestOptions) error {
	pathWithQuery := options.Path
	url := c.baseURL + pathWithQuery
	if options.Query != nil {
		// Build query string from struct
		queryBytes, _ := json.Marshal(options.Query)
		var qMap map[string]interface{}
		json.Unmarshal(queryBytes, &qMap)
		if len(qMap) > 0 {
			params := urlpkg.Values{}
			for k, v := range qMap {
				if v == nil {
					continue
				}
				switch val := v.(type) {
				case []interface{}:
					strVals := make([]string, len(val))
					for i, e := range val {
						strVals[i] = fmt.Sprintf("%v", e)
					}
					params.Set(k, strings.Join(strVals, ","))
				default:
					params.Set(k, fmt.Sprintf("%v", val))
				}
			}
			if qs := params.Encode(); qs != "" {
				pathWithQuery += "?" + params.Encode()
				url += "?" + qs
			}
		}
	}

	var body io.Reader
	if options.Body != nil {
		jsonBody, err := json.Marshal(options.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(options.Method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Generate Bron JWT
	bodyStr := ""
	if options.Body != nil {
		jsonBody, _ := json.Marshal(options.Body)
		bodyStr = string(jsonBody)
	}

	// Parse JWK to get kid
	var jwk map[string]interface{}
	json.Unmarshal([]byte(c.apiKey), &jwk)
	kid, _ := jwk["kid"].(string)

	token, err := auth.GenerateBronJwt(auth.BronJwtOptions{
		Method:     options.Method,
		Path:       pathWithQuery,
		Body:       bodyStr,
		Kid:        kid,
		PrivateKey: c.apiKey,
	})
	if err != nil {
		return fmt.Errorf("failed to generate JWT: %w", err)
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "ApiKey "+token)
	req.Header.Set("User-Agent", "Bron SDK Go/"+version.SDK_VERSION)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
