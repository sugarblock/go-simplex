package v2

type EventRequest struct {
	EventId string `json:"event_id"`
}

type EventDeleteResponse struct {
	Status *string `json:"status"`
}

type EventResponse struct {
	Events *[]Event `json:"events"`
}

type Event struct {
	EventId   *string  `json:"event_id"`
	Name      *string  `json:"name"`
	Payment   *Payment `json:"payment"`
	Timestamp *string  `json:"timestamp"`
}

type Payment struct {
	Id                *string       `json:"id"`
	Status            *string       `json:"status"`
	CreatedAt         *string       `json:"created_at"`
	UpdatedAt         *string       `json:"updated_at"`
	FiatTotalAmount   *DigitalMoney `json:"fiat_total_amount"`
	CryptoTotalAmount *DigitalMoney `json:"crypto_total_amount"`
	PartnerEndUserId  *string       `json:"partner_end_user_id"`
}
