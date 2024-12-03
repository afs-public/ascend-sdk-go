# AvailableRestriction

Available Restriction on an Account.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Description`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | The description of the restriction.                                            | A full outbound transfer of assets was initiated through ACATS or other means. |
| `DisplayName`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | The human readable representation of the restriction code.                     | Account Transfer Request Received - Full Outbound                              |
| `RestrictionCode`                                                              | **string*                                                                      | :heavy_minus_sign:                                                             | The restriction code.                                                          | ACAT_FULL_OUTBOUND                                                             |