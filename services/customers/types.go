package customers

type CustomerAddress struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type IndividualData struct {
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	BirthDate   string          `json:"birthDate"`
	Nationality string          `json:"nationality"`
	NationalID  string          `json:"nationalId"`
	Address     CustomerAddress `json:"address"`
	Email       string          `json:"email"`
	Phone       string          `json:"phone"`
}

type FeatureAccess struct {
	ID               string   `json:"id"`
	Type             string   `json:"type"`
	CardType         string   `json:"cardType,omitempty"`
	Status           string   `json:"status"`
	RejectionReasons []string `json:"rejectionReasons"`
}

type Sumsub struct {
	URL string `json:"url"`
}

type Customer struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Status         string         `json:"status"`
	IndividualData IndividualData `json:"individualData"`
	FeaturesAccess []FeatureAccess `json:"featuresAccess"`
	Sumsub         Sumsub         `json:"sumsub"`
}

type CreateIndividualCustomerRequest struct {
	Type        string          `json:"type"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	BirthDate   string          `json:"birthDate"`
	Nationality string          `json:"nationality"`
	NationalID  string          `json:"nationalId,omitempty"`
	Address     CustomerAddress `json:"address"`
	Email       string          `json:"email"`
	Phone       string          `json:"phone"`
}

type CreateIndividualCustomerData struct {
	ID               string         `json:"id"`
	Type             string         `json:"type"`
	Status           string         `json:"status"`
	IndividualData   IndividualData `json:"individualData"`
	RejectionReasons []string       `json:"rejectionReasons"`
	CanResubmit      bool           `json:"canResubmit"`
}

type CreateCustomerViaSumsubRequest struct {
	Type             string `json:"type"`
	SumsubShareToken string `json:"sumsubShareToken"`
}

type CreateCustomerViaSumsubData struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Sumsub Sumsub `json:"sumsub"`
}

type RequestFeatureAccessRequest struct {
	AccessType string `json:"accessType"`
	CardType   string `json:"cardType,omitempty"`
}

type RequestFeatureAccessData struct {
	AccessType string `json:"accessType"`
	CustomerID string `json:"customerId"`
	Status     string `json:"status"`
}
