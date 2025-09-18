# Retirements
(*Retirements*)

## Overview

### Available Operations

* [ListContributionSummaries](#listcontributionsummaries) - List Contribution Summaries
* [RetrieveContributionConstraints](#retrievecontributionconstraints) - Retrieve Contribution Constraints
* [ListDistributionSummaries](#listdistributionsummaries) - List Distribution Summaries
* [RetrieveDistributionConstraints](#retrievedistributionconstraints) - Retrieve Distribution Constraints

## ListContributionSummaries

Lists the aggregated retirement contribution summaries by tax year

### Example Usage

<!-- UsageSnippet language="go" operationID="RetirementConstraints_ListContributionSummaries" method="get" path="/transfers/v1/accounts/{account_id}/contributionSummaries" -->
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

    res, err := s.Retirements.ListContributionSummaries(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", ascendsdkgo.Int(2), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZ3hh"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListContributionSummariesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `accountID`                                                                                                               | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The account id.                                                                                                           | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                |
| `pageSize`                                                                                                                | **int*                                                                                                                    | :heavy_minus_sign:                                                                                                        | Number of contribution summaries to get (partitioned by tax year) Default = 2 (current year and prior year), maximum = 10 | 2                                                                                                                         |
| `pageToken`                                                                                                               | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | When paginating, this is used to retrieve a specific page from the overall response                                       | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZ3hh                                                                  |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.RetirementConstraintsListContributionSummariesResponse](../../models/operations/retirementconstraintslistcontributionsummariesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RetrieveContributionConstraints

Retrieves retirement contribution constraints for an account

### Example Usage

<!-- UsageSnippet language="go" operationID="RetirementConstraints_RetrieveContributionConstraints" method="post" path="/transfers/v1/accounts/{account_id}:retrieveContributionConstraints" -->
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

    res, err := s.Retirements.RetrieveContributionConstraints(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.RetrieveContributionConstraintsRequestCreate{
        Mechanism: components.MechanismAch,
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ContributionConstraints != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `accountID`                                                                                                                        | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | The account id.                                                                                                                    | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                         |
| `retrieveContributionConstraintsRequestCreate`                                                                                     | [components.RetrieveContributionConstraintsRequestCreate](../../models/components/retrievecontributionconstraintsrequestcreate.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.RetirementConstraintsRetrieveContributionConstraintsResponse](../../models/operations/retirementconstraintsretrievecontributionconstraintsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDistributionSummaries

Lists the aggregated retirement distribution summaries by tax year

### Example Usage

<!-- UsageSnippet language="go" operationID="RetirementConstraints_ListDistributionSummaries" method="get" path="/transfers/v1/accounts/{account_id}/distributionSummaries" -->
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

    res, err := s.Retirements.ListDistributionSummaries(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", ascendsdkgo.Int(2), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZ3hh"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDistributionSummariesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `accountID`                                                                                                               | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The account id.                                                                                                           | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                |
| `pageSize`                                                                                                                | **int*                                                                                                                    | :heavy_minus_sign:                                                                                                        | Number of distribution summaries to get (partitioned by tax year) Default = 2 (current year and prior year), maximum = 10 | 2                                                                                                                         |
| `pageToken`                                                                                                               | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | When paginating, this is used to retrieve a specific page from the overall response                                       | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZ3hh                                                                  |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.RetirementConstraintsListDistributionSummariesResponse](../../models/operations/retirementconstraintslistdistributionsummariesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RetrieveDistributionConstraints

Retrieves retirement distribution constraints for an account

### Example Usage

<!-- UsageSnippet language="go" operationID="RetirementConstraints_RetrieveDistributionConstraints" method="post" path="/transfers/v1/accounts/{account_id}:retrieveDistributionConstraints" -->
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

    res, err := s.Retirements.RetrieveDistributionConstraints(ctx, "01H8FM6EXVH77SAW3TC8KAWMES", components.RetrieveDistributionConstraintsRequestCreate{
        Mechanism: components.RetrieveDistributionConstraintsRequestCreateMechanismAch,
        Name: "accounts/01H8FM6EXVH77SAW3TC8KAWMES",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DistributionConstraints != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        | Example                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |                                                                                                                                    |
| `accountID`                                                                                                                        | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | The account id.                                                                                                                    | 01H8FM6EXVH77SAW3TC8KAWMES                                                                                                         |
| `retrieveDistributionConstraintsRequestCreate`                                                                                     | [components.RetrieveDistributionConstraintsRequestCreate](../../models/components/retrievedistributionconstraintsrequestcreate.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |                                                                                                                                    |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |                                                                                                                                    |

### Response

**[*operations.RetirementConstraintsRetrieveDistributionConstraintsResponse](../../models/operations/retirementconstraintsretrievedistributionconstraintsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |