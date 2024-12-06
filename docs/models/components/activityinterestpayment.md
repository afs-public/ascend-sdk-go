# ActivityInterestPayment

Object containing metadata for interest payments


## Fields

| Field                                                                                                                                                       | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 | Example                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CashRate`                                                                                                                                                  | [*components.ActivityInterestPaymentCashRate](../../models/components/activityinterestpaymentcashrate.md)                                                   | :heavy_minus_sign:                                                                                                                                          | The rate (raw value, not a percentage, example: 50% will be .5 in this field) at which cash will be disbursed to the shareholder                            | {<br/>"value": "0.25"<br/>}                                                                                                                                 |
| `CorporateActionGeneralInformation`                                                                                                                         | [*components.ActivityInterestPaymentCorporateActionGeneralInformation](../../models/components/activityinterestpaymentcorporateactiongeneralinformation.md) | :heavy_minus_sign:                                                                                                                                          | Common fields for corporate actions                                                                                                                         |                                                                                                                                                             |
| `PaymentDate`                                                                                                                                               | [*components.ActivityInterestPaymentPaymentDate](../../models/components/activityinterestpaymentpaymentdate.md)                                             | :heavy_minus_sign:                                                                                                                                          | The anticipated payment date at the depository                                                                                                              | {<br/>"day": 14,<br/>"month": 5,<br/>"year": 2024<br/>}                                                                                                     |
| `RecordDate`                                                                                                                                                | [*components.ActivityInterestPaymentRecordDate](../../models/components/activityinterestpaymentrecorddate.md)                                               | :heavy_minus_sign:                                                                                                                                          | The date on which positions are recorded in order to calculate entitlement                                                                                  | {<br/>"day": 14,<br/>"month": 5,<br/>"year": 2024<br/>}                                                                                                     |
| `Settled`                                                                                                                                                   | [*components.ActivityInterestPaymentSettled](../../models/components/activityinterestpaymentsettled.md)                                                     | :heavy_minus_sign:                                                                                                                                          | The accounts settled position for which the corporate action was paid                                                                                       | {<br/>"value": "0.25"<br/>}                                                                                                                                 |