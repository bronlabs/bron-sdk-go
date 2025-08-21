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

	"math/rand"
	"strconv"

	"github.com/bronlabs/bron-sdk-go/sdk/auth"
)

type RequestOptions struct {
	Method string
	Path   string
	Body   interface{}
	Query  interface{}
}

// APIError is returned for non-2xx responses
type APIError struct {
	Status    int
	Code      string
	Message   string
	RequestID string
}

func (e *APIError) Error() string { return e.Message }

type RetryPolicy struct {
	Max  int
	Base time.Duration
}

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	signer     func(auth.BronJwtOptions) (string, error)
	clock      func() time.Time
	retry      *RetryPolicy
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

func (c *Client) SetRetryPolicy(p RetryPolicy) { c.retry = &p }

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

	var bodyBytes []byte
	if options.Body != nil {
		jb, err := json.Marshal(options.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyBytes = jb
	}

	attempt := 0
	for {
		var body io.Reader
		if bodyBytes != nil {
			body = bytes.NewBuffer(append([]byte(nil), bodyBytes...))
		}

		req, err := http.NewRequestWithContext(ctx, options.Method, url, body)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		bodyStr := ""
		if bodyBytes != nil {
			bodyStr = string(bodyBytes)
		}
		var jwk map[string]interface{}
		json.Unmarshal([]byte(c.apiKey), &jwk)
		kid, _ := jwk["kid"].(string)
		issuedAt := c.clock().Unix()
		token, err := c.signer(auth.BronJwtOptions{Method: options.Method, Path: pathWithQuery, Body: bodyStr, Kid: kid, PrivateKey: c.apiKey, Iat: &issuedAt})
		if err != nil {
			return fmt.Errorf("failed to generate JWT: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "ApiKey "+token)
		req.Header.Set("User-Agent", "Bron SDK Go/"+SDK_VERSION)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			// network error: retry only if policy allows and method idempotent
			if c.shouldRetry(options.Method, 0) && attempt < c.maxAttempts() {
				if waitErr := c.sleepBackoff(ctx, attempt, 0); waitErr == nil {
					attempt++
					continue
				}
			}
			return fmt.Errorf("failed to execute request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			// Retry on 429/5xx for idempotent methods
			if (resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500) && c.shouldRetry(options.Method, resp.StatusCode) && attempt < c.maxAttempts() {
				if waitErr := c.sleepFromRetryAfterOrBackoff(ctx, resp, attempt); waitErr == nil {
					attempt++
					continue
				}
			}
			b, _ := io.ReadAll(resp.Body)
			return c.parseAPIError(resp, b)
		}

		if result != nil {
			if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
				return fmt.Errorf("failed to decode response: %w", err)
			}
		}
		return nil
	}
}

func (c *Client) maxAttempts() int {
	if c.retry == nil || c.retry.Max <= 0 {
		return 0
	}
	return c.retry.Max
}

func (c *Client) shouldRetry(method string, status int) bool {
	// Only retry idempotent GET by default
	return strings.EqualFold(method, http.MethodGet)
}

func (c *Client) sleepFromRetryAfterOrBackoff(ctx context.Context, resp *http.Response, attempt int) error {
	if ra := resp.Header.Get("Retry-After"); ra != "" {
		if secs, err := strconv.Atoi(ra); err == nil {
			return c.sleepCtx(ctx, time.Duration(secs)*time.Second)
		}
		if t, err := http.ParseTime(ra); err == nil {
			d := time.Until(t)
			if d > 0 {
				return c.sleepCtx(ctx, d)
			}
		}
	}
	return c.sleepBackoff(ctx, attempt, resp.StatusCode)
}

func (c *Client) sleepBackoff(ctx context.Context, attempt int, _ int) error {
	if c.retry == nil || c.retry.Base <= 0 {
		return nil
	}
	delay := c.retry.Base << attempt
	// jitter: +/- 50%
	j := time.Duration(rand.Int63n(int64(delay))) - delay/2
	return c.sleepCtx(ctx, delay+j)
}

func (c *Client) sleepCtx(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func (c *Client) parseAPIError(resp *http.Response, body []byte) error {
	reqID := resp.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = resp.Header.Get("X-Request-Id")
	}
	if reqID == "" {
		reqID = resp.Header.Get("X-Correlation-ID")
	}
	type errPayload struct {
		Code             string `json:"code"`
		Message          string `json:"message"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		RequestID        string `json:"requestId"`
	}
	var p errPayload
	_ = json.Unmarshal(body, &p)
	msg := p.Message
	if msg == "" {
		msg = p.Error
	}
	if msg == "" {
		msg = p.ErrorDescription
	}
	if msg == "" {
		msg = string(body)
	}
	code := p.Code
	if code == "" {
		code = http.StatusText(resp.StatusCode)
	}
	if reqID == "" {
		reqID = p.RequestID
	}
	return &APIError{Status: resp.StatusCode, Code: code, Message: msg, RequestID: reqID}
}
