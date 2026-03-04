# AlternativeInvestments
(*AlternativeInvestments*)

## Overview

### Available Operations

* [ListAlternativeInvestments](#listalternativeinvestments) - List Alternative Investment Assets
* [GetAlternativeInvestment](#getalternativeinvestment) - Get Alternative Investment Asset

## ListAlternativeInvestments

Retrieves a list of available alternative investment assets and their details.  Replace `{asset_id}` in the endpoint path with a dash to act as a wild card.  Ex:`/trading/v1/assets/-/alternativeInvestments`

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeInvestments_ListAlternativeInvestments" method="get" path="/trading/v1/assets/{asset_id}/alternativeInvestments" -->
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

    res, err := s.AlternativeInvestments.ListAlternativeInvestments(ctx, "-", ascendsdkgo.Int(50), ascendsdkgo.String("eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ=="), ascendsdkgo.String("state == 'OPEN' && type == 'PRIVATE_EQUITY_FUND'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAlternativeInvestmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                                                  | Required                                                                                                                                                                                                                                                                              | Description                                                                                                                                                                                                                                                                           | Example                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                       |
| `assetID`                                                                                                                                                                                                                                                                             | *string*                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                    | The asset id.                                                                                                                                                                                                                                                                         | -                                                                                                                                                                                                                                                                                     |
| `pageSize`                                                                                                                                                                                                                                                                            | **int*                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | The maximum number of orders to return.  - Max value =100  - Values above 100 will be coerced to 100.  - If the specified value is greater than the number of orders, the service will return fewer than the specified value.  - If unspecified, at most 100 orders will be returned. | 50                                                                                                                                                                                                                                                                                    |
| `pageToken`                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | For pagination, provide the page token, received from a previous `ListAlternativeInvestments` call, to retrieve the subsequent page.  All other parameters provided to `ListAlternativeInvestments` must match the request that provided the page token.                              | eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ==                                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | A CEL string to filter results. All fields from the response can be used as filters.<br/><br/> See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) guide for syntax details and examples.                                                 | state == 'OPEN' && type == 'PRIVATE_EQUITY_FUND'                                                                                                                                                                                                                                      |
| `opts`                                                                                                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | The options for this request.                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                       |

### Response

**[*operations.AlternativeInvestmentsListAlternativeInvestmentsResponse](../../models/operations/alternativeinvestmentslistalternativeinvestmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAlternativeInvestment

Retrieves the specified alternative investment asset details.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeInvestments_GetAlternativeInvestment" method="get" path="/trading/v1/assets/{asset_id}/alternativeInvestment" -->
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

    res, err := s.AlternativeInvestments.GetAlternativeInvestment(ctx, "123")
    if err != nil {
        log.Fatal(err)
    }
    if res.AlternativeInvestment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 123                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AlternativeInvestmentsGetAlternativeInvestmentResponse](../../models/operations/alternativeinvestmentsgetalternativeinvestmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |