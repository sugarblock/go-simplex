package v2

type PaymentRequest struct {
	AccountDetail     AccountDetail `json:"account_details"`
	TransactionDetail `json:"transaction_details"`
}

type PaymentResponse struct {
	IsKycUpdateRequired *bool `json:"is_kyc_update_required"`
}

type AccountDetail struct {
	AppProviderId  string      `json:"app_provider_id"`
	AppVersionId   string      `json:"app_version_id"`
	AppEndUserId   string      `json:"app_end_user_id"`
	AppInstallDate string      `json:"app_install_date"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	SignupLogin    SignupLogin `json:"signup_login"`
}

type SignupLogin struct {
	Location           string `json:"location"`
	Uaid               string `json:"uaid"`
	AcceptLanguage     string `json:"accept_language"`
	HttpAcceptLanguage string `json:"http_accept_language"`
	UserAgent          string `json:"user_agent"`
	CookieSessionId    string `json:"cookie_session_id"`
	Timestamp          string `json:"timestamp"`
	Ip                 string `json:"ip"`
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
