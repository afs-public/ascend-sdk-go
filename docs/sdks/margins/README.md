# Margins
(*Margins*)

## Overview

### Available Operations

* [GetBuyingPower](#getbuyingpower) - Get Buying Power

## GetBuyingPower

Gets the buying power for an account

### Example Usage

```go
package main

import(
	"ascend-sdk/models/components"
	ascendsdk "ascend-sdk"
	"context"
	"log"
)

func main() {
    s := ascendsdk.New(
        ascendsdk.WithSecurity(components.Security{
            APIKey: ascendsdk.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 404, 500 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |