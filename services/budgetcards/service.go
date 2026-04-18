package budgetcards

import (
	"context"
	"fmt"
	"net/url"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/budget-cards"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// Issue issues a new budget card.
func (s *Service) Issue(ctx context.Context, req IssueBudgetCardRequest) (*types.ApiResponse[BudgetCard], error) {
	out := &types.ApiResponse[BudgetCard]{}
	if err := s.c.Post(ctx, basePath+"/issue", req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetByID returns a budget card by ID.
func (s *Service) GetByID(ctx context.Context, cardID string) (*types.ApiResponse[BudgetCard], error) {
	out := &types.ApiResponse[BudgetCard]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) Freeze(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/freeze", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) Unfreeze(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/unfreeze", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) DeleteCard(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Delete(ctx, fmt.Sprintf("%s/%s/delete", basePath, url.PathEscape(cardID)), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) SetPin(ctx context.Context, cardID, pin string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"pin": pin}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/pin", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) UpdateEmail(ctx context.Context, cardID, email string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"email": email}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/email", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) UpdatePhoneNumber(ctx context.Context, cardID, phone string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"phone": phone}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/phone", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) SetVelocityLimits(ctx context.Context, cardID string, req SetVelocityLimitsRequest) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/velocity", basePath, url.PathEscape(cardID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) GetSensitive(ctx context.Context, cardID string) (*types.ApiResponse[GetSensitiveData], error) {
	out := &types.ApiResponse[GetSensitiveData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/sensitive", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
