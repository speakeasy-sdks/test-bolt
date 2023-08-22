# Identifier

A type and value combination that defines the identifier used to detect
the existence of an account.



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `IdentifierType`                                                                         | [IdentifierIdentifierType](../../models/shared/identifieridentifiertype.md)              | :heavy_check_mark:                                                                       | The type of identifier                                                                   | email                                                                                    |
| `IdentifierValue`                                                                        | *string*                                                                                 | :heavy_check_mark:                                                                       | The value of the identifier. The value must be valid for the specified `identifier_type` | alice@example.com                                                                        |