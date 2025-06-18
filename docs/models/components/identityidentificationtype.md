# IdentityIdentificationType

**Field Dependencies:**

An SSN or ITIN is required when `check_types` is `DATABASE`

Required if `check_types` is `DATABASE`.

Otherwise, must be empty.


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `IdentityIdentificationTypeTypeUnspecified` | TYPE_UNSPECIFIED                            |
| `IdentityIdentificationTypeSsn`             | SSN                                         |
| `IdentityIdentificationTypePassport`        | PASSPORT                                    |
| `IdentityIdentificationTypeDriversLicense`  | DRIVERS_LICENSE                             |
| `IdentityIdentificationTypeItin`            | ITIN                                        |