# ZenoPay

ZenoPay is a Go package for integrating with the ZenoPay payment API. It provides methods to initiate payments and check payment statuses.

## Installation

To install the package, run:

```bash
go get github.com/hekimapro/zeno
```

## Features

- Initiate a payment
- Check payment status

## Requirements

- Go 1.18 or later
- ZenoPay API credentials (`APIKey`, `SecretKey`, `AccountID`)

## Usage

### 1. Import the package

```go
import "github.com/hekimapro/zeno"
```

### 2. Initialize ZenoPay

```go
package main

import (
	"fmt"
	"log"

	"github.com/hekimapro/zeno"
)

func main() {
	// Initialize ZenoPay with API credentials
	zenoPay := zeno.NewZenoPay(zeno.ZenoPayOptionsType{
		APIKey:    "your_api_key",
		SecretKey: "your_secret_key",
		AccountID: "your_account_id",
	})

	// Example usage of Pay function
	orderID, err := zenoPay.Pay(zeno.PaymentOptionsType{
		CustomerName:        "John Doe",
		CustomerEmail:       "johndoe@example.com",
		CustomerPhoneNumber: "1234567890",
		AmountToCharge:      100.50,
		CallbackURL:         "https://your-callback-url.com",
	})

	if err != nil {
		log.Fatalf("Error initiating payment: %s", err.Error())
	}
	fmt.Printf("Payment initiated successfully. Order ID: %s\n", orderID)

	// Example usage of CheckPaymentStatus function
	status, err := zenoPay.CheckPaymentStatus(orderID)
	if err != nil {
		log.Fatalf("Error checking payment status: %s", err.Error())
	}
	fmt.Printf("Payment status: %s\n", status)
}
```

## API Reference

### ZenoPay Options

| Field       | Type   | Description                |
|-------------|--------|----------------------------|
| `APIKey`    | string | Your ZenoPay API key       |
| `SecretKey` | string | Your ZenoPay secret key    |
| `AccountID` | string | Your ZenoPay account ID    |

### Payment Options

| Field                 | Type    | Description                            |
|-----------------------|---------|----------------------------------------|
| `CustomerName`        | string  | Name of the customer                   |
| `CustomerEmail`       | string  | Email address of the customer          |
| `CustomerPhoneNumber` | string  | Phone number of the customer           |
| `AmountToCharge`      | float64 | Payment amount                         |
| `CallbackURL`         | string  | URL to receive payment status updates  |

### Functions

#### `Pay`

- **Description:** Initiates a payment.
- **Parameters:**
  - `PaymentOptionsType`
- **Returns:**
  - `string` - Order ID
  - `error` - Error message, if any

#### `CheckPaymentStatus`

- **Description:** Checks the payment status for a given order ID.
- **Parameters:**
  - `string` - Order ID
- **Returns:**
  - `string` - Payment status
  - `error` - Error message, if any

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Replace `"your_api_key"`, `"your_secret_key"`, and `"your_account_id"` with your actual ZenoPay API credentials to get started.
```