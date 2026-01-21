# PositionJournals
(*PositionJournals*)

## Overview

### Available Operations

* [CreatePositionJournal](#createpositionjournal) - Create Position Journal
* [GetPositionJournal](#getpositionjournal) - Get Position Journal
* [CancelPositionJournal](#cancelpositionjournal) - Cancel Position Journal

## CreatePositionJournal

Creates a position journal

### Example Usage

<!-- UsageSnippet language="go" operationID="PositionJournals_CreatePositionJournal" method="post" path="/transfers/v1/positionJournals" -->
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

    res, err := s.PositionJournals.CreatePositionJournal(ctx, components.PositionJournalCreate{
        ClientTransferID: "113bw03-49f8-4525-934c-560fb39dg2kd",
        DestinationAccount: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
        Identifier: "AAPL",
        IdentifierType: components.IdentifierTypeSymbol,
        Quantity: components.DecimalCreate{},
        SourceAccount: "accounts/01H8FM6EXVH77SAW3TC8KAWMES",
        Type: components.PositionJournalCreateTypeReward,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PositionJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.PositionJournalCreate](../../models/components/positionjournalcreate.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PositionJournalsCreatePositionJournalResponse](../../models/operations/positionjournalscreatepositionjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPositionJournal

Gets an existing position journal

### Example Usage

<!-- UsageSnippet language="go" operationID="PositionJournals_GetPositionJournal" method="get" path="/transfers/v1/positionJournals/{positionJournal_id}" -->
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

    res, err := s.PositionJournals.GetPositionJournal(ctx, "20230817000319")
    if err != nil {
        log.Fatal(err)
    }
    if res.PositionJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `positionJournalID`                                      | *string*                                                 | :heavy_check_mark:                                       | The positionJournal id.                                  | 20230817000319                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PositionJournalsGetPositionJournalResponse](../../models/operations/positionjournalsgetpositionjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelPositionJournal

Cancels an existing position journal

### Example Usage

<!-- UsageSnippet language="go" operationID="PositionJournals_CancelPositionJournal" method="post" path="/transfers/v1/positionJournals/{positionJournal_id}:cancel" -->
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

    res, err := s.PositionJournals.CancelPositionJournal(ctx, "20240717000319", components.CancelPositionJournalRequestCreate{
        Name: "positionJournals/20240717000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PositionJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `positionJournalID`                                                                                            | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The positionJournal id.                                                                                        | 20240717000319                                                                                                 |
| `cancelPositionJournalRequestCreate`                                                                           | [components.CancelPositionJournalRequestCreate](../../models/components/cancelpositionjournalrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.PositionJournalsCancelPositionJournalResponse](../../models/operations/positionjournalscancelpositionjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |