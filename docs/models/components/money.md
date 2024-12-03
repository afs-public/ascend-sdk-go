# Money

Object containing currency/ price information for the trade lot


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CurrencyCode`                                              | **string*                                                   | :heavy_minus_sign:                                          | Currency code of the price                                  | USD                                                         |
| `Price`                                                     | [*components.LotPrice](../../models/components/lotprice.md) | :heavy_minus_sign:                                          | Price of the trade lot                                      | {<br/>"value": "0.25"<br/>}                                 |