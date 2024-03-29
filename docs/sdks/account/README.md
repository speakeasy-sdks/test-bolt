# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [AccountAddPaymentMethod](#accountaddpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [AccountAddressDelete](#accountaddressdelete) - Delete an existing address
* [AccountExists](#accountexists) - Determine the existence of a Bolt account
* [AccountGet](#accountget) - Retrieve account details

## AccountAddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from
your backend because authentication requires the use of your private key.<br />
**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function,
which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


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


    operationSecurity := operations.AccountAddPaymentMethodSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
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
                                ID: "D4g3h5tBuVYK9",
                            },
                    ),
                    Bin: "411111",
                    Expiration: "2025-03",
                    ID: testbolt.String("X5h6j8uLpVGK0"),
                    Last4: "1004",
                    Network: shared.PaymentMethodCreditCardNetworkVisa,
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                    Type: shared.TypeCredit,
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

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AccountAddPaymentMethodRequest](../../pkg/models/operations/accountaddpaymentmethodrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.AccountAddPaymentMethodSecurity](../../pkg/models/operations/accountaddpaymentmethodsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.AccountAddPaymentMethodResponse](../../pkg/models/operations/accountaddpaymentmethodresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AccountAddressDelete

Delete an existing address. Deleting an address does not invalidate transactions or
shipments that are associated with it.


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
    res, err := s.Account.AccountAddressDelete(ctx, operations.AccountAddressDeleteRequest{
        XPublishableKey: "<value>",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountAddressDeleteRequest](../../pkg/models/operations/accountaddressdeleterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.AccountAddressDeleteResponse](../../pkg/models/operations/accountaddressdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 422                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AccountExists

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"context"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"log"
)

func main() {
    s := testbolt.New()

    ctx := context.Background()
    res, err := s.Account.AccountExists(ctx, operations.AccountExistsRequest{
        XPublishableKey: "<value>",
        Identifier: shared.Identifier{
            IdentifierType: shared.IdentifierTypeEmail,
            IdentifierValue: "alice@example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AccountExistsRequest](../../pkg/models/operations/accountexistsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.AccountExistsResponse](../../pkg/models/operations/accountexistsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 422                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AccountGet

Retrieve a shopper's account details, such as addresses and payment information

### Example Usage

```go
package main

import(
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := testbolt.New()


    operationSecurity := operations.AccountGetSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.AccountGet(ctx, operations.AccountGetRequest{
        XPublishableKey: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AccountGetRequest](../../pkg/models/operations/accountgetrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.AccountGetSecurity](../../pkg/models/operations/accountgetsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.AccountGetResponse](../../pkg/models/operations/accountgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
