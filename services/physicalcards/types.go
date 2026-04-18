package physicalcards

type ShippingData struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PhoneCode    string `json:"phoneCode"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
}

type OrderPhysicalCardRequest struct {
	CardType     string       `json:"cardType"`
	ShippingData ShippingData `json:"shippingData"`
}

type ActivatePhysicalCardRequest struct {
	CardType string `json:"cardType"`
	Code     string `json:"code"`
}
