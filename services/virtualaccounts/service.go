package virtualaccounts

import (
	"context"
	"fmt"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/virtual-account"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// GetList returns all virtual accounts for the customer.
func (s *Service) GetList(ctx context.Context, customerID string) (*types.ApiResponse[[]VirtualAccount], error) {
	out := &types.ApiResponse[[]VirtualAccount]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s", basePath, url.PathEscape(customerID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetEligibility checks whether the customer is eligible for virtual accounts.
func (s *Service) GetEligibility(ctx context.Context, customerID string) (*types.ApiResponse[EligibilityData], error) {
	out := &types.ApiResponse[EligibilityData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/eligibility", basePath, url.PathEscape(customerID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a virtual account for the customer.
func (s *Service) Create(ctx context.Context, customerID string, req CreateVirtualAccountRequest) (*types.ApiResponse[CreateVirtualAccountData], error) {
	out := &types.ApiResponse[CreateVirtualAccountData]{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/create", basePath, url.PathEscape(customerID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
