package utils

type MccDescription struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ValidateIbanData struct {
	Valid bool `json:"valid"`
}

type BankBySwiftData struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Address string `json:"address"`
	Swift   string `json:"swift"`
}
