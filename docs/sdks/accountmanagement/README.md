# AccountManagement
(*AccountManagement*)

## Overview

### Available Operations

* [ListAccounts](#listaccounts) - List Accounts
* [UpdateAccount](#updateaccount) - Update Account
* [AddParty](#addparty) - Add Party
* [UpdateParty](#updateparty) - Update Party
* [ReplaceParty](#replaceparty) - Replace Party
* [RemoveParty](#removeparty) - Remove Party
* [CloseAccount](#closeaccount) - Close Account
* [CreateTrustedContact](#createtrustedcontact) - Create Trusted Contact
* [UpdateTrustedContact](#updatetrustedcontact) - Update Trusted Contact
* [DeleteTrustedContact](#deletetrustedcontact) - Delete Trusted Contact
* [CreateInterestedParty](#createinterestedparty) - Create Interested Party
* [UpdateInterestedParty](#updateinterestedparty) - Update Interested Party
* [DeleteInterestedParty](#deleteinterestedparty) - Delete Interested Party
* [ListAvailableRestrictions](#listavailablerestrictions) - List Available Restrictions
* [CreateRestriction](#createrestriction) - Create Restriction
* [EndRestriction](#endrestriction) - End Restriction

## ListAccounts

Gets a list of Accounts based on search criteria.

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"github.com/afs-public/ascend-sdk-go/models/operations"
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
    res, err := s.AccountManagement.ListAccounts(ctx, operations.AccountsListAccountsRequest{
        PageSize: ascendsdkgo.Int(25),
        PageToken: ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"),
        OrderBy: ascendsdkgo.String("state"),
        Filter: ascendsdkgo.String("account_number == \"R9AHY8P\""),
        View: operations.ViewFull.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAccountsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountsListAccountsRequest](../../models/operations/accountslistaccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AccountsListAccountsResponse](../../models/operations/accountslistaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccount

UPDATE Updates an Account.

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
    res, err := s.AccountManagement.UpdateAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.AccountRequestUpdate{
        AcceptsIssuerDirectCommunication: ascendsdkgo.Bool(false),
        Advised: ascendsdkgo.Bool(true),
        CatAccountHolderType: components.AccountRequestUpdateCatAccountHolderTypeIIndividual.ToPointer(),
        Identifiers: []components.IdentifierUpdate{
            components.IdentifierUpdate{
                Type: components.IdentifierUpdateTypeOriginatingAccountID.ToPointer(),
            },
        },
        InterestedParties: []components.InterestedPartyUpdate{
            components.InterestedPartyUpdate{
                Recipient: ascendsdkgo.String("John Dough"),
                StatementDeliveryPreference: components.InterestedPartyUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
            components.InterestedPartyUpdate{
                Recipient: ascendsdkgo.String("John Dough"),
                StatementDeliveryPreference: components.InterestedPartyUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
            components.InterestedPartyUpdate{
                Recipient: ascendsdkgo.String("John Dough"),
                StatementDeliveryPreference: components.InterestedPartyUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
        },
        InvestmentProfile: &components.InvestmentProfileUpdate{
            AccountGoals: &components.AccountGoalsUpdate{
                InvestmentObjective: components.AccountGoalsUpdateInvestmentObjectiveGrowth.ToPointer(),
                LiquidityNeeds: components.AccountGoalsUpdateLiquidityNeedsVeryImportant.ToPointer(),
                RiskTolerance: components.AccountGoalsUpdateRiskToleranceMedium.ToPointer(),
                TimeHorizon: components.AccountGoalsUpdateTimeHorizonAverage.ToPointer(),
            },
            CustomerProfile: &components.CustomerProfileUpdate{
                AnnualIncomeRangeUsd: components.CustomerProfileUpdateAnnualIncomeRangeUsdFrom100KTo200K.ToPointer(),
                FederalTaxBracket: ascendsdkgo.Float64(1.5),
                InvestmentExperience: components.CustomerProfileUpdateInvestmentExperienceGood.ToPointer(),
                LiquidNetWorthRangeUsd: components.CustomerProfileUpdateLiquidNetWorthRangeUsdFrom100KTo200K.ToPointer(),
                TotalNetWorthRangeUsd: components.CustomerProfileUpdateTotalNetWorthRangeUsdFrom100KTo200K.ToPointer(),
            },
        },
        Managed: ascendsdkgo.Bool(true),
        Parties: []components.PartyRequestUpdate{
            components.PartyRequestUpdate{
                EmailAddress: ascendsdkgo.String("example@domain.com"),
                ProspectusDeliveryPreference: components.PartyRequestUpdateProspectusDeliveryPreferenceDigital.ToPointer(),
                ProxyDeliveryPreference: components.PartyRequestUpdateProxyDeliveryPreferenceDigital.ToPointer(),
                RelationType: components.PartyRequestUpdateRelationTypePrimaryOwner.ToPointer(),
                StatementDeliveryPreference: components.PartyRequestUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TaxDocumentDeliveryPreference: components.PartyRequestUpdateTaxDocumentDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.PartyRequestUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
            components.PartyRequestUpdate{
                EmailAddress: ascendsdkgo.String("example@domain.com"),
                ProspectusDeliveryPreference: components.PartyRequestUpdateProspectusDeliveryPreferenceDigital.ToPointer(),
                ProxyDeliveryPreference: components.PartyRequestUpdateProxyDeliveryPreferenceDigital.ToPointer(),
                RelationType: components.PartyRequestUpdateRelationTypePrimaryOwner.ToPointer(),
                StatementDeliveryPreference: components.PartyRequestUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TaxDocumentDeliveryPreference: components.PartyRequestUpdateTaxDocumentDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.PartyRequestUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
            components.PartyRequestUpdate{
                EmailAddress: ascendsdkgo.String("example@domain.com"),
                ProspectusDeliveryPreference: components.PartyRequestUpdateProspectusDeliveryPreferenceDigital.ToPointer(),
                ProxyDeliveryPreference: components.PartyRequestUpdateProxyDeliveryPreferenceDigital.ToPointer(),
                RelationType: components.PartyRequestUpdateRelationTypePrimaryOwner.ToPointer(),
                StatementDeliveryPreference: components.PartyRequestUpdateStatementDeliveryPreferenceDigital.ToPointer(),
                TaxDocumentDeliveryPreference: components.PartyRequestUpdateTaxDocumentDeliveryPreferenceDigital.ToPointer(),
                TradeConfirmationDeliveryPreference: components.PartyRequestUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
            },
        },
        PrimaryRegisteredRepID: ascendsdkgo.String("01HB7N66WW02WG3B6B9W29K0HF"),
        TaxProfile: &components.AccountTaxProfileUpdate{
            CostBasisLotDisposalMethod: components.AccountTaxProfileUpdateCostBasisLotDisposalMethodCostBasisLotDisposalFifo.ToPointer(),
            Section475Election: ascendsdkgo.Bool(true),
        },
        TrustedContacts: []components.TrustedContactUpdate{
            components.TrustedContactUpdate{
                EmailAddress: ascendsdkgo.String("example@email.com"),
                FamilyName: ascendsdkgo.String("Doe"),
                GivenName: ascendsdkgo.String("John"),
                MiddleNames: ascendsdkgo.String("Larry"),
            },
            components.TrustedContactUpdate{
                EmailAddress: ascendsdkgo.String("example@email.com"),
                FamilyName: ascendsdkgo.String("Doe"),
                GivenName: ascendsdkgo.String("John"),
                MiddleNames: ascendsdkgo.String("Larry"),
            },
        },
        WrapFeeBilled: ascendsdkgo.Bool(true),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `accountID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The account id.                                                                    | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                         |
| `accountRequestUpdate`                                                             | [components.AccountRequestUpdate](../../models/components/accountrequestupdate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.AccountsUpdateAccountResponse](../../models/operations/accountsupdateaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AddParty

Adds a party to an account

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
    res, err := s.AccountManagement.AddParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.AddPartyRequestCreate{
        AuthorizedByPartyIds: []string{
            "8096110d-fb55-4f9d-b883-b84f0b70d3ea",
            "8096110d-fb55-4f9d-b883-b84f0b70d3rb",
        },
        Parent: "accounts/01HC3MAQ4DR9QN1V8MJ4CN1HMK",
        Party: components.PartyRequestCreate{
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Party != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `accountID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | The account id.                                                                      | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                           |
| `addPartyRequestCreate`                                                              | [components.AddPartyRequestCreate](../../models/components/addpartyrequestcreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.AccountsAddPartyResponse](../../models/operations/accountsaddpartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateParty

Updates a Party.

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
    res, err := s.AccountManagement.UpdateParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "a58ddb02-3954-4249-a7d5-1d408def12cf", components.PartyRequestUpdate{
        EmailAddress: ascendsdkgo.String("example@domain.com"),
        ProspectusDeliveryPreference: components.PartyRequestUpdateProspectusDeliveryPreferenceDigital.ToPointer(),
        ProxyDeliveryPreference: components.PartyRequestUpdateProxyDeliveryPreferenceDigital.ToPointer(),
        RelationType: components.PartyRequestUpdateRelationTypePrimaryOwner.ToPointer(),
        StatementDeliveryPreference: components.PartyRequestUpdateStatementDeliveryPreferenceDigital.ToPointer(),
        TaxDocumentDeliveryPreference: components.PartyRequestUpdateTaxDocumentDeliveryPreferenceDigital.ToPointer(),
        TradeConfirmationDeliveryPreference: components.PartyRequestUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Party != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `accountID`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The account id.                                                                | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                     |
| `partyID`                                                                      | *string*                                                                       | :heavy_check_mark:                                                             | The party id.                                                                  | a58ddb02-3954-4249-a7d5-1d408def12cf                                           |
| `partyRequestUpdate`                                                           | [components.PartyRequestUpdate](../../models/components/partyrequestupdate.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AccountsUpdatePartyResponse](../../models/operations/accountsupdatepartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ReplaceParty

Replaces a party on an account

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
    res, err := s.AccountManagement.ReplaceParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "8096110d-fb55-4f9d-b883-b84f0b70d3ea", components.ReplacePartyRequestCreate{
        AuthorizedByPartyIds: []string{
            "8096110d-fb55-4f9d-b883-b84f0b70d3ea",
            "8096110d-fb55-4f9d-b883-b84f0b70d3rb",
        },
        Name: "accounts/01HC3MAQ4DR9QN1V8MJ4CN1HMK/parties/8096110d-fb55-4f9d-b883-b84f0b70d3ea",
        Party: components.PartyRequestCreate{
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Party != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `accountID`                                                                                  | *string*                                                                                     | :heavy_check_mark:                                                                           | The account id.                                                                              | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                   |
| `partyID`                                                                                    | *string*                                                                                     | :heavy_check_mark:                                                                           | The party id.                                                                                | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                                                         |
| `replacePartyRequestCreate`                                                                  | [components.ReplacePartyRequestCreate](../../models/components/replacepartyrequestcreate.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.AccountsReplacePartyResponse](../../models/operations/accountsreplacepartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## RemoveParty

Remove a party from an account

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
    res, err := s.AccountManagement.RemoveParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "8096110d-fb55-4f9d-b883-b84f0b70d3ea", components.RemovePartyRequestCreate{
        AuthorizedByPartyIds: []string{
            "8096110d-fb55-4f9d-b883-b84f0b70d3ea",
            "8096110d-fb55-4f9d-b883-b84f0b70d3rb",
        },
        Name: "accounts/01HC3MAQ4DR9QN1V8MJ4CN1HMK/parties/8096110d-fb55-4f9d-b883-b84f0b70d3ea",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The account id.                                                                            | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                 |
| `partyID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | The party id.                                                                              | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                                                       |
| `removePartyRequestCreate`                                                                 | [components.RemovePartyRequestCreate](../../models/components/removepartyrequestcreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.AccountsRemovePartyResponse](../../models/operations/accountsremovepartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CloseAccount

CUSTOM Places an ACCT_MAINT_CLOSURE_PREP_BY_CORRESPONDENT restriction on the Account ready for closure.

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
    res, err := s.AccountManagement.CloseAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.CloseAccountRequestCreate{
        ReplacedFdid: ascendsdkgo.String("JY3E1OC2J3IA"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `accountID`                                                                                  | *string*                                                                                     | :heavy_check_mark:                                                                           | The account id.                                                                              | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                   |
| `closeAccountRequestCreate`                                                                  | [components.CloseAccountRequestCreate](../../models/components/closeaccountrequestcreate.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.AccountsCloseAccountResponse](../../models/operations/accountscloseaccountresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| sdkerrors.Status             | 400, 403, 404, 409, 500, 503 | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CreateTrustedContact

Creates a new Trusted Contact for an account.

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
    res, err := s.AccountManagement.CreateTrustedContact(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.TrustedContactCreate{
        EmailAddress: ascendsdkgo.String("example@email.com"),
        FamilyName: "Doe",
        GivenName: "John",
        MiddleNames: ascendsdkgo.String("Larry"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `accountID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The account id.                                                                    | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                         |
| `trustedContactCreate`                                                             | [components.TrustedContactCreate](../../models/components/trustedcontactcreate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.AccountsCreateTrustedContactResponse](../../models/operations/accountscreatetrustedcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTrustedContact

Updates a Trusted Contact.

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
    res, err := s.AccountManagement.UpdateTrustedContact(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "8096110d-fb55-4f9d-b883-b84f0b70d3ea", components.TrustedContactUpdate{
        EmailAddress: ascendsdkgo.String("example@email.com"),
        FamilyName: ascendsdkgo.String("Doe"),
        GivenName: ascendsdkgo.String("John"),
        MiddleNames: ascendsdkgo.String("Larry"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrustedContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `accountID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The account id.                                                                    | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                         |
| `trustedContactID`                                                                 | *string*                                                                           | :heavy_check_mark:                                                                 | The trustedContact id.                                                             | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                                               |
| `trustedContactUpdate`                                                             | [components.TrustedContactUpdate](../../models/components/trustedcontactupdate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.AccountsUpdateTrustedContactResponse](../../models/operations/accountsupdatetrustedcontactresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## DeleteTrustedContact

DELETE Deletes a Trusted Contact for an Account.

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
    res, err := s.AccountManagement.DeleteTrustedContact(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "8096110d-fb55-4f9d-b883-b84f0b70d3ea")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                               |
| `trustedContactID`                                       | *string*                                                 | :heavy_check_mark:                                       | The trustedContact id.                                   | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountsDeleteTrustedContactResponse](../../models/operations/accountsdeletetrustedcontactresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CreateInterestedParty

Creates an Interested Party record for an Account.

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
    res, err := s.AccountManagement.CreateInterestedParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.InterestedPartyCreate{
        MailingAddress: components.PostalAddressCreate{},
        Recipient: "John Dough",
        StatementDeliveryPreference: components.StatementDeliveryPreferenceDigital.ToPointer(),
        TradeConfirmationDeliveryPreference: components.TradeConfirmationDeliveryPreferenceDigital.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InterestedParty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `accountID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | The account id.                                                                      | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                           |
| `interestedPartyCreate`                                                              | [components.InterestedPartyCreate](../../models/components/interestedpartycreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.AccountsCreateInterestedPartyResponse](../../models/operations/accountscreateinterestedpartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateInterestedParty

Updates an Interested Party.

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
    res, err := s.AccountManagement.UpdateInterestedParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "ecf44f2f-7030-48ed-b937-c40891ee10c8", components.InterestedPartyUpdate{
        Recipient: ascendsdkgo.String("John Dough"),
        StatementDeliveryPreference: components.InterestedPartyUpdateStatementDeliveryPreferenceDigital.ToPointer(),
        TradeConfirmationDeliveryPreference: components.InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.InterestedParty != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `accountID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | The account id.                                                                      | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                           |
| `interestedPartyID`                                                                  | *string*                                                                             | :heavy_check_mark:                                                                   | The interestedParty id.                                                              | ecf44f2f-7030-48ed-b937-c40891ee10c8                                                 |
| `interestedPartyUpdate`                                                              | [components.InterestedPartyUpdate](../../models/components/interestedpartyupdate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.AccountsUpdateInterestedPartyResponse](../../models/operations/accountsupdateinterestedpartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## DeleteInterestedParty

Deletes an Interested Party associated from an Account.

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
    res, err := s.AccountManagement.DeleteInterestedParty(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "8096110d-fb55-4f9d-b883-b84f0b70d3ea")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                               |
| `interestedPartyID`                                      | *string*                                                 | :heavy_check_mark:                                       | The interestedParty id.                                  | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountsDeleteInterestedPartyResponse](../../models/operations/accountsdeleteinterestedpartyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAvailableRestrictions

Gets a list of possible Restrictions that can be placed on an Account based on Enrollments.

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
    res, err := s.AccountManagement.ListAvailableRestrictions(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAvailableRestrictionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountsListAvailableRestrictionsResponse](../../models/operations/accountslistavailablerestrictionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CreateRestriction

Applies a Restriction to an account that suspends one or more Entitlements.

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
    res, err := s.AccountManagement.CreateRestriction(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.RestrictionCreate{
        CreateReason: "Some reason for creating",
        EndedReason: ascendsdkgo.String("Some reason for removing"),
        RestrictionCode: "MARGIN_CALL_VIOLATION_REG_T",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Restriction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `accountID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | The account id.                                                              | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                   |
| `restrictionCreate`                                                          | [components.RestrictionCreate](../../models/components/restrictioncreate.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.AccountsCreateRestrictionResponse](../../models/operations/accountscreaterestrictionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 409, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## EndRestriction

Ends a Restriction on an Account.

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
    res, err := s.AccountManagement.EndRestriction(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", "FRAUD_SUSPENDED_BY_CORRESPONDENT", components.EndRestrictionRequestCreate{
        Reason: "Reason for ending the restriction",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01HC3MAQ4DR9QN1V8MJ4CN1HMK                                                                       |
| `restrictionID`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | The restriction id.                                                                              | FRAUD_SUSPENDED_BY_CORRESPONDENT                                                                 |
| `endRestrictionRequestCreate`                                                                    | [components.EndRestrictionRequestCreate](../../models/components/endrestrictionrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AccountsEndRestrictionResponse](../../models/operations/accountsendrestrictionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |