# Investigations
(*Investigations*)

## Overview

### Available Operations

* [GetInvestigation](#getinvestigation) - Get Investigations
* [UpdateInvestigation](#updateinvestigation) - Update Investigation 
* [ListInvestigations](#listinvestigations) - List Investigations
* [LinkDocuments](#linkdocuments) - Link Documents
* [GetWatchlistItem](#getwatchlistitem) - Get Watchlist Item

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
    res, err := s.Investigations.UpdateInvestigation(ctx, "01HEWVF4ZSNKYRP293J53ASJCJ", components.InvestigationUpdate{
        Comment: ascendsdkgo.String("Updating Watchlist matches"),
        CorrespondentID: ascendsdkgo.String("01HAT5GANHSZ8E8J0RAHQ8BK9K"),
        Entity: &components.EntityUpdate{
            DbaNames: []string{
                "Some Company Alias",
            },
            EmailAddresses: []string{
                "jdough@domain.com",
            },
            Identifications: []components.EntityIdentificationUpdate{
                components.EntityIdentificationUpdate{
                    AdministrativeArea: ascendsdkgo.String("TX"),
                    DocumentReferenceIds: []string{
                        "0f01ae1f-d24c-4171-8f3f-c0b820bf3044",
                    },
                    RegionCode: ascendsdkgo.String("US"),
                    Type: components.EntityIdentificationUpdateTypeEin.ToPointer(),
                    Value: ascendsdkgo.String("666-12-3456"),
                },
            },
            LegalName: ascendsdkgo.String("Enterprises, Inc"),
            OperatingRegionCodes: []string{
                "US",
            },
            PhoneNumbers: []string{
                "214-765-1010",
            },
            RegistrationRegion: ascendsdkgo.String("US"),
        },
        IdentityVerification: components.InvestigationUpdateIdentityVerificationPassed.ToPointer(),
        IdentityVerificationScope: components.InvestigationUpdateIdentityVerificationScopePerformedByApex.ToPointer(),
        Person: &components.PersonUpdate{
            BirthRegionCode: ascendsdkgo.String("US"),
            CitizenshipRegionCodes: []string{
                "US",
            },
            DocumentIds: []string{
                "0f01ae1f-d24c-4171-8f3f-c0b820bf3044",
            },
            EmailAddresses: []string{
                "jdough@domain.com",
            },
            FamilyName: ascendsdkgo.String("Dough"),
            GivenName: ascendsdkgo.String("John"),
            Identifications: []components.PersonIdentificationUpdate{
                components.PersonIdentificationUpdate{
                    AdministrativeArea: ascendsdkgo.String("TX"),
                    DocumentReferenceIds: []string{
                        "0f01ae1f-d24c-4171-8f3f-c0b820bf3044",
                    },
                    RegionCode: ascendsdkgo.String("US"),
                    Type: components.PersonIdentificationUpdateTypeSsn.ToPointer(),
                    Value: ascendsdkgo.String("666-12-3456"),
                },
            },
            MiddleNames: ascendsdkgo.String("Jacob"),
            NameSuffix: components.PersonUpdateNameSuffixJr.ToPointer(),
            PhoneNumbers: []string{
                "214-765-1010",
            },
            ProvidedIdentityVerification: &components.ProvidedIdentityVerificationUpdate{
                AddressVerified: ascendsdkgo.Bool(true),
                BirthDateVerified: ascendsdkgo.Bool(true),
                ExternalCaseID: ascendsdkgo.String("123456"),
                IdentityVerificationDocumentIds: []string{
                    "0f01ae1f-d24c-4171-8f3f-c0b820bf3044",
                },
                NameVerified: ascendsdkgo.Bool(true),
                ProvidedIdentityVerificationID: ascendsdkgo.String("123456"),
                RawVendorDataDocumentID: ascendsdkgo.String("123456"),
                TaxIDVerified: ascendsdkgo.Bool(true),
                Vendor: ascendsdkgo.String("Your identity verification Vendor"),
            },
        },
        WatchlistMatches: []components.WatchlistMatchUpdate{
            components.WatchlistMatchUpdate{
                ExcludeFromScreening: ascendsdkgo.Bool(false),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Investigation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `investigationID`                                                                | *string*                                                                         | :heavy_check_mark:                                                               | The investigation id.                                                            | 01HEWVF4ZSNKYRP293J53ASJCJ                                                       |
| `investigationUpdate`                                                            | [components.InvestigationUpdate](../../models/components/investigationupdate.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

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
    res, err := s.Investigations.ListInvestigations(ctx, ascendsdkgo.Int(100), nil, ascendsdkgo.String("investigation_subject.person_investigation.given_name == 'Jane' && investigation_subject.person_investigation.family_name == 'Dough'"))
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
| `filter`                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                 | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> ListInvestigationStatesResponse.investigation_states | investigation_subject.person_investigation.given_name == 'Jane' && investigation_subject.person_investigation.family_name == 'Dough'                                                                                                               |
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