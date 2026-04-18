package physicalcards

import (
	"context"
	"fmt"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/physical-cards"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// Order orders a physical card.
func (s *Service) Order(ctx context.Context, cardID string, req OrderPhysicalCardRequest) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/order", basePath, url.PathEscape(cardID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Activate activates a physical card.
func (s *Service) Activate(ctx context.Context, cardID string, req ActivatePhysicalCardRequest) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/activate", basePath, url.PathEscape(cardID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
