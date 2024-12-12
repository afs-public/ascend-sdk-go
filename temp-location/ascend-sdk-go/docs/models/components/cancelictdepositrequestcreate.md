# CancelIctDepositRequestCreate

Request to cancel an existing ICT deposit


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Name`                                                                              | *string*                                                                            | :heavy_check_mark:                                                                  | Full name of the ICT deposit resource, which contains account id and ICT deposit id | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictDeposits/20240321000472                      |
| `Reason`                                                                            | **string*                                                                           | :heavy_minus_sign:                                                                  | Reason why the ICT deposit is being canceled                                        | User Request                                                                        |