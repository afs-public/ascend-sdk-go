# WireWithdrawalStateTaxWithholding

The state tax withholding.


## Fields

| Field                                                                                                                                                   | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             | Example                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Amount`                                                                                                                                                | [*components.WireWithdrawalIraDistributionStateTaxWithholdingAmount](../../models/components/wirewithdrawaliradistributionstatetaxwithholdingamount.md) | :heavy_minus_sign:                                                                                                                                      | Fixed USD amount to withhold for taxes.                                                                                                                 | {<br/>"value": "1.23"<br/>}                                                                                                                             |
| `Percentage`                                                                                                                                            | [*components.WireWithdrawalIraDistributionPercentage](../../models/components/wirewithdrawaliradistributionpercentage.md)                               | :heavy_minus_sign:                                                                                                                                      | Percentage of total disbursement amount to withhold for taxes.                                                                                          | {<br/>"value": "11.25"<br/>}                                                                                                                            |