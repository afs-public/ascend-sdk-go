# InvestorDocs
(*InvestorDocs*)

## Overview

### Available Operations

* [BatchCreateUploadLinks](#batchcreateuploadlinks) - Batch Create Upload Links
* [ListDocuments](#listdocuments) - List Documents

## BatchCreateUploadLinks

Create a batch of signed links that can be used to upload files.

### Example Usage

<!-- UsageSnippet language="go" operationID="InvestorCommunicationService_BatchCreateUploadLinks" method="post" path="/investordocs/v1/uploadLinks:batchCreate" -->
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

    res, err := s.InvestorDocs.BatchCreateUploadLinks(ctx, components.BatchCreateUploadLinksRequestCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.BatchCreateUploadLinksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [components.BatchCreateUploadLinksRequestCreate](../../models/components/batchcreateuploadlinksrequestcreate.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.InvestorCommunicationServiceBatchCreateUploadLinksResponse](../../models/operations/investorcommunicationservicebatchcreateuploadlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDocuments

List documents that match search parameters.

### Example Usage

<!-- UsageSnippet language="go" operationID="InvestorCommunicationService_ListDocuments" method="get" path="/investordocs/v1/documents" -->
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

    res, err := s.InvestorDocs.ListDocuments(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `pageSize`                                                                                                                                            | **int*                                                                                                                                                | :heavy_minus_sign:                                                                                                                                    | The maximum number of items to return; The service may return fewer than this value                                                                   |
| `pageToken`                                                                                                                                           | **string*                                                                                                                                             | :heavy_minus_sign:                                                                                                                                    | Token used to get a specific page of results                                                                                                          |
| `filter`                                                                                                                                              | **string*                                                                                                                                             | :heavy_minus_sign:                                                                                                                                    | CEL filter to be applied against the documents; Providing a correspondent to search for is required; Only one correspondent can be searched at a time |
| `opts`                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                              | :heavy_minus_sign:                                                                                                                                    | The options for this request.                                                                                                                         |

### Response

**[*operations.InvestorCommunicationServiceListDocumentsResponse](../../models/operations/investorcommunicationservicelistdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |