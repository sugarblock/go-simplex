package v2

type PaymentRequest struct {
	AccountDetail     AccountDetail `json:"account_details" validate:"required"`
	TransactionDetail `json:"transaction_details" validate:"required"`
}

type PaymentResponse struct {
	IsKycUpdateRequired *bool `json:"is_kyc_update_required"`
}

type AccountDetail struct {
	AppProviderId  string      `json:"app_provider_id"`
	AppVersionId   string      `json:"app_version_id" validate:"required"`
	AppEndUserId   string      `json:"app_end_user_id" validate:"required"`
	AppInstallDate string      `json:"app_install_date"`
	Email          string      `json:"email" validate:"email"`
	Phone          string      `json:"phone" validate:"e164"`
	SignupLogin    SignupLogin `json:"signup_login" validate:"required"`
}

type SignupLogin struct {
	Location           string `json:"location"`
	Uaid               string `json:"uaid"`
	AcceptLanguage     string `json:"accept_language"`
	HttpAcceptLanguage string `json:"http_accept_language"`
	UserAgent          string `json:"user_agent"`
	CookieSessionId    string `json:"cookie_session_id"`
	Timestamp          string `json:"timestamp"`
	Ip                 string `json:"ip" validate:"required,ip4_addr"`
}

type TransactionDetail struct {
	PaymentDetails PaymentDetails `json:"payment_details" validate:"required"`
}
type PaymentDetails struct {
	QuoteId            string            `json:"quote_id" validate:"required"`
	PaymentId          string            `json:"payment_id" validate:"required"`
	OrderId            string            `json:"order_id" validate:"required"`
	DestinationWallet  DestinationWallet `json:"destination_wallet" validate:"required"`
	OriginalHttpRefUrl string            `json:"original_http_ref_url" validate:"required"`
}

type DestinationWallet struct {
	Currency string `json:"currency" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Tag      string `json:"tag"`
}
