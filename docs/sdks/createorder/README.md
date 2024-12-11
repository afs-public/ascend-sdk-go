# CreateOrder
(*CreateOrder*)

## Overview

### Available Operations

* [CreateOrder](#createorder) - Create Order
* [GetOrder](#getorder) - Get Order
* [CancelOrder](#cancelorder) - Cancel Order

## CreateOrder

Creates a new order for equity or fixed income securities.

 Equity quantities may be for fractional shares, whole shares, or notional dollar amounts. Fixed income orders may be specified in face value currency amounts, with prices expressed in conventional "percentage of par" values.

 Upon successful submission, if the request is a duplicate, returns the existing order in its current state in the system. If the request is not a duplicate, returns the summary of the newly submitted order.

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
    res, err := s.CreateOrder.CreateOrder(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", components.OrderCreate{
        AssetType: components.AssetTypeEquity,
        BrokerCapacity: components.BrokerCapacityAgency.ToPointer(),
        ClientOrderID: "a6d5258b-6b23-478a-8145-98e79d60427a",
        ClientReceivedTime: types.MustNewTimeFromString("[object Object]"),
        Commission: &components.CommissionCreate{
            Type: components.CommissionCreateTypeAmount,
            Value: components.DecimalCreate{},
        },
        CurrencyCode: ascendsdkgo.String("USD"),
        Identifier: "SBUX",
        IdentifierType: components.IdentifierTypeSymbol,
        LimitPrice: &components.LimitPriceCreate{
            Price: components.DecimalCreate{},
            Type: components.LimitPriceCreateTypePricePerUnit,
        },
        OrderDate: components.DateCreate{},
        OrderType: components.OrderTypeMarket,
        Side: components.SideBuy,
        SpecialReportingInstructions: []components.SpecialReportingInstructions{
            components.SpecialReportingInstructionsRisklessPrincipal,
            components.SpecialReportingInstructionsWithRights,
        },
        StopPrice: &components.StopPriceCreate{
            Price: components.DecimalCreate{},
            Type: components.StopPriceCreateTypePricePerUnit,
        },
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 401, 403, 409, 500, 503 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetOrder

Gets an order by order ID.

 Upon successful submission, returns the details of the queried order.

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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 401, 403, 404, 500, 503 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CancelOrder

Submits an order cancellation request by order ID. Confirmation of order cancellation requests are provided through asynchronous events.

 Upon successful submission, returns the details of the order pending cancellation.

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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 401, 403, 404, 500, 503 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |