package v2

type EventRequest struct {
	EventId *string `json:"event_id,omitempty"`
}

type EventDeleteResponse struct {
	Status *string `json:"status,omitempty"`
}

type EventResponse struct {
	Events *[]Event `json:"events,omitempty"`
}

type Event struct {
	EventId   *string  `json:"event_id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Payment   *Payment `json:"payment,omitempty"`
	Timestamp *string  `json:"timestamp,omitempty"`
}

type Payment struct {
	Id                *string       `json:"id,omitempty"`
	Status            *string       `json:"status,omitempty"`
	CreatedAt         *string       `json:"created_at,omitempty"`
	UpdatedAt         *string       `json:"updated_at,omitempty"`
	FiatTotalAmount   *DigitalMoney `json:"fiat_total_amount,omitempty"`
	CryptoTotalAmount *DigitalMoney `json:"crypto_total_amount,omitempty"`
	PartnerEndUserId  *string       `json:"partner_end_user_id,omitempty"`
}
