# BeneficiaryEnrollmentMetadata

Metadata for the BENEFICIARY_DESIGNATION enrollment type.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ContingentBeneficiaries`                                                                        | [][components.Beneficiary](../../models/components/beneficiary.md)                               | :heavy_minus_sign:                                                                               | Contingent Beneficiary list is optional, with a maximum of five contingent beneficiaries.        |
| `PrimaryBeneficiaries`                                                                           | [][components.Beneficiary](../../models/components/beneficiary.md)                               | :heavy_minus_sign:                                                                               | At least one primary beneficiary must be provided, with a maximum of five primary beneficiaries. |