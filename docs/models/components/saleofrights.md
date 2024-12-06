# SaleOfRights

Object containing more information about a sale_of_rights


## Fields

| Field                                                                                                                                           | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     | Example                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `CashRate`                                                                                                                                      | [*components.EntrySaleOfRightsCashRate](../../models/components/entrysaleofrightscashrate.md)                                                   | :heavy_minus_sign:                                                                                                                              | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which cash will be disbursed to the shareholder                | {<br/>"value": "0.25"<br/>}                                                                                                                     |
| `CorporateActionGeneralInformation`                                                                                                             | [*components.EntrySaleOfRightsCorporateActionGeneralInformation](../../models/components/entrysaleofrightscorporateactiongeneralinformation.md) | :heavy_minus_sign:                                                                                                                              | Common fields for corporate actions                                                                                                             |                                                                                                                                                 |
| `PaymentDate`                                                                                                                                   | [*components.EntrySaleOfRightsPaymentDate](../../models/components/entrysaleofrightspaymentdate.md)                                             | :heavy_minus_sign:                                                                                                                              | The anticipated payment date at the depository.                                                                                                 | {<br/>"day": 14,<br/>"month": 5,<br/>"year": 2024<br/>}                                                                                         |
| `RecordDate`                                                                                                                                    | [*components.EntrySaleOfRightsRecordDate](../../models/components/entrysaleofrightsrecorddate.md)                                               | :heavy_minus_sign:                                                                                                                              | The date on which positions are recorded in order to calculate entitlement                                                                      | {<br/>"day": 14,<br/>"month": 5,<br/>"year": 2024<br/>}                                                                                         |
| `Settled`                                                                                                                                       | [*components.EntrySaleOfRightsSettled](../../models/components/entrysaleofrightssettled.md)                                                     | :heavy_minus_sign:                                                                                                                              | Corresponds to the position's settled quantity                                                                                                  | {<br/>"value": "0.25"<br/>}                                                                                                                     |