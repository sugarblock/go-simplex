package v2

type QuoteRequest struct {
	EndUserId         string   `json:"end_user_id" validate:"required"`
	DigitalCurrency   string   `json:"digital_currency" validate:"required"`
	FiatCurrency      string   `json:"fiat_currency" validate:"required"`
	RequestedCurrency string   `json:"requested_currency" validate:"required"`
	RequestedAmount   float32  `json:"requested_amount" validate:"required"`
	WalletId          string   `json:"wallet_id" validate:"required"`
	ClientIp          string   `json:"client_ip" validate:"required"`
	PaymentMethods    []string `json:"payment_methods,omitempty"`
}

type QuoteResponse struct {
	HttpStatus                 int           `json:"httpStatus"`
	Message                    string        `json:"message"`
	UserId                     string        `json:"user_id"`
	QuoteId                    string        `json:"quote_id"`
	WalletId                   string        `json:"wallet_id"`
	DigitalMoney               *DigitalMoney `json:"digital_money"`
	FiatMoney                  *FiatMoney    `json:"fiat_money"`
	ValidUntil                 string        `json:"valid_until"`
	SupportedDigitalCurrencies []string      `json:"supported_digital_currencies"`
}

type DigitalMoney struct {
	Currency string  `json:"currency" validate:"required,alpha"`
	Amount   float32 `json:"amount" validate:"required,numeric"`
}

type FiatMoney struct {
	Currency    string  `json:"currency" validate:"required,alpha"`
	BaseAmount  float32 `json:"base_amount" validate:"required,numeric"`
	TotalAmount float32 `json:"total_amount" validate:"required,numeric"`
}
