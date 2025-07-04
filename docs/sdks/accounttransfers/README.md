# AccountTransfers
(*AccountTransfers*)

## Overview

### Available Operations

* [CreateTransfer](#createtransfer) - Create Transfer
* [ListTransfers](#listtransfers) - List Transfers
* [AcceptTransfer](#accepttransfer) - Accept Transfer
* [RejectTransfer](#rejecttransfer) - Reject Transfer
* [GetTransfer](#gettransfer) - Get Transfer

## CreateTransfer

Creates a transfer

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
    res, err := s.AccountTransfers.CreateTransfer(ctx, "00000000-0000-0000-0000-000000000002", "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.TransferCreate{
        Deliverer: components.TransferAccountCreate{},
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AcatsTransfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                | Example                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |                                                                                                                                                                                                            |
| `correspondentID`                                                                                                                                                                                          | *string*                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                         | The correspondent id.                                                                                                                                                                                      | 00000000-0000-0000-0000-000000000002                                                                                                                                                                       |
| `accountID`                                                                                                                                                                                                | *string*                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                         | The account id.                                                                                                                                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                 |
| `transferCreate`                                                                                                                                                                                           | [components.TransferCreate](../../models/components/transfercreate.md)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                         | N/A                                                                                                                                                                                                        |                                                                                                                                                                                                            |
| `requestID`                                                                                                                                                                                                | **string*                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                         | A client-specified ID for the account transfer; no specific pattern is imposed. This field is used for idempotency to ensure that repeated requests with the same ID do not result in duplicate transfers. | ABC-123                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                         | The options for this request.                                                                                                                                                                              |                                                                                                                                                                                                            |

### Response

**[*operations.AccountTransfersCreateTransferResponse](../../models/operations/accounttransferscreatetransferresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTransfers

Lists existing transfers using a CEL filter.

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
    res, err := s.AccountTransfers.ListTransfers(ctx, operations.AccountTransfersListTransfersRequest{
        CorrespondentID: "00000000-0000-0000-0000-000000000002",
        AccountID: "01FAKEACCOUNT",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTransfersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.AccountTransfersListTransfersRequest](../../models/operations/accounttransferslisttransfersrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.AccountTransfersListTransfersResponse](../../models/operations/accounttransferslisttransfersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AcceptTransfer

Accept one side (incoming/outgoing) of an internal Ascend transfer. When both the incoming side and outgoing side of the transfer have been accepted then bookkeeping is done immediately. If neither, or only one side of a transfer is accepted, then both sides of the internal perform bookkeeping one full settlement day after they went into the bookkeeping queue.

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
    res, err := s.AccountTransfers.AcceptTransfer(ctx, "00000000-0000-0000-0000-000000000002", "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "00000000-0000-0000-0000-000000000000", components.AcceptTransferRequestCreate{
        Name: "correspondents/00000000-0000-0000-0000-000000000002/accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/transfers/00000000-0000-0000-0000-000000000000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AcceptTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `correspondentID`                                                                                | *string*                                                                                         | :heavy_check_mark:                                                                               | The correspondent id.                                                                            | 00000000-0000-0000-0000-000000000002                                                             |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                       |
| `transferID`                                                                                     | *string*                                                                                         | :heavy_check_mark:                                                                               | The transfer id.                                                                                 | 00000000-0000-0000-0000-000000000000                                                             |
| `acceptTransferRequestCreate`                                                                    | [components.AcceptTransferRequestCreate](../../models/components/accepttransferrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AccountTransfersAcceptTransferResponse](../../models/operations/accounttransfersaccepttransferresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RejectTransfer

Reject one side (incoming/outgoing) of an internal Ascend transfer. Rejecting one side automatically moves the contra side of the transfer to be rejected as well.

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
    res, err := s.AccountTransfers.RejectTransfer(ctx, "00000000-0000-0000-0000-000000000002", "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "00000000-0000-0000-0000-000000000000", components.RejectTransferRequestCreate{
        Name: "correspondents/00000000-0000-0000-0000-000000000002/accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/transfers/00000000-0000-0000-0000-000000000000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RejectTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `correspondentID`                                                                                | *string*                                                                                         | :heavy_check_mark:                                                                               | The correspondent id.                                                                            | 00000000-0000-0000-0000-000000000002                                                             |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                       |
| `transferID`                                                                                     | *string*                                                                                         | :heavy_check_mark:                                                                               | The transfer id.                                                                                 | 00000000-0000-0000-0000-000000000000                                                             |
| `rejectTransferRequestCreate`                                                                    | [components.RejectTransferRequestCreate](../../models/components/rejecttransferrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AccountTransfersRejectTransferResponse](../../models/operations/accounttransfersrejecttransferresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTransfer

Gets an existing transfer

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
    res, err := s.AccountTransfers.GetTransfer(ctx, "00000000-0000-0000-0000-000000000002", "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "00000000-0000-0000-0000-000000000000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AcatsTransfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `correspondentID`                                        | *string*                                                 | :heavy_check_mark:                                       | The correspondent id.                                    | 00000000-0000-0000-0000-000000000002                     |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `transferID`                                             | *string*                                                 | :heavy_check_mark:                                       | The transfer id.                                         | 00000000-0000-0000-0000-000000000000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountTransfersGetTransferResponse](../../models/operations/accounttransfersgettransferresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |