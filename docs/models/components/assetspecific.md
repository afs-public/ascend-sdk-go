# AssetSpecific

Asset Specific fields returned from the request.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `DayTrade`                                                          | [*components.DayTrade](../../models/components/daytrade.md)         | :heavy_minus_sign:                                                  | Day Trade Asset Buying Power figures returned from the request.     |                                                                     |
| `NonDayTrade`                                                       | [*components.NonDayTrade](../../models/components/nondaytrade.md)   | :heavy_minus_sign:                                                  | Non Day Trade Asset Buying Power figures returned from the request. |                                                                     |
| `PriceAmount`                                                       | [*components.PriceAmount](../../models/components/priceamount.md)   | :heavy_minus_sign:                                                  | The current price of the asset in USD returned from the request.    | {<br/>"value": "100.00"<br/>}                                       |