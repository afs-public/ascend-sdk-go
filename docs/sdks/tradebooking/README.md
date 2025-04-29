# TradeBooking
(*TradeBooking*)

## Overview

### Available Operations

* [CreateTrade](#createtrade) - Create Trade
* [GetTrade](#gettrade) - Get Trade
* [CompleteTrade](#completetrade) - Complete Trade
* [CancelTrade](#canceltrade) - Cancel Trade
* [RebookTrade](#rebooktrade) - Rebook Trade
* [CreateExecution](#createexecution) - Create Execution
* [GetExecution](#getexecution) - Get Execution
* [CancelExecution](#cancelexecution) - Cancel Execution
* [RebookExecution](#rebookexecution) - Rebook Execution

## CreateTrade

Creates a trade with one or more executions. Combination of (account_id, client_order_id, and the process_date (determined by Booking service)) determines the uniqueness of the trade.

 Upon successful submission, returns the created trade and its details including Booking service enriched details.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/types"
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
    res, err := s.TradeBooking.CreateTrade(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", components.TradeCreate{
        AccountID: "02HASWB2DTMRT3DAM45P56J2T2",
        BrokerCapacity: components.TradeCreateBrokerCapacityAgency,
        ClientOrderID: "00be5285-0623-4560-8c58-f05af2c56ba0",
        Executions: []components.ExecutionCreate{
            components.ExecutionCreate{
                ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                ExternalID: "0H06HAP3A3Y",
                Price: components.DecimalCreate{},
                Quantity: components.DecimalCreate{},
            },
            components.ExecutionCreate{
                ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                ExternalID: "0H06HAP3A3Y",
                Price: components.DecimalCreate{},
                Quantity: components.DecimalCreate{},
            },
            components.ExecutionCreate{
                ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                ExternalID: "0H06HAP3A3Y",
                Price: components.DecimalCreate{},
                Quantity: components.DecimalCreate{},
            },
        },
        Identifier: "AAPL",
        IdentifierType: components.TradeCreateIdentifierTypeSymbol,
        RouteType: components.RouteTypeMngd,
        Side: components.TradeCreateSideBuy,
        SourceApplication: "Trading-App",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Trade != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `accountID`                                                      | *string*                                                         | :heavy_check_mark:                                               | The account id.                                                  | 01FAKEACCOUNT1TYKWEYRH8S2K                                       |
| `tradeCreate`                                                    | [components.TradeCreate](../../models/components/tradecreate.md) | :heavy_check_mark:                                               | N/A                                                              |                                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.BookingCreateTradeResponse](../../models/operations/bookingcreatetraderesponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 409, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetTrade

Gets a trade and all executions by trade_id.

 Upon successful submission, returns the trade details and all the execution by trade_id.

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
    res, err := s.TradeBooking.GetTrade(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", "01FAKETRADEIDPROVIDEDFROMCREATETRADE")
    if err != nil {
        log.Fatal(err)
    }
    if res.Trade != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01FAKEACCOUNT1TYKWEYRH8S2K                               |
| `tradeID`                                                | *string*                                                 | :heavy_check_mark:                                       | The trade id.                                            | 01FAKETRADEIDPROVIDEDFROMCREATETRADE                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BookingGetTradeResponse](../../models/operations/bookinggettraderesponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CompleteTrade

Complete a Trade by closing and generating any fees and withholdings if necessary. Once this endpoint returns an OK, the combination of details that generated the Trade (account_id, client_order_id, and the process_date) cannot be reused.

 Upon successful submission, returns completed trade details and all the executions. Trades that are left open will be automatically closed nightly before Ledger's EOD.

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
    res, err := s.TradeBooking.CompleteTrade(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", components.CompleteTradeRequestCreate{
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/trades/01J0XX2KDN3M9QKFKRE2HYSCQM",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CompleteTradeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `accountID`                                                                                    | *string*                                                                                       | :heavy_check_mark:                                                                             | The account id.                                                                                | 02HASWB2DTMRT3DAM45P56J2T2                                                                     |
| `tradeID`                                                                                      | *string*                                                                                       | :heavy_check_mark:                                                                             | The trade id.                                                                                  | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                     |
| `completeTradeRequestCreate`                                                                   | [components.CompleteTradeRequestCreate](../../models/components/completetraderequestcreate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.BookingCompleteTradeResponse](../../models/operations/bookingcompletetraderesponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CancelTrade

Cancel a trade and all the executions using the original trade_id. CancelTrade will either cancel everything, or nothing at all if a failure occurs.

 Upon successful submission, returns an empty response.

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
    res, err := s.TradeBooking.CancelTrade(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", "01FAKETRADEIDPROVIDEDFROMCREATETRADE", components.CancelTradeRequestCreate{
        Name: "accounts/01FAKEACCOUNT1TYKWEYRH8S2K/trades/01FAKETRADEIDPROVIDEDFROMCREATETRADE",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelTradeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The account id.                                                                            | 01FAKEACCOUNT1TYKWEYRH8S2K                                                                 |
| `tradeID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | The trade id.                                                                              | 01FAKETRADEIDPROVIDEDFROMCREATETRADE                                                       |
| `cancelTradeRequestCreate`                                                                 | [components.CancelTradeRequestCreate](../../models/components/canceltraderequestcreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.BookingCancelTradeResponse](../../models/operations/bookingcanceltraderesponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## RebookTrade

Rebook a trade by the original trade_id. The entire original trade's executions are rebooked using the executions provided in the request. If applicable, fees and backup withholdings will be re-calculated.

 Upon successful submission, returns the rebooked trade details and all the executions.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/types"
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
    res, err := s.TradeBooking.RebookTrade(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", components.RebookTradeRequestCreate{
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/trades/01J0XX2KDN3M9QKFKRE2HYSCQM",
        Trade: components.TradeCreate{
            AccountID: "02HASWB2DTMRT3DAM45P56J2T2",
            BrokerCapacity: components.TradeCreateBrokerCapacityAgency,
            ClientOrderID: "00be5285-0623-4560-8c58-f05af2c56ba0",
            Executions: []components.ExecutionCreate{
                components.ExecutionCreate{
                    ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                    ExternalID: "0H06HAP3A3Y",
                    Price: components.DecimalCreate{},
                    Quantity: components.DecimalCreate{},
                },
                components.ExecutionCreate{
                    ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                    ExternalID: "0H06HAP3A3Y",
                    Price: components.DecimalCreate{},
                    Quantity: components.DecimalCreate{},
                },
                components.ExecutionCreate{
                    ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
                    ExternalID: "0H06HAP3A3Y",
                    Price: components.DecimalCreate{},
                    Quantity: components.DecimalCreate{},
                },
            },
            Identifier: "AAPL",
            IdentifierType: components.TradeCreateIdentifierTypeSymbol,
            RouteType: components.RouteTypeMngd,
            Side: components.TradeCreateSideBuy,
            SourceApplication: "Trading-App",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RebookTradeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The account id.                                                                            | 02HASWB2DTMRT3DAM45P56J2T2                                                                 |
| `tradeID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | The trade id.                                                                              | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                 |
| `rebookTradeRequestCreate`                                                                 | [components.RebookTradeRequestCreate](../../models/components/rebooktraderequestcreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.BookingRebookTradeResponse](../../models/operations/bookingrebooktraderesponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CreateExecution

Create a new execution under an existing trade that is open.

 Upon successful submission, returns the created execution and its details.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/types"
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
    res, err := s.TradeBooking.CreateExecution(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", "01FAKETRADEIDPROVIDEDFROMCREATETRADE", components.ExecutionCreate{
        ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
        ExternalID: "0H06HAP3A3Y",
        Price: components.DecimalCreate{},
        Quantity: components.DecimalCreate{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Execution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `accountID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | The account id.                                                          | 01FAKEACCOUNT1TYKWEYRH8S2K                                               |
| `tradeID`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | The trade id.                                                            | 01FAKETRADEIDPROVIDEDFROMCREATETRADE                                     |
| `executionCreate`                                                        | [components.ExecutionCreate](../../models/components/executioncreate.md) | :heavy_check_mark:                                                       | N/A                                                                      |                                                                          |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.BookingCreateExecutionResponse](../../models/operations/bookingcreateexecutionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetExecution

Gets an execution by execution_id.

 Upon successful submission, returns the execution details by execution_id.

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
    res, err := s.TradeBooking.GetExecution(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", "01FAKETRADEIDPROVIDEDFROMCREATETRADE", "01FAKEEXECUTONIDPROVIDEDFROMBOOKINGAPI")
    if err != nil {
        log.Fatal(err)
    }
    if res.Execution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01FAKEACCOUNT1TYKWEYRH8S2K                               |
| `tradeID`                                                | *string*                                                 | :heavy_check_mark:                                       | The trade id.                                            | 01FAKETRADEIDPROVIDEDFROMCREATETRADE                     |
| `executionID`                                            | *string*                                                 | :heavy_check_mark:                                       | The execution id.                                        | 01FAKEEXECUTONIDPROVIDEDFROMBOOKINGAPI                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BookingGetExecutionResponse](../../models/operations/bookinggetexecutionresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CancelExecution

Cancel an execution using the original execution_id. If applicable, fees and backup withholdings will be re-calculated.

 Upon successful submission, returns the execution that was canceled.

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
    res, err := s.TradeBooking.CancelExecution(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", "02G0XX2KDN3M9QKFKRE2HYSCMY", components.CancelExecutionRequestCreate{
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/trades/01J0XX2KDN3M9QKFKRE2HYSCQM/executions/02G0XX2KDN3M9QKFKRE2HYSCMY",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelExecutionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `accountID`                                                                                        | *string*                                                                                           | :heavy_check_mark:                                                                                 | The account id.                                                                                    | 02HASWB2DTMRT3DAM45P56J2T2                                                                         |
| `tradeID`                                                                                          | *string*                                                                                           | :heavy_check_mark:                                                                                 | The trade id.                                                                                      | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                         |
| `executionID`                                                                                      | *string*                                                                                           | :heavy_check_mark:                                                                                 | The execution id.                                                                                  | 02G0XX2KDN3M9QKFKRE2HYSCMY                                                                         |
| `cancelExecutionRequestCreate`                                                                     | [components.CancelExecutionRequestCreate](../../models/components/cancelexecutionrequestcreate.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.BookingCancelExecutionResponse](../../models/operations/bookingcancelexecutionresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## RebookExecution

Rebook an execution by the original execution_id. If applicable, fees and backup withholdings will be re-calculated.

 Upon successful submission, returns the rebooked execution details.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/types"
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
    res, err := s.TradeBooking.RebookExecution(ctx, "02HASWB2DTMRT3DAM45P56J2T2", "01J0XX2KDN3M9QKFKRE2HYSCQM", "02G0XX2KDN3M9QKFKRE2HYSCMY", components.RebookExecutionRequestCreate{
        Execution: components.ExecutionCreate{
            ExecutionTime: types.MustNewTimeFromString("2024-07-17T12:00:00Z"),
            ExternalID: "0H06HAP3A3Y",
            Price: components.DecimalCreate{},
            Quantity: components.DecimalCreate{},
        },
        Name: "accounts/02HASWB2DTMRT3DAM45P56J2T2/trades/01J0XX2KDN3M9QKFKRE2HYSCQM/executions/02G0XX2KDN3M9QKFKRE2HYSCMY",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RebookExecutionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `accountID`                                                                                        | *string*                                                                                           | :heavy_check_mark:                                                                                 | The account id.                                                                                    | 02HASWB2DTMRT3DAM45P56J2T2                                                                         |
| `tradeID`                                                                                          | *string*                                                                                           | :heavy_check_mark:                                                                                 | The trade id.                                                                                      | 01J0XX2KDN3M9QKFKRE2HYSCQM                                                                         |
| `executionID`                                                                                      | *string*                                                                                           | :heavy_check_mark:                                                                                 | The execution id.                                                                                  | 02G0XX2KDN3M9QKFKRE2HYSCMY                                                                         |
| `rebookExecutionRequestCreate`                                                                     | [components.RebookExecutionRequestCreate](../../models/components/rebookexecutionrequestcreate.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.BookingRebookExecutionResponse](../../models/operations/bookingrebookexecutionresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |