# BeneficiaryEnrollmentMetadataCreate

Enrollment metadata for the BENEFICIARY_DESIGNATION enrollment type.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ContingentBeneficiaries`                                                                        | [][components.BeneficiaryCreate](../../models/components/beneficiarycreate.md)                   | :heavy_minus_sign:                                                                               | Contingent Beneficiary list is optional, with a maximum of five contingent beneficiaries.        |
| `PrimaryBeneficiaries`                                                                           | [][components.BeneficiaryCreate](../../models/components/beneficiarycreate.md)                   | :heavy_check_mark:                                                                               | At least one primary beneficiary must be provided, with a maximum of five primary beneficiaries. |