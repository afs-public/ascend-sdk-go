# IraSepEnrollmentMetadata

Metadata for the SEP_IRA_REGISTRATION enrollment type


## Fields

| Field                                                                                                                                                           | Type                                                                                                                                                            | Required                                                                                                                                                        | Description                                                                                                                                                     | Example                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DividendReinvestmentPlan`                                                                                                                                      | [*components.EnrollmentIraSepEnrollmentMetadataDividendReinvestmentPlan](../../models/components/enrollmentirasepenrollmentmetadatadividendreinvestmentplan.md) | :heavy_minus_sign:                                                                                                                                              | Option to auto-enroll in Dividend Reinvestment; defaults to true                                                                                                | DIVIDEND_REINVESTMENT_ENROLL                                                                                                                                    |
| `FdicCashSweep`                                                                                                                                                 | [*components.EnrollmentIraSepEnrollmentMetadataFdicCashSweep](../../models/components/enrollmentirasepenrollmentmetadatafdiccashsweep.md)                       | :heavy_minus_sign:                                                                                                                                              | Option to auto-enroll in FDIC cash sweep; defaults to true                                                                                                      | FDIC_CASH_SWEEP_ENROLL                                                                                                                                          |