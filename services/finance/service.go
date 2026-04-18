package finance

import (
	"context"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/finance"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// GetCryptoDepositAddress returns a crypto deposit address for the given network/currency.
func (s *Service) GetCryptoDepositAddress(ctx context.Context, network, currency string) (*types.ApiResponse[GetCryptoDepositAddressData], error) {
	q := url.Values{}
	q.Set("network", network)
	q.Set("currency", currency)
	out := &types.ApiResponse[GetCryptoDepositAddressData]{}
	if err := s.c.Get(ctx, basePath+"/deposit", q, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
