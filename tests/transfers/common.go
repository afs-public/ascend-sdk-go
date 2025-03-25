package transfers

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/google/uuid"
)

const Wire_Deposit_ID = "20250218014356"

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
	pendingIctWithdrawalId        *string
	pendingIctDepositId           *string
	deceasedAccountId             string
	deceasedBankRelationshipId    string
	accountId                     string
	feeId                         *string
	creditId                      *string
	wireId                        *string
	cashJournalId                 *string
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
				AccountID: fixture.deceasedAccountId,
				Title:     "Default Bank",
			},
		},
	}
	response, err := fixture.sdk.InstantCashTransferICT.CreateIctDeposit(fixture.ctx, fixture.deceasedAccountId, ictDepositRequest)
	if err != nil {
		return "", err
	}
	depositId := strings.Split(*response.IctDeposit.Name, "/")
	return depositId[len(depositId)-1], nil
}

func getAchDeposit(accountID string, bankRelationship string, fixture Fixtures) (string, error) {
	achDepositRequest := components.AchDepositCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("1000.00"),
		},
		BankRelationship: "accounts/" + accountID + "/bankRelationships/" + bankRelationship,
		ClientTransferID: uuid.New().String(),
	}
	response, err := fixture.sdk.ACHTransfers.CreateAchDeposit(fixture.ctx, accountID, achDepositRequest)
	if err != nil {
		return "", err
	}
	achDepositName := response.AchDeposit.Name
	depositId := strings.Split(*achDepositName, "/")
	return depositId[len(depositId)-1], nil
}

func getAchWithdrawal(accountID string, bankRelationship string, fixture Fixtures) (string, error) {
	achWithdrawalRequest := components.AchWithdrawalCreate{
		Amount: &components.DecimalCreate{
			Value: ascendsdk.String("1.00"),
		},
		BankRelationship: "accounts/" + accountID + "/bankRelationships/" + bankRelationship,
		ClientTransferID: uuid.New().String(),
	}
	response, err := fixture.sdk.ACHTransfers.CreateAchWithdrawal(fixture.ctx, accountID, achWithdrawalRequest)
	if err != nil {
		return "", err
	}
	achWithdrawalName := response.AchWithdrawal.Name
	withdrawalId := strings.Split(*achWithdrawalName, "/")
	return withdrawalId[len(withdrawalId)-1], nil
}

