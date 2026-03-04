# PreIPOInterests
(*PreIPOInterests*)

## Overview

### Available Operations

* [CreatePreIpoCompanyInterest](#createpreipocompanyinterest) - Create Pre IPO Company Interest
* [ListPreIpoCompanyInterests](#listpreipocompanyinterests) - List Pre IPO Company Interests
* [GetPreIpoCompanyInterest](#getpreipocompanyinterest) - Get Pre IPO Company Interest
* [UpdatePreIpoCompanyInterest](#updatepreipocompanyinterest) - Update Pre IPO Company Interest
* [DeletePreIpoCompanyInterest](#deletepreipocompanyinterest) - Delete Pre IPO Company Interest

## CreatePreIpoCompanyInterest

Creates a Pre IPO Company Interest.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyInterests_CreatePreIpoCompanyInterest" method="post" path="/trading/v1/accounts/{account_id}/preIpoCompanyInterests" -->
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

    res, err := s.PreIPOInterests.CreatePreIpoCompanyInterest(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", components.PreIpoCompanyInterestCreate{
        Amount: components.DecimalCreate{},
        PreIpoCompanyID: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PreIpoCompanyInterest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                       |
| `preIpoCompanyInterestCreate`                                                                    | [components.PreIpoCompanyInterestCreate](../../models/components/preipocompanyinterestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.PreIpoCompanyInterestsCreatePreIpoCompanyInterestResponse](../../models/operations/preipocompanyinterestscreatepreipocompanyinterestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 409 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPreIpoCompanyInterests

Lists Pre IPO Company Interests.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyInterests_ListPreIpoCompanyInterests" method="get" path="/trading/v1/accounts/{account_id}/preIpoCompanyInterests" -->
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

    res, err := s.PreIPOInterests.ListPreIpoCompanyInterests(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", ascendsdkgo.Int(50), ascendsdkgo.String("eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ=="))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPreIpoCompanyInterestsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                        |
| `accountID`                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                     | The account id.                                                                                                                                                                                                                                                        | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                                             | **int*                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                     | The maximum number of Pre IPO Company Interests to return. The service may return fewer than this value. If unspecified, at most 100 Pre IPO Company Interests will be returned. The maximum value is 100; values above 100 will be coerced to 100.                    | 50                                                                                                                                                                                                                                                                     |
| `pageToken`                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                     | A page token, received from a previous `ListPreIpoCompanyInterestsRequest` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListPreIpoCompanyInterestsRequest` must match the call that provided the page token. | eyJzaXplIjoxMCwib2Zmc2V0IjoxMDAsInBhcmVudElkIjoicGFyZW50SWQifQ==                                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                        |

### Response

**[*operations.PreIpoCompanyInterestsListPreIpoCompanyInterestsResponse](../../models/operations/preipocompanyinterestslistpreipocompanyinterestsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPreIpoCompanyInterest

Gets a Pre IPO Company Interest.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyInterests_GetPreIpoCompanyInterest" method="get" path="/trading/v1/accounts/{account_id}/preIpoCompanyInterests/{preIpoCompanyInterest_id}" -->
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

    res, err := s.PreIPOInterests.GetPreIpoCompanyInterest(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "3fa85f64-5717-4562-b3fc-2c963f66afa6")
    if err != nil {
        log.Fatal(err)
    }
    if res.PreIpoCompanyInterest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01JR8YQT40WAQT8S95ZQGS1QN0                               |
| `preIpoCompanyInterestID`                                | *string*                                                 | :heavy_check_mark:                                       | The preIpoCompanyInterest id.                            | 3fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PreIpoCompanyInterestsGetPreIpoCompanyInterestResponse](../../models/operations/preipocompanyinterestsgetpreipocompanyinterestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePreIpoCompanyInterest

Updates a Pre IPO Company Interest.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyInterests_UpdatePreIpoCompanyInterest" method="patch" path="/trading/v1/accounts/{account_id}/preIpoCompanyInterests/{preIpoCompanyInterest_id}" -->
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

    res, err := s.PreIPOInterests.UpdatePreIpoCompanyInterest(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "3fa85f64-5717-4562-b3fc-2c963f66afa6", components.PreIpoCompanyInterestUpdate{}, ascendsdkgo.String("[object Object]"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PreIpoCompanyInterest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                       |
| `preIpoCompanyInterestID`                                                                        | *string*                                                                                         | :heavy_check_mark:                                                                               | The preIpoCompanyInterest id.                                                                    | 3fa85f64-5717-4562-b3fc-2c963f66afa6                                                             |
| `preIpoCompanyInterestUpdate`                                                                    | [components.PreIpoCompanyInterestUpdate](../../models/components/preipocompanyinterestupdate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `updateMask`                                                                                     | **string*                                                                                        | :heavy_minus_sign:                                                                               | The list of fields to update. The only updatable field is `amount`.                              | {<br/>"update_mask": "amount"<br/>}                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.PreIpoCompanyInterestsUpdatePreIpoCompanyInterestResponse](../../models/operations/preipocompanyinterestsupdatepreipocompanyinterestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePreIpoCompanyInterest

Deletes a Pre IPO Company Interest.

### Example Usage

<!-- UsageSnippet language="go" operationID="PreIpoCompanyInterests_DeletePreIpoCompanyInterest" method="delete" path="/trading/v1/accounts/{account_id}/preIpoCompanyInterests/{preIpoCompanyInterest_id}" -->
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

    res, err := s.PreIPOInterests.DeletePreIpoCompanyInterest(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", "3fa85f64-5717-4562-b3fc-2c963f66afa6")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01JR8YQT40WAQT8S95ZQGS1QN0                               |
| `preIpoCompanyInterestID`                                | *string*                                                 | :heavy_check_mark:                                       | The preIpoCompanyInterest id.                            | 3fa85f64-5717-4562-b3fc-2c963f66afa6                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.PreIpoCompanyInterestsDeletePreIpoCompanyInterestResponse](../../models/operations/preipocompanyinterestsdeletepreipocompanyinterestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |