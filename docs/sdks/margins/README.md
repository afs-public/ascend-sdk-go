# Margins
(*Margins*)

## Overview

### Available Operations

* [GetBuyingPower](#getbuyingpower) - Get Buying Power

## GetBuyingPower

Gets the buying power for an account

### Example Usage

<!-- UsageSnippet language="go" operationID="MarginsRealTime_GetBuyingPower" method="get" path="/buyingpower/v1/accounts/{account_id}/buyingPower" -->
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

    res, err := s.Margins.GetBuyingPower(ctx, "01HMS9S15AKBHBD8GPW33P2PMH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BuyingPower != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HMS9S15AKBHBD8GPW33P2PMH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.MarginsRealTimeGetBuyingPowerResponse](../../models/operations/marginsrealtimegetbuyingpowerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |