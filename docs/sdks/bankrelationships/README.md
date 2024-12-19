# BankRelationships
(*BankRelationships*)

## Overview

### Available Operations

* [CreateBankRelationship](#createbankrelationship) - Create Bank Relationship
* [ListBankRelationships](#listbankrelationships) - List Bank Relationships
* [GetBankRelationship](#getbankrelationship) - Get Bank Relationship
* [UpdateBankRelationship](#updatebankrelationship) - Update Bank Relationship
* [CancelBankRelationship](#cancelbankrelationship) - Cancel Bank Relationship
* [VerifyMicroDeposits](#verifymicrodeposits) - Verify Micro Deposits
* [ReissueMicroDeposits](#reissuemicrodeposits) - Reissue Micro Deposits

## CreateBankRelationship

Creates a bank relationship.

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
    res, err := s.BankRelationships.CreateBankRelationship(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.BankRelationshipCreate{
        Nickname: "My Primary Bank",
        VerificationMethod: components.VerificationMethodMicroDeposit,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `accountID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | The account id.                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                             |
| `bankRelationshipCreate`                                                               | [components.BankRelationshipCreate](../../models/components/bankrelationshipcreate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.BankRelationshipsCreateBankRelationshipResponse](../../models/operations/bankrelationshipscreatebankrelationshipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBankRelationships

Lists bank relationships for an account.

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
    res, err := s.BankRelationships.ListBankRelationships(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListBankRelationshipsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                    | Example                                                                                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                                                             | The context to use for the request.                                                                                                                                                                                                            |                                                                                                                                                                                                                                                |
| `accountID`                                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                             | The account id.                                                                                                                                                                                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                                                     |
| `pageSize`                                                                                                                                                                                                                                     | **int*                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                             | The maximum number of bank relationships to return. The service may return fewer than this value. If unspecified, at most 50 bank relationships will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.        | 100                                                                                                                                                                                                                                            |
| `pageToken`                                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                             | A page token, received from a previous `ListBankRelationships` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListBankRelationships` must match the call that provided the page token. | CMFRGgYQup3BhQgaCSkAQCKS7AAAAA==                                                                                                                                                                                                               |
| `state`                                                                                                                                                                                                                                        | [*operations.State](../../models/operations/state.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                             | The state of bank relationships to filter by. Unspecified returns relationships of all states.                                                                                                                                                 | APPROVED                                                                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                             | The options for this request.                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                |

### Response

**[*operations.BankRelationshipsListBankRelationshipsResponse](../../models/operations/bankrelationshipslistbankrelationshipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBankRelationship

Gets an existing bank relationship.

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
    res, err := s.BankRelationships.GetBankRelationship(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e")
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `bankRelationshipID`                                     | *string*                                                 | :heavy_check_mark:                                       | The bankRelationship id.                                 | 651ef9de0dee00240813e60e                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BankRelationshipsGetBankRelationshipResponse](../../models/operations/bankrelationshipsgetbankrelationshipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateBankRelationship

Updates an existing bank relationship.

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
    res, err := s.BankRelationships.UpdateBankRelationship(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e", components.BankRelationshipUpdate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `accountID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | The account id.                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                             |
| `bankRelationshipID`                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | The bankRelationship id.                                                               | 651ef9de0dee00240813e60e                                                               |
| `bankRelationshipUpdate`                                                               | [components.BankRelationshipUpdate](../../models/components/bankrelationshipupdate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `updateMask`                                                                           | **string*                                                                              | :heavy_minus_sign:                                                                     | The field of the bank relationship to update. Only `nickname` is supported.            | {<br/>"update_mask": {<br/>"paths": [<br/>"nickname"<br/>]<br/>}<br/>}                 |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.BankRelationshipsUpdateBankRelationshipResponse](../../models/operations/bankrelationshipsupdatebankrelationshipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelBankRelationship

Cancels an existing bank relationship.

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
    res, err := s.BankRelationships.CancelBankRelationship(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e", components.CancelBankRelationshipRequestCreate{
        Comment: "User Request",
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `accountID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The account id.                                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                       |
| `bankRelationshipID`                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The bankRelationship id.                                                                                         | 651ef9de0dee00240813e60e                                                                                         |
| `cancelBankRelationshipRequestCreate`                                                                            | [components.CancelBankRelationshipRequestCreate](../../models/components/cancelbankrelationshiprequestcreate.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.BankRelationshipsCancelBankRelationshipResponse](../../models/operations/bankrelationshipscancelbankrelationshipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## VerifyMicroDeposits

Verifies a pending bank relationship with the `MICRO_DEPOSIT` verification method. Successful verification of the micro deposit amounts will result in the relationship moving to the `APPROVED` state. The micro deposits will be taken back from the bank account.

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
    res, err := s.BankRelationships.VerifyMicroDeposits(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e", components.VerifyMicroDepositsRequestCreate{
        Amounts: components.MicroDepositAmountsCreate{
            Amount1: components.DecimalCreate{},
            Amount2: components.DecimalCreate{},
        },
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `accountID`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The account id.                                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                 |
| `bankRelationshipID`                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The bankRelationship id.                                                                                   | 651ef9de0dee00240813e60e                                                                                   |
| `verifyMicroDepositsRequestCreate`                                                                         | [components.VerifyMicroDepositsRequestCreate](../../models/components/verifymicrodepositsrequestcreate.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.BankRelationshipsVerifyMicroDepositsResponse](../../models/operations/bankrelationshipsverifymicrodepositsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ReissueMicroDeposits

Reissues micro deposits after micro deposit verification has failed. The user should have received a message that new micro deposits should be reissued.

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
    res, err := s.BankRelationships.ReissueMicroDeposits(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e", components.ReissueMicroDepositsRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankRelationship != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `accountID`                                                                                                  | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The account id.                                                                                              | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                   |
| `bankRelationshipID`                                                                                         | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The bankRelationship id.                                                                                     | 651ef9de0dee00240813e60e                                                                                     |
| `reissueMicroDepositsRequestCreate`                                                                          | [components.ReissueMicroDepositsRequestCreate](../../models/components/reissuemicrodepositsrequestcreate.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.BankRelationshipsReissueMicroDepositsResponse](../../models/operations/bankrelationshipsreissuemicrodepositsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |