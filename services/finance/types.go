package finance

type GetCryptoDepositAddressData struct {
	Address  string `json:"address"`
	Network  string `json:"network"`
	Currency string `json:"currency"`
}
