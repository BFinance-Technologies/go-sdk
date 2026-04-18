//go:build integration
// +build integration

package integration

import "github.com/BFinance-Technologies/go-sdk/types"

func paginationParams(page, limit int) *types.PaginationParams {
	return &types.PaginationParams{Page: &page, Limit: &limit}
}
