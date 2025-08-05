package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bronlabs/bron-sdk-go/src/utils"
)

type RequestOptions struct {
	Method string
	Path   string
	Body   interface{}
	Query  map[string]interface{}
}

type Client struct {
	baseUrl    string
	apiKeyJwk  string
	httpClient *http.Client
}

func NewClient(baseUrl, apiKeyJwk string) *Client {
	return &Client{
		baseUrl:   baseUrl,
		apiKeyJwk: apiKeyJwk,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *Client) Request(result interface{}, options RequestOptions) error {
	fullPath := options.Path

	if options.Query != nil && len(options.Query) > 0 {
		params := url.Values{}
		for key, value := range options.Query {
			switch v := value.(type) {
			case []string:
				params.Set(key, strings.Join(v, ","))
			case string:
				params.Set(key, v)
			default:
				if str, err := json.Marshal(v); err == nil {
					params.Set(key, string(str))
				}
			}
		}
		fullPath += "?" + params.Encode()
	}

	requestURL := h.baseUrl + fullPath

	// Generate JWT for authentication
	privateKey, kid, err := utils.ParseJwkEcPrivateKey(h.apiKeyJwk)
	if err != nil {
		return fmt.Errorf("failed to parse JWK: %w", err)
	}

	bodyStr := ""
	if options.Body != nil {
		bodyBytes, err := json.Marshal(options.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal body: %w", err)
		}
		bodyStr = string(bodyBytes)
	}

	jwt, err := utils.GenerateBronJwt(utils.BronJwtOptions{
		Method:     options.Method,
		Path:       fullPath,
		Body:       bodyStr,
		Kid:        kid,
		PrivateKey: privateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to generate JWT: %w", err)
	}

	var body io.Reader
	if options.Body != nil {
		bodyBytes, _ := json.Marshal(options.Body)
		body = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(options.Method, requestURL, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "ApiKey "+jwt)
	if options.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
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