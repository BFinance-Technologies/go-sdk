package esim

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/BFinance-Technologies/go-sdk/httpclient"
	"github.com/BFinance-Technologies/go-sdk/types"
)

const basePath = "/esim"

type Service struct {
	c *httpclient.Client
}

func New(c *httpclient.Client) *Service {
	return &Service{c: c}
}

func (p *GetPackagesParams) toQuery() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Unlimited != "" {
		q.Set("inlimited", p.Unlimited)
	}
	if p.SortBy != "" {
		q.Set("sortBy", p.SortBy)
	}
	if p.SortType != "" {
		q.Set("sortType", p.SortType)
	}
	if p.Page != nil {
		q.Set("page", strconv.Itoa(*p.Page))
	}
	return q
}

// GetCountries returns the list of countries with eSIM coverage.
func (s *Service) GetCountries(ctx context.Context) (*types.ApiResponse[GetCountriesData], error) {
	out := &types.ApiResponse[GetCountriesData]{}
	if err := s.c.Get(ctx, basePath+"/countries", nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetRegions returns the list of eSIM regions.
func (s *Service) GetRegions(ctx context.Context) (*types.ApiResponse[GetRegionsData], error) {
	out := &types.ApiResponse[GetRegionsData]{}
	if err := s.c.Get(ctx, basePath+"/regions", nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetCountryPackages returns packages available for a given country.
func (s *Service) GetCountryPackages(ctx context.Context, countryCode string, params *GetPackagesParams) (*types.ApiResponse[GetPackagesData], error) {
	out := &types.ApiResponse[GetPackagesData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/packages/country/%s", basePath, url.PathEscape(countryCode)), params.toQuery(), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetGlobalPackages returns global eSIM packages.
func (s *Service) GetGlobalPackages(ctx context.Context, params *GetPackagesParams) (*types.ApiResponse[GetPackagesData], error) {
	out := &types.ApiResponse[GetPackagesData]{}
	if err := s.c.Get(ctx, basePath+"/packages/global", params.toQuery(), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetRegionalPackages returns packages for a specific region.
func (s *Service) GetRegionalPackages(ctx context.Context, regionID string, params *GetPackagesParams) (*types.ApiResponse[GetPackagesData], error) {
	out := &types.ApiResponse[GetPackagesData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/packages/regions/%s", basePath, url.PathEscape(regionID)), params.toQuery(), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetPackageDetails returns details for a specific package.
func (s *Service) GetPackageDetails(ctx context.Context, packageID string) (*types.ApiResponse[PackageDetails], error) {
	out := &types.ApiResponse[PackageDetails]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/packages/%s", basePath, url.PathEscape(packageID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// PurchasePackage purchases a package.
func (s *Service) PurchasePackage(ctx context.Context, packageID string, req *PurchasePackageRequest) (*types.ApiResponse[PurchasePackageData], error) {
	out := &types.ApiResponse[PurchasePackageData]{}
	var body any
	if req != nil {
		body = req
	}
	if err := s.c.Post(ctx, fmt.Sprintf("%s/packages/%s", basePath, url.PathEscape(packageID)), body, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetDetails returns details for a purchased eSIM.
func (s *Service) GetDetails(ctx context.Context, esimID string) (*types.ApiResponse[GetEsimDetailsData], error) {
	out := &types.ApiResponse[GetEsimDetailsData]{}
	if err := s.c.Get(ctx, fmt.Sprintf("%s/details/%s", basePath, url.PathEscape(esimID)), nil, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
