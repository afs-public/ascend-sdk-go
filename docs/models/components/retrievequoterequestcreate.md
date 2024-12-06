# RetrieveQuoteRequestCreate

Request object for retrieving fixed income quotes


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `AssetType`                                                                                                                                                                                                                                                                                                                                                  | [components.RetrieveQuoteRequestCreateAssetType](../../models/components/retrievequoterequestcreateassettype.md)                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The type of asset being evaluated in cost calculations                                                                                                                                                                                                                                                                                                       | FIXED_INCOME                                                                                                                                                                                                                                                                                                                                                 |
| `BrokerCapacity`                                                                                                                                                                                                                                                                                                                                             | [*components.RetrieveQuoteRequestCreateBrokerCapacity](../../models/components/retrievequoterequestcreatebrokercapacity.md)                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | Capacity used in determining bid and ask prices. Defaults to "AGENCY" if no value specified.                                                                                                                                                                                                                                                                 | AGENCY                                                                                                                                                                                                                                                                                                                                                       |
| `Identifier`                                                                                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | Identifier of the asset (of the type specified in `identifier_type`).                                                                                                                                                                                                                                                                                        | 3.78331e+07                                                                                                                                                                                                                                                                                                                                                  |
| `IdentifierType`                                                                                                                                                                                                                                                                                                                                             | [components.RetrieveQuoteRequestCreateIdentifierType](../../models/components/retrievequoterequestcreateidentifiertype.md)                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The identifier type of the asset being sought                                                                                                                                                                                                                                                                                                                | CUSIP                                                                                                                                                                                                                                                                                                                                                        |
| `MinimumQuantity`                                                                                                                                                                                                                                                                                                                                            | [*components.DecimalCreate](../../models/components/decimalcreate.md)                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].<br/><br/> [BigDecimal]:<br/> https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html<br/> [decimal.Decimal]: https://docs.python.org/3/library/decimal.html |                                                                                                                                                                                                                                                                                                                                                              |
| `Parent`                                                                                                                                                                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The parent resource where this order will be created. Format: accounts/{account_id}                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                              |