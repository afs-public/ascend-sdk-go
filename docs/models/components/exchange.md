# Exchange

Object containing metadata for exchanges


## Fields

| Field                                                                                                                                   | Type                                                                                                                                    | Required                                                                                                                                | Description                                                                                                                             | Example                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `CashRate`                                                                                                                              | [*components.EntryExchangeCashRate](../../models/components/entryexchangecashrate.md)                                                   | :heavy_minus_sign:                                                                                                                      | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which cash will be disbursed to the shareholder        | {<br/>"value": "0.25"<br/>}                                                                                                             |
| `CorporateActionGeneralInformation`                                                                                                     | [*components.EntryExchangeCorporateActionGeneralInformation](../../models/components/entryexchangecorporateactiongeneralinformation.md) | :heavy_minus_sign:                                                                                                                      | Common fields for corporate actions                                                                                                     |                                                                                                                                         |
| `StockRate`                                                                                                                             | [*components.EntryStockRate](../../models/components/entrystockrate.md)                                                                 | :heavy_minus_sign:                                                                                                                      | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which shares will be disbursed to the shareholder      | {<br/>"value": "0.25"<br/>}                                                                                                             |
| `Type`                                                                                                                                  | [*components.EntryExchangeType](../../models/components/entryexchangetype.md)                                                           | :heavy_minus_sign:                                                                                                                      | Corresponds to whether the event is CASH \| STOCK \| CASH_AND_STOCK                                                                     | CASH                                                                                                                                    |