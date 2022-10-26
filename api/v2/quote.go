package v2

type QuoteRequest struct {
	EndUserId         string   `json:"end_user_id"`
	DigitalCurrency   string   `json:"digital_currency"`
	FiatCurrency      string   `json:"fiat_currency"`
	RequestedCurrency string   `json:"requested_currency"`
	RequestedAmount   float32  `json:"requested_amount"`
	WalletId          string   `json:"wallet_id"`
	ClientIp          string   `json:"client_ip"`
	PaymentMethods    []string `json:"payment_methods,omitempty"`
}

type QuoteResponse struct {
	UserId                     *string       `json:"user_id"`
	QuoteId                    *string       `json:"quote_id"`
	WalletId                   *string       `json:"wallet_id"`
	DigitalMoney               *DigitalMoney `json:"digital_money"`
	FiatMoney                  *FiatMoney    `json:"fiat_money"`
	ValidUntil                 *string       `json:"valid_until"`
	SupportedDigitalCurrencies *[]string     `json:"supported_digital_currencies"`
}

type DigitalMoney struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type FiatMoney struct {
	Currency    string  `json:"currency"`
	BaseAmount  float32 `json:"base_amount"`
	TotalAmount float32 `json:"total_amount"`
}
