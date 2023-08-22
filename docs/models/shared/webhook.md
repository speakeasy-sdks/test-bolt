# Webhook

The webhook was successfully created


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `CreatedAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | The time at which the webhook was created           | 2017-07-21T17:32:28Z                                |
| `Event`                                             | [WebhookEvent](../../models/shared/webhookevent.md) | :heavy_check_mark:                                  | N/A                                                 |                                                     |
| `ID`                                                | **string*                                           | :heavy_minus_sign:                                  | The webhook's ID                                    | wh_za7VbYcSQU2zRgGQXQAm-g                           |
| `URL`                                               | *string*                                            | :heavy_check_mark:                                  | The webhook's URL                                   | https://www.example.com/webhook                     |