# IdentityReportedDeceased

Whether or not the identity has been reported as deceased This is determined by parsing the vendor response for deceased indicators from the SSA Death Master File Equifax-specific indicators: reason codes "90" (SSN Death Indicator) or "SQ" (SSN reported as deceased) null/unset = not checked or unable to determine, false = checked and not deceased, true = deceased


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `IdentityReportedDeceasedDeceasedUnspecified` | DECEASED_UNSPECIFIED                          |
| `IdentityReportedDeceasedDeceased`            | DECEASED                                      |
| `IdentityReportedDeceasedNotDeceased`         | NOT_DECEASED                                  |
| `IdentityReportedDeceasedUnknown`             | UNKNOWN                                       |