# AccountTransfers
(*AccountTransfers*)

## Overview

### Available Operations

* [CreateTransfer](#createtransfer) - Create Transfer
* [GetTransfer](#gettransfer) - Get Transfer

## CreateTransfer

Creates a transfer

### Example Usage

```go
package main

import(
	"ascend-sdk/models/components"
	ascendsdk "ascend-sdk"
	"context"
	"log"
)

func main() {
    s := ascendsdk.New(
        ascendsdk.WithSecurity(components.Security{
            APIKey: ascendsdk.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
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
        Assets: []components.AssetCreate{
            components.AssetCreate{
                Identifier: "US37733W2044",
                Position: components.PositionCreate{
                    Quantity: components.DecimalCreate{},
                },
                Type: components.AssetCreateTypeCusip,
            },
            components.AssetCreate{
                Identifier: "US37733W2044",
                Position: components.PositionCreate{
                    Quantity: components.DecimalCreate{},
                },
                Type: components.AssetCreateTypeCusip,
            },
            components.AssetCreate{
                Identifier: "US37733W2044",
                Position: components.PositionCreate{
                    Quantity: components.DecimalCreate{},
                },
                Type: components.AssetCreateTypeCusip,
            },
        },
        Comment: ascendsdk.String("From XYZ Brokerage"),
        Deliverer: components.TransferAccountCreate{
            ApexAccountID: "01H8FB90ZRRFWXB4XC2JPJ1D4Y",
            ExternalAccount: components.ExternalAccountCreate{
                AccountNumber: "1234567890",
                ParticipantNumber: "987",
            },
        },
        OriginalControlNumber: ascendsdk.String("20240201123456"),
    }, ascendsdk.String("ABC-123"))
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

## GetTransfer

Gets an existing transfer

### Example Usage

```go
package main

import(
	"ascend-sdk/models/components"
	ascendsdk "ascend-sdk"
	"context"
	"log"
)

func main() {
    s := ascendsdk.New(
        ascendsdk.WithSecurity(components.Security{
            APIKey: ascendsdk.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
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