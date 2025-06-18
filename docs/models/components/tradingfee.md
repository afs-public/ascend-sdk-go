# TradingFee

A fee that applies to an order


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Amount`                                                                           | [*components.TradingFeeAmount](../../models/components/tradingfeeamount.md)        | :heavy_minus_sign:                                                                 | The amount of the fee. This is a monetary value in the same currency as the order. | {<br/>"value": "5.00"<br/>}                                                        |
| `Type`                                                                             | [*components.TradingFeeType](../../models/components/tradingfeetype.md)            | :heavy_minus_sign:                                                                 | The type of fee being specified. Only the type of "BROKER_FEE" is supported.       | BROKER_FEE                                                                         |