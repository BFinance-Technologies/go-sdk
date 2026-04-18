package virtualaccounts

type DepositInstructions struct {
	Currency          string   `json:"currency"`
	IBAN              string   `json:"iban"`
	BIC               string   `json:"bic"`
	AccountHolderName string   `json:"accountHolderName"`
	BankName          string   `json:"bankName"`
	BankAddress       string   `json:"bankAddress"`
	BankRoutingNumber string   `json:"bankRoutingNumber"`
	BankAccountNumber string   `json:"bankAccountNumber"`
	BeneficiaryName   string   `json:"beneficiaryName"`
	BeneficiaryAddr   string   `json:"beneficiaryAddress"`
	PaymentRails      []string `json:"paymentRails"`
	CLABE             string   `json:"clabe"`
	PIXCode           string   `json:"pixCode"`
}

type VirtualAccount struct {
	ID                  string              `json:"id"`
	Status              string              `json:"status"`
	DepositInstructions DepositInstructions `json:"depositInstructions"`
}

type EligibilityData struct {
	Eligible    bool     `json:"eligibile"`
	EligibleFor []string `json:"eligibileFor"`
}

type Destination struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Address  string `json:"address"`
}

type CreateVirtualAccountRequest struct {
	Type        string      `json:"type"`
	Destination Destination `json:"destination"`
}

type CreateVirtualAccountData struct {
	Status              string              `json:"status"`
	DepositInstructions DepositInstructions `json:"depositInstructions"`
}
