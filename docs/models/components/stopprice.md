# StopPrice

The stop price for this order. Only allowed for equities, when the side is SELL.


## Fields

| Field                                                                                                                                                                                                           | Type                                                                                                                                                                                                            | Required                                                                                                                                                                                                        | Description                                                                                                                                                                                                     | Example                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Price`                                                                                                                                                                                                         | [*components.OrderStopPricePrice](../../models/components/orderstoppriceprice.md)                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                              | The stop price which must be greater than zero if provided. For equity orders in the USD currency, up to 2 decimal places are allowed for prices above $1 and up to 4 decimal places for prices at or below $1. | {<br/>"value": "88.132"<br/>}                                                                                                                                                                                   |
| `Type`                                                                                                                                                                                                          | [*components.OrderStopPriceType](../../models/components/orderstoppricetype.md)                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                              | The type of this price, which must PRICE_PER_UNIT for equity orders. (Fixed income and mutual fund assets do not support stop orders.)                                                                          | PRICE_PER_UNIT                                                                                                                                                                                                  |