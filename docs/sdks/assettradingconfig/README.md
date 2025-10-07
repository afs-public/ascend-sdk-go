# AssetTradingConfig
(*AssetTradingConfig*)

## Overview

### Available Operations

* [GetAssetTradingConfig](#getassettradingconfig) - Get Asset Trading Config
* [ListAssetTradingConfigs](#listassettradingconfigs) - List Asset Trading Configs

## GetAssetTradingConfig

Gets an asset trading config by asset_id `/assettradingconfig/v1/correspondents/{correspondent_id}/assets/{asset_id}/tradingConfig`

### Example Usage

<!-- UsageSnippet language="go" operationID="AssetTradingConfigService_GetAssetTradingConfig" method="get" path="/assettradingconfig/v1/correspondents/{correspondent_id}/assets/{asset_id}/tradingConfig" -->
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

    res, err := s.AssetTradingConfig.GetAssetTradingConfig(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", "612")
    if err != nil {
        log.Fatal(err)
    }
    if res.AssetTradingConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `correspondentID`                                        | *string*                                                 | :heavy_check_mark:                                       | The correspondent id.                                    | 01HBRQ5BW6ZAY4BNWP4GWRD80X                               |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 612                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AssetTradingConfigServiceGetAssetTradingConfigResponse](../../models/operations/assettradingconfigservicegetassettradingconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAssetTradingConfigs

Retrieve a list of asset trading configs `/assettradingconfig/v1/correspondents/{correspondent_id}/assets/-/tradingConfigs`

### Example Usage

<!-- UsageSnippet language="go" operationID="AssetTradingConfigService_ListAssetTradingConfigs" method="get" path="/assettradingconfig/v1/correspondents/{correspondent_id}/assets/{asset_id}/tradingConfigs" -->
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

    res, err := s.AssetTradingConfig.ListAssetTradingConfigs(ctx, operations.AssetTradingConfigServiceListAssetTradingConfigsRequest{
        CorrespondentID: "01HBRQ5BW6ZAY4BNWP4GWRD80X",
        AssetID: "",
        PageSize: ascendsdkgo.Int(100),
        PageToken: ascendsdkgo.String("Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA=="),
        Filter: ascendsdkgo.String("symbol == 'SBUX' && asset_type == 'EQUITY'"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAssetTradingConfigsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `request`                                                                                                                                                | [operations.AssetTradingConfigServiceListAssetTradingConfigsRequest](../../models/operations/assettradingconfigservicelistassettradingconfigsrequest.md) | :heavy_check_mark:                                                                                                                                       | The request object to use for the request.                                                                                                               |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.AssetTradingConfigServiceListAssetTradingConfigsResponse](../../models/operations/assettradingconfigservicelistassettradingconfigsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500, 503, 504      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |