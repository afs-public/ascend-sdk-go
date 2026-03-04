# PreIPOResearchDocuments
(*PreIPOResearchDocuments*)

## Overview

### Available Operations

* [ListPreIpoCompanyResearchDocuments](#listpreipocompanyresearchdocuments) - List Pre IPO Company Research Documents
* [GetPreIpoCompanyResearchDocument](#getpreipocompanyresearchdocument) - Get Pre IPO Company Research Document
* [DownloadPreIpoCompanyResearchDocument](#downloadpreipocompanyresearchdocument) - Download Pre IPO Company Research Document

## ListPreIpoCompanyResearchDocuments

Lists Pre IPO Company Research Documents.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyResearchDocuments_ListPreIpoCompanyResearchDocuments" method="get" path="/trading/v1/preIpoCompanies/{preIpoCompany_id}/researchDocuments" -->
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

    res, err := s.PreIPOResearchDocuments.ListPreIpoCompanyResearchDocuments(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6", ascendsdkgo.Int(50), ascendsdkgo.String("eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ=="), ascendsdkgo.String("type == 'MARKET' && relation == 'SUBJECT'"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPreIpoCompanyResearchDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                        |
| `preIpoCompanyID`                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                     | The preIpoCompany id.                                                                                                                                                                                                                                                                  | 3fa85f64-5717-4562-b3fc-2c963f66afa6                                                                                                                                                                                                                                                   |
| `pageSize`                                                                                                                                                                                                                                                                             | **int*                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | The maximum number of Pre IPO Company Research Documents to return. The service may return fewer than this value. If unspecified, at most 100 Pre IPO Company Research Documents will be returned. The maximum value is 100; values above 100 will be coerced to 100.                  | 50                                                                                                                                                                                                                                                                                     |
| `pageToken`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | A page token, received from a previous `ListPreIpoCompanyResearchDocumentsRequest` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListPreIpoCompanyResearchDocumentsRequest` must match the call that provided the page token. | eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ==                                                                                                                                                                                                                       |
| `filter`                                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | A CEL string to filter results. Filterable fields:<br/> - type<br/> - relation<br/> Only `&&` and `==` operators are allowed.<br/> See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search)<br/> page in Guides for more information;           | type == 'MARKET' && relation == 'SUBJECT'                                                                                                                                                                                                                                              |
| `opts`                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.PreIpoCompanyResearchDocumentsListPreIpoCompanyResearchDocumentsResponse](../../models/operations/preipocompanyresearchdocumentslistpreipocompanyresearchdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPreIpoCompanyResearchDocument

Gets a Pre IPO Company Research Document.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyResearchDocuments_GetPreIpoCompanyResearchDocument" method="get" path="/trading/v1/preIpoCompanies/{preIpoCompany_id}/researchDocuments/{researchDocument_id}" -->
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

    res, err := s.PreIPOResearchDocuments.GetPreIpoCompanyResearchDocument(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6", "019af223-10eb-7ea4-861f-a1c283e035a2")
    if err != nil {
        log.Fatal(err)
    }
    if res.PreIpoCompanyResearchDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `preIpoCompanyID`                                        | *string*                                                 | :heavy_check_mark:                                       | The preIpoCompany id.                                    | 3fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `researchDocumentID`                                     | *string*                                                 | :heavy_check_mark:                                       | The researchDocument id.                                 | 019af223-10eb-7ea4-861f-a1c283e035a2                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PreIpoCompanyResearchDocumentsGetPreIpoCompanyResearchDocumentResponse](../../models/operations/preipocompanyresearchdocumentsgetpreipocompanyresearchdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DownloadPreIpoCompanyResearchDocument

Downloads a Pre IPO Company Research Document.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyResearchDocuments_DownloadPreIpoCompanyResearchDocument" method="get" path="/trading/v1/preIpoCompanies/{preIpoCompany_id}/researchDocuments/{researchDocument_id}:download" -->
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

    res, err := s.PreIPOResearchDocuments.DownloadPreIpoCompanyResearchDocument(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6", "019af223-10eb-7ea4-861f-a1c283e035a2")
    if err != nil {
        log.Fatal(err)
    }
    if res.DownloadPreIpoCompanyResearchDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `preIpoCompanyID`                                        | *string*                                                 | :heavy_check_mark:                                       | The preIpoCompany id.                                    | 3fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `researchDocumentID`                                     | *string*                                                 | :heavy_check_mark:                                       | The researchDocument id.                                 | 019af223-10eb-7ea4-861f-a1c283e035a2                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PreIpoCompanyResearchDocumentsDownloadPreIpoCompanyResearchDocumentResponse](../../models/operations/preipocompanyresearchdocumentsdownloadpreipocompanyresearchdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |