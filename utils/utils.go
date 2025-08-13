package utils

import (
	"fmt"

	"github.com/hekimapro/utils/helpers"
	"github.com/hekimapro/utils/request"
)

// GetHeaders retrieves the API key from environment variables and constructs the request headers.
//
// Steps:
// 1. Fetch the ZENOPAY_API_KEY from environment using helpers.GetENVValue.
// 2. If the key is missing, return an error indicating the missing environment variable.
// 3. Construct the headers map with "x-api-key" set to the API key.
// 4. Return the headers for use in API requests.
//
// Returns:
// - Pointer to request.Headers containing the API key header.
// - Error if the API key is missing.
func GetHeaders() (*request.Headers, error) {
	APIKEY := helpers.GetENVValue("zenopay api key")
	if APIKEY == "" {
		return nil, helpers.CreateError("ZENOPAY_API_KEY is missing in .env file")
	}

	headers := &request.Headers{"x-api-key": APIKEY}
	return headers, nil
}

// GetURL returns the full API endpoint URL for the given endpoint name.
//
// Steps:
// 1. Define the base URL for the Zenopay API.
// 2. Match the endpointName with known endpoints:
//   - "push" => USSD push
//   - "status" => Order status check
//   - "cashin" => Local wallet cash-in
//   - "checkout" => Checkout payment
//
// 3. Return the full URL for the selected endpoint.
// 4. If endpointName does not match any known endpoint, return the base URL.
//
// Parameters:
// - endpointName: Name of the endpoint to generate the URL for.
//
// Returns:
// - Full URL string to use in API requests.
func GetURL(endpointName string) string {
	baseURL := "https://zenoapi.com/api/payments"

	switch endpointName {
	case "push":
		return fmt.Sprintf("%s/mobile_money_tanzania", baseURL)
	case "status":
		return fmt.Sprintf("%s/order-status", baseURL)
	case "cashin":
		return fmt.Sprintf("%s/walletcashin/process", baseURL)
	case "checkout":
		return fmt.Sprintf("%s/checkout", baseURL)
	default:
		return baseURL
	}
}
