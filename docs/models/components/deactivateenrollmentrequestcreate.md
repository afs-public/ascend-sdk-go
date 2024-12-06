# DeactivateEnrollmentRequestCreate

The request for deactivating an Enrollment on an account.


## Fields

| Field                                                                                                                                     | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               | Example                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `EnrollmentID`                                                                                                                            | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | A system-generated unique identifier referencing a single instance of an enrollment;                                                      | 22951598-70e2-46f1-bb32-38e8da7a5cdb                                                                                                      |
| `EnrollmentType`                                                                                                                          | [*components.DeactivateEnrollmentRequestCreateEnrollmentType](../../models/components/deactivateenrollmentrequestcreateenrollmenttype.md) | :heavy_minus_sign:                                                                                                                        | Describes the name of the enrollment; Expressed as an enum                                                                                | CASH_FDIC_CASH_SWEEP                                                                                                                      |