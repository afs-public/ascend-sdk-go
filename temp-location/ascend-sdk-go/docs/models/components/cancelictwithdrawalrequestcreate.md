# CancelIctWithdrawalRequestCreate

Request to cancel an existing ICT withdrawal


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Name`                                                                                    | *string*                                                                                  | :heavy_check_mark:                                                                        | Full name of the ICT withdrawal resource, which contains account id and ICT withdrawal id | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictWithdrawals/20240321000472                         |
| `Reason`                                                                                  | **string*                                                                                 | :heavy_minus_sign:                                                                        | Reason why the ICT withdrawal is being canceled                                           | User Request                                                                              |