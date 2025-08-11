# BasketOrders
(*BasketOrders*)

## Overview

### Available Operations

* [CreateBasket](#createbasket) - Create Basket
* [AddOrders](#addorders) - Add Orders
* [GetBasket](#getbasket) - Get Basket
* [SubmitBasket](#submitbasket) - Submit Basket
* [ListBasketOrders](#listbasketorders) - List Basket Orders
* [ListCompressedOrders](#listcompressedorders) - List Compressed Orders

## CreateBasket

Creates an empty basket

 Upon successful submission, if the request is a duplicate, returns the existing basket in its current state in the system. If the request is not a duplicate, returns the summary of the newly created basket.

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_CreateBasket" method="post" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets" -->
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

    res, err := s.BasketOrders.CreateBasket(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", components.BasketCreate{
        ClientBasketID: "39347a0d-860b-48e8-a04d-511133f057e3",
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Basket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `correspondentID`                                                  | *string*                                                           | :heavy_check_mark:                                                 | The correspondent id.                                              | 01HPMZZM6RKMVZA1JQ63RQKJRP                                         |
| `basketCreate`                                                     | [components.BasketCreate](../../models/components/basketcreate.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |                                                                    |

### Response

**[*operations.BasketOrdersServiceCreateBasketResponse](../../models/operations/basketordersservicecreatebasketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 409 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddOrders

Adds a list of basket orders to a basket

 Upon successful submission, returns the basket with a new total count of orders within the basket

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_AddOrders" method="post" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets/{basket_id}:addOrders" -->
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

    res, err := s.BasketOrders.AddOrders(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", "fffd326-72fa-4d2b-bd1f-45384fe5d521", components.AddOrdersRequestCreate{
        BasketOrders: []components.BasketOrderCreate{},
        Name: "correspondents/01HPMZZM6RKMVZA1JQ63RQKJRP/baskets/fffd326-72fa-4d2b-bd1f-45384fe5d521",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Basket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `correspondentID`                                                                      | *string*                                                                               | :heavy_check_mark:                                                                     | The correspondent id.                                                                  | 01HPMZZM6RKMVZA1JQ63RQKJRP                                                             |
| `basketID`                                                                             | *string*                                                                               | :heavy_check_mark:                                                                     | The basket id.                                                                         | fffd326-72fa-4d2b-bd1f-45384fe5d521                                                    |
| `addOrdersRequestCreate`                                                               | [components.AddOrdersRequestCreate](../../models/components/addordersrequestcreate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.BasketOrdersServiceAddOrdersResponse](../../models/operations/basketordersserviceaddordersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 404, 409 | application/json        |
| sdkerrors.Status        | 500, 503                | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetBasket

Gets a basket by basket ID.

 Upon successful submission, returns the details of the queried basket

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_GetBasket" method="get" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets/{basket_id}" -->
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

    res, err := s.BasketOrders.GetBasket(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", "fffd326-72fa-4d2b-bd1f-45384fe5d521")
    if err != nil {
        log.Fatal(err)
    }
    if res.Basket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `correspondentID`                                        | *string*                                                 | :heavy_check_mark:                                       | The correspondent id.                                    | 01HPMZZM6RKMVZA1JQ63RQKJRP                               |
| `basketID`                                               | *string*                                                 | :heavy_check_mark:                                       | The basket id.                                           | fffd326-72fa-4d2b-bd1f-45384fe5d521                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BasketOrdersServiceGetBasketResponse](../../models/operations/basketordersservicegetbasketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SubmitBasket

Submits a basket for execution in the market

 Upon successful submission, if the request is a duplicate, returns the existing basket in its current state in the system. If the request is not a duplicate, returns the summary of the newly submitted basket in a SUBMITTED state

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_SubmitBasket" method="post" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets/{basket_id}:submit" -->
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

    res, err := s.BasketOrders.SubmitBasket(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", "fffd326-72fa-4d2b-bd1f-45384fe5d521", components.SubmitBasketRequestCreate{
        Name: "correspondents/01HPMZZM6RKMVZA1JQ63RQKJRP/baskets/fffd326-72fa-4d2b-bd1f-45384fe5d521",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Basket != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `correspondentID`                                                                            | *string*                                                                                     | :heavy_check_mark:                                                                           | The correspondent id.                                                                        | 01HPMZZM6RKMVZA1JQ63RQKJRP                                                                   |
| `basketID`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | The basket id.                                                                               | fffd326-72fa-4d2b-bd1f-45384fe5d521                                                          |
| `submitBasketRequestCreate`                                                                  | [components.SubmitBasketRequestCreate](../../models/components/submitbasketrequestcreate.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.BasketOrdersServiceSubmitBasketResponse](../../models/operations/basketordersservicesubmitbasketresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBasketOrders

Gets a list of basket orders within a basket.

 Upon successful submission, returns a list of basket orders for the basket. If the list of basket orders becomes too large, a token is returned to retrieve the next page of basket orders.

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_ListBasketOrders" method="get" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets/{basket_id}/basketOrders" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
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

    res, err := s.BasketOrders.ListBasketOrders(ctx, operations.BasketOrdersServiceListBasketOrdersRequest{
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
        BasketID: "fffd326-72fa-4d2b-bd1f-45384fe5d521",
        PageSize: ascendsdkgo.Int(25),
        PageToken: ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"),
        ShowRemoved: ascendsdkgo.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListBasketOrdersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.BasketOrdersServiceListBasketOrdersRequest](../../models/operations/basketordersservicelistbasketordersrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.BasketOrdersServiceListBasketOrdersResponse](../../models/operations/basketordersservicelistbasketordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCompressedOrders

Gets a list of compressed orders within a basket.

 Upon successful submission, returns a list of compressed orders for the basket. If the basket has not been submitted yet, this list will be empty. If the list of compressed orders becomes too large, a token is returned to retrieve the next page of compressed orders.

### Example Usage

<!-- UsageSnippet language="go" operationID="BasketOrdersService_ListCompressedOrders" method="get" path="/baskettrading/v1/correspondents/{correspondent_id}/baskets/{basket_id}/compressedOrders" -->
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

    res, err := s.BasketOrders.ListCompressedOrders(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", "fffd326-72fa-4d2b-bd1f-45384fe5d521", ascendsdkgo.Int(25), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCompressedOrdersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                          |                                                                                                                                                                                                                                              |
| `correspondentID`                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                           | The correspondent id.                                                                                                                                                                                                                        | 01HPMZZM6RKMVZA1JQ63RQKJRP                                                                                                                                                                                                                   |
| `basketID`                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                           | The basket id.                                                                                                                                                                                                                               | fffd326-72fa-4d2b-bd1f-45384fe5d521                                                                                                                                                                                                          |
| `pageSize`                                                                                                                                                                                                                                   | **int*                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                           | The maximum number of compressed orders to return. The service may return fewer than this value. If unspecified, at most 1000 compressed orders will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.      | 25                                                                                                                                                                                                                                           |
| `pageToken`                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                           | A page token, received from a previous `ListCompressedOrders` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListCompressedOrders` must match the call that provided the page token. | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                |                                                                                                                                                                                                                                              |

### Response

**[*operations.BasketOrdersServiceListCompressedOrdersResponse](../../models/operations/basketordersservicelistcompressedordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |