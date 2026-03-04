# AlternativeAccountAccreditation
(*AlternativeAccountAccreditation*)

## Overview

### Available Operations

* [GetAccountAccreditation](#getaccountaccreditation) - Get Account Accreditation
* [SetAccountAccreditationType](#setaccountaccreditationtype) - Set Account Accreditation

## GetAccountAccreditation

Gets the accreditation properties for the specified account.

### Example Usage

<!-- UsageSnippet language="go" operationID="AccountAccreditationService_GetAccountAccreditation" method="get" path="/trading/v1/accounts/{account_id}/accreditation" -->
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

    res, err := s.AlternativeAccountAccreditation.GetAccountAccreditation(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0")
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountAccreditation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01JR8YQT40WAQT8S95ZQGS1QN0                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountAccreditationServiceGetAccountAccreditationResponse](../../models/operations/accountaccreditationservicegetaccountaccreditationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetAccountAccreditationType

Sets the accreditation type for an account. Accounts must be accredited to participate in alternative investment orders.

### Example Usage

<!-- UsageSnippet language="go" operationID="AccountAccreditationService_SetAccountAccreditationType" method="post" path="/trading/v1/accounts/{account_id}/accreditation:setType" -->
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

    res, err := s.AlternativeAccountAccreditation.SetAccountAccreditationType(ctx, "01JR8YQT40WAQT8S95ZQGS1QN0", components.SetAccountAccreditationTypeRequestCreate{
        AccreditationType: components.SetAccountAccreditationTypeRequestCreateAccreditationTypeNetWorthGt1M,
        Name: "accounts/01JR8YQT40WAQT8S95ZQGS1QN0/accreditation",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountAccreditation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `accountID`                                                                                                                | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The account id.                                                                                                            | 01JR8YQT40WAQT8S95ZQGS1QN0                                                                                                 |
| `setAccountAccreditationTypeRequestCreate`                                                                                 | [components.SetAccountAccreditationTypeRequestCreate](../../models/components/setaccountaccreditationtyperequestcreate.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |                                                                                                                            |

### Response

**[*operations.AccountAccreditationServiceSetAccountAccreditationTypeResponse](../../models/operations/accountaccreditationservicesetaccountaccreditationtyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |