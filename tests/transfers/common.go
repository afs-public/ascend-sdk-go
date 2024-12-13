package transfers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/google/uuid"
)

type Fixtures struct {
	t                             *testing.T
	sdk                           *ascendsdk.SDK
	ctx                           context.Context
	enrolledDepositAccountId      string
	enrolledWithdrawalAccountId   string
	completedWithdrawalAccountId  *string
	pendingDepositAchAccountId    *string
	pendingWithdrawalAchAccountId *string
	bankRelationshipDepositId     string
	bankRelationshipWithdrawalId  string
	ictWithdrawalAccountId        string
	ictDepositAccountId           string
	pendingIctWithdrawalId        *string
	pendingIctDepositId           *string
	deceasedAccountId             string
	deceasedBankRelationshipId    string
}

func getMicrodepositAmounts(sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) (components.MicroDepositAmounts, error) {
	res, err := sdk.TestSimulation.GetMicroDepositAmounts(ctx, accountId, bankRelationshipId)
	if err != nil {
		return components.MicroDepositAmounts{}, err
	}
	return *res.MicroDepositAmounts, nil
}

func getPendingIctDeposit(fixture Fixtures) (string, error) {
	ictDepositRequest := components.IctDepositCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("1000.00"),
		},
		ClientTransferID: uuid.New().String(),
		Program:          components.ProgramBrokerPartner,
		TravelRule: components.IctDepositTravelRuleCreate{
			IndividualOriginatingParty: &components.TravelRulePartyCreate{
				Address: components.PostalAddressCreate{
					AddressLines:       []string{"19409 Sherilyn Courts"},
					RegionCode:         ascendsdk.String("US"),
					PostalCode:         ascendsdk.String("97035"),
					AdministrativeArea: ascendsdk.String("OR"),
					Locality:           ascendsdk.String("Portland"),
				},
				FamilyName: "Jacob",
				GivenNames: []string{"Bob"},
			},
			IndividualRecipientParty: &components.TravelRulePartyCreate{
				Address: components.PostalAddressCreate{
					AddressLines:       []string{"19409 Sherilyn Courts"},
					RegionCode:         ascendsdk.String("US"),
					PostalCode:         ascendsdk.String("97035"),
					AdministrativeArea: ascendsdk.String("OR"),
					Locality:           ascendsdk.String("Portland"),
				},
				FamilyName: "Jacob",
				GivenNames: []string{"Bob"},
			},
			OriginatingInstitution: components.InstitutionCreate{
				AccountID: fixture.ictDepositAccountId,
				Title:     "Default Bank",
			},
		},
	}
	response, err := fixture.sdk.InstantCashTransferICT.CreateIctDeposit(fixture.ctx, fixture.ictDepositAccountId, ictDepositRequest)
	if err != nil {
		return "", err
	}
	depositId := strings.Split(*response.IctDeposit.Name, "/")
	return depositId[len(depositId)-1], nil
}

func getPendingAchDeposit(fixture Fixtures) (string, error) {
	achDepositRequest := components.AchDepositCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("1000.00"),
		},
		BankRelationship: "accounts/" + fixture.enrolledDepositAccountId + "/bankRelationships/" + fixture.bankRelationshipDepositId,
		ClientTransferID: uuid.New().String(),
	}
	response, err := fixture.sdk.ACHTransfers.CreateAchDeposit(fixture.ctx, fixture.enrolledDepositAccountId, achDepositRequest)
	if err != nil {
		return "", err
	}
	achDepositName := response.AchDeposit.Name
	depositId := strings.Split(*achDepositName, "/")
	return depositId[len(depositId)-1], nil
}

func getPendingAchWithdrawal(fixture Fixtures) (string, error) {
	achWithdrawalRequest := components.AchWithdrawalCreate{
		Amount: &components.DecimalCreate{
			Value: ascendsdk.String("1.00"),
		},
		BankRelationship: "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + fixture.bankRelationshipWithdrawalId,
		ClientTransferID: uuid.New().String(),
	}
	response, err := fixture.sdk.ACHTransfers.CreateAchWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, achWithdrawalRequest)
	if err != nil {
		return "", err
	}
	achWithdrawalId := response.AchWithdrawal.Name
	withdrawalId := strings.Split(*achWithdrawalId, "/")
	return withdrawalId[len(withdrawalId)-1], nil
}

