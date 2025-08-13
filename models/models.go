package models

// USSDPushRequest represents the payload for initiating a USSD push transaction.
type USSDPushRequest struct {
	Amount               float64     `json:"amount"`
	OrderID              string      `json:"order_id"`
	CustomerName         string      `json:"buyer_name"`
	CustomerPhoneNumber  string      `json:"buyer_phone"`
	CustomerEmailAddress string      `json:"buyer_email"`
	WebhookURL           string      `json:"webhook_url,omitempty"`
	Metadata             interface{} `json:"metadata,omitempty"`
}

// USSDPushResponse represents the response from a USSD push request.
type USSDPushResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	OrderID    string `json:"order_id,omitempty"`   // Not present on error
	ResultCode string `json:"resultcode,omitempty"` // Not present on error
}

// CheckStatus represents the response structure when checking the status of a USSD transaction.
type CheckStatus struct {
	Result     string                    `json:"result"`
	Message    string                    `json:"message"`
	Reference  string                    `json:"reference"`
	ResultCode string                    `json:"resultcode"`
	Data       []CheckStatusResponseData `json:"data"`
}

// CheckStatusResponseData holds detailed information for each transaction status in a status check response.
type CheckStatusResponseData struct {
	OrderID       string `json:"order_id"`
	CreationDate  string `json:"creation_date"`
	Amount        string `json:"string"`
	PaymentStatus string `json:"payment_status"`
	TransactionID string `json:"transid"`
	Channel       string `json:"channel"`
	Reference     string `json:"reference"`
	MSISDN        string `json:"msisdn"`
}

// USSDWebhookData represents the payload sent to the webhook for a USSD transaction update.
type USSDWebhookData struct {
	OrderID       string      `json:"order_id"`
	PaymentStatus string      `json:"payment_status"`
	Reference     string      `json:"reference"`
	Metadata      interface{} `json:"metadata"`
}

// SendMoneyRequest represents the payload for initiating a cash-in transaction.
type SendMoneyRequest struct {
	TransactionID string  `json:"transid"`
	UtilityCode   string  `json:"utilitycode"` // Example: "CASHIN"
	PhoneNumber   string  `json:"utilityref"`
	Amount        float64 `json:"amount"`
	PIN           int     `json:"pin"`
}

// SendMoneyResponse represents the response from a successful send money transaction.
type SendMoneyResponse struct {
	Status          string                       `json:"status"`
	Message         string                       `json:"message"`
	AmountSent      float64                      `json:"amount_sent_to_customer"`
	TotalDeducted   float64                      `json:"total_deducted"`
	NewBalance      float64                      `json:"new_balance"`
	ZenoPayResponse SendMoneyZenoPayResponseData `json:"zenopay_response"`
}

// SendMoneyZenoPayResponseData holds additional provider-specific response details for a send money transaction.
type SendMoneyZenoPayResponseData struct {
	Reference     string        `json:"reference"`
	TransactionID string        `json:"transid"`
	ResultCode    string        `json:"resultcode"`
	Result        string        `json:"result"`
	Message       string        `json:"message"`
	Data          []interface{} `json:"data"`
}

// SendMoneyErrorResponse represents the error response returned for a send money transaction.
type SendMoneyErrorResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Errors  SendMoneyErrors `json:"errors"`
}

// SendMoneyErrors holds detailed error information for a send money request.
type SendMoneyErrors struct {
	UtilityRef []string `json:"utilityref"`
}

// CheckoutRequest represents the payload for a multi-currency checkout transaction.
type CheckoutRequest struct {
	Amount               float64     `json:"amount"`
	Currency             string      `json:"currency"`
	RedirectURL          string      `json:"redirect_url"`
	CustomerName         string      `json:"buyer_name"`
	CustomerPhoneNumber  string      `json:"buyer_phone"`
	CustomerEmailAddress string      `json:"buyer_email"`
	WebhookURL           string      `json:"webhook_url"`
	Metadata             interface{} `json:"metadata"`
	OrderID              string      `json:"order_id"`
}

// CheckoutResponse represents the response for a multi-currency checkout transaction.
type CheckoutResponse struct {
	PaymentLink          string `json:"payment_link"`
	TransactionReference string `json:"tx_ref"`
	Error                string `json:"error,omitempty"`
}
