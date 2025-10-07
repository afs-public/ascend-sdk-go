package helpers

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
)

var EnrollmentPrincipalApproverID = "01HMESE8WMDNTTWJ2BAEG3TZWA"

const WITHDRAWAL_ACCOUNT_ID = "01JHGTEPC6ZTAHCFRH2MD3VJJT"

func SetupAscendSDK() (*ascendsdk.SDK, error) {
	privateKey := os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY")
	apiKey := os.Getenv("API_KEY")
	name := os.Getenv("SERVICE_ACCOUNT_CREDS_NAME")
	organization := os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION")

	url := "https://uat.apexapis.com"
	accountType := "serviceAccount"

	if privateKey == "" || apiKey == "" || name == "" || organization == "" {
		return nil, fmt.Errorf("environment variables are not set properly")
	}

	return ascendsdk.New(
		ascendsdk.WithServerURL(url),
		ascendsdk.WithSecurity(components.Security{
			APIKey: ascendsdk.String(apiKey),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   privateKey,
				Name:         name,
				Organization: organization,
				Type:         accountType,
			},
		})), nil
}

func CreateLegalNaturalPersonId(s *ascendsdk.SDK, ctx context.Context) (*string, error) {
	createLegalNaturalPersonRequest := components.LegalNaturalPersonCreate{
		BirthDate: components.DateCreate{
			Year:  ascendsdk.Int(1990),
			Month: ascendsdk.Int(1),
			Day:   ascendsdk.Int(1),
		},
		CitizenshipCountries: []string{"US"},
		CorrespondentID:      os.Getenv("CORRESPONDENT_ID"),
		FamilyName:           "Jacob",
		GivenName:            "Bob",
		PersonalAddress: components.PostalAddressCreate{
			Locality:           ascendsdk.String("Portland"),
			RegionCode:         ascendsdk.String("US"),
			PostalCode:         ascendsdk.String("97209"),
			AdministrativeArea: ascendsdk.String("OR"),
			AddressLines:       []string{"19409 Sherilyn Courts"},
		},
		PoliticallyExposedImmediateFamilyNames: []string{},
		TaxID:                                  ascendsdk.String("874-45-6789"),
		TaxIDType:                              components.TaxIDType.ToPointer(components.TaxIDTypeTaxIDTypeSsn),
		TaxProfile: components.TaxProfileCreate{
			FederalTaxClassification: components.FederalTaxClassificationIndivSolepropOrSinglememberllc,
			UsTinStatus:              components.UsTinStatusPassing,
			IrsFormType:              components.IrsFormTypeW9,
			LegalTaxRegionCode:       "US",
		},
		Employment: components.EmploymentCreate{
			Occupation:       ascendsdk.String("Software Engineer"),
			EmploymentStatus: components.EmploymentStatusEmployed,
			EmployerAddress: &components.PostalAddressCreate{
				AdministrativeArea: ascendsdk.String("OR"),
				RegionCode:         ascendsdk.String("US"),
				PostalCode:         ascendsdk.String("97209"),
				Locality:           ascendsdk.String("Portland"),
				AddressLines:       []string{"1234 NW 12th Ave"},
			},
		},
		IdentityVerificationResult: &components.IdentityVerificationResultCreate{
			AddressVerified:   true,
			BirthDateVerified: true,
			ExecutionDate: components.DateCreate{
				Year:  ascendsdk.Int(2021),
				Month: ascendsdk.Int(1),
				Day:   ascendsdk.Int(1),
			},
			NameVerified:            true,
			TaxIDVerified:           true,
			ExternalCaseID:          "123456",
			Vendor:                  "Super Security Service",
			RawVendorDataDocumentID: ascendsdk.String("04eb923b-793d-481d-98c4-bb16f17378ea"),
		},
	}

	res, err := s.PersonManagement.CreateLegalNaturalPerson(ctx, createLegalNaturalPersonRequest)

	if err != nil {
		return nil, fmt.Errorf("failed to create legal natural person: %w", err)
	}

	return res.LegalNaturalPerson.LegalNaturalPersonID, nil
}

func CreateAccountIdWithLNP(s *ascendsdk.SDK, ctx context.Context, lnpID *string) (*string, error) {
	createAccountRequest := components.AccountRequestCreate{
		AccountGroupID:  os.Getenv("ACCOUNT_GROUP_ID"),
		CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
		Parties: []components.PartyRequestCreate{
			{
				LegalNaturalPersonID: lnpID,
				RelationType:         components.RelationTypePrimaryOwner,
				EmailAddress:         "mail@example.com",
				PhoneNumber: components.PhoneNumberCreate{
					E164Number: ascendsdk.String("+15035550123"),
					Extension:  ascendsdk.String("123"),
				},
				MailingAddress: components.PostalAddressCreate{
					AddressLines:       []string{"1234 NW 12th Ave"},
					RegionCode:         ascendsdk.String("US"),
					PostalCode:         ascendsdk.String("97209"),
					AdministrativeArea: ascendsdk.String("OR"),
					Locality:           ascendsdk.String("Portland"),
				},
			},
		}}

	res, err := s.AccountCreation.CreateAccount(ctx, createAccountRequest)

	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return res.Account.AccountID, nil
}

