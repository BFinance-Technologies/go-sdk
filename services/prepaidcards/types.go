package prepaidcards

import (
	"encoding/json"
	"strconv"
	"time"
)

// FlexibleString accepts both JSON strings and JSON numbers and stores
// them as a string. Useful because some BFinance endpoints (e.g. the
// card "sensitive" payload) return the PAN as a number in test data.
type FlexibleString string

func (f *FlexibleString) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*f = FlexibleString(s)
		return nil
	}
	// Raw JSON number — keep its textual form so we don't lose digits
	// to float64 rounding on long PANs.
	if _, err := strconv.ParseFloat(string(b), 64); err != nil {
		return err
	}
	*f = FlexibleString(string(b))
	return nil
}

type PrepaidCardItem struct {
	ID               string `json:"id"`
	MaskedCardNumber string `json:"maskedCardNumber"`
	Currency         string `json:"currency"`
	Status           string `json:"status"`
	ExternalCardID   string `json:"externalCardId"`
}

type GetPrepaidCardsData struct {
	Cards []PrepaidCardItem `json:"cards"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
}

type IssuePrepaidCardRequest struct {
	TypeID         string  `json:"typeId"`
	InitialBalance float64 `json:"initialBalance"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Label          string  `json:"label,omitempty"`
}

type IssuedPrepaidCard struct {
	ID               string  `json:"id"`
	CardNumber       string  `json:"cardNumber"`
	CardExpire       string  `json:"cardExpire"`
	CardCVC          string  `json:"cardCVC"`
	CardBalance      float64 `json:"cardBalance"`
	Currency         string  `json:"currency"`
	MaskedCardNumber string  `json:"maskedCardNumber"`
	Brand            string  `json:"brand"`
	Label            string  `json:"label"`
}

type IssuePrepaidCardData struct {
	Card IssuedPrepaidCard `json:"card"`
}

type ReissuePrepaidCardRequest struct {
	CardID         string  `json:"cardId"`
	InitialBalance float64 `json:"initialBalance"`
}

type ReissuedPrepaidCard struct {
	ID               string  `json:"id"`
	CardExpire       string  `json:"cardExpire"`
	CardBalance      float64 `json:"cardBalance"`
	Currency         string  `json:"currency"`
	MaskedCardNumber string  `json:"maskedCardNumber"`
	Brand            string  `json:"brand"`
	Label            string  `json:"label"`
}

type ReissuePrepaidCardData struct {
	Card ReissuedPrepaidCard `json:"card"`
}

type PrepaidCardBalance struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// PrepaidCardSensitive holds raw card credentials. `Number` is declared as
// a FlexibleString because the API sometimes returns the PAN as a JSON
// number instead of a string.
type PrepaidCardSensitive struct {
	Number FlexibleString `json:"number"`
	Expire string         `json:"expire"`
	CVC    string         `json:"cvc"`
}

type CardHolder struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type PrepaidCardDetails struct {
	ID               string               `json:"id"`
	MaskedCardNumber string               `json:"maskedCardNumber"`
	Sensitive        PrepaidCardSensitive `json:"sensetive"`
	Currency         string               `json:"currency"`
	Status           string               `json:"status"`
	ExternalCardID   string               `json:"externalCardId,omitempty"`
	Label            string               `json:"label,omitempty"`
	Balance          PrepaidCardBalance   `json:"balance"`
	CardHolder       *CardHolder          `json:"cardHolder,omitempty"`
	Email            string               `json:"email,omitempty"`
	Phone            string               `json:"phone,omitempty"`
}

type CardTransaction struct {
	ID            string    `json:"id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	USDAmount     float64   `json:"usdAmount"`
	Merchant      string    `json:"merchant"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	Date          time.Time `json:"date"`
	TransactionID string    `json:"transactionId"`
	DeclineReason *string   `json:"declineReason"`
}

type GetTransactionsData struct {
	Transactions []CardTransaction `json:"transactions"`
}

type GetSensitiveData struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
	CVC    string `json:"cvc"`
}

type GenerateTopUpAddressAmount struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Fee float64 `json:"fee"`
}

type GenerateTopUpAddressData struct {
	Address  string                     `json:"address"`
	Network  string                     `json:"network"`
	Currency string                     `json:"currency"`
	QRImage  string                     `json:"qrImage"`
	Amount   GenerateTopUpAddressAmount `json:"amount"`
}

type SpendingLimitItem struct {
	Limit float64 `json:"limit"`
	Used  float64 `json:"used"`
}

type SpendingLimitItemWithoutUsed struct {
	Limit float64 `json:"limit"`
}

type SpendingLimitsCategory struct {
	Daily       SpendingLimitItem            `json:"daily"`
	Monthly     SpendingLimitItem            `json:"monthly"`
	Transaction SpendingLimitItemWithoutUsed `json:"transaction"`
}

type GetSpendingLimitsData struct {
	Ecommerce SpendingLimitsCategory `json:"ecommerce"`
	POS       SpendingLimitsCategory `json:"pos"`
}

type SetSpendingLimitsRequest struct {
	Type  string  `json:"type"`
	Limit float64 `json:"limit"`
}

type SetSpendingLimitsData struct {
	Type  string  `json:"type"`
	Limit float64 `json:"limit"`
}
