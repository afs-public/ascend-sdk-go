# IdentityIdentificationCreateType

**Field Dependencies:**

An SSN or ITIN is required when `check_types` is `DATABASE`

Required if `check_types` is `DATABASE`.

Otherwise, must be empty.


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `IdentityIdentificationCreateTypeTypeUnspecified` | TYPE_UNSPECIFIED                                  |
| `IdentityIdentificationCreateTypeSsn`             | SSN                                               |
| `IdentityIdentificationCreateTypePassport`        | PASSPORT                                          |
| `IdentityIdentificationCreateTypeDriversLicense`  | DRIVERS_LICENSE                                   |
| `IdentityIdentificationCreateTypeItin`            | ITIN                                              |
| `IdentityIdentificationCreateTypeIDCard`          | ID_CARD                                           |