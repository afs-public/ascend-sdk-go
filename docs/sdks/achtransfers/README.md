# ACHTransfers
(*ACHTransfers*)

## Overview

### Available Operations

* [CreateAchDeposit](#createachdeposit) - Create ACH Deposit
* [GetAchDeposit](#getachdeposit) - Get ACH Deposit
* [CancelAchDeposit](#cancelachdeposit) - Cancel ACH Deposit
* [CreateAchWithdrawal](#createachwithdrawal) - Create ACH Withdrawal
* [GetAchWithdrawal](#getachwithdrawal) - Get ACH Withdrawal
* [CancelAchWithdrawal](#cancelachwithdrawal) - Cancel ACH Withdrawal

## CreateAchDeposit

Creates an ACH deposit.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_CreateAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.CreateAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.AchDepositCreate{
        Amount: components.DecimalCreate{},
        BankRelationship: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
        ClientTransferID: "179dcd33-49f8-4615-989c-560fb387c4fd",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `accountID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | The account id.                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                 |
| `achDepositCreate`                                                         | [components.AchDepositCreate](../../models/components/achdepositcreate.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.AchDepositsCreateAchDepositResponse](../../models/operations/achdepositscreateachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAchDeposit

Gets an existing ACH deposit.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_GetAchDeposit" method="get" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.GetAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319")
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `achDepositID`                                           | *string*                                                 | :heavy_check_mark:                                       | The achDeposit id.                                       | 20230817000319                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AchDepositsGetAchDepositResponse](../../models/operations/achdepositsgetachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelAchDeposit

Cancels an existing ACH deposit.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_CancelAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}:cancel" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.CancelAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.CancelAchDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The account id.                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                           |
| `achDepositID`                                                                                       | *string*                                                                                             | :heavy_check_mark:                                                                                   | The achDeposit id.                                                                                   | 20230817000319                                                                                       |
| `cancelAchDepositRequestCreate`                                                                      | [components.CancelAchDepositRequestCreate](../../models/components/cancelachdepositrequestcreate.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.AchDepositsCancelAchDepositResponse](../../models/operations/achdepositscancelachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAchWithdrawal

Creates an ACH withdrawal.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_CreateAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.CreateAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.AchWithdrawalCreate{
        BankRelationship: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
        ClientTransferID: "179dcd33-49f8-4615-989c-560fb387c4fd",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `accountID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | The account id.                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                       |
| `achWithdrawalCreate`                                                            | [components.AchWithdrawalCreate](../../models/components/achwithdrawalcreate.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.AchWithdrawalsCreateAchWithdrawalResponse](../../models/operations/achwithdrawalscreateachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAchWithdrawal

Gets an existing ACH withdrawal.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_GetAchWithdrawal" method="get" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.GetAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726")
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `achWithdrawalID`                                        | *string*                                                 | :heavy_check_mark:                                       | The achWithdrawal id.                                    | 20230620500726                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AchWithdrawalsGetAchWithdrawalResponse](../../models/operations/achwithdrawalsgetachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelAchWithdrawal

Cancels an existing ACH withdrawal.

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_CancelAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}:cancel" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.ACHTransfers.CancelAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726", components.CancelAchWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawals/20230620500726",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `accountID`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The account id.                                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                 |
| `achWithdrawalID`                                                                                          | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The achWithdrawal id.                                                                                      | 20230620500726                                                                                             |
| `cancelAchWithdrawalRequestCreate`                                                                         | [components.CancelAchWithdrawalRequestCreate](../../models/components/cancelachwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.AchWithdrawalsCancelAchWithdrawalResponse](../../models/operations/achwithdrawalscancelachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |