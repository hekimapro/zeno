package models

// PaymentOptionsType defines payment details.
type PaymentOptionsType struct {
	CustomerName        string  `json:"customer_name" validate:"required"`
	CustomerEmail       string  `json:"customer_email" validate:"required,email"`
	CustomerPhoneNumber string  `json:"customer_phone_number" validate:"required,len=10|len=12"`
	AmountToCharge      float64 `json:"amount_to_charge" validate:"required,gt=0"`
	CallbackURL         string  `json:"callback_url" validate:"required,url"`
}

// PaymentResponseType defines the response for a payment request.
type PaymentResponseType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	OrderID string `json:"order_id"`
}
