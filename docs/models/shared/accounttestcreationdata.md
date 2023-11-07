# AccountTestCreationData


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `DeactivateAt`                                         | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    | 2017-07-21T17:32:28Z                                   |
| `EmailState`                                           | [shared.EmailState](../../models/shared/emailstate.md) | :heavy_check_mark:                                     | N/A                                                    | unverified                                             |
| `HasAddress`                                           | **bool*                                                | :heavy_minus_sign:                                     | N/A                                                    | true                                                   |
| `IsMigrated`                                           | **bool*                                                | :heavy_minus_sign:                                     | N/A                                                    | true                                                   |
| `PhoneState`                                           | [shared.PhoneState](../../models/shared/phonestate.md) | :heavy_check_mark:                                     | N/A                                                    | verified                                               |