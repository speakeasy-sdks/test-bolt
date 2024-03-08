<!-- Start SDK Example Usage [usage] -->
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
		APIKey: "<YOUR_API_KEY_HERE>",
		Oauth:  "Bearer <YOUR_ACCESS_TOKEN_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.AccountAddPaymentMethod(ctx, operations.AccountAddPaymentMethodRequest{
		XPublishableKey: "<value>",
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
<!-- End SDK Example Usage [usage] -->