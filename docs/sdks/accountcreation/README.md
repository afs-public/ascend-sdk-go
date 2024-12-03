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
	"ascend-sdk/models/components"
	ascendsdk "ascend-sdk"
	"context"
	"log"
)

func main() {
    s := ascendsdk.New(
        ascendsdk.WithSecurity(components.Security{
            APIKey: ascendsdk.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
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
        AcceptsIssuerDirectCommunication: ascendsdk.Bool(false),
        AccountGroupID: "01ARZ3NDEKTSV4RRFFQ69G5FAV",
        Advised: ascendsdk.Bool(true),
        CatAccountHolderType: components.CatAccountHolderTypeIIndividual.ToPointer(),
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
        Identifiers: []components.IdentifierCreate{
            components.IdentifierCreate{
                Type: components.IdentifierCreateTypeOriginatingAccountID,
                Value: "12345678",
            },
        },
        InterestedParties: []components.InterestedPartyCreate{
            components.InterestedPartyCreate{
                MailingAddress: components.PostalAddressCreate{},
                Recipient: "John Dough",
                StatementDeliveryPreference: components.StatementDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.TradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
        },
        InvestmentProfile: &components.InvestmentProfileCreate{
            AccountGoals: &components.AccountGoalsCreate{
                InvestmentObjective: components.InvestmentObjectiveGrowth.ToPointer(),
                LiquidityNeeds: components.LiquidityNeedsVeryImportant.ToPointer(),
                RiskTolerance: components.RiskToleranceMedium.ToPointer(),
                TimeHorizon: components.TimeHorizonAverage.ToPointer(),
            },
            CustomerProfile: &components.CustomerProfileCreate{
                AnnualIncomeRangeUsd: components.AnnualIncomeRangeUsdFrom100KTo200K.ToPointer(),
                FederalTaxBracket: ascendsdk.Float64(1.5),
                InvestmentExperience: components.InvestmentExperienceGood.ToPointer(),
                LiquidNetWorthRangeUsd: components.LiquidNetWorthRangeUsdFrom100KTo200K.ToPointer(),
                TotalNetWorthRangeUsd: components.TotalNetWorthRangeUsdFrom100KTo200K.ToPointer(),
            },
        },
        Managed: ascendsdk.Bool(true),
        Parties: []components.PartyRequestCreate{
            components.PartyRequestCreate{
                EmailAddress: "example@domain.com",
                MailingAddress: components.PostalAddressCreate{},
                PhoneNumber: components.PhoneNumberCreate{},
                ProspectusDeliveryPreference: components.ProspectusDeliveryPreferenceDigital.ToPointer(),
                ProxyDeliveryPreference: components.ProxyDeliveryPreferenceDigital.ToPointer(),
                RelationType: components.RelationTypePrimaryOwner,
                StatementDeliveryPreference: components.PartyRequestCreateStatementDeliveryPreferenceDigital.ToPointer(),
                TaxDocumentDeliveryPreference: components.TaxDocumentDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.PartyRequestCreateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
            components.PartyRequestCreate{
                EmailAddress: "example@domain.com",
                MailingAddress: components.PostalAddressCreate{},
                PhoneNumber: components.PhoneNumberCreate{},
                ProspectusDeliveryPreference: components.ProspectusDeliveryPreferenceDigital.ToPointer(),
                ProxyDeliveryPreference: components.ProxyDeliveryPreferenceDigital.ToPointer(),
                RelationType: components.RelationTypePrimaryOwner,
                StatementDeliveryPreference: components.PartyRequestCreateStatementDeliveryPreferenceDigital.ToPointer(),
                TaxDocumentDeliveryPreference: components.TaxDocumentDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.PartyRequestCreateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
        },
        PrimaryRegisteredRepID: ascendsdk.String("01HB7N66WW02WG3B6B9W29K0HF"),
        TaxProfile: &components.AccountTaxProfileCreate{
            CostBasisLotDisposalMethod: components.CostBasisLotDisposalMethodCostBasisLotDisposalFifo.ToPointer(),
            Section475Election: ascendsdk.Bool(true),
        },
        TrustedContacts: []components.TrustedContactCreate{
            components.TrustedContactCreate{
                EmailAddress: ascendsdk.String("example@email.com"),
                FamilyName: "Doe",
                GivenName: "John",
                MiddleNames: ascendsdk.String("Larry"),
            },
            components.TrustedContactCreate{
                EmailAddress: ascendsdk.String("example@email.com"),
                FamilyName: "Doe",
                GivenName: "John",
                MiddleNames: ascendsdk.String("Larry"),
            },
        },
        WrapFeeBilled: ascendsdk.Bool(true),
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
	"ascend-sdk/models/components"
	ascendsdk "ascend-sdk"
	"context"
	"ascend-sdk/models/operations"
	"log"
)

func main() {
    s := ascendsdk.New(
        ascendsdk.WithSecurity(components.Security{
            APIKey: ascendsdk.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.AccountCreation.GetAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", operations.QueryParamViewFull.ToPointer())
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