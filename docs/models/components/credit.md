# Credit

Object containing more information about the credit being paid


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AdditionalInstructions`                                                  | **string*                                                                 | :heavy_minus_sign:                                                        | Free form text field providing additional information about a transaction | FDIC sweep interest payment                                               |
| `CreditType`                                                              | [*components.CreditType](../../models/components/credittype.md)           | :heavy_minus_sign:                                                        | Provides more details on the type of credit                               | WRITE_OFF                                                                 |
| `Taxable`                                                                 | **bool*                                                                   | :heavy_minus_sign:                                                        | Indicates whether the credit is taxable                                   | false                                                                     |