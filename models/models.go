package models

// USSD push request struct
type USSDPushRequest struct {
	Amount               float64     `json:"amount"`
	OrderID              string      `json:"order_id"`
	CustomerName         string      `json:"buyer_name"`
	CustomerPhoneNumber  string      `json:"buyer_phone"`
	CustomerEmailAddress string      `json:"buyer_email"`
	WebhookURL           string      `json:"webhook_url,omitempty"`
	Metadata             interface{} `json:"metadata,omitempty"`
}

// USSD push response struct
type USSDPushResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	OrderID    string `json:"order_id,omitempty"`   // not present on error
	ResultCode string `json:"resultcode,omitempty"` // not present on error
}

// USSD push status check struct
type USSDPushStatusCheckResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	Reference  string `json:"reference"`
	ResultCode string `json:"resultcode"`
	Data       []PushStatusCheckResponseData
}

// status check response data struct
type PushStatusCheckResponseData struct {
	OrderID       string `json:"order_id"`
	CreationDate  string `json:"creation_date"`
	Amount        string `json:"string"`
	PaymentStatus string `json:"payment_status"`
	TransactionID string `json:"transid"`
	Channel       string `json:"channel"`
	Reference     string `json:"reference"`
	MSISDN        string `json:"msisdn"`
}

// USSD Webhook
type USSDWebhookData struct {
	OrderID       string      `json:"order_id"`
	PaymentStatus string      `json:"payment_status"`
	Reference     string      `json:"reference"`
	Metadata      interface{} `json:"metadata"`
}

// send money struct
type SendMoneyRequest struct {
	TransactionID string  `json:"transid"`
	UtilityCode   string  `json:"utilitycode"` // "CASHIN",
	PhoneNumber   string  `json:"utilityref"`
	Amount        float64 `json:"amount"`
	PIN           int     `json:"pin"`
}

// send money response struct
type SendMoneyResponse struct {
	Status          string                       `json:"status"`
	Message         string                       `json:"message"`
	AmountSent      float64                      `json:"amount_sent_to_customer"`
	TotalDeducted   float64                      `json:"total_deducted"`
	NewBalance      float64                      `json:"new_balance"`
	ZenoPayResponse SendMoneyZenoPayResponseData `json:"zenopay_response"`
}

// send money zeno response data struct
type SendMoneyZenoPayResponseData struct {
	Reference     string        `json:"reference"`
	TransactionID string        `json:"transid"`
	ResultCode    string        `json:"resultcode"`
	Result        string        `json:"result"`
	Message       string        `json:"message"`
	Data          []interface{} `json:"data"`
}

// send money zeno error response
type SendMoneyErrorResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Errors  SendMoneyErrors `json:"errors"`
}

// send money errors
type SendMoneyErrors struct {
	UtilityRef []string `json:"utilityref"`
}
