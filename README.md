# ZenoPay PRO Go SDK

A Go package for interacting with the **ZenoPay API**. It allows you to send USSD push requests, check transaction status, send money locally or internationally, and manage checkout payments. Supports both synchronous API calls and webhook-based notifications.

---

## Features

- Push USSD requests to customer phones
- Check payment or transaction status
- Send money to mobile payment accounts
- Send money internationally
- Initiate checkout payments
- Webhook support for payment notifications

---

## Installation

```bash
go get github.com/hekimapro/zeno
```

---

## Environment Setup

Before making any request, set the `ZENOPAY_API_KEY` in your `.env` file:

```env
ZENOPAY_API_KEY = your_api_key_here
```

You can generate an API key from your Zenopay application or dashboard.

---

## Package Imports

```go
import (
	"github.com/hekimapro/zeno"
	"github.com/hekimapro/zeno/models"
)
```

- Use `zeno.FunctionName()` to call any API function
- Use `models` to access request and response types

---

## API Functions Usage

### 1. Push USSD

Send a USSD push request to the customer.

```go

request := &models.USSDPushRequest{
	Amount:      1000,
	CustomerName: "Hekima Peter",
	CustomerEmailAddress: "info@hekima.pro",
	CustomerPhoneNumber: "0756628215",
	OrderID:   "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7",
	WebhookURL: "https://hekima.pro/api/payment/callback",
	Metadata: [],
}

response, err := zeno.PushUSSD(request)
if err != nil {
	log.Fatal(err)
}

fmt.Println("USSD Push Response:", response)
```


---

### 2. Check Status

Check the status of a transaction by `orderID`.

```go
orderID := "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7"

response, err := zeno.CheckStatus(orderID)
if err != nil {
	log.Fatal(err)
}

fmt.Println(response)
```

### 3. Send Money

Send money to a mobile payment account.

```go
request := &models.SendMoneyRequest{
	UtilityCode: "CASHIN",
	TransactionID: "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7"
	PhoneNumber: "0756628215",
	Amount:      5000,
	PIN:   1234,
}

response, err := zeno.SendMoney(request)
if err != nil {
	log.Fatal(err)
}

fmt.Println(response)
```


### 4. Checkout

Initiate a checkout payment.

```go
request := &models.CheckoutRequest{
	PhoneNumber: "254712345678",
	Amount:      10000,
	Currency:    "KES",
	RedirectURL: "https://hekima.pro/payment/redirect",
	CustomerName: "Hekima Peter",
	CustomerEmailAddress: "info@hekima.pro",
	CustomerPhoneNumber: "0756628215",
	OrderID:   "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7",
	WebhookURL: "https://hekima.pro/api/payment/callback",
	Metadata: [],
}

response, err := zeno.Checkout(request)
if err != nil {
	log.Fatal(err)
}

fmt.Println(response)
```

---

### 5. Webhook

Zenopay sends POST requests to your configured webhook endpoint with a x-api-key.

**Example Webhook Handler in Go:**

```go
http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
	var payload models.WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Payment Event:", payload.Event)
	fmt.Println("Transaction Status:", payload.Data.Status)
	w.WriteHeader(http.StatusOK)
})
```
---

## Models

### Request Models

- **USSDPushRequest** — For sending USSD push requests
- **SendMoneyRequest** — For sending local payments
- **CheckoutRequest** — For sending international payments
- **CheckoutRequest** — For initiating checkout
- **WebhookPayload** — For handling webhook events

### Response Models

- **USSDPushResponse** — Response for USSD push
- **CheckStatus** — Transaction status details
- **SendMoneyResponse** — Response for local payments
- **CheckoutResponse** — Response for international payments
- **CheckoutResponse** — Response for checkout requests

---

## Support

For assistance, contact:\
Email: [support@zenoapi.com](mailto\:support@zenoapi.com)\
Website: [https://zenoapi.com](https://zenoapi.com)

Developer Contact:\
Email: [info@hekima.pro](mailto\:info@hekima.pro)\
Phone: +255 756 628 215

---

## License

MIT License

---

## Contribution

This package is open for contributions. Fork the repo, make changes, and submit a pull request.

