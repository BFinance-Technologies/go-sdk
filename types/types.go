package types

import "time"

// ApiResponse represents a generic API response wrapper.
type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	ErrorID string `json:"errorId,omitempty"`
}

// MessageResponse represents a simple status+message response.
type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PaginationParams holds pagination options for list endpoints.
type PaginationParams struct {
	Page  *int `json:"page,omitempty" url:"page,omitempty"`
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
}

// RequestOptions holds per-request options.
type RequestOptions struct {
	Headers        map[string]string
	IdempotencyKey string
}

// DateResponse can be either a string or time.Time depending on context.
type DateResponse = time.Time