func getPendingIctWithdrawal(fixture Fixtures) (string, error) {
	ictWithdrawalRequest := components.IctWithdrawalCreate{
		Amount: &components.DecimalCreate{
			Value: ascendsdk.String("10.00"),
		},
		ClientTransferID: uuid.New().String(),
		Program:          components.IctWithdrawalCreateProgramBankingPartner,
		TravelRule: components.IctWithdrawalTravelRuleCreate{
			RecipientInstitution: components.InstitutionCreate{
				AccountID: fixture.ictWithdrawalAccountId,
				Title:     "Default Bank",
			},
		},
	}
	response, err := fixture.sdk.InstantCashTransferICT.CreateIctWithdrawal(fixture.ctx,
		fixture.ictWithdrawalAccountId, ictWithdrawalRequest)
	if err != nil {
		return "", err
	}
	withdrawalId := strings.Split(*response.IctWithdrawal.Name, "/")
	return withdrawalId[len(withdrawalId)-1], nil
}

func getCompletedWithdrawalId(fixture Fixtures) (*string, error) {
	res, err := fixture.sdk.BankRelationships.ListBankRelationships(fixture.ctx, fixture.enrolledWithdrawalAccountId, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	for _, relationship := range res.ListBankRelationshipsResponse.BankRelationships {
		if *relationship.State.State == components.BankRelationshipStateStateApproved {
			cancelRequest := components.CancelBankRelationshipRequestCreate{
				Name:    "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + *relationship.Name,
				Comment: "Canceling Bank User Request",
			}
			_, err := fixture.sdk.BankRelationships.CancelBankRelationship(fixture.ctx, fixture.enrolledWithdrawalAccountId, *relationship.Name, cancelRequest)
			if err != nil {
				return nil, err
			}
		}
	}
	bankRelationshipRequest := components.BankRelationshipCreate{
		BankAccount: &components.BankAccountCreate{
			AccountNumber: *ascendsdk.String(strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)),
			Owner:         "TEST123",
			RoutingNumber: "112203216",
			Type:          components.BankAccountCreateTypeSavings,
		},
		Nickname:           "ACH TEST",
		VerificationMethod: components.VerificationMethodMicroDeposit,
	}
	resCreateBankRelationship, err := fixture.sdk.BankRelationships.CreateBankRelationship(
		fixture.ctx, fixture.enrolledWithdrawalAccountId, bankRelationshipRequest, nil)
	if err != nil {
		return nil, err
	}
	bankRelationshipId := strings.Split(*resCreateBankRelationship.BankRelationship.Name, "/")[3]
	time.Sleep(15 * time.Second)
	microDepositAmounts, err := getMicrodepositAmounts(fixture.sdk, fixture.ctx, fixture.enrolledWithdrawalAccountId, bankRelationshipId)
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	microDepositsRequest := components.VerifyMicroDepositsRequestCreate{
		Amounts: components.MicroDepositAmountsCreate{
			Amount1: components.DecimalCreate{
				Value: microDepositAmounts.Amount1.Value,
			},
			Amount2: components.DecimalCreate{
				Value: microDepositAmounts.Amount2.Value,
			},
		},
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + bankRelationshipId,
	}
	_, err = fixture.sdk.BankRelationships.VerifyMicroDeposits(fixture.ctx, fixture.enrolledWithdrawalAccountId, bankRelationshipId, microDepositsRequest)
	if err != nil {
		return nil, err
	}
	time.Sleep(5 * time.Second)
	achWithdrawalRequest := components.AchWithdrawalCreate{
		BankRelationship: "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + bankRelationshipId,
		Amount:           &components.DecimalCreate{Value: ascendsdk.String("0.01")},
		ClientTransferID: uuid.New().String(),
		FullDisbursement: ascendsdk.Bool(false),
		Memo:             ascendsdk.String("ACH"),
	}
	response, err := fixture.sdk.ACHTransfers.CreateAchWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, achWithdrawalRequest)
	if err != nil {
		return nil, err
	}
	withdrawalId := strings.Split(*response.AchWithdrawal.Name, "/")
	return &withdrawalId[len(withdrawalId)-1], nil
}

func isNotDuringTradingHours() bool {
	//init the loc
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		// Handle error appropriately (e.g., log and return false)
		fmt.Println("Error loading location:", err)
		return false
	}
	now := time.Now().In(loc)
	hour := now.Hour()
	return hour < 6 || hour > 15
}
