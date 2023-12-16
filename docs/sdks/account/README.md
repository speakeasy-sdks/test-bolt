# Account
(*Account*)

## Overview

Account endpoints allow you to view and manage shoppers' accounts. For example,
you can add or remove addresses and payment information.


### Available Operations

* [AccountAddPaymentMethod](#accountaddpaymentmethod) - Add a payment method to a shopper's Bolt account Wallet.
* [AccountAddressCreate](#accountaddresscreate) - Add an address
* [AccountAddressDelete](#accountaddressdelete) - Delete an existing address
* [AccountAddressEdit](#accountaddressedit) - Edit an existing address
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
        XPublishableKey: "string",
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
| sdkerrors.SDKError | 400-600            | */*                |

## AccountAddressCreate

Add an address to the shopper's account

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


    operationSecurity := operations.AccountAddressCreateSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddressCreate(ctx, operations.AccountAddressCreateRequest{
        XPublishableKey: "string",
        AddressListing: shared.AddressListing{
            Company: testbolt.String("ACME Corporation"),
            CountryCode: "US",
            Email: testbolt.String("alice@example.com"),
            FirstName: "Alice",
            ID: "D4g3h5tBuVYK9",
            IsDefault: testbolt.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: testbolt.String("+14155550199"),
            PostalCode: "94105",
            Region: testbolt.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: testbolt.String("c/o Shipping Department"),
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AccountAddressCreateRequest](../../pkg/models/operations/accountaddresscreaterequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.AccountAddressCreateSecurity](../../pkg/models/operations/accountaddresscreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.AccountAddressCreateResponse](../../pkg/models/operations/accountaddresscreateresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.AccountAddressCreateResponseBody | 400                                        | application/json                           |
| sdkerrors.SDKError                         | 400-600                                    | */*                                        |

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
	"net/http"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Account.AccountAddressDelete(ctx, operations.AccountAddressDeleteRequest{
        XPublishableKey: "string",
        ID: "D4g3h5tBuVYK9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

## AccountAddressEdit

Edit an existing address on the shopper's account. This does not edit addresses
that are already associated with other resources, such as transactions or
shipments.


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


    operationSecurity := operations.AccountAddressEditSecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
            Oauth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.AccountAddressEdit(ctx, operations.AccountAddressEditRequest{
        XPublishableKey: "string",
        AddressListing: shared.AddressListing{
            Company: testbolt.String("ACME Corporation"),
            CountryCode: "US",
            Email: testbolt.String("alice@example.com"),
            FirstName: "Alice",
            ID: "D4g3h5tBuVYK9",
            IsDefault: testbolt.Bool(true),
            LastName: "Baker",
            Locality: "San Francisco",
            Phone: testbolt.String("+14155550199"),
            PostalCode: "94105",
            Region: testbolt.String("CA"),
            StreetAddress1: "535 Mission St, Ste 1401",
            StreetAddress2: testbolt.String("c/o Shipping Department"),
        },
        ID: "D4g3h5tBuVYK9",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddressListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountAddressEditRequest](../../pkg/models/operations/accountaddresseditrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.AccountAddressEditSecurity](../../pkg/models/operations/accountaddresseditsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.AccountAddressEditResponse](../../pkg/models/operations/accountaddresseditresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.AccountAddressEditResponseBody | 400                                      | application/json                         |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |

## AccountExists

Determine whether or not an identifier is associated with an existing Bolt account.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"context"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Account.AccountExists(ctx, operations.AccountExistsRequest{
        XPublishableKey: "string",
        Identifier: shared.Identifier{
            IdentifierType: shared.IdentifierTypeEmail,
            IdentifierValue: "alice@example.com",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

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
        XPublishableKey: "string",
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
| sdkerrors.SDKError | 400-600            | */*                |
