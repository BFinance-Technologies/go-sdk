//go:build integration
// +build integration

// Package integration runs against the real BFinance API when a valid
// BFINANCE_API_TOKEN is set. Run with:
//
//	BFINANCE_API_TOKEN=... go test -tags=integration -v ./integration/...
package integration

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	bfinance "github.com/BFinance-Technologies/go-sdk"
	"github.com/BFinance-Technologies/go-sdk/httpclient"
)

func newClient(t *testing.T) *bfinance.Client {
	t.Helper()
	token := os.Getenv("BFINANCE_API_TOKEN")
	if token == "" {
		t.Skip("BFINANCE_API_TOKEN not set")
	}
	return bfinance.New(bfinance.Config{
		APIToken:  token,
		TimeoutMs: 15000,
	})
}

func withCtx(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	return context.WithTimeout(context.Background(), 20*time.Second)
}

func dumpErr(t *testing.T, err error) {
	t.Helper()
	var httpErr *httpclient.HttpError
	if errors.As(err, &httpErr) {
		t.Logf("HttpError: status=%d message=%q details=%v", httpErr.StatusCode, httpErr.Message, httpErr.Details)
		return
	}
	var netErr *httpclient.NetworkError
	if errors.As(err, &netErr) {
		t.Logf("NetworkError: %s", netErr.Message)
		return
	}
	t.Logf("err: %T %v", err, err)
}

// skipIfForbidden marks the test as skipped when the token lacks access to the
// endpoint. We still want to exercise the SDK's request/response plumbing, but
// 403 means the endpoint is gated server-side, not that the SDK is broken.
func skipIfForbidden(t *testing.T, err error) bool {
	t.Helper()
	var httpErr *httpclient.HttpError
	if errors.As(err, &httpErr) && httpErr.StatusCode == 403 {
		t.Skipf("skipped: token lacks access (403 %s)", httpErr.Message)
		return true
	}
	return false
}

func TestBalance_GetUserBalance(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Balance.GetUserBalance(ctx)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetUserBalance: %v", err)
	}
	t.Logf("balance: status=%s eur=%v usd=%v", resp.Status, resp.Data.EUR, resp.Data.USD)
}

func TestPrepaidCards_GetList(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	page, limit := 1, 5
	resp, err := c.PrepaidCards.GetList(ctx, nil)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetList(nil): %v", err)
	}
	t.Logf("prepaid (nil params): status=%s count=%d", resp.Status, len(resp.Data.Cards))

	typed, err := c.PrepaidCards.GetList(ctx, paginationParams(page, limit))
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetList(typed): %v", err)
	}
	t.Logf("prepaid (page=%d limit=%d): count=%d page=%d limit=%d",
		page, limit, len(typed.Data.Cards), typed.Data.Page, typed.Data.Limit)
}

func TestEsim_Countries(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Esim.GetCountries(ctx)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetCountries: %v", err)
	}
	t.Logf("esim countries: status=%s count=%d", resp.Status, len(resp.Data.Countries))
	if len(resp.Data.Countries) > 0 {
		c0 := resp.Data.Countries[0]
		t.Logf("first country: iso=%s name=%q", c0.ISOName, c0.Name)
	}
}

func TestEsim_Regions(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Esim.GetRegions(ctx)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetRegions: %v", err)
	}
	t.Logf("esim regions: status=%s count=%d", resp.Status, len(resp.Data.Regions))
}

func TestEsim_GlobalPackages(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Esim.GetGlobalPackages(ctx, nil)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetGlobalPackages: %v", err)
	}
	t.Logf("esim global packages: status=%s count=%d page=%d totalPages=%d",
		resp.Status, len(resp.Data.Packages), resp.Data.Page, resp.Data.TotalPages)
}

func TestUtils_ValidateIban(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	// Known-valid example IBAN (documentation sample from ECBS).
	resp, err := c.Utils.ValidateIban(ctx, "DE89370400440532013000")
	if err != nil {
		if skipIfForbidden(t, err) {
			return
		}
		dumpErr(t, err)
		t.Fatalf("ValidateIban: %v", err)
	}
	t.Logf("validateIban valid-sample: status=%s valid=%v", resp.Status, resp.Data.Valid)
	if !resp.Data.Valid {
		t.Errorf("expected DE89370400440532013000 to be valid, got false")
	}

	resp2, err := c.Utils.ValidateIban(ctx, "not-an-iban")
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("ValidateIban(invalid): %v", err)
	}
	t.Logf("validateIban invalid-sample: status=%s valid=%v", resp2.Status, resp2.Data.Valid)
	if resp2.Data.Valid {
		t.Errorf("expected 'not-an-iban' to be invalid, got true")
	}
}

