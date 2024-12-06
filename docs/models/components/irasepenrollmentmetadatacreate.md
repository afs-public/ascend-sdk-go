# IRASEPEnrollmentMetadataCreate

Enrollment metadata for IRA SEP_IRA accounts enrollment type


## Fields

| Field                                                                                                                                                   | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             | Example                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DividendReinvestmentPlan`                                                                                                                              | [*components.IRASEPEnrollmentMetadataCreateDividendReinvestmentPlan](../../models/components/irasepenrollmentmetadatacreatedividendreinvestmentplan.md) | :heavy_minus_sign:                                                                                                                                      | Option to auto-enroll in Dividend Reinvestment; defaults to true                                                                                        | DIVIDEND_REINVESTMENT_ENROLL                                                                                                                            |
| `FdicCashSweep`                                                                                                                                         | [*components.IRASEPEnrollmentMetadataCreateFdicCashSweep](../../models/components/irasepenrollmentmetadatacreatefdiccashsweep.md)                       | :heavy_minus_sign:                                                                                                                                      | Option to auto-enroll in FDIC cash sweep; defaults to true                                                                                              | FDIC_CASH_SWEEP_ENROLL                                                                                                                                  |