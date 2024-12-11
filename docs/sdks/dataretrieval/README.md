# DataRetrieval
(*DataRetrieval*)

## Overview

### Available Operations

* [ListSnapshots](#listsnapshots) - List Snapshots

## ListSnapshots

Returns details of a list of snapshots

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
    res, err := s.DataRetrieval.ListSnapshots(ctx, ascendsdkgo.String("snapshot_id==\"daily_accounts\"&&process_date==date(\"2023-09-30\")"), ascendsdkgo.Int(5000), ascendsdkgo.String("ZXhhbXBsZQo"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSnapshotsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                        | Required                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                 | Example                                                                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                          | The context to use for the request.                                                                                                                                                                                         |                                                                                                                                                                                                                             |
| `filter`                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                          | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `snapshot_id`<br/> `process_date` | snapshot_id=="daily_accounts"&&process_date==date("2023-09-30")                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                  | **int*                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                          | The number of snapshots to be returned per page. Defaults to 500. Maximum is 1000.                                                                                                                                          | 5000                                                                                                                                                                                                                        |
| `pageToken`                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                          | The token used to retrieve a page of snapshots.                                                                                                                                                                             | ZXhhbXBsZQo                                                                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                          | The options for this request.                                                                                                                                                                                               |                                                                                                                                                                                                                             |

### Response

**[*operations.SnapshotsListSnapshotsResponse](../../models/operations/snapshotslistsnapshotsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 403, 500           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |