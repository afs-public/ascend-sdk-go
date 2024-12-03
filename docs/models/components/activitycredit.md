# ActivityCredit

Object containing more information about the credit being paid


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `AdditionalInstructions`                                                        | **string*                                                                       | :heavy_minus_sign:                                                              | Free form text field providing additional information about a transaction       | FDIC sweep interest payment                                                     |
| `CreditType`                                                                    | [*components.ActivityCreditType](../../models/components/activitycredittype.md) | :heavy_minus_sign:                                                              | Further detail describing the type of credit                                    | WRITE_OFF                                                                       |
| `Taxable`                                                                       | **bool*                                                                         | :heavy_minus_sign:                                                              | No longer applicable                                                            | false                                                                           |