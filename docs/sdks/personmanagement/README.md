# PersonManagement
(*PersonManagement*)

## Overview

### Available Operations

* [CreateLegalNaturalPerson](#createlegalnaturalperson) - Create Legal Natural Person
* [ListLegalNaturalPersons](#listlegalnaturalpersons) - List Legal Natural Persons
* [GetLegalNaturalPerson](#getlegalnaturalperson) - Get Legal Natural Persons
* [UpdateLegalNaturalPerson](#updatelegalnaturalperson) - Update Legal Natural Person
* [AssignLargeTrader](#assignlargetrader) - Assign Large Trader
* [EndLargeTraderLegalNaturalPerson](#endlargetraderlegalnaturalperson) - End Large Trader
* [CreateLegalEntity](#createlegalentity) - Create Legal Entity
* [ListLegalEntities](#listlegalentities) - List Legal Entity
* [GetLegalEntity](#getlegalentity) - Get Legal Entity
* [UpdateLegalEntity](#updatelegalentity) - Update Legal Entity
* [AssignLargeTraderLegalEntity](#assignlargetraderlegalentity) - Assign Entity Large Trader
* [EndLargeTrader](#endlargetrader) - End Entity Large Trader

## CreateLegalNaturalPerson

Creates a Legal Natural Person.

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
    res, err := s.PersonManagement.CreateLegalNaturalPerson(ctx, components.LegalNaturalPersonCreate{
        BirthDate: components.DateCreate{},
        CitizenshipCountries: []string{
            "US",
            "CA",
        },
        ControlPersonCompanySymbols: ascendsdkgo.String("AAPL, GOOL"),
        CorrespondentEmployee: ascendsdkgo.Bool(false),
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
        DeathDate: &components.DateCreate{},
        Employment: components.EmploymentCreate{
            Employer: ascendsdkgo.String("Apex Fintech Solutions"),
            EmploymentStatus: components.EmploymentStatusEmployed,
            Occupation: "Software Engineer",
            StartYear: ascendsdkgo.Int(2019),
        },
        FamilyName: "Doe",
        FinraAssociatedEntity: ascendsdkgo.String("Entity Name"),
        ForeignIdentification: &components.ForeignIdentificationCreate{
            ExpirationDate: &components.DateCreate{},
            Ftin: true,
            IdentificationNumber: "M1C1W7GQSK",
            IssueDate: &components.DateCreate{},
            IssuingRegionCode: "CA",
            Type: components.TypePassport.ToPointer(),
        },
        GivenName: "John",
        IdentityVerificationResult: &components.IdentityVerificationResultCreate{
            AddressVerified: true,
            BirthDateVerified: true,
            ExecutionDate: components.DateCreate{},
            ExternalCaseID: "6526280",
            IdentityVerificationDocumentIds: []string{
                "d257c455-f355-493d-9c8f-06f8ace5d5fd",
                "6ace3020-24a2-45c4-9f3b-752101073127",
            },
            NameVerified: true,
            RawVendorDataDocumentID: ascendsdkgo.String("7a63073a-e694-4a38-b6e0-552044b503f2"),
            TaxIDVerified: true,
            Vendor: "Super Security Service",
        },
        LargeTrader: &components.LargeTraderCreate{
            EffectiveDate: components.DateCreate{},
            LargeTraderID: "123412341234",
        },
        MaritalStatus: components.MaritalStatusSingle.ToPointer(),
        MiddleNames: ascendsdkgo.String("Smith"),
        NameSuffix: components.NameSuffixJr.ToPointer(),
        NaturalPersonFdd: &components.NaturalPersonFddCreate{
            CustomerNonReferralSource: ascendsdkgo.String("Introduced through mobile app."),
            CustomerReferralSource: &components.CustomerReferralSourceCreate{
                Name: "John Doe",
                RelationshipToApplicant: "Friend",
                RelationshipYearsWithApplicant: 5,
                RelationshipYearsWithBroker: 2,
            },
            EmploymentAndEmployerDescription: "I am a line cook at a fine dining restaurant with 55 employees.",
            NegativeNews: components.NegativeNewsCreate{
                NegativeNewsAgainstRelatedParties: true,
                NegativeNewsAgainstRelatedPartiesDescription: ascendsdkgo.String("Juan was indicated in numerous publications but not involved with Japan's misappropriation of taxpayer funds in 2013."),
            },
            OtherSourcesOfWealth: components.OtherSourcesOfWealthCreate{
                ApplicantHasOtherSourcesOfWealth: true,
                OtherSourcesOfWealth: ascendsdkgo.String("I also have a small business selling handmade jewelry."),
                OtherSourcesOfWealthVerification: ascendsdkgo.String("I have a business license and tax returns to verify my business."),
            },
        },
        NonCitizenResidency: &components.NonCitizenResidencyCreate{
            ResidencyStatus: components.ResidencyStatusUsPermanentResident,
        },
        PersonalAddress: components.PostalAddressCreate{},
        PoliticallyExposedImmediateFamilyNames: []string{
            "Sue Doe",
        },
        PoliticallyExposedOrganization: ascendsdkgo.String("PEAK6, Apex Clearing"),
        SubjectToBackupWithholding: ascendsdkgo.Bool(false),
        TaxID: ascendsdkgo.String("987-65-4321"),
        TaxIDType: components.TaxIDTypeTaxIDTypeSsn.ToPointer(),
        TaxProfile: components.TaxProfileCreate{
            FederalTaxClassification: components.FederalTaxClassificationCCorporation,
            IrsFormType: components.IrsFormTypeIrsFormTypeUnspecified,
            LegalTaxRegionCode: "US",
            UsTinStatus: components.UsTinStatusPassing,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalNaturalPerson != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.LegalNaturalPersonCreate](../../models/components/legalnaturalpersoncreate.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.AccountsCreateLegalNaturalPersonResponse](../../models/operations/accountscreatelegalnaturalpersonresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalNaturalPersons

Gets a list of Legal Natural Person records based on search criteria.

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
    res, err := s.PersonManagement.ListLegalNaturalPersons(ctx, ascendsdkgo.Int(25), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), nil, ascendsdkgo.String("legal_natural_person_id == \"e6716139-da77-46d1-9f15-13599161db0b\""))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListLegalNaturalPersonsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                   | Example                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                                    | **int*                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                            | The maximum number of legal natural persons to return. The service may return fewer than this value. If unspecified, at most 25 legal natural persons will be returned. The maximum value is 100; values above 100 will be coerced to 100.                                                                                                                                                                                    | 25                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `pageToken`                                                                                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                            | A page token, received from a previous `ListLegalNaturalPersons` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListLegalNaturalPersons` must match the call that provided the page token.                                                                                                                                                                  | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                                                                                                                                                                                                      |
| `orderBy`                                                                                                                                                                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                            | The order in which legal natural persons are listed.                                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `filter`                                                                                                                                                                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                            | A CEL string to filter results; Use `upperAscii()` for case-insensitive searches; E.g. `given_name.upperAscii()=="rUsTy".upperAscii()`; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `legal_natural_person_id`<br/> `correspondent_id`<br/> `given_name`<br/> `family_name`<br/> `tax_id`<br/> `tax_id_type`<br/> `investigation_id` | legal_natural_person_id == "e6716139-da77-46d1-9f15-13599161db0b"                                                                                                                                                                                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                                                                                                                                                                               |

### Response

**[*operations.AccountsListLegalNaturalPersonsResponse](../../models/operations/accountslistlegalnaturalpersonsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalNaturalPerson

Get Legal Natural Person

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
    res, err := s.PersonManagement.GetLegalNaturalPerson(ctx, "e6716139-da77-46d1-9f15-13599161db0b")
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalNaturalPerson != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `legalNaturalPersonID`                                   | *string*                                                 | :heavy_check_mark:                                       | The legalNaturalPerson id.                               | e6716139-da77-46d1-9f15-13599161db0b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountsGetLegalNaturalPersonResponse](../../models/operations/accountsgetlegalnaturalpersonresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateLegalNaturalPerson

Updates a Legal Natural Person.

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
    res, err := s.PersonManagement.UpdateLegalNaturalPerson(ctx, "e6716139-da77-46d1-9f15-13599161db0b", components.LegalNaturalPersonUpdate{
        CitizenshipCountries: []string{
            "US",
            "CA",
        },
        ControlPersonCompanySymbols: ascendsdkgo.String("AAPL, GOOL"),
        CorrespondentEmployee: ascendsdkgo.Bool(false),
        CorrespondentID: ascendsdkgo.String("01HPMZZM6RKMVZA1JQ63RQKJRP"),
        Employment: &components.EmploymentUpdate{
            Employer: ascendsdkgo.String("Apex Fintech Solutions"),
            EmploymentStatus: components.EmploymentUpdateEmploymentStatusEmployed.ToPointer(),
            Occupation: ascendsdkgo.String("Software Engineer"),
            StartYear: ascendsdkgo.Int(2019),
        },
        FamilyName: ascendsdkgo.String("Doe"),
        FinraAssociatedEntity: ascendsdkgo.String("Entity Name"),
        ForeignIdentification: &components.ForeignIdentificationUpdate{
            Ftin: ascendsdkgo.Bool(true),
            IdentificationNumber: ascendsdkgo.String("M1C1W7GQSK"),
            IssuingRegionCode: ascendsdkgo.String("CA"),
            Type: components.ForeignIdentificationUpdateTypePassport.ToPointer(),
        },
        GivenName: ascendsdkgo.String("John"),
        IdentityVerificationResult: &components.IdentityVerificationResultUpdate{
            AddressVerified: ascendsdkgo.Bool(true),
            BirthDateVerified: ascendsdkgo.Bool(true),
            ExternalCaseID: ascendsdkgo.String("6526280"),
            IdentityVerificationDocumentIds: []string{
                "d257c455-f355-493d-9c8f-06f8ace5d5fd",
                "6ace3020-24a2-45c4-9f3b-752101073127",
            },
            NameVerified: ascendsdkgo.Bool(true),
            RawVendorDataDocumentID: ascendsdkgo.String("7a63073a-e694-4a38-b6e0-552044b503f2"),
            TaxIDVerified: ascendsdkgo.Bool(true),
            Vendor: ascendsdkgo.String("Super Security Service"),
        },
        LargeTrader: &components.LargeTraderUpdate{
            LargeTraderID: ascendsdkgo.String("123412341234"),
        },
        MaritalStatus: components.LegalNaturalPersonUpdateMaritalStatusSingle.ToPointer(),
        MiddleNames: ascendsdkgo.String("Smith"),
        NameSuffix: components.LegalNaturalPersonUpdateNameSuffixJr.ToPointer(),
        NaturalPersonFdd: &components.NaturalPersonFddUpdate{
            CustomerNonReferralSource: ascendsdkgo.String("Introduced through mobile app."),
            CustomerReferralSource: &components.CustomerReferralSourceUpdate{
                Name: ascendsdkgo.String("John Doe"),
                RelationshipToApplicant: ascendsdkgo.String("Friend"),
                RelationshipYearsWithApplicant: ascendsdkgo.Int(5),
                RelationshipYearsWithBroker: ascendsdkgo.Int(2),
            },
            EmploymentAndEmployerDescription: ascendsdkgo.String("I am a line cook at a fine dining restaurant with 55 employees."),
            NegativeNews: &components.NegativeNewsUpdate{
                NegativeNewsAgainstRelatedParties: ascendsdkgo.Bool(true),
                NegativeNewsAgainstRelatedPartiesDescription: ascendsdkgo.String("Juan was indicated in numerous publications but not involved with Japan's misappropriation of taxpayer funds in 2013."),
            },
            OtherSourcesOfWealth: &components.OtherSourcesOfWealthUpdate{
                ApplicantHasOtherSourcesOfWealth: ascendsdkgo.Bool(true),
                OtherSourcesOfWealth: ascendsdkgo.String("I also have a small business selling handmade jewelry."),
                OtherSourcesOfWealthVerification: ascendsdkgo.String("I have a business license and tax returns to verify my business."),
            },
        },
        NonCitizenResidency: &components.NonCitizenResidencyUpdate{
            ResidencyStatus: components.NonCitizenResidencyUpdateResidencyStatusUsPermanentResident.ToPointer(),
        },
        PoliticallyExposedImmediateFamilyNames: []string{
            "Sue Doe",
        },
        PoliticallyExposedOrganization: ascendsdkgo.String("PEAK6, Apex Clearing"),
        SubjectToBackupWithholding: ascendsdkgo.Bool(false),
        TaxID: ascendsdkgo.String("987-65-4321"),
        TaxIDType: components.LegalNaturalPersonUpdateTaxIDTypeTaxIDTypeEin.ToPointer(),
        TaxProfile: &components.TaxProfileUpdate{
            FederalTaxClassification: components.TaxProfileUpdateFederalTaxClassificationCCorporation.ToPointer(),
            IrsFormType: components.TaxProfileUpdateIrsFormTypeIrsFormTypeUnspecified.ToPointer(),
            LegalTaxRegionCode: ascendsdkgo.String("US"),
            UsTinStatus: components.TaxProfileUpdateUsTinStatusPassing.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalNaturalPerson != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `legalNaturalPersonID`                                                                     | *string*                                                                                   | :heavy_check_mark:                                                                         | The legalNaturalPerson id.                                                                 | e6716139-da77-46d1-9f15-13599161db0b                                                       |
| `legalNaturalPersonUpdate`                                                                 | [components.LegalNaturalPersonUpdate](../../models/components/legalnaturalpersonupdate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.AccountsUpdateLegalNaturalPersonResponse](../../models/operations/accountsupdatelegalnaturalpersonresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AssignLargeTrader

Assigns a person's Large Trader ID.

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
    res, err := s.PersonManagement.AssignLargeTrader(ctx, "e6716139-da77-46d1-9f15-13599161db0b", components.AssignLargeTraderRequestCreate{
        LargeTraderID: "1234567890",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LargeTrader != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `legalNaturalPersonID`                                                                                 | *string*                                                                                               | :heavy_check_mark:                                                                                     | The legalNaturalPerson id.                                                                             | e6716139-da77-46d1-9f15-13599161db0b                                                                   |
| `assignLargeTraderRequestCreate`                                                                       | [components.AssignLargeTraderRequestCreate](../../models/components/assignlargetraderrequestcreate.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.AccountsAssignLargeTraderResponse](../../models/operations/accountsassignlargetraderresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## EndLargeTraderLegalNaturalPerson

Removes a person's Large Trader ID.

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
    res, err := s.PersonManagement.EndLargeTraderLegalNaturalPerson(ctx, "e6716139-da77-46d1-9f15-13599161db0b", components.EndLargeTraderRequestCreate{
        EndReason: components.EndReasonEventReasonOther,
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
| `legalNaturalPersonID`                                                                           | *string*                                                                                         | :heavy_check_mark:                                                                               | The legalNaturalPerson id.                                                                       | e6716139-da77-46d1-9f15-13599161db0b                                                             |
| `endLargeTraderRequestCreate`                                                                    | [components.EndLargeTraderRequestCreate](../../models/components/endlargetraderrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AccountsEndLargeTraderLegalNaturalPersonResponse](../../models/operations/accountsendlargetraderlegalnaturalpersonresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CreateLegalEntity

Creates a Legal Entity.

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
    res, err := s.PersonManagement.CreateLegalEntity(ctx, components.LegalEntityCreate{
        AccreditedInvestor: ascendsdkgo.Bool(false),
        Adviser: ascendsdkgo.Bool(false),
        BrokerDealer: ascendsdkgo.Bool(false),
        BusinessIndustrialClassification: components.BusinessIndustrialClassificationFinanceInsuranceAndRealEstate.ToPointer(),
        CorporateStructure: components.CorporateStructureCorporationCCorp.ToPointer(),
        CorrespondentID: "01HPMZZM6RKMVZA1JQ63RQKJRP",
        EntityDueDiligence: &components.EntityDueDiligenceCreate{
            EntityIssuesBearerShares: false,
            NegativeNews: components.NegativeNewsCreate{
                NegativeNewsAgainstRelatedParties: true,
                NegativeNewsAgainstRelatedPartiesDescription: ascendsdkgo.String("Juan was indicated in numerous publications but not involved with Japan's misappropriation of taxpayer funds in 2013."),
            },
        },
        EntityName: "Acme, Inc.",
        EntityType: components.EntityTypeCorporation,
        ExemptCustomerReason: components.ExemptCustomerReasonNonBankListedEntity.ToPointer(),
        ExemptVerifyingBeneficialOwners: ascendsdkgo.Bool(false),
        ForTheBenefitOf: ascendsdkgo.String("John Dough"),
        ForeignFinancialInstitution: ascendsdkgo.Bool(false),
        FormationDate: &components.DateCreate{},
        LargeTrader: &components.LargeTraderCreate{
            EffectiveDate: components.DateCreate{},
            LargeTraderID: "123412341234",
        },
        LegalAddress: components.PostalAddressCreate{},
        LeiCode: ascendsdkgo.String("12340012345678911000"),
        OperatingRegions: []string{
            "US",
            "CA",
        },
        RegistrationRegion: "US",
        RegulatedInvestmentCompany: ascendsdkgo.Bool(false),
        RelatedDocumentIds: []string{
            "fb3f181c-f2fb-4bc2-b75a-79302c634ae5",
        },
        RevocableTrust: ascendsdkgo.Bool(false),
        SubjectToBackupWithholding: ascendsdkgo.Bool(false),
        TaxID: "987-65-4321",
        TaxIDType: components.LegalEntityCreateTaxIDTypeTaxIDTypeItin,
        TaxProfile: components.TaxProfileCreate{
            FederalTaxClassification: components.FederalTaxClassificationCCorporation,
            IrsFormType: components.IrsFormTypeIrsFormTypeUnspecified,
            LegalTaxRegionCode: "US",
            UsTinStatus: components.UsTinStatusPassing,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.LegalEntityCreate](../../models/components/legalentitycreate.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AccountsCreateLegalEntityResponse](../../models/operations/accountscreatelegalentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalEntities

Gets a list of Legal Entity records based on search criteria.

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
    res, err := s.PersonManagement.ListLegalEntities(ctx, ascendsdkgo.Int(25), ascendsdkgo.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListLegalEntitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                  |
| `pageSize`                                                                                                                                                                                                                                                                       | **int*                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The maximum number of legal entities to return. The service may return fewer than this value. If unspecified, at most 25 legal entities will be returned. The maximum value is 100; values above 100 will be coerced to 100.                                                     | 25                                                                                                                                                                                                                                                                               |
| `pageToken`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A page token, received from a previous `ListLegalEntities` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListLegalEntities` must match the call that provided the page token.                                 | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                                                         |
| `orderBy`                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The order in which legal entities are listed.                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                  |
| `filter`                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A CEL string to filter results; Use `upperAscii()` for case-insensitive searches; E.g. `entity_name.upperAscii()=="AcMe,InC".upperAscii()`; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; |                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The options for this request.                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                  |

### Response

**[*operations.AccountsListLegalEntitiesResponse](../../models/operations/accountslistlegalentitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 500, 503 | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalEntity

Get Legal Entity

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
    res, err := s.PersonManagement.GetLegalEntity(ctx, "e6716139-da77-46d1-9f15-13599161db0b")
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `legalEntityID`                                          | *string*                                                 | :heavy_check_mark:                                       | The legalEntity id.                                      | e6716139-da77-46d1-9f15-13599161db0b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AccountsGetLegalEntityResponse](../../models/operations/accountsgetlegalentityresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## UpdateLegalEntity

Updates a Legal Entity.

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
    res, err := s.PersonManagement.UpdateLegalEntity(ctx, "42567868-9373-4872-9d24-2e33f6c19b75", components.LegalEntityUpdate{
        AccreditedInvestor: ascendsdkgo.Bool(false),
        Adviser: ascendsdkgo.Bool(false),
        BrokerDealer: ascendsdkgo.Bool(false),
        BusinessIndustrialClassification: components.LegalEntityUpdateBusinessIndustrialClassificationFinanceInsuranceAndRealEstate.ToPointer(),
        CorporateStructure: components.LegalEntityUpdateCorporateStructureCorporationCCorp.ToPointer(),
        CorrespondentID: ascendsdkgo.String("01HPMZZM6RKMVZA1JQ63RQKJRP"),
        EntityDueDiligence: &components.EntityDueDiligenceUpdate{
            EntityIssuesBearerShares: ascendsdkgo.Bool(false),
            NegativeNews: &components.NegativeNewsUpdate{
                NegativeNewsAgainstRelatedParties: ascendsdkgo.Bool(true),
                NegativeNewsAgainstRelatedPartiesDescription: ascendsdkgo.String("Juan was indicated in numerous publications but not involved with Japan's misappropriation of taxpayer funds in 2013."),
            },
        },
        EntityName: ascendsdkgo.String("Acme, Inc."),
        EntityType: components.LegalEntityUpdateEntityTypeCorporation.ToPointer(),
        ExemptCustomerReason: components.LegalEntityUpdateExemptCustomerReasonNonBankListedEntity.ToPointer(),
        ExemptVerifyingBeneficialOwners: ascendsdkgo.Bool(false),
        ForTheBenefitOf: ascendsdkgo.String("John Dough"),
        ForeignFinancialInstitution: ascendsdkgo.Bool(false),
        LargeTrader: &components.LargeTraderUpdate{
            LargeTraderID: ascendsdkgo.String("123412341234"),
        },
        LeiCode: ascendsdkgo.String("12340012345678911000"),
        OperatingRegions: []string{
            "US",
            "CA",
        },
        RegistrationRegion: ascendsdkgo.String("US"),
        RegulatedInvestmentCompany: ascendsdkgo.Bool(false),
        RelatedDocumentIds: []string{
            "fb3f181c-f2fb-4bc2-b75a-79302c634ae5",
        },
        RevocableTrust: ascendsdkgo.Bool(false),
        SubjectToBackupWithholding: ascendsdkgo.Bool(false),
        TaxID: ascendsdkgo.String("987-65-4321"),
        TaxIDType: components.LegalEntityUpdateTaxIDTypeTaxIDTypeEin.ToPointer(),
        TaxProfile: &components.TaxProfileUpdate{
            FederalTaxClassification: components.TaxProfileUpdateFederalTaxClassificationCCorporation.ToPointer(),
            IrsFormType: components.TaxProfileUpdateIrsFormTypeIrsFormTypeUnspecified.ToPointer(),
            LegalTaxRegionCode: ascendsdkgo.String("US"),
            UsTinStatus: components.TaxProfileUpdateUsTinStatusPassing.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `legalEntityID`                                                              | *string*                                                                     | :heavy_check_mark:                                                           | The legalEntity id.                                                          | 42567868-9373-4872-9d24-2e33f6c19b75                                         |
| `legalEntityUpdate`                                                          | [components.LegalEntityUpdate](../../models/components/legalentityupdate.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.AccountsUpdateLegalEntityResponse](../../models/operations/accountsupdatelegalentityresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AssignLargeTraderLegalEntity

Assigns a person's Large Trader ID.

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
    res, err := s.PersonManagement.AssignLargeTraderLegalEntity(ctx, "e6716139-da77-46d1-9f15-13599161db0b", components.AssignLargeTraderRequestCreate{
        LargeTraderID: "1234567890",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LargeTrader != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `legalEntityID`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | The legalEntity id.                                                                                    | e6716139-da77-46d1-9f15-13599161db0b                                                                   |
| `assignLargeTraderRequestCreate`                                                                       | [components.AssignLargeTraderRequestCreate](../../models/components/assignlargetraderrequestcreate.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.AccountsAssignLargeTraderLegalEntityResponse](../../models/operations/accountsassignlargetraderlegalentityresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## EndLargeTrader

Removes a person's Large Trader ID.

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
    res, err := s.PersonManagement.EndLargeTrader(ctx, "e6716139-da77-46d1-9f15-13599161db0b", components.EndLargeTraderRequestCreate{
        EndReason: components.EndReasonEventReasonCreated,
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
| `legalEntityID`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | The legalEntity id.                                                                              | e6716139-da77-46d1-9f15-13599161db0b                                                             |
| `endLargeTraderRequestCreate`                                                                    | [components.EndLargeTraderRequestCreate](../../models/components/endlargetraderrequestcreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AccountsEndLargeTrader1Response](../../models/operations/accountsendlargetrader1response.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |