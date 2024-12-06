# BidAskPrice

The definition of a price value and its calculation method as returned in quote data


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 | Example                                                                     |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Price`                                                                     | [*components.BidAskPricePrice](../../models/components/bidaskpriceprice.md) | :heavy_minus_sign:                                                          | The price value                                                             | {<br/>"value": "97.43"<br/>}                                                |
| `Type`                                                                      | [*components.BidAskPriceType](../../models/components/bidaskpricetype.md)   | :heavy_minus_sign:                                                          | The calculation type of this price                                          | PERCENTAGE_OF_PAR                                                           |