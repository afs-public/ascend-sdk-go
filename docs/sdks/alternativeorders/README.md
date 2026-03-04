# AlternativeOrders
(*AlternativeOrders*)

## Overview

### Available Operations

* [CreateAlternativeOrder](#createalternativeorder) - Create Alternative Order
* [ListAlternativeOrders](#listalternativeorders) - List Alternative Orders
* [GetAlternativeOrder](#getalternativeorder) - Get Alternative Order
* [RetrievePendingInvestorActions](#retrievependinginvestoractions) - Get Pending Investor Actions
* [SettleAlternativeOrder](#settlealternativeorder) - Simulate Alternative Order Booking

## CreateAlternativeOrder

Creates an order for an alternative investment asset.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeOrders_CreateAlternativeOrder" method="post" path="/trading/v1/accounts/{account_id}/alternativeOrders" -->
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

    res, err := s.AlternativeOrders.CreateAlternativeOrder(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", components.AlternativeOrderCreate{
        ClientOrderID: "f5074670-1b25-439f-9b5c-702027660800",
        Identifier: "8395",
        IdentifierType: components.AlternativeOrderCreateIdentifierTypeAssetID,
        Side: components.AlternativeOrderCreateSideBuy,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlternativeOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `accountID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | The account id.                                                                        | 01JR8YQT40WAQT8S95ZQGS1QN0                                                             |
| `alternativeOrderCreate`                                                               | [components.AlternativeOrderCreate](../../models/components/alternativeordercreate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.AlternativeOrdersCreateAlternativeOrderResponse](../../models/operations/alternativeorderscreatealternativeorderresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 404, 409 | application/json        |
| sdkerrors.Status        | 500                     | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAlternativeOrders

Retrieves a list of alternative investment orders for the specified account.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeOrders_ListAlternativeOrders" method="get" path="/trading/v1/accounts/{account_id}/alternativeOrders" -->
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

    res, err := s.AlternativeOrders.ListAlternativeOrders(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", ascendsdkgo.Int(100), ascendsdkgo.String("cGFnZV90b2tlbgo="), ascendsdkgo.String("side == 'BUY' && order_status == 'FILLED'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAlternativeOrdersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                                                        | Example                                                                                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                 | The account id.                                                                                                                                                                                                                                                                    | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                                                                                                                                                                                                         |
| `pageSize`                                                                                                                                                                                                                                                                         | **int*                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                 | The maximum number of orders to return. - Max value = 100 - Values above 100 will be coerced to 100. - If the specified value is greater than the number of orders, the service will return fewer than the specified value. - If unspecified, at most 100 orders will be returned. | 100                                                                                                                                                                                                                                                                                |
| `pageToken`                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                 | For pagination, provide the page token, received from a previous `ListAlternativeOrders` call, to retrieve the subsequent page. All other parameters provided to `ListAlternativeOrders` must match the request that provided the page token.                                      | cGFnZV90b2tlbgo=                                                                                                                                                                                                                                                                   |
| `filter`                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                 | A CEL string to filter results. All fields from the response can be used as filters.<br/><br/> See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) guide for syntax details and examples.                                              | side == 'BUY' && order_status == 'FILLED'                                                                                                                                                                                                                                          |
| `opts`                                                                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                    |

### Response

**[*operations.AlternativeOrdersListAlternativeOrdersResponse](../../models/operations/alternativeorderslistalternativeordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAlternativeOrder

Retrieves the details for the specified alternative investment order.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeOrders_GetAlternativeOrder" method="get" path="/trading/v1/accounts/{account_id}/alternativeOrders/{alternativeOrder_id}" -->
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

    res, err := s.AlternativeOrders.GetAlternativeOrder(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "01H5TSDAD9MQ98B8KF36J3Q8P9")
    if err != nil {
        log.Fatal(err)
    }
    if res.AlternativeOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01JR8YQT40WAQT8S95ZQGS1QN0                               |
| `alternativeOrderID`                                     | *string*                                                 | :heavy_check_mark:                                       | The alternativeOrder id.                                 | 01H5TSDAD9MQ98B8KF36J3Q8P9                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AlternativeOrdersGetAlternativeOrderResponse](../../models/operations/alternativeordersgetalternativeorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RetrievePendingInvestorActions

Retrieves the links for any order documents that are awaiting signature and the `party_id` of the party responsible for signing them.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeOrders_RetrievePendingInvestorActions" method="get" path="/trading/v1/accounts/{account_id}/alternativeOrders/{alternativeOrder_id}:retrievePendingInvestorActions" -->
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

    res, err := s.AlternativeOrders.RetrievePendingInvestorActions(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "01H5TSDAD9MQ98B8KF36J3Q8P9")
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrievePendingInvestorActionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01JR8YQT40WAQT8S95ZQGS1QN0                               |
| `alternativeOrderID`                                     | *string*                                                 | :heavy_check_mark:                                       | The alternativeOrder id.                                 | 01H5TSDAD9MQ98B8KF36J3Q8P9                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AlternativeOrdersRetrievePendingInvestorActionsResponse](../../models/operations/alternativeordersretrievependinginvestoractionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SettleAlternativeOrder

Simulates settlement process for an alternative order. For use in UAT environment only.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeOrders_SettleAlternativeOrder" method="post" path="/trading/v1/accounts/{account_id}/alternativeOrders/{alternativeOrder_id}:settle" -->
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

    res, err := s.AlternativeOrders.SettleAlternativeOrder(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "01H5TSDAD9MQ98B8KF36J3Q8P9", components.SettleAlternativeOrderRequestCreate{
        Name: "accounts/01JR8YQT40WAQT8S95ZQGS1QN0/alternativeOrders/01H5TSDAD9MQ98B8KF36J3Q8P9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlternativeOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `accountID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The account id.                                                                                                  | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                                       |
| `alternativeOrderID`                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The alternativeOrder id.                                                                                         | 01H5TSDAD9MQ98B8KF36J3Q8P9                                                                                       |
| `settleAlternativeOrderRequestCreate`                                                                            | [components.SettleAlternativeOrderRequestCreate](../../models/components/settlealternativeorderrequestcreate.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.AlternativeOrdersSettleAlternativeOrderResponse](../../models/operations/alternativeorderssettlealternativeorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |