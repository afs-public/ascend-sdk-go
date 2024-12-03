# EntryFee

Object containing more information about the fee being charged


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AdditionalInstructions`                                                  | **string*                                                                 | :heavy_minus_sign:                                                        | Free form text field providing additional information about a transaction | Fee instruction                                                           |
| `Type`                                                                    | [*components.EntryFeeType](../../models/components/entryfeetype.md)       | :heavy_minus_sign:                                                        | Enum providing additional information about the type of fee being charged | LIQUIDITY                                                                 |