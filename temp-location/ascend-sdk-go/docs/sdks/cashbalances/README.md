# CashBalances
(*CashBalances*)

## Overview

### Available Operations

* [CalculateCashBalance](#calculatecashbalance) - Get Cash Balance

## CalculateCashBalance

Calculates the cash balance for an account.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.CashBalances.CalculateCashBalance(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", operations.MechanismAch.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.CalculateCashBalanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `accountID`                                                                                                                  | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The account id.                                                                                                              | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                   |
| `mechanism`                                                                                                                  | [*operations.Mechanism](../../models/operations/mechanism.md)                                                                | :heavy_minus_sign:                                                                                                           | The withdraw mechanism to calculate the balance for. The mechanism determines what account activity will affect the balance. | ACH                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*operations.CashBalancesCalculateCashBalanceResponse](../../models/operations/cashbalancescalculatecashbalanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |