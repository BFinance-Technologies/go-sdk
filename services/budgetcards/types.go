package budgetcards

type BudgetCardSensitive struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
	CVC    string `json:"cvc"`
}

type BudgetCard struct {
	ID               string              `json:"id"`
	MaskedCardNumber string              `json:"maskedCardNumber"`
	Currency         string              `json:"currency"`
	Status           string              `json:"status"`
	ExternalCardID   string              `json:"externalCardId"`
	Sensitive        BudgetCardSensitive `json:"sensitive"`
}

type IssueBudgetCardRequest struct {
	TypeID    string `json:"typeId"`
	BudgetID  string `json:"budgetId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Label     string `json:"label,omitempty"`
}

type SetVelocityLimitsRequest struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type GetSensitiveData struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
	CVC    string `json:"cvc"`
}
