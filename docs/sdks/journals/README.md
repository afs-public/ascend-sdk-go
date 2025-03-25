# Journals
(*Journals*)

## Overview

### Available Operations

* [RetrieveCashJournalConstraints](#retrievecashjournalconstraints) - Retrieve Cash Journal Constraints
* [CreateCashJournal](#createcashjournal) - Create Cash Journal
* [GetCashJournal](#getcashjournal) - Get Cash Journal
* [CancelCashJournal](#cancelcashjournal) - Cancel Cash Journal
* [CheckPartyType](#checkpartytype) - Retrieve Cash Journal Party

## RetrieveCashJournalConstraints

Retrieves retirement contribution and distribution constraints for a cash journal transfer

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
    res, err := s.Journals.RetrieveCashJournalConstraints(ctx, components.RetrieveCashJournalConstraintsRequestCreate{
        DestinationAccount: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
        SourceAccount: "accounts/01H8FM6EXVH77SAW3TC8KAWMES",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournalConstraints != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [components.RetrieveCashJournalConstraintsRequestCreate](../../models/components/retrievecashjournalconstraintsrequestcreate.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.RetirementConstraintsRetrieveCashJournalConstraintsResponse](../../models/operations/retirementconstraintsretrievecashjournalconstraintsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCashJournal

Creates a cash journal

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
    res, err := s.Journals.CreateCashJournal(ctx, components.CashJournalCreate{
        ClientTransferID: "113bw03-49f8-4525-934c-560fb39dg2kd",
        DestinationAccount: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
        SourceAccount: "accounts/01H8FM6EXVH77SAW3TC8KAWMES",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.CashJournalCreate](../../models/components/cashjournalcreate.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CashJournalsCreateCashJournalResponse](../../models/operations/cashjournalscreatecashjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCashJournal

Gets an existing cash journal

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
    res, err := s.Journals.GetCashJournal(ctx, "20230817000319")
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `cashJournalID`                                          | *string*                                                 | :heavy_check_mark:                                       | The cashJournal id.                                      | 20230817000319                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CashJournalsGetCashJournalResponse](../../models/operations/cashjournalsgetcashjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelCashJournal

Cancels an existing cash journal

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
    res, err := s.Journals.CancelCashJournal(ctx, "20240717000319", components.CancelCashJournalRequestCreate{
        Name: "cashJournals/20240717000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `cashJournalID`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | The cashJournal id.                                                                                    | 20240717000319                                                                                         |
| `cancelCashJournalRequestCreate`                                                                       | [components.CancelCashJournalRequestCreate](../../models/components/cancelcashjournalrequestcreate.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.CashJournalsCancelCashJournalResponse](../../models/operations/cashjournalscancelcashjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CheckPartyType

Determines whether a potential cash journal will be considered first party or third party given a source and destination account

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
    res, err := s.Journals.CheckPartyType(ctx, components.CheckPartyTypeRequestCreate{
        DestinationAccount: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
        SourceAccount: "accounts/01H8FM6EXVH77SAW3TC8KAWMES",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckPartyTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CheckPartyTypeRequestCreate](../../models/components/checkpartytyperequestcreate.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CashJournalsCheckPartyTypeResponse](../../models/operations/cashjournalscheckpartytyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |