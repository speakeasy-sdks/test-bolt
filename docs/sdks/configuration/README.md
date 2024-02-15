# Configuration
(*Configuration*)

## Overview

Merchant configuration endpoints allow you to retrieve and configure merchant-level
configuration, such as callback URLs, identifiers, secrets, etc...


### Available Operations

* [MerchantCallbacksGet](#merchantcallbacksget) - Retrieve callback URLs for the merchant
* [MerchantCallbacksUpdate](#merchantcallbacksupdate) - Update callback URLs for the merchant
* [MerchantIdentifiersGet](#merchantidentifiersget) - Retrieve identifiers for the merchant

## MerchantCallbacksGet

Return callback URLs configured on the merchant such as OAuth URLs.


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
    res, err := s.Configuration.MerchantCallbacksGet(ctx, operations.MerchantCallbacksGetRequest{
        XPublishableKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.MerchantCallbacksGetRequest](../../pkg/models/operations/merchantcallbacksgetrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.MerchantCallbacksGetResponse](../../pkg/models/operations/merchantcallbacksgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MerchantCallbacksUpdate

Update and configure callback URLs on the merchant such as OAuth URLs.


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
    res, err := s.Configuration.MerchantCallbacksUpdate(ctx, operations.MerchantCallbacksUpdateRequest{
        XPublishableKey: "<value>",
        CallbackUrls: shared.CallbackUrls{
            AccountPage: testbolt.String("https://www.example.com/account"),
            BaseDomain: testbolt.String("https://www.example.com/"),
            ConfirmationRedirect: testbolt.String("https://www.example.com/bolt/redirect"),
            CreateOrder: testbolt.String("https://www.example.com/bolt/order"),
            Debug: testbolt.String("https://www.example.com/bolt/debug"),
            GetAccount: testbolt.String("https://www.example.com/bolt/account"),
            MobileAppDomain: testbolt.String("https://m.example.com/"),
            OauthLogout: testbolt.String("https://www.example.com/bolt/logout"),
            OauthRedirect: testbolt.String("https://www.example.com/bolt/oauth"),
            PrivacyPolicy: testbolt.String("https://www.example.com/privacy-policy"),
            ProductInfo: testbolt.String("https://www.example.com/bolt/product"),
            RemoteAPI: testbolt.String("https://www.example.com/bolt/remote-api"),
            Shipping: testbolt.String("https://www.example.com/bolt/shipping"),
            SupportPage: testbolt.String("https://www.example.com/help"),
            Tax: testbolt.String("https://www.example.com/bolt/tax"),
            TermsOfService: testbolt.String("https://www.example.com/terms-of-service"),
            UniversalMerchantAPI: testbolt.String("https://www.example.com/bolt/merchant-api"),
            UpdateCart: testbolt.String("https://www.example.com/bolt/cart"),
            ValidateAdditionalAccountData: testbolt.String("https://www.example.com/bolt/validate-account"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CallbackUrls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.MerchantCallbacksUpdateRequest](../../pkg/models/operations/merchantcallbacksupdaterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.MerchantCallbacksUpdateResponse](../../pkg/models/operations/merchantcallbacksupdateresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.MerchantCallbacksUpdateResponseBody | 400                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## MerchantIdentifiersGet

Return several identifiers for the merchant, such as public IDs, publishable keys, signing secrets, etc...

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"context"
	"log"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Configuration.MerchantIdentifiersGet(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Identifiers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.MerchantIdentifiersGetResponse](../../pkg/models/operations/merchantidentifiersgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
