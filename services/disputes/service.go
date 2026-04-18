package disputes

import (
	"context"
	"fmt"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/disputes"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// Create creates a new dispute.
func (s *Service) Create(ctx context.Context, req CreateDisputeRequest) (*types.ApiResponse[CreateDisputeData], error) {
	out := &types.ApiResponse[CreateDisputeData]{}
	if err := s.c.Post(ctx, basePath, req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetStatus returns the current status of a dispute.
func (s *Service) GetStatus(ctx context.Context, disputeID string) (*types.ApiResponse[GetDisputeStatusData], error) {
	out := &types.ApiResponse[GetDisputeStatusData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/status", basePath, url.PathEscape(disputeID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Cancel cancels a dispute.
func (s *Service) Cancel(ctx context.Context, disputeID string) (*types.ApiResponse[CancelDisputeData], error) {
	out := &types.ApiResponse[CancelDisputeData]{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/cancel", basePath, url.PathEscape(disputeID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
