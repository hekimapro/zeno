package zeno

import (
	"encoding/json"
	"fmt"

	"github.com/hekimapro/utils/log"
	"github.com/hekimapro/utils/request"
	"github.com/hekimapro/zeno/models"
	"github.com/hekimapro/zeno/utils"
)

// PushUSSD sends a USSD push request to the configured endpoint and returns the parsed response.
// It retrieves the API URL and headers, sends the HTTP request, and unmarshals the JSON response.
//
// Parameters:
//   - requestBody: *models.USSDPushRequest with the following fields:
//     Amount               (float64)  - Transaction amount
//     OrderID              (string)   - Unique order identifier
//     CustomerName         (string)   - Buyer’s full name
//     CustomerPhoneNumber  (string)   - Buyer’s phone number (will be formatted before sending)
//     CustomerEmailAddress (string)   - Buyer’s email address
//     WebhookURL           (string)   - Optional webhook URL for transaction status updates
//     Metadata             (any)      - Optional custom metadata
//
// Returns:
//   - *models.USSDPushResponse with fields:
//     Status     (string) - Request status
//     Message    (string) - Response message
//     OrderID    (string) - Returned order ID (if successful)
//     ResultCode (string) - Provider-specific result code (if successful)
//   - error if any step fails.
func PushUSSD(requestBody *models.USSDPushRequest) (*models.USSDPushResponse, error) {
	// Get the configured endpoint URL for USSD push requests.
	URL := utils.GetURL("push")

	// Get required HTTP headers for authentication and content type.
	headers, err := utils.GetHeaders()
	if err != nil {
		return nil, err
	}

	fmt.Println(headers)
	// Send the HTTP POST request with payload and headers to the API endpoint.
	raw, err := request.Post(URL, requestBody, headers)
	if err != nil {
		return nil, err
	}

	// Convert the raw JSON response into the USSDPushResponse struct.
	var responseBody models.USSDPushResponse
	if err := json.Unmarshal(raw, &responseBody); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Return the parsed API response.
	return &responseBody, nil
}

// CheckStatus retrieves the status of a given order from the configured endpoint.
// It fetches the API URL and headers, sends a GET request, and parses the JSON response.
//
// Parameters:
//   - orderID: (string) Unique order identifier to check.
//
// Returns:
//   - *models.CheckStatus with fields:
//     Result     (string) - Status result
//     Message    (string) - Response message
//     Reference  (string) - Transaction reference
//     ResultCode (string) - Provider-specific result code
//     Data       ([]CheckStatusResponseData) - List of detailed status records, each containing:
//     OrderID, CreationDate, Amount, PaymentStatus, TransactionID, Channel, Reference, MSISDN
//   - error if any step fails.
func CheckStatus(orderID string) (*models.CheckStatus, error) {
	// Build the full status check URL by appending the order ID as a query parameter.
	URL := fmt.Sprintf("%s?order_id=%s", utils.GetURL("status"), orderID)

	// Get required HTTP headers for authentication and content type.
	headers, err := utils.GetHeaders()
	if err != nil {
		return nil, err
	}

	// Send the HTTP GET request to the API endpoint.
	raw, err := request.Get(URL, headers)
	if err != nil {
		return nil, err
	}

	// Convert the raw JSON response into the CheckStatus struct.
	var responseBody models.CheckStatus
	if err := json.Unmarshal(raw, &responseBody); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Return the parsed API response containing the order status.
	return &responseBody, nil
}

// SendMoney sends a cash-in transaction request to the configured endpoint.
// It retrieves the API URL and headers, sends the HTTP POST request, and parses the JSON response.
//
// Parameters:
//   - requestBody: *models.SendMoneyRequest with the following fields:
//     TransactionID (string)  - Unique transaction ID
//     UtilityCode   (string)  - Provider utility code (e.g., "CASHIN")
//     PhoneNumber   (string)  - Recipient phone number (will be formatted before sending)
//     Amount        (float64) - Amount to send
//     PIN           (int)     - Security PIN
//
// Returns:
//   - *models.SendMoneyResponse with fields:
//     Status        (string)  - Request status
//     Message       (string)  - Response message
//     AmountSent    (float64) - Amount successfully sent
//     TotalDeducted (float64) - Total amount deducted from balance
//     NewBalance    (float64) - New account balance
//     ZenoPayResponse (SendMoneyZenoPayResponseData) - Provider-specific details
//   - error if any step fails.
func SendMoney(requestBody *models.SendMoneyRequest) (*models.SendMoneyResponse, error) {
	// Get the configured endpoint URL for sending money (cash-in).
	URL := utils.GetURL("cashin")

	// Get required HTTP headers for authentication and content type.
	headers, err := utils.GetHeaders()
	if err != nil {
		return nil, err
	}

	// Send the HTTP POST request with payload and headers to the API endpoint.
	raw, err := request.Post(URL, requestBody, headers)
	if err != nil {
		return nil, err
	}

	// Convert the raw JSON response into the SendMoneyResponse struct.
	var responseBody models.SendMoneyResponse
	if err := json.Unmarshal(raw, &responseBody); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Return the parsed API response.
	return &responseBody, nil
}

// Checkout sends a checkout request to the configured endpoint.
// It retrieves the API URL and headers, sends the HTTP POST request, and parses the JSON response.
//
// Parameters:
//   - requestBody: *models.CheckoutRequest with the following fields:
//     Amount               (float64)  - Transaction amount
//     Currency             (string)   - Currency code (e.g., "USD", "TZS")
//     RedirectURL          (string)   - URL to redirect after payment
//     CustomerName         (string)   - Buyer’s full name
//     CustomerPhoneNumber  (string)   - Buyer’s phone number (will be formatted before sending)
//     CustomerEmailAddress (string)   - Buyer’s email address
//     WebhookURL           (string)   - Webhook URL for transaction status updates
//     Metadata             (any)      - Custom metadata
//     OrderID              (string)   - Unique order identifier
//
// Returns:
//   - *models.CheckoutResponse with fields:
//     PaymentLink          (string) - URL to complete the payment
//     TransactionReference (string) - Provider transaction reference
//     Error                (string) - Optional error message
//   - error if any step fails.
func Checkout(requestBody *models.CheckoutRequest) (*models.CheckoutResponse, error) {
	// Get the configured endpoint URL for checkout requests.
	URL := utils.GetURL("checkout")

	// Get required HTTP headers for authentication and content type.
	headers, err := utils.GetHeaders()
	if err != nil {
		return nil, err
	}

	// Send the HTTP POST request with payload and headers to the API endpoint.
	raw, err := request.Post(URL, requestBody, headers)
	if err != nil {
		return nil, err
	}

	// Convert the raw JSON response into the CheckoutResponse struct.
	var responseBody models.CheckoutResponse
	if err := json.Unmarshal(raw, &responseBody); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// Return the parsed API response.
	return &responseBody, nil
}
