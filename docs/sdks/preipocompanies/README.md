# PreIPOCompanies
(*PreIPOCompanies*)

## Overview

### Available Operations

* [ListPreIpoCompanies](#listpreipocompanies) - List Pre IPO Company
* [GetPreIpoCompany](#getpreipocompany) - Get Pre IPO Company

## ListPreIpoCompanies

Lists Pre IPO Companies.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanies_ListPreIpoCompanies" method="get" path="/trading/v1/preIpoCompanies" -->
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

    res, err := s.PreIPOCompanies.ListPreIpoCompanies(ctx, ascendsdkgo.Int(50), ascendsdkgo.String("eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ=="))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPreIpoCompaniesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                          |
| `pageSize`                                                                                                                                                                                                                                               | **int*                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                       | The maximum number of Pre IPO Companies to return. The service may return fewer than this value. If unspecified, at most 100 Pre IPO Companies will be returned. The maximum value is 100; values above 100 will be coerced to 100.                      | 50                                                                                                                                                                                                                                                       |
| `pageToken`                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                       | A page token, received from a previous `ListPreIpoCompaniesRequest` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListPreIpoCompaniesRequest` must match the call that provided the page token. | eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ==                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                          |

### Response

**[*operations.PreIpoCompaniesListPreIpoCompaniesResponse](../../models/operations/preipocompanieslistpreipocompaniesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPreIpoCompany

Gets a Pre IPO Company.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanies_GetPreIpoCompany" method="get" path="/trading/v1/preIpoCompanies/{preIpoCompany_id}" -->
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

    res, err := s.PreIPOCompanies.GetPreIpoCompany(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
    if err != nil {
        log.Fatal(err)
    }
    if res.PreIpoCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `preIpoCompanyID`                                        | *string*                                                 | :heavy_check_mark:                                       | The preIpoCompany id.                                    | 3fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PreIpoCompaniesGetPreIpoCompanyResponse](../../models/operations/preipocompaniesgetpreipocompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |