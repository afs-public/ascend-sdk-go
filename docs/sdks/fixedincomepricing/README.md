# FixedIncomePricing
(*FixedIncomePricing*)

## Overview

### Available Operations

* [PreviewOrderCost](#previewordercost) - Preview Order Cost
* [RetrieveQuote](#retrievequote) - Retrieve Quote
* [RetrieveFixedIncomeMarks](#retrievefixedincomemarks) - Retrieve Fixed Income Marks

## PreviewOrderCost

Returns a calculation estimating the costs involved in ordering a given quantity of a Fixed Income asset at a specified limit price.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderPriceService_PreviewOrderCost" method="post" path="/trading/v1/accounts/{account_id}/orders:previewOrderCost" -->
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

    res, err := s.FixedIncomePricing.PreviewOrderCost(ctx, "<id>", components.OrderCostPreviewRequestCreate{
        AssetType: components.OrderCostPreviewRequestCreateAssetTypeFixedIncome,
        Identifier: "37833100",
        IdentifierType: components.OrderCostPreviewRequestCreateIdentifierTypeCusip,
        LimitPrice: components.LimitPriceCreate{
            Price: components.DecimalCreate{},
            Type: components.LimitPriceCreateTypePricePerUnit,
        },
        Parent: "<value>",
        Quantity: components.DecimalCreate{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderCostPreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The account id.                                                                                      |
| `orderCostPreviewRequestCreate`                                                                      | [components.OrderCostPreviewRequestCreate](../../models/components/ordercostpreviewrequestcreate.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.OrderPriceServicePreviewOrderCostResponse](../../models/operations/orderpriceservicepreviewordercostresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RetrieveQuote

Returns quote information containing the best bid/ask for the given Fixed Income asset.

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderPriceService_RetrieveQuote" method="post" path="/trading/v1/accounts/{account_id}/orders:retrieveAssetQuote" -->
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

    res, err := s.FixedIncomePricing.RetrieveQuote(ctx, "<id>", components.RetrieveQuoteRequestCreate{
        AssetType: components.RetrieveQuoteRequestCreateAssetTypeFixedIncome,
        Identifier: "37833100",
        IdentifierType: components.RetrieveQuoteRequestCreateIdentifierTypeCusip,
        Parent: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `accountID`                                                                                    | *string*                                                                                       | :heavy_check_mark:                                                                             | The account id.                                                                                |
| `retrieveQuoteRequestCreate`                                                                   | [components.RetrieveQuoteRequestCreate](../../models/components/retrievequoterequestcreate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.OrderPriceServiceRetrieveQuoteResponse](../../models/operations/orderpriceserviceretrievequoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RetrieveFixedIncomeMarks

Returns marks for a specified set of Fixed Income assets (up to 100 per request)

### Example Usage

<!-- UsageSnippet language="go" operationID="OrderPriceService_RetrieveFixedIncomeMarks" method="post" path="/trading/v1/correspondents/{correspondent_id}/prices:retrieveFixedIncomeMarks" -->
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

    res, err := s.FixedIncomePricing.RetrieveFixedIncomeMarks(ctx, "<id>", components.RetrieveFixedIncomeMarksRequestCreate{
        Parent: "<value>",
        SecurityIdentifiers: []components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
            components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
                Identifier: "US0378331005",
                IdentifierType: components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeIsin,
            },
            components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
                Identifier: "38259P508",
                IdentifierType: components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeCusip,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RetrieveFixedIncomeMarksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `correspondentID`                                                                                                    | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The correspondent id.                                                                                                |
| `retrieveFixedIncomeMarksRequestCreate`                                                                              | [components.RetrieveFixedIncomeMarksRequestCreate](../../models/components/retrievefixedincomemarksrequestcreate.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.OrderPriceServiceRetrieveFixedIncomeMarksResponse](../../models/operations/orderpriceserviceretrievefixedincomemarksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500, 503           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |