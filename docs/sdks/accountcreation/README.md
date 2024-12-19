# AccountCreation
(*AccountCreation*)

## Overview

### Available Operations

* [CreateAccount](#createaccount) - Create Account
* [GetAccount](#getaccount) - Get Account

## CreateAccount

CREATE Creates an account.

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
    res, err := s.AccountCreation.CreateAccount(ctx, components.AccountRequestCreate{
        AccountGroupID: "01ARZ3NDEKTSV4RRFFQ69G5FAV",
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
        Parties: []components.PartyRequestCreate{
            components.PartyRequestCreate{
                EmailAddress: "example@domain.com",
                MailingAddress: components.PostalAddressCreate{},
                PhoneNumber: components.PhoneNumberCreate{},
                RelationType: components.RelationTypePrimaryOwner,
            },
            components.PartyRequestCreate{
                EmailAddress: "example@domain.com",
                MailingAddress: components.PostalAddressCreate{},
                PhoneNumber: components.PhoneNumberCreate{},
                RelationType: components.RelationTypePrimaryOwner,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.AccountRequestCreate](../../models/components/accountrequestcreate.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AccountsCreateAccountResponse](../../models/operations/accountscreateaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccount

READ Get Account

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
    res, err := s.AccountCreation.GetAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |                                                                         |
| `accountID`                                                             | *string*                                                                | :heavy_check_mark:                                                      | The account id.                                                         | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                              |
| `view`                                                                  | [*operations.QueryParamView](../../models/operations/queryparamview.md) | :heavy_minus_sign:                                                      | The view to return. Defaults to `FULL`.                                 | FULL                                                                    |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |                                                                         |

### Response

**[*operations.AccountsGetAccountResponse](../../models/operations/accountsgetaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |