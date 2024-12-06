# ForeignIndividualAccountEnrollmentMetadata

Metadata for the REGISTRATION_INDIVIDUAL_FOREIGN type


## Fields

| Field                                                                                                                                                                                               | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         | Example                                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `DividendReinvestmentPlan`                                                                                                                                                                          | [*components.EnrollmentForeignIndividualAccountEnrollmentMetadataDividendReinvestmentPlan](../../models/components/enrollmentforeignindividualaccountenrollmentmetadatadividendreinvestmentplan.md) | :heavy_minus_sign:                                                                                                                                                                                  | Option to auto-enroll in Dividend Reinvestment; defaults to true                                                                                                                                    | DIVIDEND_REINVESTMENT_ENROLL                                                                                                                                                                        |
| `FdicCashSweep`                                                                                                                                                                                     | [*components.EnrollmentForeignIndividualAccountEnrollmentMetadataFdicCashSweep](../../models/components/enrollmentforeignindividualaccountenrollmentmetadatafdiccashsweep.md)                       | :heavy_minus_sign:                                                                                                                                                                                  | Option to auto-enroll in FDIC cash sweep; defaults to true                                                                                                                                          | FDIC_CASH_SWEEP_ENROLL                                                                                                                                                                              |
| `ForeignNaturalPersonAccountEnrollmentMetadata`                                                                                                                                                     | [*components.ForeignNaturalPersonAccountEnrollmentMetadata](../../models/components/foreignnaturalpersonaccountenrollmentmetadata.md)                                                               | :heavy_minus_sign:                                                                                                                                                                                  | Enrollment metadata for Accounts that have a foreign Legal Natural Person owner.                                                                                                                    |                                                                                                                                                                                                     |