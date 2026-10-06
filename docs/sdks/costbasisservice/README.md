# CostBasisService
(*CostBasisService*)

## Overview

### Available Operations

* [SearchClosedLots](#searchclosedlots) - Search Closed Lots
* [SearchOpenLots](#searchopenlots) - Search Open Lots

## SearchClosedLots

SearchClosedLots returns a list of closed lots for a given account and date range. Please note that if no trade date ranges are provided, all the closed lots using the MTD range will be returned.

### Example Usage

<!-- UsageSnippet language="go" operationID="CostBasisService_SearchClosedLots" method="post" path="/costbasis/v1/accounts/{account_id}/closedLots:search" -->
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

    res, err := s.CostBasisService.SearchClosedLots(ctx, "01J71HKJ1K1GX5C0EWZ4BCPACB", components.SearchClosedLotsRequestCreate{
        Parent: "accounts/01J71HKJ1K1GX5C0EWZ4BCPACB",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchClosedLotsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The account id.                                                                                      | 01J71HKJ1K1GX5C0EWZ4BCPACB                                                                           |
| `searchClosedLotsRequestCreate`                                                                      | [components.SearchClosedLotsRequestCreate](../../models/components/searchclosedlotsrequestcreate.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.CostBasisServiceSearchClosedLotsResponse](../../models/operations/costbasisservicesearchclosedlotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SearchOpenLots

SearchOpenLots returns a list of open lots for a given account

### Example Usage

<!-- UsageSnippet language="go" operationID="CostBasisService_SearchOpenLots" method="post" path="/costbasis/v1/accounts/{account_id}/openLots:search" -->
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

    res, err := s.CostBasisService.SearchOpenLots(ctx, "01J71HKJ1K1GX5C0EWZ4BCPACB", components.SearchOpenLotsRequestCreate{
        Parent: "accounts/01J71HKJ1K1GX5C0EWZ4BCPACB",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchOpenLotsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01J71HKJ1K1GX5C0EWZ4BCPACB                                                                       |
| `searchOpenLotsRequestCreate`                                                                    | [components.SearchOpenLotsRequestCreate](../../models/components/searchopenlotsrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.CostBasisServiceSearchOpenLotsResponse](../../models/operations/costbasisservicesearchopenlotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |