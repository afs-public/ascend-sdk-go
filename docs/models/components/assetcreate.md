# AssetCreate

The asset being transferred If cash, the asset_id is the currency code (e.g. USD) and the position is the amount


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Identifier`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | The asset identifier                                                     | US37733W2044                                                             |
| `Position`                                                               | [components.PositionCreate](../../models/components/positioncreate.md)   | :heavy_check_mark:                                                       | The position or amount of the asset                                      |                                                                          |
| `Type`                                                                   | [components.AssetCreateType](../../models/components/assetcreatetype.md) | :heavy_check_mark:                                                       | The asset identifier type                                                | CUSIP                                                                    |