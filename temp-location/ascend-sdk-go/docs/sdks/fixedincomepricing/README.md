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
    res, err := s.FixedIncomePricing.PreviewOrderCost(ctx, "<id>", components.OrderCostPreviewRequestCreate{
        AssetType: components.OrderCostPreviewRequestCreateAssetTypeFixedIncome,
        BrokerCapacity: components.OrderCostPreviewRequestCreateBrokerCapacityAgency.ToPointer(),
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## RetrieveQuote

Returns quote information containing the best bid/ask for the given Fixed Income asset.

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
    res, err := s.FixedIncomePricing.RetrieveQuote(ctx, "<id>", components.RetrieveQuoteRequestCreate{
        AssetType: components.RetrieveQuoteRequestCreateAssetTypeFixedIncome,
        BrokerCapacity: components.RetrieveQuoteRequestCreateBrokerCapacityAgency.ToPointer(),
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## RetrieveFixedIncomeMarks

Returns marks for a specified set of Fixed Income assets (up to 100 per request)

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
    res, err := s.FixedIncomePricing.RetrieveFixedIncomeMarks(ctx, "<id>", components.RetrieveFixedIncomeMarksRequestCreate{
        Parent: "<value>",
        SecurityIdentifiers: []components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
            components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
                Identifier: "37833100",
                IdentifierType: components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeCusip,
            },
            components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
                Identifier: "37833100",
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 401, 403, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |