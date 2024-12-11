# FeesAndCredits
(*FeesAndCredits*)

## Overview

### Available Operations

* [CreateFee](#createfee) - Create Fee
* [GetFee](#getfee) - Get Fee
* [CancelFee](#cancelfee) - Cancel Fee
* [CreateCredit](#createcredit) - Create Credit
* [GetCredit](#getcredit) - Get Credit
* [CancelCredit](#cancelcredit) - Cancel Credit

## CreateFee

Create a fee

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
    res, err := s.FeesAndCredits.CreateFee(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.TransfersFeeCreate{
        Amount: components.DecimalCreate{},
        ClientTransferID: "179dcd33-49f8-4615-989c-560fb387c4fd",
        Description: ascendsdkgo.String("Fee charged for platform access"),
        Type: components.TransfersFeeCreateTypePlatform,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersFee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `accountID`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The account id.                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                     |
| `transfersFeeCreate`                                                           | [components.TransfersFeeCreate](../../models/components/transfersfeecreate.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.FeesCreateFeeResponse](../../models/operations/feescreatefeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFee

Retrieve an existing fee

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
    res, err := s.FeesAndCredits.GetFee(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230823123456")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersFee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `feeID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The fee id.                                              | 20230823123456                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FeesGetFeeResponse](../../models/operations/feesgetfeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelFee

Cancel an existing fee

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
    res, err := s.FeesAndCredits.CancelFee(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230823123456", components.CancelFeeRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/fees/20230823123456",
        Reason: ascendsdkgo.String("Reverse fee charge"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersFee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `accountID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | The account id.                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                             |
| `feeID`                                                                                | *string*                                                                               | :heavy_check_mark:                                                                     | The fee id.                                                                            | 20230823123456                                                                         |
| `cancelFeeRequestCreate`                                                               | [components.CancelFeeRequestCreate](../../models/components/cancelfeerequestcreate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.FeesCancelFeeResponse](../../models/operations/feescancelfeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCredit

Create a credit

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
    res, err := s.FeesAndCredits.CreateCredit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.TransfersCreditCreate{
        Amount: components.DecimalCreate{},
        ClientTransferID: "179dcd33-49f8-4615-989c-560fb387c4fd",
        Description: ascendsdkgo.String("Credit given as promotion"),
        Type: components.TransfersCreditCreateTypePromotional,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersCredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `accountID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | The account id.                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                           |
| `transfersCreditCreate`                                                              | [components.TransfersCreditCreate](../../models/components/transferscreditcreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.CreditsCreateCreditResponse](../../models/operations/creditscreatecreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCredit

Retrieve an existing credit

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
    res, err := s.FeesAndCredits.GetCredit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230823123456")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersCredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `creditID`                                               | *string*                                                 | :heavy_check_mark:                                       | The credit id.                                           | 20230823123456                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CreditsGetCreditResponse](../../models/operations/creditsgetcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelCredit

Cancel an existing credit

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
    res, err := s.FeesAndCredits.CancelCredit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230823123456", components.CancelCreditRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/credits/20230823123456",
        Reason: ascendsdkgo.String("Reverse previous credit"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransfersCredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `accountID`                                                                                  | *string*                                                                                     | :heavy_check_mark:                                                                           | The account id.                                                                              | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                   |
| `creditID`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | The credit id.                                                                               | 20230823123456                                                                               |
| `cancelCreditRequestCreate`                                                                  | [components.CancelCreditRequestCreate](../../models/components/cancelcreditrequestcreate.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.CreditsCancelCreditResponse](../../models/operations/creditscancelcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |