package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultBaseURL = "https://business.bfinance.app"
	APIPrefix      = "/external/api"
)

// LogLevel controls the verbosity of request logging.
type LogLevel string

const (
	LogNone  LogLevel = "none"
	LogError LogLevel = "error"
	LogInfo  LogLevel = "info"
	LogDebug LogLevel = "debug"
)

// LoggingConfig controls HTTP request/response logging.
type LoggingConfig struct {
	Level          LogLevel
	Logger         *log.Logger
	IncludeBody    bool
	IncludeHeaders bool
}

// Config holds the HTTP client configuration.
type Config struct {
	APIToken  string
	BaseURL   string
	TimeoutMs int
	Headers   map[string]string
	Logging   *LoggingConfig
}

// Client is the low-level HTTP client for the BFinance API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiToken   string
	headers    map[string]string
	logging    *LoggingConfig
}

// New creates a new HTTP client.
func New(cfg Config) *Client {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	timeout := cfg.TimeoutMs
	if timeout == 0 {
		timeout = 3000
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		},
		baseURL:  baseURL + APIPrefix,
		apiToken: cfg.APIToken,
		headers:  cfg.Headers,
		logging:  cfg.Logging,
	}
}

// RequestOptions holds per-request options.
type RequestOptions struct {
	Headers        map[string]string
	IdempotencyKey string
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, params url.Values, opts *RequestOptions, result any) error {
	return c.do(ctx, http.MethodGet, path, params, nil, opts, result)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body any, opts *RequestOptions, result any) error {
	return c.do(ctx, http.MethodPost, path, nil, body, opts, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, path string, opts *RequestOptions, result any) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil, opts, result)
}

func (c *Client) do(ctx context.Context, method, path string, params url.Values, body any, opts *RequestOptions, result any) error {
	fullURL := c.baseURL + path
	if params != nil && len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return &NetworkError{Message: err.Error()}
	}

	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	if opts != nil {
		for k, v := range opts.Headers {
			req.Header.Set(k, v)
		}
		if opts.IdempotencyKey != "" {
			req.Header.Set("Idempotency-Key", opts.IdempotencyKey)
		}
	}

	c.logRequest(method, fullURL)

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		return &NetworkError{Message: err.Error()}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &NetworkError{Message: fmt.Sprintf("read response body: %s", err)}
	}

	c.logResponse(resp.StatusCode, elapsed)

	if resp.StatusCode >= 400 {
		var apiErr struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			ErrorID string `json:"errorId"`
			Data    any    `json:"data"`
		}
		_ = json.Unmarshal(respBody, &apiErr)
		msg := apiErr.Message
		if msg == "" {
			msg = http.StatusText(resp.StatusCode)
		}
		return &HttpError{
			Message:    msg,
			StatusCode: resp.StatusCode,
			Details:    apiErr.Data,
		}
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	}

	return nil
}

func (c *Client) logRequest(method, url string) {
	if c.logging == nil || c.logging.Level == LogNone {
		return
	}
	logger := c.logging.Logger
	if logger == nil {
		logger = log.Default()
	}
	if c.logging.Level == LogInfo || c.logging.Level == LogDebug {
		logger.Printf("[bfinance] --> %s %s", method, url)
	}
}

func (c *Client) logResponse(status int, elapsed time.Duration) {
	if c.logging == nil || c.logging.Level == LogNone {
		return
	}
	logger := c.logging.Logger
	if logger == nil {
		logger = log.Default()
	}
	if c.logging.Level == LogError && status >= 400 {
		logger.Printf("[bfinance] <-- %d (%s)", status, elapsed)
	} else if c.logging.Level == LogInfo || c.logging.Level == LogDebug {
		logger.Printf("[bfinance] <-- %d (%s)", status, elapsed)
	}
}
