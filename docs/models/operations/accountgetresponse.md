# AccountGetResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `Account`                                              | [*shared.Account](../../models/shared/account.md)      | :heavy_minus_sign:                                     | The account details were successfully retrieved        |