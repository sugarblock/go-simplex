package v2

type QuoteRequest struct {
	EndUserId         *string   `json:"end_user_id,omitempty"`
	DigitalCurrency   *string   `json:"digital_currency,omitempty"`
	FiatCurrency      *string   `json:"fiat_currency,omitempty"`
	RequestedCurrency *string   `json:"requested_currency,omitempty"`
	RequestedAmount   *float32  `json:"requested_amount,omitempty"`
	WalletId          *string   `json:"wallet_id,omitempty"`
	ClientIp          *string   `json:"client_ip,omitempty"`
	PaymentMethods    *[]string `json:"payment_methods,omitempty"`
}

type QuoteResponse struct {
	UserId                     *string       `json:"user_id,omitempty"`
	QuoteId                    *string       `json:"quote_id,omitempty"`
	WalletId                   *string       `json:"wallet_id,omitempty"`
	DigitalMoney               *DigitalMoney `json:"digital_money,omitempty"`
	FiatMoney                  *FiatMoney    `json:"fiat_money,omitempty"`
	ValidUntil                 *string       `json:"valid_until,omitempty"`
	SupportedDigitalCurrencies *[]string     `json:"supported_digital_currencies,omitempty"`
	Error                      *string       `json:"error,omitempty"`
}

type DigitalMoney struct {
	Currency *string  `json:"currency,omitempty"`
	Amount   *float32 `json:"amount,omitempty"`
}

type FiatMoney struct {
	Currency    *string  `json:"currency,omitempty"`
	BaseAmount  *float32 `json:"base_amount,omitempty"`
	TotalAmount *float32 `json:"total_amount,omitempty"`
}
