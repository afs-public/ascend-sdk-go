# CreateOrder
(*CreateOrder*)

## Overview

### Available Operations

* [CreateOrder](#createorder) - Create Order
* [GetOrder](#getorder) - Get Order
* [CancelOrder](#cancelorder) - Cancel Order
* [SetExtraReportingData](#setextrareportingdata) - Set Extra Reporting Data
* [ListCorrespondentOrders](#listcorrespondentorders) - List Correspondent Orders

## CreateOrder

Creates a new order for equity or fixed income securities.

 Equity quantities may be for fractional shares, whole shares, or notional dollar amounts. Fixed income orders may be specified in face value currency amounts, with prices expressed in conventional "percentage of par" values.

 Upon successful submission, if the request is a duplicate, returns the existing order in its current state in the system. If the request is not a duplicate, returns the summary of the newly submitted order.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderService_CreateOrder" method="post" path="/trading/v1/accounts/{account_id}/orders" -->
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

    res, err := s.CreateOrder.CreateOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", components.OrderCreate{
        AssetType: components.AssetTypeEquity,
        ClientOrderID: "a6d5258b-6b23-478a-8145-98e79d60427a",
        Identifier: "SBUX",
        IdentifierType: components.IdentifierTypeSymbol,
        OrderDate: components.DateCreate{},
        OrderType: components.OrderTypeMarket,
        Side: components.SideBuy,
        TimeInForce: components.TimeInForceDay,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `accountID`                                                      | *string*                                                         | :heavy_check_mark:                                               | The account id.                                                  | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                       |
| `orderCreate`                                                    | [components.OrderCreate](../../models/components/ordercreate.md) | :heavy_check_mark:                                               | N/A                                                              |                                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.OrderServiceCreateOrderResponse](../../models/operations/orderservicecreateorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 409 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOrder

Gets an order by order ID.

 Upon successful submission, returns the details of the queried order.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderService_GetOrder" method="get" path="/trading/v1/accounts/{account_id}/orders/{order_id}" -->
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

    res, err := s.CreateOrder.GetOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "ebb0c9b5-2c74-45c9-a4ab-40596b778706")
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HBRQ5BW6ZAY4BNWP4GWRD80X                               |
| `orderID`                                                | *string*                                                 | :heavy_check_mark:                                       | The order id.                                            | ebb0c9b5-2c74-45c9-a4ab-40596b778706                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.OrderServiceGetOrderResponse](../../models/operations/orderservicegetorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelOrder

Submits an order cancellation request by order ID. Confirmation of order cancellation requests are provided through asynchronous events.

 Upon successful submission, returns the details of the order pending cancellation.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderService_CancelOrder" method="post" path="/trading/v1/accounts/{account_id}/orders/{order_id}:cancel" -->
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

    res, err := s.CreateOrder.CancelOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "ebb0c9b5-2c74-45c9-a4ab-40596b778706", components.CancelOrderRequestCreate{
        Name: "accounts/01HBRQ5BW6ZAY4BNWP4GWRD80X/orders/ebb0c9b5-2c74-45c9-a4ab-40596b778706",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The account id.                                                                            | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                                 |
| `orderID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | The order id.                                                                              | ebb0c9b5-2c74-45c9-a4ab-40596b778706                                                       |
| `cancelOrderRequestCreate`                                                                 | [components.CancelOrderRequestCreate](../../models/components/cancelorderrequestcreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.OrderServiceCancelOrderResponse](../../models/operations/orderservicecancelorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetExtraReportingData

Sets extra reporting data to an existing order. Any SetExtraReportingDataRequest must include the name of the order and the cancel_confirmed_time

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderService_SetExtraReportingData" method="post" path="/trading/v1/accounts/{account_id}/orders/{order_id}:setExtraReportingData" -->
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

    res, err := s.CreateOrder.SetExtraReportingData(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "ebb0c9b5-2c74-45c9-a4ab-40596b778706", components.SetExtraReportingDataRequestCreate{
        CancelConfirmedTime: types.MustNewTimeFromString("2025-12-13T15:28:17.262732Z"),
        Name: "accounts/01HBRQ5BW6ZAY4BNWP4GWRD80X/orders/ebb0c9b5-2c74-45c9-a4ab-40596b778706",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                                                     |
| `orderID`                                                                                                      | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The order id.                                                                                                  | ebb0c9b5-2c74-45c9-a4ab-40596b778706                                                                           |
| `setExtraReportingDataRequestCreate`                                                                           | [components.SetExtraReportingDataRequestCreate](../../models/components/setextrareportingdatarequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.OrderServiceSetExtraReportingDataResponse](../../models/operations/orderservicesetextrareportingdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCorrespondentOrders

Lists orders matching the specified filter criteria. Results are paginated and sorted in the reverse order of their creation.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderService_ListCorrespondentOrders" method="get" path="/trading/v1/correspondents/{correspondent_id}/orders" -->
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

    res, err := s.CreateOrder.ListCorrespondentOrders(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", ascendsdkgo.String("open && order_date >= date('2025-05-10')"), ascendsdkgo.Int(50), ascendsdkgo.String("CiAKGjBpNDd2Nmp2Zml2cXRwYjBpOXA"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCorrespondentOrdersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `correspondentID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | The correspondent id.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `filter`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | CEL filter string expressing what orders should be listed. The only properties available for filtering are the boolean `open` and `order_date`. Each of these represent fields on the Orders object, and more details about each can be found attached to the properties.<br/><br/> If `open` is not provided, both "open" and "not open" orders will be returned. All `order_date` searches are limited to orders within the most recent 365 days. If no `order_date` is specified, the default will search between now and 365 days ago. | open && order_date >= date('2025-05-10')                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | **int*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | The number of records to return in a single page. The maximum page size is 100. If a value is not provided, the default of 100 will be used. If a value less than one, or greater than the maximum, is provided, the default value will be used.                                                                                                                                                                                                                                                                                   | 50                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `pageToken`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | The token for the next page of content to fetch. When paginating, all other parameters provided to `ListOrders` must match the call that provided the page token.                                                                                                                                                                                                                                                                                                                                                                  | CiAKGjBpNDd2Nmp2Zml2cXRwYjBpOXA                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |

### Response

**[*operations.OrderServiceListCorrespondentOrdersResponse](../../models/operations/orderservicelistcorrespondentordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |