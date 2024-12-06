# EnrollmentPlannedActivity

Details the customer's intended trading and banking-related activities at the time of account application; informs risk checks and forms a baseline for anomalous activity detection


## Fields

| Field                                                                                                                                                   | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             | Example                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ForeignBondTradingDetails`                                                                                                                             | [*components.EnrollmentForeignBondTradingDetails](../../models/components/enrollmentforeignbondtradingdetails.md)                                       | :heavy_minus_sign:                                                                                                                                      | The foreign bond trading countries details                                                                                                              |                                                                                                                                                         |
| `LowPricedSecurities`                                                                                                                                   | [*components.EnrollmentLowPricedSecurities](../../models/components/enrollmentlowpricedsecurities.md)                                                   | :heavy_minus_sign:                                                                                                                                      | The account anticipates trading in securities trading for less than $5 per share and are typically traded over-the-counter (OTC) or through pink sheets |                                                                                                                                                         |
| `PrimaryAccountActivityType`                                                                                                                            | [*components.EnrollmentPrimaryAccountActivityType](../../models/components/enrollmentprimaryaccountactivitytype.md)                                     | :heavy_minus_sign:                                                                                                                                      | The primary account activity type                                                                                                                       | ACTIVE_TRADING                                                                                                                                          |
| `WithdrawalFrequency`                                                                                                                                   | [*components.EnrollmentWithdrawalFrequency](../../models/components/enrollmentwithdrawalfrequency.md)                                                   | :heavy_minus_sign:                                                                                                                                      | The frequency by which cash is anticipated to be withdrawn from the account                                                                             | FREQUENT                                                                                                                                                |