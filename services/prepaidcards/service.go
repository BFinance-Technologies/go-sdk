package prepaidcards

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/prepaid-cards"

// Service exposes the prepaid cards endpoints.
type Service struct {
	c *httpclient.Client
}

// New creates a new PrepaidCards service.
func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// GetList returns a paginated list of prepaid cards.
func (s *Service) GetList(ctx context.Context, params *types.PaginationParams) (*types.ApiResponse[GetPrepaidCardsData], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
	}
	out := &types.ApiResponse[GetPrepaidCardsData]{}
	if err := s.c.Get(ctx, basePath+"/list", q, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Issue issues a new prepaid card.
func (s *Service) Issue(ctx context.Context, payload IssuePrepaidCardRequest) (*types.ApiResponse[IssuePrepaidCardData], error) {
	out := &types.ApiResponse[IssuePrepaidCardData]{}
	if err := s.c.Post(ctx, basePath+"/issue", payload, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Reissue re-issues a prepaid card.
func (s *Service) Reissue(ctx context.Context, payload ReissuePrepaidCardRequest) (*types.ApiResponse[ReissuePrepaidCardData], error) {
	out := &types.ApiResponse[ReissuePrepaidCardData]{}
	if err := s.c.Post(ctx, basePath+"/reissue", payload, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetDetails returns details for a specific prepaid card.
func (s *Service) GetDetails(ctx context.Context, cardID string) (*types.ApiResponse[PrepaidCardDetails], error) {
	out := &types.ApiResponse[PrepaidCardDetails]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetTransactions lists card transactions.
func (s *Service) GetTransactions(ctx context.Context, cardID string, params *types.PaginationParams) (*types.ApiResponse[GetTransactionsData], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
	}
	out := &types.ApiResponse[GetTransactionsData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/transactions", basePath, url.PathEscape(cardID)), q, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetSensitive returns sensitive card data (PAN, expiry, CVC).
func (s *Service) GetSensitive(ctx context.Context, cardID string) (*types.ApiResponse[GetSensitiveData], error) {
	out := &types.ApiResponse[GetSensitiveData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/sensetive", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Freeze freezes a prepaid card.
func (s *Service) Freeze(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/freeze", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Unfreeze unfreezes a prepaid card.
func (s *Service) Unfreeze(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/unfreeze", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCard deletes a prepaid card.
func (s *Service) DeleteCard(ctx context.Context, cardID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	if err := s.c.Delete(ctx, fmt.Sprintf("%s/%s/delete", basePath, url.PathEscape(cardID)), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateEmail updates the email attached to a card.
func (s *Service) UpdateEmail(ctx context.Context, cardID, email string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"email": email}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/email", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatePhoneNumber updates the phone number on the card.
func (s *Service) UpdatePhoneNumber(ctx context.Context, cardID, phone string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"phone": phone}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/phone", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// SetPin sets the card PIN.
func (s *Service) SetPin(ctx context.Context, cardID, pin string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"pin": pin}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/pin", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// TopUp tops up the card balance. idempotencyKey is optional.
func (s *Service) TopUp(ctx context.Context, cardID string, amount float64, idempotencyKey string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]float64{"amount": amount}
	var opts *httpclient.RequestOptions
	if idempotencyKey != "" {
		opts = &httpclient.RequestOptions{IdempotencyKey: idempotencyKey}
	}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/topup", basePath, url.PathEscape(cardID)), body, opts, out); err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawFunds withdraws funds from a card. idempotencyKey is optional.
func (s *Service) WithdrawFunds(ctx context.Context, cardID string, amount float64, idempotencyKey string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]float64{"amount": amount}
	var opts *httpclient.RequestOptions
	if idempotencyKey != "" {
		opts = &httpclient.RequestOptions{IdempotencyKey: idempotencyKey}
	}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/withdraw", basePath, url.PathEscape(cardID)), body, opts, out); err != nil {
		return nil, err
	}
	return out, nil
}

// ApproveTransaction approves a pending transaction.
func (s *Service) ApproveTransaction(ctx context.Context, cardID, actionID string) (*types.MessageResponse, error) {
	out := &types.MessageResponse{}
	body := map[string]string{"actionId": actionID}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/transactions/approve", basePath, url.PathEscape(cardID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateTopUpAddress generates a crypto top-up address.
func (s *Service) GenerateTopUpAddress(ctx context.Context, cardID, currency, network string) (*types.ApiResponse[GenerateTopUpAddressData], error) {
	q := url.Values{}
	q.Set("currency", currency)
	q.Set("network", network)
	out := &types.ApiResponse[GenerateTopUpAddressData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/address", basePath, url.PathEscape(cardID)), q, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetSpendingLimits returns the spending limits on the card.
func (s *Service) GetSpendingLimits(ctx context.Context, cardID string) (*types.ApiResponse[GetSpendingLimitsData], error) {
	out := &types.ApiResponse[GetSpendingLimitsData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/%s/limits", basePath, url.PathEscape(cardID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// SetSpendingLimits sets a spending limit.
func (s *Service) SetSpendingLimits(ctx context.Context, cardID string, req SetSpendingLimitsRequest) (*types.ApiResponse[SetSpendingLimitsData], error) {
	out := &types.ApiResponse[SetSpendingLimitsData]{}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/%s/limits", basePath, url.PathEscape(cardID)), req, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
