# ActivityRedemptionPartial

Object containing metadata for partial redemption


## Fields

| Field                                                                                                                                                           | Type                                                                                                                                                            | Required                                                                                                                                                        | Description                                                                                                                                                     | Example                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Action`                                                                                                                                                        | [*components.ActivityRedemptionPartialAction](../../models/components/activityredemptionpartialaction.md)                                                       | :heavy_minus_sign:                                                                                                                                              | Corresponds to whether the entry is incoming or outgoing                                                                                                        | INCOMING                                                                                                                                                        |
| `CashRate`                                                                                                                                                      | [*components.ActivityRedemptionPartialCashRate](../../models/components/activityredemptionpartialcashrate.md)                                                   | :heavy_minus_sign:                                                                                                                                              | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which cash will be disbursed to the shareholder                                | {<br/>"value": "0.25"<br/>}                                                                                                                                     |
| `CorporateActionGeneralInformation`                                                                                                                             | [*components.ActivityRedemptionPartialCorporateActionGeneralInformation](../../models/components/activityredemptionpartialcorporateactiongeneralinformation.md) | :heavy_minus_sign:                                                                                                                                              | Common fields for corporate actions                                                                                                                             |                                                                                                                                                                 |
| `PaymentDate`                                                                                                                                                   | [*components.ActivityRedemptionPartialPaymentDate](../../models/components/activityredemptionpartialpaymentdate.md)                                             | :heavy_minus_sign:                                                                                                                                              | The anticipated payment date at the depository                                                                                                                  | {<br/>"day": 14,<br/>"month": 5,<br/>"year": 2024<br/>}                                                                                                         |
| `Quantity`                                                                                                                                                      | [*components.ActivityRedemptionPartialQuantity](../../models/components/activityredemptionpartialquantity.md)                                                   | :heavy_minus_sign:                                                                                                                                              | Corresponds to the position's trade quantity                                                                                                                    | {<br/>"value": "100.00"<br/>}                                                                                                                                   |