# Payments
(*Payments*)

## Overview

Use the Payments API to tokenize and process alternative payment methods including Paypal with Bolt. This API is for the Bolt
Accounts package.


### Available Operations

* [GuestPaymentsInitialize](#guestpaymentsinitialize) - Initialize a Bolt payment for guest shoppers
* [PaymentsInitialize](#paymentsinitialize) - Initialize a Bolt payment for logged in shoppers

## GuestPaymentsInitialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for guest shoppers.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"context"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"log"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Payments.GuestPaymentsInitialize(ctx, operations.GuestPaymentsInitializeRequest{
        XPublishableKey: "string",
        GuestPaymentMethodInitializeRequest: shared.GuestPaymentMethodInitializeRequest{
            Cart: shared.Cart{
                Amounts: shared.Amounts{
                    Currency: "USD",
                    Tax: testbolt.Int64(900),
                    Total: 900,
                },
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amounts: shared.Amounts{
                            Currency: "USD",
                            Tax: testbolt.Int64(900),
                            Total: 900,
                        },
                        Code: testbolt.String("SUMMER10DISCOUNT"),
                        DetailsURL: testbolt.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: testbolt.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: testbolt.String("Large tote with Bolt logo."),
                        ImageURL: testbolt.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testbolt.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceAddressReferenceID(
                                shared.AddressReferenceID{
                                    DotTag: shared.AddressReferenceIDTagID,
                                    ID: "D4g3h5tBuVYK9",
                                },
                        ),
                        Carrier: testbolt.String("FedEx"),
                        Cost: &shared.Amounts{
                            Currency: "USD",
                            Tax: testbolt.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: shared.CreateGuestPaymentMethodInitializeRequestPaymentMethodPaymentMethodPaypal(
                    shared.PaymentMethodPaypal{
                        DotTag: shared.PaymentMethodPaypalTagPaypal,
                        Cancel: "www.example.com/handle_paypal_cancel",
                        Success: "www.example.com/handle_paypal_success",
                    },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethodInitializeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GuestPaymentsInitializeRequest](../../pkg/models/operations/guestpaymentsinitializerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GuestPaymentsInitializeResponse](../../pkg/models/operations/guestpaymentsinitializeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PaymentsInitialize

Initialize a Bolt payment token that will be used to reference this payment to
Bolt when it is updated or finalized for logged in shoppers.


### Example Usage

```go
package main

import(
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"log"
)

func main() {
    s := testbolt.New()


    operationSecurity := operations.PaymentsInitializeSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Payments.PaymentsInitialize(ctx, operations.PaymentsInitializeRequest{
        XPublishableKey: "string",
        PaymentMethodInitializeRequest: shared.PaymentMethodInitializeRequest{
            Cart: shared.Cart{
                Amounts: shared.Amounts{
                    Currency: "USD",
                    Tax: testbolt.Int64(900),
                    Total: 900,
                },
                Discounts: []shared.CartDiscount{
                    shared.CartDiscount{
                        Amounts: shared.Amounts{
                            Currency: "USD",
                            Tax: testbolt.Int64(900),
                            Total: 900,
                        },
                        Code: testbolt.String("SUMMER10DISCOUNT"),
                        DetailsURL: testbolt.String("https://www.example.com/SUMMER-SALE"),
                    },
                },
                DisplayID: testbolt.String("215614191"),
                Items: []shared.CartItem{
                    shared.CartItem{
                        Description: testbolt.String("Large tote with Bolt logo."),
                        ImageURL: testbolt.String("https://www.example.com/products/123456/images/1.png"),
                        Name: "Bolt Swag Bag",
                        Quantity: 1,
                        Reference: "item_100",
                        TotalAmount: 1000,
                        UnitPrice: 1000,
                    },
                },
                OrderDescription: testbolt.String("Order #1234567890"),
                OrderReference: "order_100",
                Shipments: []shared.CartShipment{
                    shared.CartShipment{
                        Address: shared.CreateAddressReferenceAddressReferenceExplicit(
                                shared.AddressReferenceExplicit{
                                    DotTag: shared.AddressReferenceExplicitTagExplicit,
                                    Company: testbolt.String("ACME Corporation"),
                                    CountryCode: "US",
                                    Email: testbolt.String("alice@example.com"),
                                    FirstName: "Alice",
                                    ID: "D4g3h5tBuVYK9",
                                    LastName: "Baker",
                                    Locality: "San Francisco",
                                    Phone: testbolt.String("+14155550199"),
                                    PostalCode: "94105",
                                    Region: testbolt.String("CA"),
                                    StreetAddress1: "535 Mission St, Ste 1401",
                                    StreetAddress2: testbolt.String("c/o Shipping Department"),
                                },
                        ),
                        Carrier: testbolt.String("FedEx"),
                        Cost: &shared.Amounts{
                            Currency: "USD",
                            Tax: testbolt.Int64(900),
                            Total: 900,
                        },
                    },
                },
            },
            PaymentMethod: shared.CreatePaymentMethodInitializeRequestPaymentMethodPaymentMethodSavedPaymentMethod(
                    shared.PaymentMethodSavedPaymentMethod{
                        DotTag: shared.PaymentMethodSavedPaymentMethodTagSavedPaymentMethod,
                        ID: "id",
                    },
            ),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethodInitializeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PaymentsInitializeRequest](../../pkg/models/operations/paymentsinitializerequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.PaymentsInitializeSecurity](../../pkg/models/operations/paymentsinitializesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.PaymentsInitializeResponse](../../pkg/models/operations/paymentsinitializeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
