# Webhooks
(*Webhooks*)

## Overview

Set up webhooks to notify your backend of events within Bolt. These webhooks
can communicate with your OMS or other systems to keep them up to date with Bolt.


<https://help.bolt.com/get-started/during-checkout/webhooks/>
### Available Operations

* [WebhooksCreate](#webhookscreate) - Create a webhook to subscribe to certain events
* [WebhooksDelete](#webhooksdelete) - Delete an existing webhook
* [WebhooksGet](#webhooksget) - Retrieve information for a specific webhook
* [WebhooksGetAll](#webhooksgetall) - Retrieve information about all existing webhooks

## WebhooksCreate

Create a new webhook to receive notifications from Bolt about various events, such as transaction status.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/callbacks"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.WebhooksCreate(ctx, shared.WebhookInput{
        Event: shared.CreateWebhookEventEventList(
                shared.EventList{
                    DotTag: shared.EventListTagList,
                    EventList: []shared.EventListEventList{
                        shared.EventListEventListAuth,
                        shared.EventListEventListCapture,
                    },
                },
        ),
        URL: "https://www.example.com/webhook",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.WebhookInput](../../models/shared/webhookinput.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.WebhooksCreateResponse](../../models/operations/webhookscreateresponse.md), error**


## WebhooksDelete

Delete an existing webhook. You will no longer receive notifications from Bolt about its events.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.WebhooksDelete(ctx, operations.WebhooksDeleteRequest{
        ID: "wh_za7VbYcSQU2zRgGQXQAm-g",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.WebhooksDeleteRequest](../../models/operations/webhooksdeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.WebhooksDeleteResponse](../../models/operations/webhooksdeleteresponse.md), error**


## WebhooksGet

Retrieve information for an existing webhook.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.WebhooksGet(ctx, operations.WebhooksGetRequest{
        ID: "wh_za7VbYcSQU2zRgGQXQAm-g",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.WebhooksGetRequest](../../models/operations/webhooksgetrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.WebhooksGetResponse](../../models/operations/webhooksgetresponse.md), error**


## WebhooksGetAll

Retrieve information about all existing webhooks.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testbolt "github.com/speakeasy-sdks/test-bolt"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/shared"
	"github.com/speakeasy-sdks/test-bolt/pkg/models/operations"
)

func main() {
    s := testbolt.New(
        testbolt.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.WebhooksGetAll(ctx, operations.WebhooksGetAllRequest{
        XPublishableKey: "maxime",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksGetAll200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.WebhooksGetAllRequest](../../models/operations/webhooksgetallrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.WebhooksGetAllResponse](../../models/operations/webhooksgetallresponse.md), error**

