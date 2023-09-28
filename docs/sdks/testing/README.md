# Testing
(*Testing*)

## Overview

Endpoints that allow you to generate and retrieve test data to verify certain
flows in non-production environments.


### Available Operations

* [TestingAccountCreate](#testingaccountcreate) - Create a test account
* [TestingCreditCardGet](#testingcreditcardget) - Retrieve a test credit card, including its token
* [TestingShipmentTrackingCreate](#testingshipmenttrackingcreate) - Simulate a shipment tracking update

## TestingAccountCreate

Create a Bolt shopper account for testing purposes.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/types"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Testing.TestingAccountCreate(ctx, shared.AccountTestCreationDataInput{
        DeactivateAt: types.MustTimeFromString("2017-07-21T17:32:28Z"),
        EmailState: shared.AccountTestCreationDataEmailStateUnverified,
        HasAddress: testbolt.Bool(true),
        IsMigrated: testbolt.Bool(true),
        PhoneState: shared.AccountTestCreationDataPhoneStateVerified,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTestCreationData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.AccountTestCreationDataInput](../../models/shared/accounttestcreationdatainput.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.TestingAccountCreateResponse](../../models/operations/testingaccountcreateresponse.md), error**


## TestingCreditCardGet

Retrieve test credit card information. This includes its token, which is
generated against the `4111 1111 1111 1004` test card.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Testing.TestingCreditCardGet(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.TestingCreditCardGetResponse](../../models/operations/testingcreditcardgetresponse.md), error**


## TestingShipmentTrackingCreate

Simulate a shipment tracking update, such as those that Bolt would ingest from
third-party shipping providers that would allow updating tracking and delivery
information to shipments associated with orders.


### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/types"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Testing.TestingShipmentTrackingCreate(ctx, shared.ShipmentTrackingUpdate{
        DeliveryDate: types.MustTimeFromString("2014-08-23:T06:00:00Z"),
        Status: shared.ShipmentTrackingUpdateStatusInTransit,
        TrackingDetails: []shared.ShipmentTrackingUpdateTrackingDetails{
            shared.ShipmentTrackingUpdateTrackingDetails{
                CountryCode: testbolt.String("US"),
                EventDate: testbolt.String("2014-08-21:T14:24:00Z"),
                Locality: testbolt.String("San Francisco"),
                Message: testbolt.String("Billing information received"),
                PostalCode: testbolt.String("94105"),
                Region: testbolt.String("CA"),
                Status: shared.ShipmentTrackingUpdateTrackingDetailsStatusPreTransit.ToPointer(),
            },
        },
        TrackingNumber: "MockBolt-143292",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.ShipmentTrackingUpdate](../../models/shared/shipmenttrackingupdate.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.TestingShipmentTrackingCreateResponse](../../models/operations/testingshipmenttrackingcreateresponse.md), error**

