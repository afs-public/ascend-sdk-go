# ForceRejectAchDepositRequestCreate

Request to force rejection of an existing ACH deposit that is pending review. FOR TESTING ONLY!


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Name`                                                         | *string*                                                       | :heavy_check_mark:                                             | The name of the ACH deposit to force reject.                   | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319 |
| `Reason`                                                       | **string*                                                      | :heavy_minus_sign:                                             | Reason why the ACH deposit is being rejected.                  | Simulate a rejected transfer                                   |