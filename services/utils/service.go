package utils

import (
	"context"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/utils"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

// GetMccDescription returns the description for an MCC code. Pass empty string to list.
func (s *Service) GetMccDescription(ctx context.Context, mcc string) (*types.ApiResponse[MccDescription], error) {
	body := map[string]string{}
	if mcc != "" {
		body["mcc"] = mcc
	}
	out := &types.ApiResponse[MccDescription]{}
	if err := s.c.Post(ctx, basePath+"/mcc", body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// ValidateIban validates an IBAN.
func (s *Service) ValidateIban(ctx context.Context, iban string) (*types.ApiResponse[ValidateIbanData], error) {
	body := map[string]string{"iban": iban}
	out := &types.ApiResponse[ValidateIbanData]{}
	if err := s.c.Post(ctx, basePath+"/validateIban", body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetBankBySwift returns bank info for a SWIFT code.
func (s *Service) GetBankBySwift(ctx context.Context, swift string) (*types.ApiResponse[BankBySwiftData], error) {
	body := map[string]string{"swift": swift}
	out := &types.ApiResponse[BankBySwiftData]{}
	if err := s.c.Post(ctx, basePath+"/getBankBySwift", body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
