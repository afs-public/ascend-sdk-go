# Investigations
(*Investigations*)

## Overview

### Available Operations

* [GetInvestigation](#getinvestigation) - Get Investigations
* [UpdateInvestigation](#updateinvestigation) - Update Investigation 
* [ListInvestigations](#listinvestigations) - List Investigations
* [LinkDocuments](#linkdocuments) - Link Documents
* [GetWatchlistItem](#getwatchlistitem) - Get Watchlist Item
* [GetCustomerIdentification](#getcustomeridentification) - Get Identity Verification

## GetInvestigation

Use this endpoint to get the current state of an investigation by request reference UUID.

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
    res, err := s.Investigations.GetInvestigation(ctx, "01HEWVF4ZSNKYRP293J53ASJCJ")
    if err != nil {
        log.Fatal(err)
    }
    if res.Investigation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `investigationID`                                        | *string*                                                 | :heavy_check_mark:                                       | The investigation id.                                    | 01HEWVF4ZSNKYRP293J53ASJCJ                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.InvestigationServiceGetInvestigationResponse](../../models/operations/investigationservicegetinvestigationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateInvestigation

Use this endpoint to update the details of an investigation by request reference UUID.

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
    res, err := s.Investigations.UpdateInvestigation(ctx, "01HEWVF4ZSNKYRP293J53ASJCJ", components.InvestigationUpdate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Investigation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                               |                                                                                                                                                                                                                                   |
| `investigationID`                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                | The investigation id.                                                                                                                                                                                                             | 01HEWVF4ZSNKYRP293J53ASJCJ                                                                                                                                                                                                        |
| `investigationUpdate`                                                                                                                                                                                                             | [components.InvestigationUpdate](../../models/components/investigationupdate.md)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                | N/A                                                                                                                                                                                                                               |                                                                                                                                                                                                                                   |
| `updateMask`                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                | The list of fields to update. Updatable Fields:<br/>  - identity_verification<br/>  - investigation_request_state<br/>  - watchlist_matches<br/>   - watchlist_id<br/>   - watchlist_item_id<br/>   - match_state<br/>   - exclude_from_screening<br/>  - comment | {<br/>"update_mask": "identity_verification"<br/>}                                                                                                                                                                                |
| `opts`                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                     |                                                                                                                                                                                                                                   |

### Response

**[*operations.InvestigationServiceUpdateInvestigationResponse](../../models/operations/investigationserviceupdateinvestigationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListInvestigations

Use this endpoint to retrieve a list of investigation summaries based on optional search parameters

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
    res, err := s.Investigations.ListInvestigations(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInvestigationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                        | Example                                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                                |                                                                                                                                                                                                                                                    |
| `pageSize`                                                                                                                                                                                                                                         | **int*                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                 | The maximum number of records to return. Default is 50 The maximum is 200, values exceeding this will be set to 200                                                                                                                                | 100                                                                                                                                                                                                                                                |
| `pageToken`                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                 | The page token used to request a specific page of the search results                                                                                                                                                                               |                                                                                                                                                                                                                                                    |
| `filter`                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                 | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> ListInvestigationStatesResponse.investigation_states | person.given_name == 'Jane' && person.family_name == 'Dough'                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                    |

### Response

**[*operations.InvestigationServiceListInvestigationsResponse](../../models/operations/investigationservicelistinvestigationsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## LinkDocuments

Use this endpoint to update identity verification document IDs.

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
    res, err := s.Investigations.LinkDocuments(ctx, "01HEWVF4ZSNKYRP293J53ASJCJ", components.LinkDocumentsRequestCreate{
        IdentityVerificationDocumentIds: []string{
            "0f01ae1f-d24c-4171-8f3f-c0b820bf3044",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkDocumentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `investigationID`                                                                              | *string*                                                                                       | :heavy_check_mark:                                                                             | The investigation id.                                                                          | 01HEWVF4ZSNKYRP293J53ASJCJ                                                                     |
| `linkDocumentsRequestCreate`                                                                   | [components.LinkDocumentsRequestCreate](../../models/components/linkdocumentsrequestcreate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.InvestigationServiceLinkDocumentsResponse](../../models/operations/investigationservicelinkdocumentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetWatchlistItem

Gets the details of an investigation by watchlist type and valid watchlist id

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
    res, err := s.Investigations.GetWatchlistItem(ctx, "DOWJONES", "123456")
    if err != nil {
        log.Fatal(err)
    }
    if res.WatchlistItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `watchlistID`                                            | *string*                                                 | :heavy_check_mark:                                       | The watchlist id.                                        | DOWJONES                                                 |
| `itemID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The item id.                                             | 123456                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.WatchlistServiceGetWatchlistItemResponse](../../models/operations/watchlistservicegetwatchlistitemresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetCustomerIdentification

Gets a CustomerIdentification by CustomerIdentification ID.

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
    res, err := s.Investigations.GetCustomerIdentification(ctx, "01HPMZZM6RKMVZA1JQ63RQKJRP", "01HEWVF4ZSNKYRP293J53ASJCJ", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerIdentification != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                       | Type                                                                                                                                                                                            | Required                                                                                                                                                                                        | Description                                                                                                                                                                                     | Example                                                                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                              | The context to use for the request.                                                                                                                                                             |                                                                                                                                                                                                 |
| `correspondentID`                                                                                                                                                                               | *string*                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                              | The correspondent id.                                                                                                                                                                           | 01HPMZZM6RKMVZA1JQ63RQKJRP                                                                                                                                                                      |
| `customerIdentificationID`                                                                                                                                                                      | *string*                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                              | The customerIdentification id.                                                                                                                                                                  | 01HEWVF4ZSNKYRP293J53ASJCJ                                                                                                                                                                      |
| `view`                                                                                                                                                                                          | [*operations.CustomerIdentificationResultServiceGetCustomerIdentificationQueryParamView](../../models/operations/customeridentificationresultservicegetcustomeridentificationqueryparamview.md) | :heavy_minus_sign:                                                                                                                                                                              | Optional. The view to return. Defaults to BASIC.                                                                                                                                                | BASIC                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                              | The options for this request.                                                                                                                                                                   |                                                                                                                                                                                                 |

### Response

**[*operations.CustomerIdentificationResultServiceGetCustomerIdentificationResponse](../../models/operations/customeridentificationresultservicegetcustomeridentificationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |