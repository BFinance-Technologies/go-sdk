// Package bfinance is the official Go SDK for the BFinance API.
package bfinance

import (
	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/services/balance"
	"github.com/BFinance-Technologies/go-sdk/services/budgetcards"
	"github.com/BFinance-Technologies/go-sdk/services/customers"
	"github.com/BFinance-Technologies/go-sdk/services/disputes"
	"github.com/BFinance-Technologies/go-sdk/services/esim"
	"github.com/BFinance-Technologies/go-sdk/services/finance"
	"github.com/BFinance-Technologies/go-sdk/services/physicalcards"
	"github.com/BFinance-Technologies/go-sdk/services/prepaidcards"
	"github.com/BFinance-Technologies/go-sdk/services/utils"
	"github.com/BFinance-Technologies/go-sdk/services/virtualaccounts"
)

// Config holds the BFinance client configuration.
type Config struct {
	// APIToken is the bearer token used to authenticate requests. Required.
	APIToken string
	// TimeoutMs is the request timeout in milliseconds. Defaults to 3000.
	TimeoutMs int
	// BaseURL overrides the default API base URL (https://business.bfinance.app).
	BaseURL string
	// Headers are extra headers sent with every request.
	Headers map[string]string
	// Logging controls HTTP request/response logging.
	Logging *httpclient.LoggingConfig
}

// Client is the root BFinance SDK client.
type Client struct {
	http *httpclient.Client

	PrepaidCards    *prepaidcards.Service
	BudgetCards     *budgetcards.Service
	PhysicalCards   *physicalcards.Service
	Customers       *customers.Service
	VirtualAccounts *virtualaccounts.Service
	Disputes        *disputes.Service
	Esim            *esim.Service
	Balance         *balance.Service
	Finance         *finance.Service
	Utils           *utils.Service
}

// New creates a new BFinance client.
func New(cfg Config) *Client {
	httpCfg := httpclient.Config{
		APIToken:  cfg.APIToken,
		BaseURL:   cfg.BaseURL,
		TimeoutMs: cfg.TimeoutMs,
		Headers:   cfg.Headers,
		Logging:   cfg.Logging,
	}
	h := httpclient.New(httpCfg)

	return &Client{
		http:            h,
		PrepaidCards:    prepaidcards.New(h),
		BudgetCards:     budgetcards.New(h),
		PhysicalCards:   physicalcards.New(h),
		Customers:       customers.New(h),
		VirtualAccounts: virtualaccounts.New(h),
		Disputes:        disputes.New(h),
		Esim:            esim.New(h),
		Balance:         balance.New(h),
		Finance:         finance.New(h),
		Utils:           utils.New(h),
	}
}
