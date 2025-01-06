package zenopay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hekimapro/zenopay-go/models"
	"github.com/hekimapro/zenopay-go/utils"
)

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
func (z *ZenoPay) Pay(options models.PaymentOptionsType) (string, error) {
	validate := validator.New()
	if err := validate.Struct(options); err != nil {
		fmt.Println(err.Error())
		return "", utils.CreateError("invalid payment options")
	}

	data := url.Values{
		"create_order": {"1"},
		"api_key":      {z.APIKey},
		"account_id":   {z.AccountID},
		"secret_key":   {z.SecretKey},
		"amount":       {fmt.Sprintf("%f", options.AmountToCharge)},
		"buyer_name":   {options.CustomerName},
		"webhook_url":  {options.CallbackURL},
		"buyer_email":  {options.CustomerEmail},
		"buyer_phone":  {options.CustomerPhoneNumber},
	}

	body, err := z.postRequest("", data)
	if err != nil {
		return "", err
	}

	var response models.PaymentResponseType
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
func (z *ZenoPay) CheckPaymentStatus(orderID string) (string, error) {
	if orderID == "" {
		return "", utils.CreateError("order ID is required")
	}

	data := url.Values{
		"check_status": {"1"},
		"order_id":     {orderID},
	}

	body, err := z.postRequest("order-status", data)
	if err != nil {
		return "", err
	}

	var response models.PaymentStatusResponseType
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err.Error())
		return "", utils.CreateError("failed to parse status response")
	}

	return response.PaymentStatus, nil
}
