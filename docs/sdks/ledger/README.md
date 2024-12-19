# Ledger
(*Ledger*)

## Overview

### Available Operations

* [ListEntries](#listentries) - List Entries
* [ListActivities](#listactivities) - List Activities
* [ListPositions](#listpositions) - List Positions
* [GetActivity](#getactivity) - Get Activity
* [GetEntry](#getentry) - Get Entry

## ListEntries

List all Entries based on a filter

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
    res, err := s.Ledger.ListEntries(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                                               |                                                                                                                                                                                                                                                                   |
| `accountID`                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                                                | The account id.                                                                                                                                                                                                                                                   | 01FAKEACCOUNT1TYKWEYRH8S2K                                                                                                                                                                                                                                        |
| `pageSize`                                                                                                                                                                                                                                                        | **int*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                | The maximum number of entries to return. The service may return fewer than this value Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                       | 0                                                                                                                                                                                                                                                                 |
| `pageToken`                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A page token, received from a previous `ListEntries` call. Provide this to retrieve the subsequent page When paginating, all other parameters provided to `ListEntries` must match the call that provided the page token in order to maintain a stable result set | v-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAOv-CAfzbNG7ZAS8xZWYyMmM3ZS01NjdmLTBhYzgtYjZmZi1kNzYwNDI3YmI3N2Q6MjAyNC0wNi0wMgA=                                                                                                               |
| `filter`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                              | process_date == date('2024-05-11') && account_id == '01HBRQ5BW6ZAY4BNWP4GWRD80X'                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                   |

### Response

**[*operations.LedgerListEntriesResponse](../../models/operations/ledgerlistentriesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListActivities

List all Completed Activities based on a filter

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
    res, err := s.Ledger.ListActivities(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListActivitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                           | Type                                                                                                                                                                                                                                                                | Required                                                                                                                                                                                                                                                            | Description                                                                                                                                                                                                                                                         | Example                                                                                                                                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                     |
| `accountID`                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                  | The account id.                                                                                                                                                                                                                                                     | 01FAKEACCOUNT1TYKWEYRH8S2K                                                                                                                                                                                                                                          |
| `pageSize`                                                                                                                                                                                                                                                          | **int*                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                  | The maximum number of activities to return. The service may return fewer than this value Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                      | 100                                                                                                                                                                                                                                                                 |
| `pageToken`                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                  | A page token, received from a previous `ListActivity` call. Provide this to retrieve the subsequent page When paginating, all other parameters provided to `ListActivity` must match the call that provided the page token in order to maintain a stable result set | Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAI_-CAfwVsHF9ARgyMDI0LTA2LTA0OjFGQTA1MDExOjUwMDEA                                                                                                                                                |
| `filter`                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                  | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                                | subtype_category == 'TRADE' && process_date >= date('2023-07-31') && settle_date >= date('2023-08-18') && side == 'BUY' &&  activity_date >= date('2023-09-15') && asset_id == 8395                                                                                 |
| `opts`                                                                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                     |

### Response

**[*operations.LedgerListActivitiesResponse](../../models/operations/ledgerlistactivitiesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListPositions

List positions based on a filter

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
    res, err := s.Ledger.ListPositions(ctx, "01HBRQ5BW6ZAY4BNWP4GWRD80X", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPositionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                      |
| `accountID`                                                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                   | The account id.                                                                                                                                                                                                                                                      | 01HBRQ5BW6ZAY4BNWP4GWRD80X                                                                                                                                                                                                                                           |
| `pageSize`                                                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                   | The maximum number of positions to return. The service may return fewer than this value Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                        | 20                                                                                                                                                                                                                                                                   |
| `pageToken`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | This page token comes from a previous `ListPositions` call; provide this token to retrieve the subsequent page When paginating, all other parameters you include in `ListPositions` must match the call that includes the page token to maintain a stable result set | Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAOv-CAfwFIZG3AS8xZWYyMmM4Ny0zNDI5LTAyYzItODRjNC03ODdmNTJlNDY1MTE6MjAyNC0wNi0wMgA=                                                                                                                 |
| `filter`                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                                 | date >= date('2023-08-31') && asset_id == 8395                                                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                      |

### Response

**[*operations.LedgerListPositionsResponse](../../models/operations/ledgerlistpositionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetActivity

Get an activity

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
    res, err := s.Ledger.GetActivity(ctx, "01FAKEACCOUNT1TYKWEYRH8S2K", "FAKEACTIVITYID")
    if err != nil {
        log.Fatal(err)
    }
    if res.Activity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01FAKEACCOUNT1TYKWEYRH8S2K                               |
| `activityID`                                             | *string*                                                 | :heavy_check_mark:                                       | The activity id.                                         | FAKEACTIVITYID                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.LedgerGetActivityResponse](../../models/operations/ledgergetactivityresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetEntry

Get an entry

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
    res, err := s.Ledger.GetEntry(ctx, "{\"account_id\":\"\"}", "{\"entry_id\":\"\"}")
    if err != nil {
        log.Fatal(err)
    }
    if res.Entry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | {<br/>"account_id": ""<br/>}                             |
| `entryID`                                                | *string*                                                 | :heavy_check_mark:                                       | The entry id.                                            | {<br/>"entry_id": ""<br/>}                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.LedgerGetEntryResponse](../../models/operations/ledgergetentryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 500, 503, 504 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |