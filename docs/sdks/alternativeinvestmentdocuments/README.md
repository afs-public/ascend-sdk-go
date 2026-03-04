# AlternativeInvestmentDocuments
(*AlternativeInvestmentDocuments*)

## Overview

### Available Operations

* [ListAlternativeInvestmentDocuments](#listalternativeinvestmentdocuments) - List Alternative Investment Documents
* [GetAlternativeInvestmentDocument](#getalternativeinvestmentdocument) - Get Alternative Investment Document
* [DownloadAlternativeInvestmentDocument](#downloadalternativeinvestmentdocument) - Download Alternative Investment Documents

## ListAlternativeInvestmentDocuments

Retrieves a list of all investment document details for the specified asset.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeInvestmentDocuments_ListAlternativeInvestmentDocuments" method="get" path="/trading/v1/assets/{asset_id}/alternativeInvestmentDocuments" -->
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

    res, err := s.AlternativeInvestmentDocuments.ListAlternativeInvestmentDocuments(ctx, "123", ascendsdkgo.Int(10), ascendsdkgo.String("next-page-token-example"), ascendsdkgo.String("type == 'OFFERING_DOCUMENT'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAlternativeInvestmentDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                                                  | Required                                                                                                                                                                                                                                                                              | Description                                                                                                                                                                                                                                                                           | Example                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                       |
| `assetID`                                                                                                                                                                                                                                                                             | *string*                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                    | The asset id.                                                                                                                                                                                                                                                                         | 123                                                                                                                                                                                                                                                                                   |
| `pageSize`                                                                                                                                                                                                                                                                            | **int*                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | The maximum number of orders to return. - Max value = 100  - Values above 100 will be coerced to 100.  - If the specified value is greater than the number of orders, the service will return fewer than the specified value.  - If unspecified, at most 100 orders will be returned. | 10                                                                                                                                                                                                                                                                                    |
| `pageToken`                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | For pagination, provide the page token, received from a previous `ListInvestmentDocuments` call, to retrieve the subsequent page.  All other parameters provided to `ListInvestmentDocuments` must match the request that provided the page token.                                    | next-page-token-example                                                                                                                                                                                                                                                               |
| `filter`                                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | A CEL string to filter results. All fields from the response can be used as filters.<br/><br/> See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) guide for syntax details and examples.                                                 | type == 'OFFERING_DOCUMENT'                                                                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                    | The options for this request.                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                       |

### Response

**[*operations.AlternativeInvestmentDocumentsListAlternativeInvestmentDocumentsResponse](../../models/operations/alternativeinvestmentdocumentslistalternativeinvestmentdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAlternativeInvestmentDocument

Retrieves a specific investment document for the specified asset.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeInvestmentDocuments_GetAlternativeInvestmentDocument" method="get" path="/trading/v1/assets/{asset_id}/alternativeInvestmentDocuments/{alternativeInvestmentDocument_id}" -->
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

    res, err := s.AlternativeInvestmentDocuments.GetAlternativeInvestmentDocument(ctx, "123", "01H7YH8QKZJ8V4Q2X8F4G6JQ2B")
    if err != nil {
        log.Fatal(err)
    }
    if res.AlternativeInvestmentDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 123                                                      |
| `alternativeInvestmentDocumentID`                        | *string*                                                 | :heavy_check_mark:                                       | The alternativeInvestmentDocument id.                    | 01H7YH8QKZJ8V4Q2X8F4G6JQ2B                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AlternativeInvestmentDocumentsGetAlternativeInvestmentDocumentResponse](../../models/operations/alternativeinvestmentdocumentsgetalternativeinvestmentdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DownloadAlternativeInvestmentDocument

Returns a URI to download the specified investment document.

### Example Usage

<!-- UsageSnippet language="go" operationID="AlternativeInvestmentDocuments_DownloadAlternativeInvestmentDocument" method="get" path="/trading/v1/assets/{asset_id}/alternativeInvestmentDocuments/{alternativeInvestmentDocument_id}:download" -->
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

    res, err := s.AlternativeInvestmentDocuments.DownloadAlternativeInvestmentDocument(ctx, "123", "01H7YH8QKZJ8V4Q2X8F4G6JQ2B")
    if err != nil {
        log.Fatal(err)
    }
    if res.DownloadAlternativeInvestmentDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 123                                                      |
| `alternativeInvestmentDocumentID`                        | *string*                                                 | :heavy_check_mark:                                       | The alternativeInvestmentDocument id.                    | 01H7YH8QKZJ8V4Q2X8F4G6JQ2B                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AlternativeInvestmentDocumentsDownloadAlternativeInvestmentDocumentResponse](../../models/operations/alternativeinvestmentdocumentsdownloadalternativeinvestmentdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |