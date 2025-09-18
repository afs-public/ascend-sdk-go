# Authentication
(*Authentication*)

## Overview

### Available Operations

* [GenerateServiceAccountToken](#generateserviceaccounttoken) - Generate Service Account Token
* [ListSigningKeys](#listsigningkeys) - List Signing Keys

## GenerateServiceAccountToken

Creates an access token for a service account.

### Example Usage

<!-- UsageSnippet language="go" operationID="Authentication_GenerateServiceAccountToken" method="post" path="/iam/v1/serviceAccounts:generateAccessToken" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := ascendsdkgo.New()

    res, err := s.Authentication.GenerateServiceAccountToken(ctx, components.GenerateServiceAccountTokenRequestCreate{
        Jws: "eyJhbGciOiAiUlMyNTYifQ.eyJuYW1lIjogImpkb3VnaCIsIm9yZ2FuaXphdGlvbiI6ICJjb3JyZXNwb25kZW50cy8xMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0NTY3ODkxMDEiLCJkYXRldGltZSI6ICIyMDI0LTAyLTA1VDIxOjAyOjI3LjkwMTE4MFoifQ.IMy3KmYoG8Ppf+7hXN7tm7J4MrNpQLGL7WCWvhh4nZWAVKkluL3/u3KC6hZ6Mb/5p7Y54CgZ68aWT2BcP5y4VtzIZR1Chm5pxbLfgE4aJuk+FnF6K3Gc3bBjOWCL58pxY2aTb0iU/exDEA1cbMDvbCzmY5kRefDvorLOqgUS/tS2MJ2jv4RlZFPlmHv5PtOruJ8xUW19gEgGhsPXYYeSHFTE1ZlaDvyXrKtpOvlf+FVc2RTuEw529LZnzwH4/eJJR3BpSpHyJTjQqiaMT3wzpXXYKfCRqnDkSSKJDzCzTb0/uWK/Lf0uafxPXk5YLdis+dbo1zNQhVVKjwnMpk1vLw",
    }, operations.AuthenticationGenerateServiceAccountTokenSecurity{
        APIKeyAuth: "<YOUR_API_KEY_HERE>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Token != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [components.GenerateServiceAccountTokenRequestCreate](../../models/components/generateserviceaccounttokenrequestcreate.md)                   | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `security`                                                                                                                                   | [operations.AuthenticationGenerateServiceAccountTokenSecurity](../../models/operations/authenticationgenerateserviceaccounttokensecurity.md) | :heavy_check_mark:                                                                                                                           | The security requirements to use for the request.                                                                                            |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.AuthenticationGenerateServiceAccountTokenResponse](../../models/operations/authenticationgenerateserviceaccounttokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401           | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSigningKeys

Gets the public signing keys used to verify JSON Web Tokens generated by this service.

### Example Usage

<!-- UsageSnippet language="go" operationID="Authentication_ListSigningKeys" method="get" path="/iam/v1/keys" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := ascendsdkgo.New()

    res, err := s.Authentication.ListSigningKeys(ctx, operations.AuthenticationListSigningKeysSecurity{
        APIKeyAuth: "<YOUR_API_KEY_HERE>",
    }, ascendsdkgo.Int(50), ascendsdkgo.String("ZXhhbXBsZQo"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSigningKeysResponse != nil {
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `security`                                                                                                           | [operations.AuthenticationListSigningKeysSecurity](../../models/operations/authenticationlistsigningkeyssecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |                                                                                                                      |
| `pageSize`                                                                                                           | **int*                                                                                                               | :heavy_minus_sign:                                                                                                   | The number of entries to return in a single page; Default = 100; Maximum = 1000                                      | 50                                                                                                                   |
| `pageToken`                                                                                                          | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Page token used for pagination; Supplying a page token returns the next page of results                              | ZXhhbXBsZQo                                                                                                          |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.AuthenticationListSigningKeysResponse](../../models/operations/authenticationlistsigningkeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |