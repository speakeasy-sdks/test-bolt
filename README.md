# github.com/speakeasy-sdks/test-bolt

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-bolt
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
)

func main() {
    s := testbolt.New()
    operationSecurity := operations.AccountAddPaymentMethodSecurity{
            APIKey: "",
            Oauth: "",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
        XPublishableKey: "maroon Silicon female",
        PaymentMethod: shared.PaymentMethod{},
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

### [Configuration](docs/sdks/configuration/README.md)

* [MerchantCallbacksGet](docs/sdks/configuration/README.md#merchantcallbacksget) - Retrieve callback URLs for the merchant
* [MerchantCallbacksUpdate](docs/sdks/configuration/README.md#merchantcallbacksupdate) - Update callback URLs for the merchant
* [MerchantIdentifiersGet](docs/sdks/configuration/README.md#merchantidentifiersget) - Retrieve identifiers for the merchant

### [Payments](docs/sdks/payments/README.md)

* [GuestPaymentsInitialize](docs/sdks/payments/README.md#guestpaymentsinitialize) - Initialize a Bolt payment for guest shoppers
* [PaymentsInitialize](docs/sdks/payments/README.md#paymentsinitialize) - Initialize a Bolt payment for logged in shoppers

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

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
