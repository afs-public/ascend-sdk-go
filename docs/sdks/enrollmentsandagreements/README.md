# EnrollmentsAndAgreements
(*EnrollmentsAndAgreements*)

## Overview

### Available Operations

* [EnrollAccount](#enrollaccount) - Enroll Account
* [ListAvailableEnrollments](#listavailableenrollments) - List Available Enrollments
* [AccountsListAvailableEnrollmentsByAccountGroup](#accountslistavailableenrollmentsbyaccountgroup) - List Available Enrollments (by Account Group)
* [DeactivateEnrollment](#deactivateenrollment) - Deactivate Enrollment
* [ListEnrollments](#listenrollments) - List Account Enrollments
* [AffirmAgreements](#affirmagreements) - Affirm Agreements
* [ListAgreements](#listagreements) - List Account Agreements
* [ListEntitlements](#listentitlements) - List Account Entitlements

## EnrollAccount

Adds an Enrollment to an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_EnrollAccount" method="post" path="/accounts/v1/accounts/{account_id}:enroll" -->
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

    res, err := s.EnrollmentsAndAgreements.EnrollAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.EnrollAccountRequestCreate{
        Enrollment: components.EnrollmentCreate{
            PrincipalApproverID: "02HB7N66WW02WL3B6B9W29K0HW",
            Type: components.EnrollmentCreateTypeRegistrationIndividual,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EnrollAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `accountID`                                                                                    | *string*                                                                                       | :heavy_check_mark:                                                                             | The account id.                                                                                | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                     |
| `enrollAccountRequestCreate`                                                                   | [components.EnrollAccountRequestCreate](../../models/components/enrollaccountrequestcreate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.AccountsEnrollAccountResponse](../../models/operations/accountsenrollaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAvailableEnrollments

Get a list of Enrollments available for an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_ListAvailableEnrollments" method="get" path="/accounts/v1/accounts/{account_id}/availableEnrollments" -->
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

    res, err := s.EnrollmentsAndAgreements.ListAvailableEnrollments(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdkgo.Int(25), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), ascendsdkgo.String("enrollment_type == \"REGISTRATION_INDIVIDUAL\""))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAvailableEnrollmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                        |
| `accountID`                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                     | The account id.                                                                                                                                                                                                                                        | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                             | **int*                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                     | The maximum number of available enrollments to return. The service may return fewer than this value. The maximum value is 100; values above 100 will be coerced to 100.                                                                                | 25                                                                                                                                                                                                                                                     |
| `pageToken`                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                     | A page token, received from a previous `ListAvailableEnrollments` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListAvailableEnrollments` must match the call that provided the page token. | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                               |
| `filter`                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                     | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `enrollment_type`                                    | enrollment_type == "REGISTRATION_INDIVIDUAL"                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                        |

### Response

**[*operations.AccountsListAvailableEnrollmentsResponse](../../models/operations/accountslistavailableenrollmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AccountsListAvailableEnrollmentsByAccountGroup

Get a list of Enrollments available for an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_ListAvailableEnrollments_1" method="get" path="/accounts/v1/accountGroups/{accountGroup_id}/availableEnrollments" -->
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

    res, err := s.EnrollmentsAndAgreements.AccountsListAvailableEnrollmentsByAccountGroup(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdkgo.Int(25), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), ascendsdkgo.String("enrollment_type == \"REGISTRATION_INDIVIDUAL\""))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAvailableEnrollmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                        |
| `accountGroupID`                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                     | The accountGroup id.                                                                                                                                                                                                                                   | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                                             | **int*                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                     | The maximum number of available enrollments to return. The service may return fewer than this value. The maximum value is 100; values above 100 will be coerced to 100.                                                                                | 25                                                                                                                                                                                                                                                     |
| `pageToken`                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                     | A page token, received from a previous `ListAvailableEnrollments` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListAvailableEnrollments` must match the call that provided the page token. | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                               |
| `filter`                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                     | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `enrollment_type`                                    | enrollment_type == "REGISTRATION_INDIVIDUAL"                                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                        |

### Response

**[*operations.AccountsListAvailableEnrollments1Response](../../models/operations/accountslistavailableenrollments1response.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeactivateEnrollment

Deactivates an Account Enrollment.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_DeactivateEnrollment" method="post" path="/accounts/v1/accounts/{account_id}/enrollments:deactivate" -->
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

    res, err := s.EnrollmentsAndAgreements.DeactivateEnrollment(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.DeactivateEnrollmentRequestCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Enrollment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `accountID`                                                                                                  | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The account id.                                                                                              | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                   |
| `deactivateEnrollmentRequestCreate`                                                                          | [components.DeactivateEnrollmentRequestCreate](../../models/components/deactivateenrollmentrequestcreate.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.AccountsDeactivateEnrollmentResponse](../../models/operations/accountsdeactivateenrollmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEnrollments

Gets a list of Enrollments for an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_ListEnrollments" method="get" path="/accounts/v1/accounts/{account_id}/enrollments" -->
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

    res, err := s.EnrollmentsAndAgreements.ListEnrollments(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdkgo.Int(5), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEnrollmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                  |                                                                                                                                                                                                                                      |
| `accountID`                                                                                                                                                                                                                          | *string*                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                   | The account id.                                                                                                                                                                                                                      | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                                                                                                                                           |
| `pageSize`                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                   | The maximum number of enrollments to return.                                                                                                                                                                                         | 5                                                                                                                                                                                                                                    |
| `pageToken`                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                   | A page token, received from a previous `ListEnrollments` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListEnrollments` must match the call that provided the page token. | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                          |
| `opts`                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                        |                                                                                                                                                                                                                                      |

### Response

**[*operations.AccountsListEnrollmentsResponse](../../models/operations/accountslistenrollmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AffirmAgreements

Affirm Agreements for an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_AffirmAgreements" method="post" path="/accounts/v1/accounts/{account_id}/agreements:affirm" -->
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

    res, err := s.EnrollmentsAndAgreements.AffirmAgreements(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.AffirmAgreementsRequestCreate{
        AccountAgreementIds: []string{
            "fa2f181c-f2fb-4bc2-b75a-79302c634ae5",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AffirmAgreementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The account id.                                                                                      | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                           |
| `affirmAgreementsRequestCreate`                                                                      | [components.AffirmAgreementsRequestCreate](../../models/components/affirmagreementsrequestcreate.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.AccountsAffirmAgreementsResponse](../../models/operations/accountsaffirmagreementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAgreements

Gets a list of Agreements on an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_ListAgreements" method="get" path="/accounts/v1/accounts/{account_id}/agreements" -->
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

    res, err := s.EnrollmentsAndAgreements.ListAgreements(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdkgo.Int(5), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAgreementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                        | Example                                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                |                                                                                                                                                                                                                                    |
| `accountID`                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                 | The account id.                                                                                                                                                                                                                    | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                                                                                                                                         |
| `pageSize`                                                                                                                                                                                                                         | **int*                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                 | The maximum number of agreements to return.                                                                                                                                                                                        | 5                                                                                                                                                                                                                                  |
| `pageToken`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                 | A page token, received from a previous `ListAgreements` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListAgreements` must match the call that provided the page token. | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                        |
| `opts`                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                      |                                                                                                                                                                                                                                    |

### Response

**[*operations.AccountsListAgreementsResponse](../../models/operations/accountslistagreementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEntitlements

Gets a list of Entitlements for an Account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Accounts_ListEntitlements" method="get" path="/accounts/v1/accounts/{account_id}/entitlements" -->
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

    res, err := s.EnrollmentsAndAgreements.ListEntitlements(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdkgo.Int(5), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                    |                                                                                                                                                                                                                                        |
| `accountID`                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                     | The account id.                                                                                                                                                                                                                        | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                                                                                                                                                             |
| `pageSize`                                                                                                                                                                                                                             | **int*                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                     | The maximum number of entitlements to return.                                                                                                                                                                                          | 5                                                                                                                                                                                                                                      |
| `pageToken`                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                     | A page token, received from a previous `ListEntitlements` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListEntitlements` must match the call that provided the page token. | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                          |                                                                                                                                                                                                                                        |

### Response

**[*operations.AccountsListEntitlementsResponse](../../models/operations/accountslistentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |