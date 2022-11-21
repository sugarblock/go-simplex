package v2

type PaymentRequest struct {
	AccountDetail     *AccountDetail     `json:"account_details,omitempty"`
	TransactionDetail *TransactionDetail `json:"transaction_details,omitempty"`
}

type PaymentResponse struct {
	IsKycUpdateRequired *bool `json:"is_kyc_update_required,omitempty"`
}

type AccountDetail struct {
	AppProviderId  *string      `json:"app_provider_id,omitempty"`
	AppVersionId   *string      `json:"app_version_id,omitempty"`
	AppEndUserId   *string      `json:"app_end_user_id,omitempty"`
	AppInstallDate *string      `json:"app_install_date,omitempty"`
	Email          *string      `json:"email,omitempty"`
	Phone          *string      `json:"phone,omitempty"`
	SignupLogin    *SignupLogin `json:"signup_login,omitempty"`
}

type SignupLogin struct {
	Location           *string `json:"location,omitempty"`
	Uaid               *string `json:"uaid,omitempty"`
	AcceptLanguage     *string `json:"accept_language,omitempty"`
	HttpAcceptLanguage *string `json:"http_accept_language,omitempty"`
	UserAgent          *string `json:"user_agent,omitempty"`
	CookieSessionId    *string `json:"cookie_session_id,omitempty"`
	Timestamp          *string `json:"timestamp,omitempty"`
	Ip                 *string `json:"ip,omitempty"`
}

type TransactionDetail struct {
	PaymentDetails PaymentDetails `json:"payment_details"`
}
type PaymentDetails struct {
	QuoteId            string            `json:"quote_id"`
	PaymentId          string            `json:"payment_id"`
	OrderId            string            `json:"order_id"`
	DestinationWallet  DestinationWallet `json:"destination_wallet"`
	OriginalHttpRefUrl string            `json:"original_http_ref_url"`
}

type DestinationWallet struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
	Tag      string `json:"tag"`
}
