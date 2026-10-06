package helpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/retry"
)

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

var EnrollmentPrincipalApproverID = "01HMESE8WMDNTTWJ2BAEG3TZWA"

const WITHDRAWAL_ACCOUNT_ID = "01JHGTEPC6ZTAHCFRH2MD3VJJT"

// Dedicated alternative-investments test account. Created 2026-08 (after the
// Monark update that routes SPV orders through Apex when the correspondent
// isn't registered for the SPV), funded, and accredited. Pre-update accounts
// like WITHDRAWAL_ACCOUNT_ID get their SPV orders rejected by Monark.
var ALTS_ACCOUNT_ID = envOr("ALTS_ACCOUNT_ID", "01M0DMB41SQR6SYZDQYZN3CJY2")

// An alternative order placed on ALTS_ACCOUNT_ID, used by get/settle tests.
var ALTS_ORDER_ID = envOr("ALTS_ORDER_ID", "01M0DMH6QGPHZGFD4T1CJMABEM")

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

	lnpID := res.LegalNaturalPerson.LegalNaturalPersonID

	// Upload CIP Results
	correspondentID := os.Getenv("CORRESPONDENT_ID")
	docType := components.IDDocumentUploadRequestCreateDocumentTypeThirdPartyCipResults
	mimeType := "application/json"
	batchSourceID := uuid.New().String()

	uploadLinkRequest := components.BatchCreateUploadLinksRequestCreate{
		CreateDocumentUploadLinkRequest: []components.CreateUploadLinkRequestCreate{
			{
				IDDocumentUploadRequest: &components.IDDocumentUploadRequestCreate{
					CorrespondentID:      correspondentID,
					DocumentType:         docType,
					LegalNaturalPersonID: *lnpID,
				},
				ClientBatchSourceID: batchSourceID,
				MimeType:            mimeType,
			},
		},
	}

	uploadLinksRes, err := s.InvestorDocs.BatchCreateUploadLinks(ctx, uploadLinkRequest)
	if err != nil {
		fmt.Printf("Failed to create upload links: %v\n", err)
		return lnpID, nil
	}

	// Check if upload links were created successfully
	if uploadLinksRes.BatchCreateUploadLinksResponse == nil || len(uploadLinksRes.BatchCreateUploadLinksResponse.UploadLink) == 0 {
		fmt.Println("No upload links returned")
		return lnpID, nil
	}

	uploadURL := uploadLinksRes.BatchCreateUploadLinksResponse.UploadLink[0].UploadLink

	// Upload the test.json file
	testFilePath := "../../../examples/test.json"

	// Check if file exists
	if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
		fmt.Printf("Test file not found at: %s\n", testFilePath)
		return lnpID, nil
	}

	// Read and validate file content
	jsonContent, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return lnpID, nil
	}

	if len(jsonContent) == 0 {
		fmt.Println("Test file is empty")
		return lnpID, nil
	}

	// Upload to signed URL
	req, err := http.NewRequest("PUT", *uploadURL, bytes.NewReader(jsonContent))
	if err != nil {
		fmt.Printf("Error creating upload request: %v\n", err)
		return lnpID, nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	uploadResp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error during upload: %v\n", err)
		return lnpID, nil
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != 200 && uploadResp.StatusCode != 201 && uploadResp.StatusCode != 204 {
		bodyBytes, _ := io.ReadAll(uploadResp.Body)
		fmt.Printf("Upload failed with status code: %d\n", uploadResp.StatusCode)
		fmt.Printf("Response body: %s\n", string(bodyBytes))
	}

	return lnpID, nil
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

// GetCorrectMicroDeposits fetches the micro-deposit amounts for a bank
// relationship. Callers create the bank relationship and call this
// immediately afterward with no delay; the amounts occasionally aren't
// queryable yet ("The requested resource does not exist") for a brief window
// after creation, observed intermittently against the real UAT environment.
// Retries with a short fixed backoff to absorb that lag instead of treating
// it as a hard failure.
func GetCorrectMicroDeposits(s *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) ([]string, error) {
	var amounts []string
	// 40 attempts x 2s: a 20x2s window was observed intermittently to be
	// insufficient against the real UAT environment. The SDK's own 504/429
	// backoff (up to 60s per call) is disabled inside the poll -- this is a
	// fast not-found-until-ready check, and stacking the two retry layers
	// multiplies the worst case into tens of minutes.
	err := retryOnTransientErrorN(func() error {
		res, opErr := s.TestSimulation.GetMicroDepositAmounts(ctx, accountId, bankRelationshipId,
			operations.WithRetries(retry.Config{Strategy: "none"}))
		if opErr != nil {
			return opErr
		}
		amounts = []string{*res.MicroDepositAmounts.GetAmount1().Value, *res.MicroDepositAmounts.GetAmount2().Value}
		return nil
	}, 40, 2*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to get micro deposits: %w", err)
	}
	return amounts, nil
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

	require.NoError(t, WaitForAccountOpen(sdk, ctx, *accountId))

	return accountId, nil
}

// WaitForAccountOpen polls until a freshly enrolled account reaches OPEN.
// A new account can take tens of seconds to open against the real UAT
// environment, and downstream resources (micro deposits, transfers, orders)
// misbehave until it does.
func WaitForAccountOpen(sdk *ascendsdk.SDK, ctx context.Context, accountId string) error {
	return retryOnTransientErrorN(func() error {
		res, err := sdk.AccountCreation.GetAccount(ctx, accountId, nil)
		if err != nil {
			return err
		}
		if res.Account == nil || res.Account.State == nil || *res.Account.State != components.AccountStateOpen {
			return fmt.Errorf("account %s has not reached OPEN state", accountId)
		}
		return nil
	}, 20, 2*time.Second)
}

// RetryOnTransientError retries op on any error, for callers that act on a
// resource (an account, bank relationship, etc.) immediately after creating
// or enrolling it. Several such resources have been observed intermittently
// unusable for a window against the real UAT environment (up to ~20s in one
// measured case) before settling -- e.g. "Permission denied on resource (or
// it might not exist)" acting on a freshly enrolled account, or a push
// subscription rejecting mutation right after creation. Scoped to
// immediately-after-creation call sites specifically, since retrying on any
// error is only safe there.
func RetryOnTransientError(op func() error) error {
	return retryOnTransientErrorN(op, 20, 2*time.Second)
}

func retryOnTransientErrorN(op func() error, maxAttempts int, retryDelay time.Duration) error {
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := op()
		if err == nil {
			return nil
		}
		lastErr = err
		if attempt < maxAttempts {
			time.Sleep(retryDelay)
		}
	}
	return lastErr
}
