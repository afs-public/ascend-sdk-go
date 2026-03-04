# OptionOrders
(*OptionOrders*)

## Overview

### Available Operations

* [CreateOptionOrder](#createoptionorder) - Create Option Order
* [GetOptionOrder](#getoptionorder) - Get Option Order
* [CancelOptionOrder](#canceloptionorder) - Cancel Option Order

## CreateOptionOrder

Creates a new option order.

 Currently only single-leg option orders are supported, but the data structures will support multi-leg orders in the future. The single-leg constraint will be repeated in this documentation, but validation rules related to the initial (future) multi-leg support are also documented.

 Upon successful submission, if the request is a duplicate, returns the existing order in its current state in the system. If the request is not a duplicate, returns the summary of the newly submitted order.

### Example Usage

<!-- UsageSnippet language="go" operationID="OptionOrderService_CreateOptionOrder" method="post" path="/trading/v1/accounts/{account_id}/optionOrders" -->
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

    res, err := s.OptionOrders.CreateOptionOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", components.OptionOrderCreate{
        BrokerCapacity: components.OptionOrderCreateBrokerCapacityAgency,
        ClientOrderID: "a6d5258b-6b23-478a-8145-98e79d60427a",
        CurrencyCode: "USD",
        Legs: []components.OptionOrderLegCreate{},
        LimitPrice: components.DecimalCreate{},
        OrderDate: components.DateCreate{},
        OrderType: components.OptionOrderCreateOrderTypeLimit,
        PriceDirection: components.PriceDirectionCredit,
        Quantity: components.DecimalCreate{},
        TimeInForce: components.OptionOrderCreateTimeInForceDay,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `accountID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | The account id.                                                              | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                   |
| `optionOrderCreate`                                                          | [components.OptionOrderCreate](../../models/components/optionordercreate.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.OptionOrderServiceCreateOptionOrderResponse](../../models/operations/optionorderservicecreateoptionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 409 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOptionOrder

Gets an option order by option order ID.

 Upon successful submission, returns the details of the queried order.

### Example Usage

<!-- UsageSnippet language="go" operationID="OptionOrderService_GetOptionOrder" method="get" path="/trading/v1/accounts/{account_id}/optionOrders/{optionOrder_id}" -->
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

    res, err := s.OptionOrders.GetOptionOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "ebb0c9b5-2c74-45c9-a4ab-40596b778706")
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HBRQ5BW6ZAY4BNWP4GWRD80X                               |
| `optionOrderID`                                          | *string*                                                 | :heavy_check_mark:                                       | The optionOrder id.                                      | ebb0c9b5-2c74-45c9-a4ab-40596b778706                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.OptionOrderServiceGetOptionOrderResponse](../../models/operations/optionorderservicegetoptionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelOptionOrder

Submits an order cancellation request by option order ID. Confirmation of order cancellation requests are provided through asynchronous events.

 Upon successful submission, returns the details of the order pending cancellation.

### Example Usage

<!-- UsageSnippet language="go" operationID="OptionOrderService_CancelOptionOrder" method="post" path="/trading/v1/accounts/{account_id}/optionOrders/{optionOrder_id}:cancel" -->
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

    res, err := s.OptionOrders.CancelOptionOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "ebb0c9b5-2c74-45c9-a4ab-40596b778706", components.CancelOptionOrderRequestCreate{
        Name: "accounts/01HBRQ5BW6ZAY4BNWP4GWRD80X/optionOrders/ebb0c9b5-2c74-45c9-a4ab-40596b778706",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `accountID`                                                                                            | *string*                                                                                               | :heavy_check_mark:                                                                                     | The account id.                                                                                        | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                                             |
| `optionOrderID`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | The optionOrder id.                                                                                    | ebb0c9b5-2c74-45c9-a4ab-40596b778706                                                                   |
| `cancelOptionOrderRequestCreate`                                                                       | [components.CancelOptionOrderRequestCreate](../../models/components/canceloptionorderrequestcreate.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.OptionOrderServiceCancelOptionOrderResponse](../../models/operations/optionorderservicecanceloptionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |