package zeno

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hekimapro/zeno/utils"
)

// PaymentStatusResponseType defines the response for checking payment status.
type PaymentStatusResponseType struct {
	Status        string `json:"status"`
	OrderID       string `json:"order_id"`
	Message       string `json:"message"`
	PaymentStatus string `json:"payment_status"`
}

// PaymentOptionsType defines payment details.
type PaymentOptionsType struct {
	CustomerName        string  `json:"customer_name"`
	CustomerEmail       string  `json:"customer_email"`
	CustomerPhoneNumber string  `json:"customer_phone_number"`
	AmountToCharge      float64 `json:"amount_to_charge"`
	CallbackURL         string  `json:"callback_url"`
}

// PaymentResponseType defines the response for a payment request.
type PaymentResponseType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	OrderID string `json:"order_id"`
}

// ZenoPay handles API interactions.
type ZenoPay struct {
	APIKey    string
	SecretKey string
	AccountID string
	BaseURL   string
}

// NewZenoPay initializes a new ZenoPay instance.
func NewZenoPay(apiKey, secretKey, accountID string) *ZenoPay {
	return &ZenoPay{
		APIKey:    apiKey,
		SecretKey: secretKey,
		AccountID: accountID,
		BaseURL:   "https://api.zeno.africa",
	}
}

// postRequest sends a POST request to the ZenoPay API.
func (z *ZenoPay) postRequest(route string, data url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", z.BaseURL, route), strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err.Error())
		return nil, utils.CreateError("failed to create request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, utils.CreateError("failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, utils.CreateError("failed to read response body")
	}

	return body, nil
}

// Pay initiates a payment and returns the order ID or an error.
func (z *ZenoPay) Pay(options PaymentOptionsType) (orderID string, err error) {

	data := url.Values{
		"create_order": {"1"},
		"api_key":      {z.APIKey},
		"account_id":   {z.AccountID},
		"secret_key":   {z.SecretKey},
		"amount":       {fmt.Sprintf("%f", options.AmountToCharge)},
		"buyer_name":   {options.CustomerName},
		"webhook_url":  {options.CallbackURL},
		"buyer_email":  {options.CustomerEmail},
		"buyer_phone":  {utils.FormatPhoneNumber(options.CustomerPhoneNumber)},
	}

	body, err := z.postRequest("", data)
	if err != nil {
		return "", err
	}

	var response PaymentResponseType
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err.Error())
		return "", utils.CreateError("failed to parse payment response")
	}

	if response.OrderID == "" {
		return "", utils.CreateError(response.Message)
	}

	return response.OrderID, nil
}

// CheckPaymentStatus checks payment status and returns the status or an error.
func (z *ZenoPay) CheckPaymentStatus(orderID string) (orderStatus string, err error) {

	data := url.Values{
		"check_status": {"1"},
		"order_id":     {orderID},
	}

	body, err := z.postRequest("order-status", data)
	if err != nil {
		return "", err
	}

	var response PaymentStatusResponseType
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err.Error())
		return "", utils.CreateError("failed to parse status response")
	}

	return response.PaymentStatus, nil
}
