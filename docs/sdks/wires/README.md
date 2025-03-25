# Wires
(*Wires*)

## Overview

### Available Operations

* [GetWireDeposit](#getwiredeposit) - Get Wire Deposit
* [CreateWireWithdrawal](#createwirewithdrawal) - Create Wire Withdrawal
* [GetWireWithdrawal](#getwirewithdrawal) - Get Wire Withdrawal
* [CancelWireWithdrawal](#cancelwirewithdrawal) - Cancel Wire Withdrawal

## GetWireDeposit

Gets an existing wire deposit

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
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
    res, err := s.Wires.GetWireDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319")
    if err != nil {
        log.Fatal(err)
    }
    if res.WireDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `wireDepositID`                                          | *string*                                                 | :heavy_check_mark:                                       | The wireDeposit id.                                      | 20230817000319                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.WireDepositsGetWireDepositResponse](../../models/operations/wiredepositsgetwiredepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateWireWithdrawal

Creates a wire withdrawal

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
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
    res, err := s.Wires.CreateWireWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.WireWithdrawalCreate{
        Beneficiary: components.WireWithdrawalBeneficiaryCreate{
            Account: "73849218650987",
        },
        ClientTransferID: "ABC-123",
        RecipientBank: components.WireWithdrawalRecipientBankCreate{
            BankID: components.RecipientBankBankIDCreate{
                ID: "ABNANL2AXXX",
                Type: components.RecipientBankBankIDCreateTypeBic,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `accountID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The account id.                                                                    | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                         |
| `wireWithdrawalCreate`                                                             | [components.WireWithdrawalCreate](../../models/components/wirewithdrawalcreate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.WireWithdrawalsCreateWireWithdrawalResponse](../../models/operations/wirewithdrawalscreatewirewithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetWireWithdrawal

Gets an existing wire withdrawal

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
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
    res, err := s.Wires.GetWireWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319")
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `wireWithdrawalID`                                       | *string*                                                 | :heavy_check_mark:                                       | The wireWithdrawal id.                                   | 20230817000319                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.WireWithdrawalsGetWireWithdrawalResponse](../../models/operations/wirewithdrawalsgetwirewithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelWireWithdrawal

Cancels an existing wire withdrawal

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
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
    res, err := s.Wires.CancelWireWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.CancelWireWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawals/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `accountID`                                                                                                  | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The account id.                                                                                              | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                   |
| `wireWithdrawalID`                                                                                           | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The wireWithdrawal id.                                                                                       | 20230817000319                                                                                               |
| `cancelWireWithdrawalRequestCreate`                                                                          | [components.CancelWireWithdrawalRequestCreate](../../models/components/cancelwirewithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.WireWithdrawalsCancelWireWithdrawalResponse](../../models/operations/wirewithdrawalscancelwirewithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |