package esim

import "time"

type EsimCountry struct {
	ISOName string `json:"isoName"`
	Name    string `json:"name"`
	FlagURL string `json:"flagUrl,omitempty"`
}

type GetCountriesData struct {
	Countries []EsimCountry `json:"countries"`
}

type EsimRegion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetRegionsData struct {
	Regions []EsimRegion `json:"regions"`
}

// PackageQuota describes a package quota value. Quantity is a float because
// the API may return fractional values (e.g. -0.001 to indicate "unlimited"
// in combination with Unlimited: true).
type PackageQuota struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type PackageQuotaWithoutUnit struct {
	Quantity float64 `json:"quantity"`
}

type EsimPackageItem struct {
	ID        string                  `json:"id"`
	Name      string                  `json:"name"`
	Price     float64                 `json:"price"`
	Data      PackageQuota            `json:"data"`
	Voice     PackageQuota            `json:"voice"`
	SMS       PackageQuotaWithoutUnit `json:"sms"`
	Validity  PackageQuota            `json:"validity"`
	Unlimited bool                    `json:"unlimited"`
}

type GetPackagesData struct {
	Packages   []EsimPackageItem `json:"packages"`
	Page       int               `json:"page"`
	TotalPages int               `json:"totalPages"`
}

type GetPackagesParams struct {
	Unlimited string `url:"inlimited,omitempty"`
	SortBy    string `url:"sortBy,omitempty"`
	SortType  string `url:"sortType,omitempty"`
	Page      *int   `url:"page,omitempty"`
}

type NetworkInfo struct {
	Name            string   `json:"name"`
	MobileStandards []string `json:"mobileStandards"`
}

type PackageCoverage struct {
	Name    string        `json:"name"`
	FlagURL string        `json:"flagUrl,omitempty"`
	Network []NetworkInfo `json:"network"`
}

type PackageDetails struct {
	ID        string                  `json:"id"`
	Price     float64                 `json:"price"`
	Unlimited bool                    `json:"unlimited"`
	Data      PackageQuota            `json:"data"`
	Voice     PackageQuota            `json:"voice"`
	SMS       PackageQuotaWithoutUnit `json:"sms"`
	Validity  PackageQuota            `json:"validity"`
	Coverage  []PackageCoverage       `json:"coverage"`
}

type PurchasePackageRequest struct {
	ExternalID string `json:"externalId,omitempty"`
}

type PurchasePackageData struct {
	ID string `json:"id"`
}

type EsimInstallation struct {
	LPA         string `json:"lpa"`
	SMDPAddress string `json:"smdpAddress"`
	ICCID       string `json:"iccid"`
}

type EsimUsageData struct {
	Initial   PackageQuota `json:"initial"`
	Remaining PackageQuota `json:"remaining"`
}

type EsimUsage struct {
	Data EsimUsageData `json:"data"`
}

type EsimInfo struct {
	Installation EsimInstallation `json:"installation"`
	CreatedAt    time.Time        `json:"createdAt"`
	Status       string           `json:"status"`
	Usage        EsimUsage        `json:"usage"`
}

type GetEsimDetailsData struct {
	Esim    EsimInfo       `json:"esim"`
	Package PackageDetails `json:"package"`
}
