# ActivityUnitSplit

Object containing metadata for unit splits


## Fields

| Field                                                                                                                                           | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     | Example                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `CorporateActionGeneralInformation`                                                                                                             | [*components.ActivityUnitSplitCorporateActionGeneralInformation](../../models/components/activityunitsplitcorporateactiongeneralinformation.md) | :heavy_minus_sign:                                                                                                                              | Common fields for corporate actions                                                                                                             |                                                                                                                                                 |
| `StockRate`                                                                                                                                     | [*components.ActivityUnitSplitStockRate](../../models/components/activityunitsplitstockrate.md)                                                 | :heavy_minus_sign:                                                                                                                              | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which shares will be disbursed to the shareholder              | {<br/>"value": "0.25"<br/>}                                                                                                                     |