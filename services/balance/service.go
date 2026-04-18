package balance

import (
	"context"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/balance"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// GetUserBalance returns the user's balance in EUR and USD.
func (s *Service) GetUserBalance(ctx context.Context) (*types.ApiResponse[GetUserBalanceData], error) {
	out := &types.ApiResponse[GetUserBalanceData]{}
	if err := s.c.Get(ctx, basePath, nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