func CreateAccountId(s *ascendsdk.SDK, ctx context.Context) (*string, error) {
	createLegalPersonId, err := CreateLegalNaturalPersonId(s, ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to create legal natural person: %w", err)
	}

	createAccountRequest := components.AccountRequestCreate{
		AccountGroupID:  os.Getenv("ACCOUNT_GROUP_ID"),
		CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
		Parties: []components.PartyRequestCreate{
			{
				LegalNaturalPersonID: createLegalPersonId,
				RelationType:         components.RelationTypePrimaryOwner,
				EmailAddress:         "mail@example.com",
				PhoneNumber: components.PhoneNumberCreate{
					E164Number: ascendsdk.String("+15035550123"),
					Extension:  ascendsdk.String("123"),
				},
				MailingAddress: components.PostalAddressCreate{
					AddressLines:       []string{"1234 NW 12th Ave"},
					RegionCode:         ascendsdk.String("US"),
					PostalCode:         ascendsdk.String("97209"),
					AdministrativeArea: ascendsdk.String("OR"),
					Locality:           ascendsdk.String("Portland"),
				},
			},
		}}

	res, err := s.AccountCreation.CreateAccount(ctx, createAccountRequest)

	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return res.Account.AccountID, nil
}

func EnrollAccountIds(s *ascendsdk.SDK, ctx context.Context, accountId string) ([]components.Agreement, error) {
	enrollmentType := components.EnrollmentCreateTypeRegistrationIndividual
	fdicCashSweep := components.IndividualEnrollmentMetadataCreateFdicCashSweepFdicCashSweepDecline

	enrollAccountRequest := components.EnrollAccountRequestCreate{
		Enrollment: components.EnrollmentCreate{
			PrincipalApproverID: EnrollmentPrincipalApproverID,
			Type:                enrollmentType,
			IndividualEnrollmentMetadata: &components.IndividualEnrollmentMetadataCreate{
				FdicCashSweep: &fdicCashSweep,
			},
		},
	}

	res, err := s.EnrollmentsAndAgreements.EnrollAccount(ctx, accountId, enrollAccountRequest)

	if err != nil {
		return nil, fmt.Errorf("failed to enroll account: %w", err)
	}

	if len(res.EnrollAccountResponse.Agreements) < 2 {
		return nil, fmt.Errorf("insufficient agreements returned")
	}

	return res.EnrollAccountResponse.Agreements, nil
}

func AffirmAgreements(s *ascendsdk.SDK, ctx context.Context, accountId string, agreements []components.Agreement) (err error) {

	var agreementIds []string

	for _, agreement := range agreements {
		agreementIds = append(agreementIds, *agreement.AgreementID)
	}

	affirmAgreementsRequest := components.AffirmAgreementsRequestCreate{
		AccountAgreementIds: agreementIds,
	}

	_, err = s.EnrollmentsAndAgreements.AffirmAgreements(ctx, accountId, affirmAgreementsRequest)

	if err != nil {
		return fmt.Errorf("failed to affirm agreements: %w", err)
	}

	return nil
}

func CreateBankRelationship(s *ascendsdk.SDK, ctx context.Context, accountId string) (*string, error) {
	bankAccountNumber := generateBankAccountNumber()
	routingNumber := "112203216"

	createBankRelationshipRequest := components.BankRelationshipCreate{
		BankAccount: &components.BankAccountCreate{
			AccountNumber: strconv.Itoa(bankAccountNumber),
			Owner:         "Test User",
			RoutingNumber: routingNumber,
			Type:          components.BankAccountCreateTypeSavings,
		},
		Nickname:           "ACH TEST",
		VerificationMethod: components.VerificationMethodMicroDeposit,
	}

	res, err := s.BankRelationships.CreateBankRelationship(ctx, accountId, createBankRelationshipRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank relationship: %w", err)
	}
	bankRelationshipId := strings.Split(*res.BankRelationship.Name, "/")[3]
	return &bankRelationshipId, nil
}

func GetCorrectMicroDeposits(s *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) ([]string, error) {
	res, err := s.TestSimulation.GetMicroDepositAmounts(ctx, accountId, bankRelationshipId)
	if err != nil {
		return nil, fmt.Errorf("failed to get micro deposits: %w", err)
	}
	return []string{*res.MicroDepositAmounts.GetAmount1().Value, *res.MicroDepositAmounts.GetAmount2().Value}, nil
}

func VerifyMicroDeposits(s *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string, amounts []string) error {
	verifyMicroDepositsRequest := components.VerifyMicroDepositsRequestCreate{
		Amounts: components.MicroDepositAmountsCreate{
			Amount1: components.DecimalCreate{
				Value: ascendsdk.String(amounts[0]),
			},
			Amount2: components.DecimalCreate{
				Value: ascendsdk.String(amounts[1]),
			},
		},
		Name: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
	}

	_, err := s.BankRelationships.VerifyMicroDeposits(ctx, accountId, bankRelationshipId, verifyMicroDepositsRequest)
	if err != nil {
		return fmt.Errorf("failed to verify micro deposits: %w", err)
	}
	return nil
}

func generateBankAccountNumber() int {
	min := 10000000
	max := 99999999
	return rand.Intn(max-min+1) + min
}

func CreateEnrolledAccount(sdk *ascendsdk.SDK, ctx context.Context, t *testing.T) (*string, error) {
	accountId, err := CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	agg, err := EnrollAccountIds(sdk, ctx, *accountId)
	require.NoError(t, err)

	err = AffirmAgreements(sdk, ctx, *accountId, agg)
	require.NoError(t, err)

	return accountId, nil
}
