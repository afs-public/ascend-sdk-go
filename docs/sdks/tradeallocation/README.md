# TradeAllocation
(*TradeAllocation*)

## Overview

### Available Operations

* [CreateTradeAllocation](#createtradeallocation) - Create Trade Allocation
* [GetTradeAllocation](#gettradeallocation) - Get Trade Allocation
* [CancelTradeAllocation](#canceltradeallocation) - Cancel Trade Allocation
* [RebookTradeAllocation](#rebooktradeallocation) - Rebook Trade Allocation

## CreateTradeAllocation

Creates a new trade allocation. These are used to allocate or distribute positions between Apex accounts.

 Upon success, returns the created trade allocation and its enriched details.

### Example Usage

<!-- UsageSnippet language="go" operationID="Booking_CreateTradeAllocation" method="post" path="/booking/v1/accounts/{account_id}/tradeAllocations" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/types"
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

    res, err := s.TradeAllocation.CreateTradeAllocation(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", components.TradeAllocationCreate{
        AssetType: components.TradeAllocationCreateAssetTypeEquity,
        BrokerCapacity: components.TradeAllocationCreateBrokerCapacityAgency,
        ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
        FromAccountID: "01HASWB2DTMRT3DAM45P56J2H3",
        Identifier: "AAPL",
        IdentifierType: components.TradeAllocationCreateIdentifierTypeSymbol,
        Price: components.DecimalCreate{},
        Quantity: components.DecimalCreate{},
        SourceApplication: "Trading-App",
        ToAccountID: "02HASWB2DTMRT3DAM45P56J2T2",
        ToSide: components.ToSideBuy,
    }, ascendsdkgo.String("8a0d35c0-428c-439e-9b03-b611530fe06f"))
    if err != nil {
        log.Fatal(err)
    }
    if res.TradeAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 | Example                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |                                                                                                                             |
| `accountID`                                                                                                                 | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | The account id.                                                                                                             | 01FAKEACCOUNT1TYKWEYRH8S2K                                                                                                  |
| `tradeAllocationCreate`                                                                                                     | [components.TradeAllocationCreate](../../models/components/tradeallocationcreate.md)                                        | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |                                                                                                                             |
| `requestID`                                                                                                                 | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | A globally unique UUID that is specific to the request. This id is used to prevent duplicate requests from being processed. | 8a0d35c0-428c-439e-9b03-b611530fe06f                                                                                        |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |                                                                                                                             |

### Response

**[*operations.BookingCreateTradeAllocationResponse](../../models/operations/bookingcreatetradeallocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404, 409 | application/json   |
| sdkerrors.Status   | 500, 503, 504      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTradeAllocation

Retrieves a trade allocation and its details.

 Upon successful submission, returns the trade allocation details.

### Example Usage

<!-- UsageSnippet language="go" operationID="Booking_GetTradeAllocation" method="get" path="/booking/v1/accounts/{account_id}/tradeAllocations/{tradeAllocation_id}" -->
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

    res, err := s.TradeAllocation.GetTradeAllocation(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM")
    if err != nil {
        log.Fatal(err)
    }
    if res.TradeAllocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 02HASWB2DTMRT3DAM45P56J2T2                               |
| `tradeAllocationID`                                      | *string*                                                 | :heavy_check_mark:                                       | The tradeAllocation id.                                  | 01J0XX2KDN3M9QKFKRE2HYSCQM                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BookingGetTradeAllocationResponse](../../models/operations/bookinggettradeallocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503, 504      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelTradeAllocation

Cancel a trade allocation using the original trade_allocation_id.

 Upon successful submission, returns an empty response. CancelTradeAllocation will either cancel everything, or nothing at all if a failure occurs.

### Example Usage

<!-- UsageSnippet language="go" operationID="Booking_CancelTradeAllocation" method="post" path="/booking/v1/accounts/{account_id}/tradeAllocations/{tradeAllocation_id}:cancel" -->
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

    res, err := s.TradeAllocation.CancelTradeAllocation(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", components.CancelTradeAllocationRequestCreate{
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/tradeAllocations/01J0XX2KDN3M9QKFKRE2HYSCQM",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelTradeAllocationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 02HASWB2DTMRT3DAM45P56J2T2                                                                                     |
| `tradeAllocationID`                                                                                            | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The tradeAllocation id.                                                                                        | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                                     |
| `cancelTradeAllocationRequestCreate`                                                                           | [components.CancelTradeAllocationRequestCreate](../../models/components/canceltradeallocationrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.BookingCancelTradeAllocationResponse](../../models/operations/bookingcanceltradeallocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503, 504      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RebookTradeAllocation

Rebook a trade allocation by the original trade_allocation_id. The allocation is rebooked by canceling the original allocation and creating a new one with the provided details.

 Upon successful submission, returns both the original and new allocation, as separate resources.

### Example Usage

<!-- UsageSnippet language="go" operationID="Booking_RebookTradeAllocation" method="post" path="/booking/v1/accounts/{account_id}/tradeAllocations/{tradeAllocation_id}:rebook" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/types"
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

    res, err := s.TradeAllocation.RebookTradeAllocation(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", components.RebookTradeAllocationRequestCreate{
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/tradeAllocations/01J0XX2KDN3M9QKFKRE2HYSCQM",
        RequestID: "8a0d35c0-428c-439e-9b03-b611530fe06f",
        TradeAllocation: components.TradeAllocationCreate{
            AssetType: components.TradeAllocationCreateAssetTypeEquity,
            BrokerCapacity: components.TradeAllocationCreateBrokerCapacityAgency,
            ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
            FromAccountID: "01HASWB2DTMRT3DAM45P56J2H3",
            Identifier: "AAPL",
            IdentifierType: components.TradeAllocationCreateIdentifierTypeSymbol,
            Price: components.DecimalCreate{},
            Quantity: components.DecimalCreate{},
            SourceApplication: "Trading-App",
            ToAccountID: "02HASWB2DTMRT3DAM45P56J2T2",
            ToSide: components.ToSideBuy,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RebookTradeAllocationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 02HASWB2DTMRT3DAM45P56J2T2                                                                                     |
| `tradeAllocationID`                                                                                            | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The tradeAllocation id.                                                                                        | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                                     |
| `rebookTradeAllocationRequestCreate`                                                                           | [components.RebookTradeAllocationRequestCreate](../../models/components/rebooktradeallocationrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.BookingRebookTradeAllocationResponse](../../models/operations/bookingrebooktradeallocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503, 504      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |