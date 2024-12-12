# Assets
(*Assets*)

## Overview

### Available Operations

* [ListAssets](#listassets) - List Assets
* [GetAsset](#getasset) - Get Asset
* [ListAssetsCorrespondent](#listassetscorrespondent) - List Assets (By Correspondent)
* [GetAssetCorrespondent](#getassetcorrespondent) - Get Asset (By Correspondent)

## ListAssets

Lists assets

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
    res, err := s.Assets.ListAssets(ctx, ascendsdkgo.String("correspondents/1234"), ascendsdkgo.Int(100), ascendsdkgo.String("Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA=="), ascendsdkgo.String("(symbol == 'IBM' && usable) || symbol == 'USD'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAssetsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                                   |
| `parent`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | The parent resource name, which is the correspondent ID.                                                                                                                                                                                                          | correspondents/1234                                                                                                                                                                                                                                               |
| `pageSize`                                                                                                                                                                                                                                                        | **int*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                | The maximum number of assets to return. The service may return fewer than this value. Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                       | 100                                                                                                                                                                                                                                                               |
| `pageToken`                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A page token, received from a previous `ListAssets` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListAssets` must match the call that provided the page token in order to maintain a stable result set. | Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA==                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                              | (symbol == 'IBM' && usable) \|\| symbol == 'USD'                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                   |

### Response

**[*operations.AssetsListAssets1Response](../../models/operations/assetslistassets1response.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetAsset

Gets assets

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
    res, err := s.Assets.GetAsset(ctx, "8395")
    if err != nil {
        log.Fatal(err)
    }
    if res.Asset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 8395                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AssetsGetAssetResponse](../../models/operations/assetsgetassetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAssetsCorrespondent

Lists assets

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
    res, err := s.Assets.ListAssetsCorrespondent(ctx, "1234", ascendsdkgo.Int(100), ascendsdkgo.String("Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA=="), ascendsdkgo.String("(symbol == 'IBM' && usable) || symbol == 'USD'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAssetsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                                   |
| `correspondentID`                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                | The correspondent id.                                                                                                                                                                                                                                             | 1234                                                                                                                                                                                                                                                              |
| `pageSize`                                                                                                                                                                                                                                                        | **int*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                | The maximum number of assets to return. The service may return fewer than this value. Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                       | 100                                                                                                                                                                                                                                                               |
| `pageToken`                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A page token, received from a previous `ListAssets` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListAssets` must match the call that provided the page token in order to maintain a stable result set. | Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA==                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                              | (symbol == 'IBM' && usable) \|\| symbol == 'USD'                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                   |

### Response

**[*operations.AssetsListAssetsCorrespondentResponse](../../models/operations/assetslistassetscorrespondentresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 500, 503, 504 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetAssetCorrespondent

Gets assets

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
    res, err := s.Assets.GetAssetCorrespondent(ctx, "8395", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Asset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `correspondentID`                                        | *string*                                                 | :heavy_check_mark:                                       | The correspondent id.                                    | 8395                                                     |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AssetsGetAssetCorrespondentResponse](../../models/operations/assetsgetassetcorrespondentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |