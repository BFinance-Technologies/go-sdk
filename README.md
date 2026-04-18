# BFinance Go SDK

Official Go SDK for the BFinance API.

## Installation

```bash
go get github.com/bfinance/bfinance-sdk-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    bfinance "github.com/bfinance/bfinance-sdk-go"
)

func main() {
    client := bfinance.New(bfinance.Config{
        APIToken:  "your-api-token",
        TimeoutMs: 5000,
    })

    ctx := context.Background()

    // List prepaid cards
    resp, err := client.PrepaidCards.GetList(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("cards: %+v\n", resp.Data.Cards)

    // Top up a card with idempotency
    _, err = client.PrepaidCards.TopUp(ctx, "card-id", 100.50, "unique-idempotency-key")
    if err != nil {
        log.Fatal(err)
    }

    // Get user balance
    bal, err := client.Balance.GetUserBalance(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("EUR: %v  USD: %v\n", bal.Data.EUR, bal.Data.USD)
}
```

## Services

- `client.PrepaidCards` — issue, manage, top-up, and inspect prepaid cards
- `client.BudgetCards` — manage budget cards
- `client.PhysicalCards` — order and activate physical cards
- `client.Customers` — create customers (direct or via Sumsub) and request feature access
- `client.VirtualAccounts` — create and list fiat virtual accounts
- `client.Disputes` — file, track, and cancel transaction disputes
- `client.Esim` — browse eSIM packages, purchase, and retrieve installation data
- `client.Balance` — read the user's EUR/USD balance
- `client.Finance` — generate crypto deposit addresses
- `client.Utils` — MCC, IBAN, and SWIFT utilities

## Errors

HTTP errors are returned as `*httpclient.HttpError` (with `StatusCode`, `Message`, `Details`).
Network-level failures are returned as `*httpclient.NetworkError`.

```go
import "errors"
import "github.com/bfinance/bfinance-sdk-go/httpclient"

_, err := client.PrepaidCards.GetDetails(ctx, "missing")
var httpErr *httpclient.HttpError
if errors.As(err, &httpErr) {
    fmt.Println(httpErr.StatusCode, httpErr.Message)
}
```

## Logging

```go
client := bfinance.New(bfinance.Config{
    APIToken: token,
    Logging: &httpclient.LoggingConfig{
        Level:       httpclient.LogDebug,
        IncludeBody: true,
    },
})
```
