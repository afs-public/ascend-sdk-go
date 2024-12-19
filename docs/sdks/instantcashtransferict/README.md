# InstantCashTransferICT
(*InstantCashTransferICT*)

## Overview

### Available Operations

* [CreateIctDeposit](#createictdeposit) - Create ICT Deposit
* [GetIctDeposit](#getictdeposit) - Get ICT Deposit
* [CancelIctDeposit](#cancelictdeposit) - Cancel ICT Deposit
* [CreateIctWithdrawal](#createictwithdrawal) - Create ICT Withdrawal
* [GetIctWithdrawal](#getictwithdrawal) - Get ICT Withdrawal
* [CancelIctWithdrawal](#cancelictwithdrawal) - Cancel ICT Withdrawal
* [LocateIctReport](#locateictreport) - Locate ICT Report

## CreateIctDeposit

Creates an ICT deposit

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
    res, err := s.InstantCashTransferICT.CreateIctDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.IctDepositCreate{
        Amount: components.DecimalCreate{},
        ClientTransferID: "ABC-123",
        Program: components.ProgramDepositOnly,
        TravelRule: components.IctDepositTravelRuleCreate{
            OriginatingInstitution: components.InstitutionCreate{
                AccountID: "<id>",
                Title: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `accountID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | The account id.                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                 |
| `ictDepositCreate`                                                         | [components.IctDepositCreate](../../models/components/ictdepositcreate.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*operations.IctDepositsCreateIctDepositResponse](../../models/operations/ictdepositscreateictdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIctDeposit

Gets an existing ICT deposit

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
    res, err := s.InstantCashTransferICT.GetIctDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472")
    if err != nil {
        log.Fatal(err)
    }
    if res.IctDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `ictDepositID`                                           | *string*                                                 | :heavy_check_mark:                                       | The ictDeposit id.                                       | 20240321000472                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.IctDepositsGetIctDepositResponse](../../models/operations/ictdepositsgetictdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelIctDeposit

Cancels an existing ICT deposit

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
    res, err := s.InstantCashTransferICT.CancelIctDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.CancelIctDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictDeposits/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The account id.                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                           |
| `ictDepositID`                                                                                       | *string*                                                                                             | :heavy_check_mark:                                                                                   | The ictDeposit id.                                                                                   | 20240321000472                                                                                       |
| `cancelIctDepositRequestCreate`                                                                      | [components.CancelIctDepositRequestCreate](../../models/components/cancelictdepositrequestcreate.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.IctDepositsCancelIctDepositResponse](../../models/operations/ictdepositscancelictdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateIctWithdrawal

Creates an ICT withdrawal

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
    res, err := s.InstantCashTransferICT.CreateIctWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.IctWithdrawalCreate{
        ClientTransferID: "20230817000319",
        Program: components.IctWithdrawalCreateProgramBrokerPartner,
        TravelRule: components.IctWithdrawalTravelRuleCreate{
            RecipientInstitution: components.InstitutionCreate{
                AccountID: "<id>",
                Title: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `accountID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | The account id.                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                       |
| `ictWithdrawalCreate`                                                            | [components.IctWithdrawalCreate](../../models/components/ictwithdrawalcreate.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.IctWithdrawalsCreateIctWithdrawalResponse](../../models/operations/ictwithdrawalscreateictwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIctWithdrawal

Gets an existing ICT withdrawal

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
    res, err := s.InstantCashTransferICT.GetIctWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472")
    if err != nil {
        log.Fatal(err)
    }
    if res.IctWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `ictWithdrawalID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ictWithdrawal id.                                    | 20240321000472                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.IctWithdrawalsGetIctWithdrawalResponse](../../models/operations/ictwithdrawalsgetictwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelIctWithdrawal

Cancels an existing ICT withdrawal

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
    res, err := s.InstantCashTransferICT.CancelIctWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.CancelIctWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictWithdrawals/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `accountID`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The account id.                                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                 |
| `ictWithdrawalID`                                                                                          | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The ictWithdrawal id.                                                                                      | 20240321000472                                                                                             |
| `cancelIctWithdrawalRequestCreate`                                                                         | [components.CancelIctWithdrawalRequestCreate](../../models/components/cancelictwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.IctWithdrawalsCancelIctWithdrawalResponse](../../models/operations/ictwithdrawalscancelictwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## LocateIctReport

Returns a signed link pointing to a recon report file for a specific ICT batch.

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
    res, err := s.InstantCashTransferICT.LocateIctReport(ctx, "01H8MCDXH4HYJJAV921BDKCC83", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LocateIctReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `correspondentID`                                                                           | *string*                                                                                    | :heavy_check_mark:                                                                          | The correspondent id.                                                                       | 01H8MCDXH4HYJJAV921BDKCC83                                                                  |
| `batchID`                                                                                   | **string*                                                                                   | :heavy_minus_sign:                                                                          | The id of the ICT batch for which to locate the report.                                     | 24114.108.2b2c1.001                                                                         |
| `programDateFilterProgram`                                                                  | [*operations.ProgramDateFilterProgram](../../models/operations/programdatefilterprogram.md) | :heavy_minus_sign:                                                                          | The ICT program for which to locate the report.                                             | BROKER_PARTNER                                                                              |
| `programDateFilterProcessDate`                                                              | [*components.DateCreate](../../models/components/datecreate.md)                             | :heavy_minus_sign:                                                                          | The process date for which to locate the report.                                            | {<br/>"process_date": {<br/>"day": 30,<br/>"month": 9,<br/>"year": 2023<br/>}<br/>}         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.IctReconReportsLocateIctReportResponse](../../models/operations/ictreconreportslocateictreportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |