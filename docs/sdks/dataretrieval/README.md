# DataRetrieval
(*DataRetrieval*)

## Overview

### Available Operations

* [ListSnapshots](#listsnapshots) - List Snapshots

## ListSnapshots

Returns details of a list of snapshots.

### Example Usage

<!-- UsageSnippet language="go" operationID="Snapshots_ListSnapshots" method="get" path="/analytics/v1/snapshots" -->
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

    res, err := s.DataRetrieval.ListSnapshots(ctx, ascendsdkgo.String("snapshot_type==\"daily_accounts\"&&process_date==date(\"2023-09-30\")"), ascendsdkgo.Int(500), ascendsdkgo.String("M_-BAwEBCVBhZ2VUb2tlbgH_ggABAgEMUnVubmluZ1RvdGFsAQQAAQZGaWx0ZXIBDAAAAAX_ggEyAA=="))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSnapshotsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                     | Type                                                                                                                                                                          | Required                                                                                                                                                                      | Description                                                                                                                                                                   | Example                                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                         | :heavy_check_mark:                                                                                                                                                            | The context to use for the request.                                                                                                                                           |                                                                                                                                                                               |
| `filter`                                                                                                                                                                      | **string*                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                            | A CEL string to filter snapshot results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; | snapshot_type=="daily_accounts"&&process_date==date("2023-09-30")                                                                                                             |
| `pageSize`                                                                                                                                                                    | **int*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                            | The number of snapshots to be returned per page. Defaults to 500. Maximum is 1000.                                                                                            | 500                                                                                                                                                                           |
| `pageToken`                                                                                                                                                                   | **string*                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                            | The token for retrieving the next page of snapshots, the value of which will have been returned in a previous response.                                                       | M_-BAwEBCVBhZ2VUb2tlbgH_ggABAgEMUnVubmluZ1RvdGFsAQQAAQZGaWx0ZXIBDAAAAAX_ggEyAA==                                                                                              |
| `opts`                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                      | :heavy_minus_sign:                                                                                                                                                            | The options for this request.                                                                                                                                                 |                                                                                                                                                                               |

### Response

**[*operations.SnapshotsListSnapshotsResponse](../../models/operations/snapshotslistsnapshotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |