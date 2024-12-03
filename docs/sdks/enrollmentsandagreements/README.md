# EnrollmentsAndAgreements
(*EnrollmentsAndAgreements*)

## Overview

### Available Operations

* [EnrollAccount](#enrollaccount) - Enroll Account
* [ListAvailableEnrollments](#listavailableenrollments) - List Available Enrollments
* [DeactivateEnrollment](#deactivateenrollment) - Deactivate Enrollment
* [ListEnrollments](#listenrollments) - List Account Enrollments
* [AffirmAgreements](#affirmagreements) - Affirm Agreements
* [ListAgreements](#listagreements) - List Account Agreements
* [ListEntitlements](#listentitlements) - List Account Entitlements

## EnrollAccount

Adds an Enrollment to an Account.

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
    res, err := s.EnrollmentsAndAgreements.EnrollAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.EnrollAccountRequestCreate{
        Enrollment: components.EnrollmentCreate{
            BeneficiaryEnrollmentMetadata: &components.BeneficiaryEnrollmentMetadataCreate{
                ContingentBeneficiaries: []components.BeneficiaryCreate{
                    components.BeneficiaryCreate{
                        BeneficiaryPercentage: 100,
                        BirthDate: &components.DateCreate{},
                        Email: ascendsdk.String("example@email.com"),
                        EntityName: ascendsdk.String("Acme, Inc."),
                        EntityType: components.BeneficiaryCreateEntityTypeCorporation.ToPointer(),
                        FamilyName: ascendsdk.String("Smith"),
                        GivenName: ascendsdk.String("Bob"),
                        MiddleNames: ascendsdk.String("Robert"),
                        RelationType: components.BeneficiaryCreateRelationTypeSpouse.ToPointer(),
                        TaxID: ascendsdk.String("123456789"),
                        TaxIDType: components.BeneficiaryCreateTaxIDTypeTaxIDTypeUnspecified.ToPointer(),
                    },
                    components.BeneficiaryCreate{
                        BeneficiaryPercentage: 100,
                        BirthDate: &components.DateCreate{},
                        Email: ascendsdk.String("example@email.com"),
                        EntityName: ascendsdk.String("Acme, Inc."),
                        EntityType: components.BeneficiaryCreateEntityTypeCorporation.ToPointer(),
                        FamilyName: ascendsdk.String("Smith"),
                        GivenName: ascendsdk.String("Bob"),
                        MiddleNames: ascendsdk.String("Robert"),
                        RelationType: components.BeneficiaryCreateRelationTypeSpouse.ToPointer(),
                        TaxID: ascendsdk.String("123456789"),
                        TaxIDType: components.BeneficiaryCreateTaxIDTypeTaxIDTypeUnspecified.ToPointer(),
                    },
                },
                PrimaryBeneficiaries: []components.BeneficiaryCreate{
                    components.BeneficiaryCreate{
                        BeneficiaryPercentage: 100,
                        BirthDate: &components.DateCreate{},
                        Email: ascendsdk.String("example@email.com"),
                        EntityName: ascendsdk.String("Acme, Inc."),
                        EntityType: components.BeneficiaryCreateEntityTypeCorporation.ToPointer(),
                        FamilyName: ascendsdk.String("Smith"),
                        GivenName: ascendsdk.String("Bob"),
                        MiddleNames: ascendsdk.String("Robert"),
                        RelationType: components.BeneficiaryCreateRelationTypeSpouse.ToPointer(),
                        TaxID: ascendsdk.String("123456789"),
                        TaxIDType: components.BeneficiaryCreateTaxIDTypeTaxIDTypeUnspecified.ToPointer(),
                    },
                    components.BeneficiaryCreate{
                        BeneficiaryPercentage: 100,
                        BirthDate: &components.DateCreate{},
                        Email: ascendsdk.String("example@email.com"),
                        EntityName: ascendsdk.String("Acme, Inc."),
                        EntityType: components.BeneficiaryCreateEntityTypeCorporation.ToPointer(),
                        FamilyName: ascendsdk.String("Smith"),
                        GivenName: ascendsdk.String("Bob"),
                        MiddleNames: ascendsdk.String("Robert"),
                        RelationType: components.BeneficiaryCreateRelationTypeSpouse.ToPointer(),
                        TaxID: ascendsdk.String("123456789"),
                        TaxIDType: components.BeneficiaryCreateTaxIDTypeTaxIDTypeUnspecified.ToPointer(),
                    },
                    components.BeneficiaryCreate{
                        BeneficiaryPercentage: 100,
                        BirthDate: &components.DateCreate{},
                        Email: ascendsdk.String("example@email.com"),
                        EntityName: ascendsdk.String("Acme, Inc."),
                        EntityType: components.BeneficiaryCreateEntityTypeCorporation.ToPointer(),
                        FamilyName: ascendsdk.String("Smith"),
                        GivenName: ascendsdk.String("Bob"),
                        MiddleNames: ascendsdk.String("Robert"),
                        RelationType: components.BeneficiaryCreateRelationTypeSpouse.ToPointer(),
                        TaxID: ascendsdk.String("123456789"),
                        TaxIDType: components.BeneficiaryCreateTaxIDTypeTaxIDTypeEin.ToPointer(),
                    },
                },
            },
            ConsentMethod: components.EnrollmentCreateConsentMethodNegativeConsentConversion.ToPointer(),
            CorporationEnrollmentMetadata: &components.CorporationEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.DividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                EddAccountEnrollmentMetadata: &components.EddAccountEnrollmentMetadataCreate{
                    DepositedFunds: components.DepositedFundsCreate{
                        InitialDepositAmount: components.DecimalCreate{},
                        InitialDepositSource: "Product Revenue",
                    },
                    DeterminedAccountRiskRating: components.DeterminedAccountRiskRatingHigh,
                    FinancialProfile: components.FinancialProfileCreate{
                        BankingRelationships: []string{
                            "<value>",
                            "<value>",
                        },
                        OtherAccounts: components.OtherAccountsCreate{
                            OwnerHasOtherAccountsAtApex: true,
                        },
                        PrimarySourceOfDepositedFunds: ascendsdk.String("Corporate Income"),
                    },
                    PlannedActivity: components.PlannedActivityCreate{
                        ForeignBondTradingDetails: components.ForeignBondTradingDetailsCreate{
                            ForeignBondTrading: true,
                            ForeignBondTradingDetail: []components.ForeignBondTradingDetailCreate{
                                components.ForeignBondTradingDetailCreate{
                                    Percentage: components.DecimalCreate{},
                                    RegionCode: "CA",
                                },
                            },
                        },
                        LowPricedSecurities: components.LowPricedSecuritiesCreate{
                            LowPricedSecurities: true,
                        },
                        PrimaryAccountActivityType: components.PrimaryAccountActivityTypeActiveTrading,
                        WithdrawalFrequency: components.WithdrawalFrequencyFrequent,
                    },
                    RelatedPepDetails: components.RelatedPepDetailsCreate{
                        DirectOrIndirectRelatedPeps: true,
                        RelatedPeps: []components.RelatedPepCreate{
                            components.RelatedPepCreate{
                                ImmediateFamilyMembers: []components.ImmediateFamilyMemberCreate{
                                    components.ImmediateFamilyMemberCreate{
                                        FamilyMemberName: "Ellen Chen",
                                        Relationship: "Daughter",
                                    },
                                },
                                Organization: "U.S. Embassy",
                                PepName: "Juan Octavio",
                                Title: "U.S. Ambassador to Japan",
                            },
                        },
                    },
                    ScopeOfBusiness: "Financial Services",
                },
                FdicCashSweep: components.FdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            CustodialEnrollmentMetadata: &components.CustodialEnrollmentMetadataCreate{
                CustodialType: components.CustodialTypeUgma,
                DividendReinvestmentPlan: components.CustodialEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.CustodialEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            EstateEnrollmentMetadata: &components.EstateEnrollmentMetadataCreate{
                CertificateOfAppointmentDocumentID: ascendsdk.String("c401f3b2-cdb5-4a6c-9f5f-aa393cf12583"),
                DividendReinvestmentPlan: components.EstateEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
            },
            ForeignIndividualAccountEnrollmentMetadata: &components.ForeignIndividualAccountEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.ForeignIndividualAccountEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.ForeignIndividualAccountEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
                ForeignNaturalPersonAccountEnrollmentMetadata: components.ForeignNaturalPersonAccountEnrollmentMetadataCreate{
                    DepositedFunds: components.DepositedFundsCreate{
                        InitialDepositAmount: components.DecimalCreate{},
                        InitialDepositSource: "Product Revenue",
                    },
                    FinancialProfile: components.FinancialProfileCreate{
                        BankingRelationships: []string{
                            "<value>",
                            "<value>",
                        },
                        OtherAccounts: components.OtherAccountsCreate{
                            OwnerHasOtherAccountsAtApex: true,
                        },
                        PrimarySourceOfDepositedFunds: ascendsdk.String("Corporate Income"),
                    },
                    PlannedActivity: components.PlannedActivityCreate{
                        ForeignBondTradingDetails: components.ForeignBondTradingDetailsCreate{
                            ForeignBondTrading: true,
                            ForeignBondTradingDetail: []components.ForeignBondTradingDetailCreate{
                                components.ForeignBondTradingDetailCreate{
                                    Percentage: components.DecimalCreate{},
                                    RegionCode: "CA",
                                },
                            },
                        },
                        LowPricedSecurities: components.LowPricedSecuritiesCreate{
                            LowPricedSecurities: true,
                        },
                        PrimaryAccountActivityType: components.PrimaryAccountActivityTypeActiveTrading,
                        WithdrawalFrequency: components.WithdrawalFrequencyFrequent,
                    },
                    RelatedPepDetails: components.RelatedPepDetailsCreate{
                        DirectOrIndirectRelatedPeps: true,
                        RelatedPeps: []components.RelatedPepCreate{
                            components.RelatedPepCreate{
                                ImmediateFamilyMembers: []components.ImmediateFamilyMemberCreate{
                                    components.ImmediateFamilyMemberCreate{
                                        FamilyMemberName: "Ellen Chen",
                                        Relationship: "Daughter",
                                    },
                                },
                                Organization: "U.S. Embassy",
                                PepName: "Juan Octavio",
                                Title: "U.S. Ambassador to Japan",
                            },
                        },
                    },
                },
            },
            IndividualEnrollmentMetadata: &components.IndividualEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IndividualEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IndividualEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            IraBeneficiaryEnrollmentMetadata: &components.IRABeneficiaryEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRABeneficiaryEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRABeneficiaryEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
                InheritedFromOwnerBirthDate: &components.DateCreate{},
                InheritedFromOwnerDeathDate: &components.DateCreate{},
                InheritedFromOwnerName: "<value>",
                InheritorIsDecedentsSpouse: false,
            },
            IraRolloverEnrollmentMetadata: &components.IRARolloverEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRARolloverEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            IraRothEnrollmentMetadata: &components.IRARothEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRARothEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRARothEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            IraSepEnrollmentMetadata: &components.IRASEPEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRASEPEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRASEPEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            IraSimpleEnrollmentMetadata: &components.IRASimpleEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRASimpleEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            IraTraditionalEnrollmentMetadata: &components.IRATraditionalEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.IRATraditionalEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.IRATraditionalEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            JointCommunityPropertyEnrollmentMetadata: &components.JointCommunityPropertyEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.JointCommunityPropertyEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.JointCommunityPropertyEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
                LegalResidencyStateOfMarriedCouple: components.LegalResidencyStateOfMarriedCoupleTx,
            },
            JointTenantsByEntiretyEnrollmentMetadata: &components.JointTenantsByEntiretyEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.JointTenantsByEntiretyEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.JointTenantsByEntiretyEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
                LegalResidencyStateOfMarriedCouple: components.JointTenantsByEntiretyEnrollmentMetadataCreateLegalResidencyStateOfMarriedCoupleTx,
            },
            JointTenantsInCommonEnrollmentMetadata: &components.JointTenantsInCommonEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.JointTenantsInCommonEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.JointTenantsInCommonEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            JointWithRightsOfSurvivorshipEnrollmentMetadata: &components.JointWithRightsOfSurvivorshipEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            LlcEnrollmentMetadata: &components.LLCEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.LLCEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                EddAccountEnrollmentMetadata: &components.EddAccountEnrollmentMetadataCreate{
                    DepositedFunds: components.DepositedFundsCreate{
                        InitialDepositAmount: components.DecimalCreate{},
                        InitialDepositSource: "Product Revenue",
                    },
                    DeterminedAccountRiskRating: components.DeterminedAccountRiskRatingHigh,
                    FinancialProfile: components.FinancialProfileCreate{
                        BankingRelationships: []string{
                            "<value>",
                            "<value>",
                            "<value>",
                        },
                        OtherAccounts: components.OtherAccountsCreate{
                            OwnerHasOtherAccountsAtApex: true,
                        },
                        PrimarySourceOfDepositedFunds: ascendsdk.String("Corporate Income"),
                    },
                    PlannedActivity: components.PlannedActivityCreate{
                        ForeignBondTradingDetails: components.ForeignBondTradingDetailsCreate{
                            ForeignBondTrading: true,
                            ForeignBondTradingDetail: []components.ForeignBondTradingDetailCreate{
                                components.ForeignBondTradingDetailCreate{
                                    Percentage: components.DecimalCreate{},
                                    RegionCode: "CA",
                                },
                            },
                        },
                        LowPricedSecurities: components.LowPricedSecuritiesCreate{
                            LowPricedSecurities: true,
                        },
                        PrimaryAccountActivityType: components.PrimaryAccountActivityTypeActiveTrading,
                        WithdrawalFrequency: components.WithdrawalFrequencyFrequent,
                    },
                    RelatedPepDetails: components.RelatedPepDetailsCreate{
                        DirectOrIndirectRelatedPeps: true,
                        RelatedPeps: []components.RelatedPepCreate{
                            components.RelatedPepCreate{
                                ImmediateFamilyMembers: []components.ImmediateFamilyMemberCreate{
                                    components.ImmediateFamilyMemberCreate{
                                        FamilyMemberName: "Ellen Chen",
                                        Relationship: "Daughter",
                                    },
                                },
                                Organization: "U.S. Embassy",
                                PepName: "Juan Octavio",
                                Title: "U.S. Ambassador to Japan",
                            },
                        },
                    },
                    ScopeOfBusiness: "Financial Services",
                },
                FdicCashSweep: components.LLCEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
            },
            OperatingEnrollmentMetadata: &components.OperatingEnrollmentMetadataCreate{
                OperatingPurpose: components.OperatingPurposeCustody,
                Subtitle: ascendsdk.String("C/F Optionality Securities"),
                TaxWithholdingMetadata: &components.TaxWithholdingMetadataCreate{
                    TaxAuthority: components.TaxAuthorityTx,
                    WithholdingType: components.WithholdingTypeBackup,
                },
            },
            PrincipalApproverID: "02HB7N66WW02WL3B6B9W29K0HW",
            TrustEnrollmentMetadata: &components.TrustEnrollmentMetadataCreate{
                DividendReinvestmentPlan: components.TrustEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll.ToPointer(),
                FdicCashSweep: components.TrustEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll.ToPointer(),
                OpenedOnBehalfOf: components.OpenedOnBehalfOfPersonalTrust,
            },
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAvailableEnrollments

Get a list of Enrollments available for an Account.

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
    res, err := s.EnrollmentsAndAgreements.ListAvailableEnrollments(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdk.Int(25), ascendsdk.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), ascendsdk.String("enrollment_type == \"REGISTRATION_INDIVIDUAL\""))
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## DeactivateEnrollment

Deactivates an Account Enrollment.

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
    res, err := s.EnrollmentsAndAgreements.DeactivateEnrollment(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", components.DeactivateEnrollmentRequestCreate{
        EnrollmentID: ascendsdk.String("22951598-70e2-46f1-bb32-38e8da7a5cdb"),
        EnrollmentType: components.DeactivateEnrollmentRequestCreateEnrollmentTypeCashFdicCashSweep.ToPointer(),
    })
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListEnrollments

Gets a list of Enrollments for an Account.

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
    res, err := s.EnrollmentsAndAgreements.ListEnrollments(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdk.Int(5), ascendsdk.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AffirmAgreements

Affirm Agreements for an Account.

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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAgreements

Gets a list of Agreements on an Account.

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
    res, err := s.EnrollmentsAndAgreements.ListAgreements(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdk.Int(5), ascendsdk.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListEntitlements

Gets a list of Entitlements for an Account.

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
    res, err := s.EnrollmentsAndAgreements.ListEntitlements(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", ascendsdk.Int(5), ascendsdk.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |