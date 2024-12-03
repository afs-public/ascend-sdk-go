# ContributionSummary

Regular and rollover contribution amounts for one tax year


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Name`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Retirement account id these contribution amounts are for                | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/contributionSummaries/2023          |
| `RegularAmount`                                                         | [*components.RegularAmount](../../models/components/regularamount.md)   | :heavy_minus_sign:                                                      | Summed contribution amounts throughout the year                         | {<br/>"value": "123.00"<br/>}                                           |
| `RolloverAmount`                                                        | [*components.RolloverAmount](../../models/components/rolloveramount.md) | :heavy_minus_sign:                                                      | Rollover contribution amount                                            | {<br/>"value": "12345.00"<br/>}                                         |
| `TaxYear`                                                               | **int*                                                                  | :heavy_minus_sign:                                                      | Tax year these contribution amounts are for                             | 2023                                                                    |