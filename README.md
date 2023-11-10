# github.com/speakeasy-sdks/test-bolt

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-bolt
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New()

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag: shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddress: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testbolt.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.TypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Account](docs/sdks/account/README.md)

* [AccountAddPaymentMethod](docs/sdks/account/README.md#accountaddpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [AccountAddressCreate](docs/sdks/account/README.md#accountaddresscreate) - Add an address
* [AccountAddressDelete](docs/sdks/account/README.md#accountaddressdelete) - Delete an existing address
* [AccountAddressEdit](docs/sdks/account/README.md#accountaddressedit) - Edit an existing address
* [AccountExists](docs/sdks/account/README.md#accountexists) - Determine the existence of a Bolt account
* [AccountGet](docs/sdks/account/README.md#accountget) - Retrieve account details

### [Payments](docs/sdks/payments/README.md)

* [GuestPaymentsInitialize](docs/sdks/payments/README.md#guestpaymentsinitialize) - Initialize a Bolt payment for guest shoppers
* [PaymentsInitialize](docs/sdks/payments/README.md#paymentsinitialize) - Initialize a Bolt payment for logged in shoppers

### [Configuration](docs/sdks/configuration/README.md)

* [MerchantCallbacksGet](docs/sdks/configuration/README.md#merchantcallbacksget) - Retrieve callback URLs for the merchant
* [MerchantCallbacksUpdate](docs/sdks/configuration/README.md#merchantcallbacksupdate) - Update callback URLs for the merchant
* [MerchantIdentifiersGet](docs/sdks/configuration/README.md#merchantidentifiersget) - Retrieve identifiers for the merchant

### [Testing](docs/sdks/testing/README.md)

* [TestingAccountCreate](docs/sdks/testing/README.md#testingaccountcreate) - Create a test account
* [TestingCreditCardGet](docs/sdks/testing/README.md#testingcreditcardget) - Retrieve a test credit card, including its token
* [TestingShipmentTrackingCreate](docs/sdks/testing/README.md#testingshipmenttrackingcreate) - Simulate a shipment tracking update

### [Webhooks](docs/sdks/webhooks/README.md)

* [WebhooksCreate](docs/sdks/webhooks/README.md#webhookscreate) - Create a webhook to subscribe to certain events
* [WebhooksDelete](docs/sdks/webhooks/README.md#webhooksdelete) - Delete an existing webhook
* [WebhooksGet](docs/sdks/webhooks/README.md#webhooksget) - Retrieve information for a specific webhook
* [WebhooksGetAll](docs/sdks/webhooks/README.md#webhooksgetall) - Retrieve information about all existing webhooks
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountAddressCreateResponseBody | 400                                        | application/json                           |
| sdkerrors.SDKError                         | 400-600                                    | */*                                        |

### Example

```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New()

	operationSecurity := operations.AccountAddressCreateSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddressCreate(ctx, operations.AccountAddressCreateRequest{
		XPublishableKey: "string",
		AddressListing: shared.AddressListing{
			Company:        testbolt.String("ACME Corporation"),
			CountryCode:    "US",
			Email:          testbolt.String("alice@example.com"),
			FirstName:      "Alice",
			ID:             "D4g3h5tBuVYK9",
			IsDefault:      testbolt.Bool(true),
			LastName:       "Baker",
			Locality:       "San Francisco",
			Phone:          testbolt.String("+14155550199"),
			PostalCode:     "94105",
			Region:         testbolt.String("CA"),
			StreetAddress1: "535 Mission St, Ste 1401",
			StreetAddress2: testbolt.String("c/o Shipping Department"),
		},
	}, operationSecurity)
	if err != nil {

		var e *sdkerrors.AccountAddressCreateResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.{username}.dev.bolt.me/v3` | `username` (default is `BL_DOMAIN`) |
| 1 | `https://{environment}.bolt.com/v3` | `environment` (default is `api-sandbox`) |

#### Example

```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New(
		testbolt.WithServerIndex(1),
	)

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag: shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddress: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testbolt.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.TypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithEnvironment testbolt.ServerEnvironment`
 * `WithUsername string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New(
		testbolt.WithServerURL("https://api.{username}.dev.bolt.me/v3"),
	)

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag: shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddress: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testbolt.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.TypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New()

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag: shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddress: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testbolt.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.TypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
	s := testbolt.New()

	operationSecurity := operations.AccountAddPaymentMethodSecurity{
		APIKey: "",
		Oauth:  "",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "string",
		PaymentMethod: shared.CreatePaymentMethodPaymentMethodCreditCard(
			shared.PaymentMethodCreditCard{
				DotTag: shared.PaymentMethodCreditCardTagCreditCard,
				BillingAddress: shared.CreateAddressReferenceAddressReferenceID(
					shared.AddressReferenceID{
						DotTag: shared.AddressReferenceIDTagID,
						ID:     "D4g3h5tBuVYK9",
					},
				),
				Bin:        "411111",
				Expiration: "2025-03",
				ID:         testbolt.String("X5h6j8uLpVGK0"),
				Last4:      "1004",
				Network:    shared.PaymentMethodCreditCardNetworkVisa,
				Token:      "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
				Type:       shared.TypeCredit,
			},
		),
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.PaymentMethod != nil {
		// handle response
	}
}

```
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
