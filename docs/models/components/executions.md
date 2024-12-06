# Executions

Details of an individual execution within an order


## Fields

| Field                                                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `AccruedInterestAmount`                                                                                                                                                                                                                                                                                                  | [*components.AccruedInterestAmount](../../models/components/accruedinterestamount.md)                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The amount of accrued interest exchanged in this execution. Will only be present for orders of Fixed Income assets.                                                                                                                                                                                                      | {<br/>"value": "25.97"<br/>}                                                                                                                                                                                                                                                                                             |
| `ExecutedPrices`                                                                                                                                                                                                                                                                                                         | [][components.ExecutedPrice](../../models/components/executedprice.md)                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The prices at which the order was executed. For Equities: there will be one price measured in PRICE_PER_UNIT (using the order currency). For Fixed Income assets: there will always be an entry measured in the PERCENTAGE_OF_PAR (100 X cost / total par value), and there may be additional entries measured in yield. | [<br/>{<br/>"price": {<br/>"value": "94.56"<br/>},<br/>"type": "PRICE_PER_UNIT"<br/>}<br/>]                                                                                                                                                                                                                              |
| `ExecutedTime`                                                                                                                                                                                                                                                                                                           | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The timestamp that this fill was transacted at the market                                                                                                                                                                                                                                                                | {<br/>"nanos": 360000000,<br/>"seconds": 1712081569<br/>}                                                                                                                                                                                                                                                                |
| `GrossCreditAmount`                                                                                                                                                                                                                                                                                                      | [*components.GrossCreditAmount](../../models/components/grosscreditamount.md)                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The net currency amount exchanged in this transaction, in the order currency. Will only be present for orders of Fixed Income assets.                                                                                                                                                                                    | {<br/>"value": "125.97"<br/>}                                                                                                                                                                                                                                                                                            |
| `PrevailingMarketPrice`                                                                                                                                                                                                                                                                                                  | [*components.PrevailingMarketPrice](../../models/components/prevailingmarketprice.md)                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The prevailing market price of the asset, without fees or commissions. Will only be present for orders of Fixed Income assets.                                                                                                                                                                                           | {<br/>"value": "101.45"<br/>}                                                                                                                                                                                                                                                                                            |
| `Quantity`                                                                                                                                                                                                                                                                                                               | [*components.Quantity](../../models/components/quantity.md)                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                       | The quantity of the order. For Equities: measured in shares. For Fixed Income assets: measured in the face value of the currency of the order.                                                                                                                                                                           | {<br/>"value": "3.591"<br/>}                                                                                                                                                                                                                                                                                             |