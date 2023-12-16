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
    res, err := s.Webhooks.WebhooksCreate(ctx, shared.WebhookInput{
        Event: shared.CreateEventEventList(
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.WebhookInput](../../pkg/models/shared/webhookinput.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.WebhooksCreateResponse](../../pkg/models/operations/webhookscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## WebhooksDelete

Delete an existing webhook. You will no longer receive notifications from Bolt about its events.

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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.WebhooksDeleteRequest](../../pkg/models/operations/webhooksdeleterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.WebhooksDeleteResponse](../../pkg/models/operations/webhooksdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 422                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## WebhooksGet

Retrieve information for an existing webhook.

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.WebhooksGetRequest](../../pkg/models/operations/webhooksgetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.WebhooksGetResponse](../../pkg/models/operations/webhooksgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 422                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## WebhooksGetAll

Retrieve information about all existing webhooks.

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
    res, err := s.Webhooks.WebhooksGetAll(ctx, operations.WebhooksGetAllRequest{
        XPublishableKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.WebhooksGetAllRequest](../../pkg/models/operations/webhooksgetallrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.WebhooksGetAllResponse](../../pkg/models/operations/webhooksgetallresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
