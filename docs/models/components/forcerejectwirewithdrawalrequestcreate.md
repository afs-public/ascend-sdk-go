# ForceRejectWireWithdrawalRequestCreate

Request to simulate the rejection of a wire withdrawal


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | The name of the wire withdrawal to force a rejection on            | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawals/20230817000319 |
| `Reason`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | The reason for the reject                                          | Simulate a rejected transfer                                       |