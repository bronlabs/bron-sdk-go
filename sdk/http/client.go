package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strings"
	"time"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
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
	signer     func(auth.BronJwtOptions) (string, error)
	clock      func() time.Time
}

func NewClient(baseURL, apiKey string) *Client {
	return NewClientWithHTTP(baseURL, apiKey, &http.Client{Timeout: 30 * time.Second})
}

func NewClientWithHTTP(baseURL, apiKey string, hc *http.Client) *Client {
	if hc == nil {
		hc = &http.Client{Timeout: 30 * time.Second}
	}
	return &Client{
		httpClient: hc,
		baseURL:    baseURL,
		apiKey:     apiKey,
		signer:     auth.GenerateBronJwt,
		clock:      time.Now,
	}
}

func (c *Client) SetSigner(s func(auth.BronJwtOptions) (string, error)) {
	if s != nil {
		c.signer = s
	}
}

func (c *Client) SetClock(cl func() time.Time) {
	if cl != nil {
		c.clock = cl
	}
}

func (c *Client) Request(result interface{}, options RequestOptions) error {
	return c.RequestWithContext(context.Background(), result, options)
}

func (c *Client) RequestWithContext(ctx context.Context, result interface{}, options RequestOptions) error {
	if ctx == nil {
		ctx = context.Background()
	}
	pathWithQuery := options.Path
	url := c.baseURL + pathWithQuery
	if options.Query != nil {
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

	req, err := http.NewRequestWithContext(ctx, options.Method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	bodyStr := ""
	if options.Body != nil {
		jsonBody, _ := json.Marshal(options.Body)
		bodyStr = string(jsonBody)
	}
	var jwk map[string]interface{}
	json.Unmarshal([]byte(c.apiKey), &jwk)
	kid, _ := jwk["kid"].(string)
	issuedAt := c.clock().Unix()
	token, err := c.signer(auth.BronJwtOptions{
		Method:     options.Method,
		Path:       pathWithQuery,
		Body:       bodyStr,
		Kid:        kid,
		PrivateKey: c.apiKey,
		Iat:        &issuedAt,
	})
	if err != nil {
		return fmt.Errorf("failed to generate JWT: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "ApiKey "+token)
	req.Header.Set("User-Agent", "Bron SDK Go/"+SDK_VERSION)
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
