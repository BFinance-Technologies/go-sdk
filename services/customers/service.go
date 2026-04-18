package customers

import (
	"context"
	"fmt"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/customers"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// CreateIndividual creates a new individual customer.
func (s *Service) CreateIndividual(ctx context.Context, req CreateIndividualCustomerRequest) (*types.ApiResponse[CreateIndividualCustomerData], error) {
	out := &types.ApiResponse[CreateIndividualCustomerData]{}
	if err := s.c.Post(ctx, basePath, req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// CreateViaSumsub creates a customer via a Sumsub share token.
func (s *Service) CreateViaSumsub(ctx context.Context, req CreateCustomerViaSumsubRequest) (*types.ApiResponse[CreateCustomerViaSumsubData], error) {
	out := &types.ApiResponse[CreateCustomerViaSumsubData]{}
	if err := s.c.Post(ctx, basePath+"/sumsub", req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetByID returns a customer by ID.
func (s *Service) GetByID(ctx context.Context, customerID string) (*types.ApiResponse[Customer], error) {
	out := &types.ApiResponse[Customer]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s", basePath, url.PathEscape(customerID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// RequestFeatureAccess requests access to a feature (cards / virtual-account) for a customer.
func (s *Service) RequestFeatureAccess(ctx context.Context, customerID string, req RequestFeatureAccessRequest) (*types.ApiResponse[RequestFeatureAccessData], error) {
	out := &types.ApiResponse[RequestFeatureAccessData]{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/request", basePath, url.PathEscape(customerID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