func TestUtils_GetBankBySwift(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Utils.GetBankBySwift(ctx, "DEUTDEFF")
	if err != nil {
		if skipIfForbidden(t, err) {
			return
		}
		dumpErr(t, err)
		t.Fatalf("GetBankBySwift: %v", err)
	}
	t.Logf("bank-by-swift DEUTDEFF: status=%s name=%q country=%q", resp.Status, resp.Data.Name, resp.Data.Country)
}

func TestUtils_GetMccDescription(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Utils.GetMccDescription(ctx, "5411")
	if err != nil {
		if skipIfForbidden(t, err) {
			return
		}
		dumpErr(t, err)
		t.Fatalf("GetMccDescription: %v", err)
	}
	t.Logf("mcc 5411: status=%s code=%s description=%q", resp.Status, resp.Data.Code, resp.Data.Description)
}

func TestPrepaidCards_Deep(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	list, err := c.PrepaidCards.GetList(ctx, paginationParams(1, 1))
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetList: %v", err)
	}
	if len(list.Data.Cards) == 0 {
		t.Skip("no prepaid cards on this account")
	}
	card := list.Data.Cards[0]
	t.Logf("using card: id=%s masked=%s currency=%s status=%s",
		card.ID, card.MaskedCardNumber, card.Currency, card.Status)

	// GetDetails
	details, err := c.PrepaidCards.GetDetails(ctx, card.ID)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetDetails: %v", err)
	}
	t.Logf("details: id=%s status=%s balance=%v %s",
		details.Data.ID, details.Data.Status,
		details.Data.Balance.Value, details.Data.Balance.Currency)

	// GetTransactions
	txs, err := c.PrepaidCards.GetTransactions(ctx, card.ID, paginationParams(1, 3))
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetTransactions: %v", err)
	}
	t.Logf("transactions: count=%d", len(txs.Data.Transactions))
	for i, tx := range txs.Data.Transactions {
		t.Logf("  tx[%d]: id=%s amount=%v %s status=%s type=%s merchant=%q",
			i, tx.ID, tx.Amount, tx.Currency, tx.Status, tx.Type, tx.Merchant)
	}

	// GetSpendingLimits — gated by card state (frozen cards 400).
	limits, err := c.PrepaidCards.GetSpendingLimits(ctx, card.ID)
	if err != nil {
		var httpErr *httpclient.HttpError
		if errors.As(err, &httpErr) && (httpErr.StatusCode == 400 || httpErr.StatusCode == 403) {
			t.Logf("GetSpendingLimits skipped for this card state: %d %q", httpErr.StatusCode, httpErr.Message)
		} else {
			dumpErr(t, err)
			t.Fatalf("GetSpendingLimits: %v", err)
		}
	} else {
		t.Logf("limits: ecom.daily=%v/%v pos.daily=%v/%v",
			limits.Data.Ecommerce.Daily.Used, limits.Data.Ecommerce.Daily.Limit,
			limits.Data.POS.Daily.Used, limits.Data.POS.Daily.Limit)
	}
}

func TestEsim_CountryAndPackageDetails(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	// Pick a country (first one).
	countries, err := c.Esim.GetCountries(ctx)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetCountries: %v", err)
	}
	if len(countries.Data.Countries) == 0 {
		t.Skip("no countries available")
	}
	iso := countries.Data.Countries[0].ISOName

	pkgs, err := c.Esim.GetCountryPackages(ctx, iso, nil)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetCountryPackages(%s): %v", iso, err)
	}
	t.Logf("country %s: %d packages", iso, len(pkgs.Data.Packages))
	if len(pkgs.Data.Packages) == 0 {
		return
	}

	p0 := pkgs.Data.Packages[0]
	t.Logf("pkg[0]: id=%s name=%q price=%v data=%v%s validity=%v%s unlimited=%v",
		p0.ID, p0.Name, p0.Price,
		p0.Data.Quantity, p0.Data.Unit,
		p0.Validity.Quantity, p0.Validity.Unit,
		p0.Unlimited)

	details, err := c.Esim.GetPackageDetails(ctx, p0.ID)
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetPackageDetails: %v", err)
	}
	t.Logf("pkg details: id=%s price=%v coverage=%d",
		details.Data.ID, details.Data.Price, len(details.Data.Coverage))
}

func TestFinance_GetCryptoDepositAddress(t *testing.T) {
	c := newClient(t)
	ctx, cancel := withCtx(t)
	defer cancel()

	resp, err := c.Finance.GetCryptoDepositAddress(ctx, "tron", "usdt")
	if err != nil {
		dumpErr(t, err)
		t.Fatalf("GetCryptoDepositAddress: %v", err)
	}
	t.Logf("deposit tron/usdt: status=%s address=%s", resp.Status, resp.Data.Address)
}
