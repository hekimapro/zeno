package models

// PaymentStatusResponseType defines the response for checking payment status.
type PaymentStatusResponseType struct {
	Status        string `json:"status"`
	OrderID       string `json:"order_id"`
	Message       string `json:"message"`
	PaymentStatus string `json:"payment_status"`
}
