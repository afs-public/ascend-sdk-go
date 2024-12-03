# Journals
(*Journals*)

## Overview

### Available Operations

* [CreateCashJournal](#createcashjournal) - Create Cash Journal
* [GetCashJournal](#getcashjournal) - Get Cash Journal
* [CancelCashJournal](#cancelcashjournal) - Cancel Cash Journal

## CreateCashJournal

Creates a cash journal

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
    res, err := s.Journals.CreateCashJournal(ctx, components.CashJournalCreate{
        ClientTransferID: "113bw03-49f8-4525-934c-560fb39dg2kd",
        DestinationAccount: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
        FullDisbursement: ascendsdk.Bool(false),
        RetirementContribution: &components.RetirementContributionCreate{
            TaxYear: 2024,
            Type: components.RetirementContributionCreateTypeRegular,
        },
        RetirementDistribution: &components.RetirementDistributionCreate{
            Type: components.RetirementDistributionCreateTypeNormal,
        },
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
    res, err := s.Journals.CancelCashJournal(ctx, "20240717000319", components.CancelCashJournalRequestCreate{
        Name: "cashJournals/20240717000319",
        Reason: ascendsdk.String("Cancel journal"),
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