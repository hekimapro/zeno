# ZenoPay PRO Go SDK

The **ZenoPay PRO Go SDK** is a robust Go package designed to interact seamlessly with the ZenoPay API. It enables developers to integrate payment functionalities such as USSD push requests, transaction status checks, local and international money transfers, and checkout payments. The SDK supports both synchronous API calls and asynchronous webhook-based notifications for real-time updates.

---

## Features

- **USSD Push Requests**: Prompt customers with USSD payment requests directly on their phones.
- **Transaction Status Checks**: Retrieve real-time status of transactions using order IDs.
- **Local Money Transfers**: Send money to mobile payment accounts within supported regions.
- **International Money Transfers**: Facilitate cross-border payments with ease.
- **Checkout Payments**: Initiate secure checkout processes for online transactions.
- **Webhook Support**: Receive real-time payment notifications via customizable webhook endpoints.
- **Error Handling**: Comprehensive error responses for robust integration.
- **Type-Safe Models**: Structured request and response models for enhanced reliability.

---

## Installation

To install the ZenoPay PRO Go SDK, run the following command:

```bash
go get github.com/hekimapro/zeno
```

Ensure you have Go 1.16 or later installed for module support.

---

## Environment Setup

Set your ZenoPay API key in a `.env` file or as an environment variable:

```env
ZENOPAY_API_KEY=your_api_key_here
```

You can obtain your API key from the [ZenoPay Dashboard](https://zenoapi.com/dashboard).

> **Note**: Ensure the `.env` file is loaded using a package like `github.com/joho/godotenv` if needed:

```go
import "github.com/joho/godotenv"

func init() {
	godotenv.Load()
}
```

---

## Package Imports

Import the necessary packages to use the SDK:

```go
import (
	"github.com/hekimapro/zeno"
	"github.com/hekimapro/zeno/models"
)
```

- Use `zeno.<FunctionName>()` to call API methods.
- Use `models` for structured request and response types.

---

## Usage Examples

### 1. Sending a USSD Push Request

Send a USSD push request to prompt a customer for payment.

```go
package main

import (
	"fmt"
	"log"
	"github.com/hekimapro/zeno"
	"github.com/hekimapro/zeno/models"
)

func main() {
	request := &models.USSDPushRequest{
		Amount:               1000,
		CustomerName:         "Hekima Peter",
		CustomerEmailAddress: "info@hekima.pro",
		CustomerPhoneNumber:  "0756628215",
		OrderID:              "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7",
		WebhookURL:           "https://hekima.pro/api/payment/callback",
		Metadata:             []string{},
	}

	response, err := zeno.PushUSSD(request)
	if err != nil {
		log.Fatalf("USSD Push failed: %v", err)
	}

	fmt.Printf("USSD Push Response: %+v\n", response)
}
```

### 2. Checking Transaction Status

Retrieve the status of a transaction using its `orderID`.

```go
package main

import (
	"fmt"
	"log"
	"github.com/hekimapro/zeno"
)

func main() {
	orderID := "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7"
	response, err := zeno.CheckStatus(orderID)
	if err != nil {
		log.Fatalf("Status check failed: %v", err)
	}

	fmt.Printf("Transaction Status: %+v\n", response)
}
```

### 3. Sending Money Locally

Transfer money to a mobile payment account.

```go
package main

import (
	"fmt"
	"log"
	"github.com/hekimapro/zeno"
	"github.com/hekimapro/zeno/models"
)

func main() {
	request := &models.SendMoneyRequest{
		UtilityCode:   "CASHIN",
		TransactionID: "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7",
		PhoneNumber:   "0756628215",
		Amount:        5000,
		PIN:           1234,
	}

	response, err := zeno.SendMoney(request)
	if err != nil {
		log.Fatalf("Send money failed: %v", err)
	}

	fmt.Printf("Send Money Response: %+v\n", response)
}
```

### 4. Initiating a Checkout Payment

Start a checkout process for online payments.

```go
package main

import (
	"fmt"
	"log"
	"github.com/hekimapro/zeno"
	"github.com/hekimapro/zeno/models"
)

func main() {
	request := &models.CheckoutRequest{
		PhoneNumber:          "254712345678",
		Amount:               10000,
		Currency:             "KES",
		RedirectURL:          "https://hekima.pro/payment/redirect",
		CustomerName:         "Hekima Peter",
		CustomerEmailAddress: "info@hekima.pro",
		CustomerPhoneNumber:  "0756628215",
		OrderID:              "0fd34cee-e4fa-44b4-bb4d-ff563bb4cfc7",
		WebhookURL:           "https://hekima.pro/api/payment/callback",
		Metadata:             []string{},
	}

	response, err := zeno.Checkout(request)
	if err != nil {
		log.Fatalf("Checkout failed: %v", err)
	}

	fmt.Printf("Checkout Response: %+v\n", response)
}
```

### 5. Handling Webhook Notifications

Set up a webhook endpoint to receive payment notifications.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/hekimapro/zeno/models"
)

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Verify API key in header
		apiKey := r.Header.Get("x-api-key")
		if apiKey != os.Getenv("ZENOPAY_API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var payload models.WebhookPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		fmt.Printf("Payment Event: %s\n", payload.Event)
		fmt.Printf("Transaction Status: %s\n", payload.Data.Status)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## Models

### Request Models

- **`USSDPushRequest`**: Configures USSD push requests with customer details and webhook URL.
- **`SendMoneyRequest`**: Defines parameters for local money transfers, including utility code and PIN.
- **`CheckoutRequest`**: Specifies details for checkout payments, including currency and redirect URL.
- **`WebhookPayload`**: Structures incoming webhook data, including event type and transaction status.

### Response Models

- **`USSDPushResponse`**: Returns details of the USSD push request, such as confirmation status.
- **`CheckStatusResponse`**: Provides transaction status details, including success or failure states.
- **`SendMoneyResponse`**: Confirms the outcome of local money transfers.
- **`CheckoutResponse`**: Details the result of checkout payment requests.

> **Tip**: Refer to the `models` package for detailed field documentation and type definitions.

---

## Error Handling

All API functions return an `error` type. Always check for errors to handle failures gracefully:

```go
if err != nil {
	log.Printf("Error: %v", err)
	// Handle error (e.g., retry, notify user, etc.)
}
```

Common errors include:
- Invalid API key
- Network issues
- Invalid request parameters
- Server-side issues (e.g., ZenoPay API downtime)

---

## Support

For technical assistance, reach out to:

- **Email**: [support@zenoapi.com](mailto:support@zenoapi.com)
- **Website**: [https://zenoapi.com](https://zenoapi.com)

Developer Contact:

- **Email**: [info@hekima.pro](mailto:info@hekima.pro)
- **Phone**: +255 756 628 215

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository: [github.com/hekimapro/zeno](https://github.com/hekimapro/zeno)
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m "Add your feature"`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a pull request

Please include tests and documentation for new features.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.