func getPendingIctWithdrawal(fixture Fixtures) (string, error) {
	ictWithdrawalRequest := components.IctWithdrawalCreate{
		Amount: &components.DecimalCreate{
			Value: ascendsdk.String("0.01"),
		},
		ClientTransferID: uuid.New().String(),
		Program:          components.IctWithdrawalCreateProgramBrokerPartner,
		TravelRule: components.IctWithdrawalTravelRuleCreate{
			RecipientInstitution: components.InstitutionCreate{
				AccountID: "09673049",
				Title:     "Default Bank",
			},
		},
	}
	fmt.Println("Creating ICT Withdrawal")
	response, err := fixture.sdk.InstantCashTransferICT.CreateIctWithdrawal(fixture.ctx,
		fixture.deceasedAccountId, ictWithdrawalRequest)
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
			cancelRelationShipId := strings.Split(*relationship.Name, "/")[3]
			cancelRequest := components.CancelBankRelationshipRequestCreate{
				Name:    "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + cancelRelationShipId,
				Comment: "Canceling Bank User Request",
			}
			_, err = fixture.sdk.BankRelationships.CancelBankRelationship(fixture.ctx, fixture.enrolledWithdrawalAccountId, cancelRelationShipId, cancelRequest)
			if err != nil {
				return nil, err
			}
		}
	}
	bankRelationShipPtr, err := helpers.CreateBankRelationship(fixture.sdk, fixture.ctx, fixture.enrolledWithdrawalAccountId)
	if err != nil {
		return nil, err
	}
	time.Sleep(15 * time.Second)
	bankRelationShip := *bankRelationShipPtr
	amounts, err := helpers.GetCorrectMicroDeposits(fixture.sdk, fixture.ctx, fixture.enrolledWithdrawalAccountId, bankRelationShip)
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	err = helpers.VerifyMicroDeposits(fixture.sdk, fixture.ctx, fixture.enrolledWithdrawalAccountId, bankRelationShip, amounts)
	if err != nil {
		return nil, err
	}
	helpers.Wait()
	achWithdrawalRequest := components.AchWithdrawalCreate{
		BankRelationship: "accounts/" + fixture.enrolledWithdrawalAccountId + "/bankRelationships/" + bankRelationShip,
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

func getWireWithdrawalId(accountId string, fixture Fixtures) (*string, error) {
	wireWithdrawalCreate := components.WireWithdrawalCreate{
		Amount: &components.DecimalCreate{
			Value: ascendsdk.String("1.00"),
		},
		ClientTransferID: uuid.New().String(),
		Beneficiary: components.WireWithdrawalBeneficiaryCreate{
			Account:      accountId,
			AccountTitle: ascendsdk.String("Test"),
			Address: &components.AddressCreate{
				City:          ascendsdk.String("Portland"),
				Country:       ascendsdk.String("USA"),
				PostalCode:    ascendsdk.String("97201"),
				State:         ascendsdk.String("OR"),
				StreetAddress: []string{"123 Main St"},
			},
			ThirdParty: ascendsdk.Bool(false),
		},
		RecipientBank: components.WireWithdrawalRecipientBankCreate{
			BankID: components.RecipientBankBankIDCreate{
				ID:   *ascendsdk.String("011000028"),
				Type: components.RecipientBankBankIDCreateType("ABA"),
			},
		},
	}
	time.Sleep(5 * time.Second)
	res, err := fixture.sdk.Wires.CreateWireWithdrawal(fixture.ctx, accountId, wireWithdrawalCreate)
	if err != nil {
		return nil, err
	}
	wireId := strings.Split(*res.WireWithdrawal.Name, "/")[3]
	return &wireId, nil
}

func isNotDuringTradingHours() bool {
	//init the loc
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// Handle error appropriately (e.g., log and return false)
		fmt.Println("Error loading location:", err)
		return false
	}
	now := time.Now().In(loc)
	currentTime := now.Hour()*60 + now.Minute()
	startTrading := 23*60 + 30 // 23:30 in minutes
	endTrading := 18 * 60      // 18:00 in minutes
	return !(currentTime >= startTrading || currentTime <= endTrading)
}

func isNotDuringICTHours() bool {
	//init the loc
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// Handle error appropriately (e.g., log and return false)
		fmt.Println("Error loading location:", err)
		return false
	}
	now := time.Now().In(loc)
	currentTime := now.Hour()*60 + now.Minute()
	startTrading := 6*60 + 30 // 6:00 in minutes
	endTrading := 15 * 60     // 15:00 in minutes
	return !(startTrading <= currentTime && currentTime <= endTrading)
}

func isNotWireHours() bool {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// Handle error appropriately (e.g., log and return false)
		fmt.Println("Error loading location:", err)
		return false
	}
	now := time.Now().In(loc)
	currentTime := now.Hour()*60 + now.Minute()
	startTrading := 7 * 60   // 7:00 in minutes
	endTrading := 14*60 + 30 // 14:30 in minutes
	return !(startTrading <= currentTime && currentTime <= endTrading)
}

func isNotJournalHours() bool {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// Handle error appropriately (e.g., log and return false)
		fmt.Println("Error loading location:", err)
		return false
	}
	now := time.Now().In(loc)
	currentTime := now.Hour()*60 + now.Minute()
	startTrading := 6 * 60 // 6:00 in minutes
	endTrading := 19 * 60  // 19:00 in minutes
	return !(startTrading <= currentTime && currentTime <= endTrading)
}
