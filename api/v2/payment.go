package v2

type PaymentRequest struct {
	AccountDetail     AccountDetail `json:"account_details" validate:"required"`
	TransactionDetail `json:"transaction_details" validate:"required"`
}

type PaymentResponse struct {
	IsKycUpdateRequired *bool `json:"is_kyc_update_required"`
}

type AccountDetail struct {
	AppProviderId  string      `json:"app_provider_id" validate:"required"`
	AppVersionId   string      `json:"app_version_id" validate:"required"`
	AppEndUserId   string      `json:"app_end_user_id" validate:"required"`
	AppInstallDate string      `json:"app_install_date" validate:"required"`
	Email          string      `json:"email" validate:"required,email"`
	Phone          string      `json:"phone" validate:"required,e164	"`
	SignupLogin    SignupLogin `json:"signup_login" validate:"required"`
}

type SignupLogin struct {
	Location           string `json:"location" validate:"required"`
	Uaid               string `json:"uaid" validate:"required"`
	AcceptLanguage     string `json:"accept_language" validate:"required"`
	HttpAcceptLanguage string `json:"http_accept_language" validate:"required"`
	UserAgent          string `json:"user_agent" validate:"required"`
	CookieSessionId    string `json:"cookie_session_id" validate:"required"`
	Timestamp          string `json:"timestamp" validate:"required,datetime"`
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
	Tag      string `json:"tag" validate:"required"`
}
