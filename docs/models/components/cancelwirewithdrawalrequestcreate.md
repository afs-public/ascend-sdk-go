# CancelWireWithdrawalRequestCreate

Request to cancel an existing wire withdrawal. The cancel will only succeed if the wire has not made it to the SendingToBank state


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | The name of the wire withdrawal to cancel                          | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawals/20230817000319 |
| `Reason`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | The reason why the wire withdrawal is being canceled               | User Request                                                       |