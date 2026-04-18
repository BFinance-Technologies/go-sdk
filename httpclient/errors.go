package httpclient

import "fmt"

// HttpError represents an HTTP error response from the API.
type HttpError struct {
	Message    string
	StatusCode int
	Details    any
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// NetworkError represents a network-level error (timeout, DNS, etc).
type NetworkError struct {
	Message string
	Details any
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("Network error: %s", e.Message)
}
