# PreIPONewsEvents
(*PreIPONewsEvents*)

## Overview

### Available Operations

* [ListPreIpoCompanyNewsEvents](#listpreipocompanynewsevents) - List Pre IPO Company News Events

## ListPreIpoCompanyNewsEvents

Lists Pre IPO Company News Events.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyNewsEvents_ListPreIpoCompanyNewsEvents" method="get" path="/trading/v1/preIpoCompanies/{preIpoCompany_id}/newsEvents" -->
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

    res, err := s.PreIPONewsEvents.ListPreIpoCompanyNewsEvents(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6", ascendsdkgo.Int(50), ascendsdkgo.String("eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ=="))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPreIpoCompanyNewsEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                          |
| `preIpoCompanyID`                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                       | The preIpoCompany id.                                                                                                                                                                                                                                                    | 3fa85f64-5717-4562-b3fc-2c963f66afa6                                                                                                                                                                                                                                     |
| `pageSize`                                                                                                                                                                                                                                                               | **int*                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                       | The maximum number of Pre IPO Company News Events to return. The service may return fewer than this value. If unspecified, at most 100 Pre IPO Company News Events will be returned. The maximum value is 100; values above 100 will be coerced to 100.                  | 50                                                                                                                                                                                                                                                                       |
| `pageToken`                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                       | A page token, received from a previous `ListPreIpoCompanyNewsEventsRequest` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListPreIpoCompanyNewsEventsRequest` must match the call that provided the page token. | eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ==                                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                          |

### Response

**[*operations.PreIpoCompanyNewsEventsListPreIpoCompanyNewsEventsResponse](../../models/operations/preipocompanynewseventslistpreipocompanynewseventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |