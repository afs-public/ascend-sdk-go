package instant_cash_transfers

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInstantCashTransfer(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	ctx := context.Background()
	accountIdPtr, err := helpers.CreateEnrolledAccount(sdk, ctx, t)
	require.NoError(t, err)
	accountId := *accountIdPtr
	account, err := sdk.AccountCreation.GetAccount(ctx, accountId, operations.QueryParamViewFull.ToPointer())
	require.NoError(t, err)

	ictDepositId := testInstantCashTransfer_TransfersCreateIctDeposit_CreateIctDeposit1(t, sdk, ctx, account.Account)
	testInstantCashTransfer_TransfersGetIctDeposit_GetIctDeposit1(t, *sdk, ctx, accountId, ictDepositId)
	testInstantCashTransfer_TransfersCancelIctDeposit_CancelIctDeposit1(t, *sdk, ctx, accountId, ictDepositId)
	ictWithdrawalId := CreateIctWithdrawal(t, sdk, ctx, accountId)
	testInstantCashTransfer_TransfersGetIctWithdrawal_GetIctWithdrawal1(t, *sdk, ctx, accountId, ictWithdrawalId)
	testInstantCashTransfer_TransfersCancelIctWithdrawal_CancelIctWithdrawal1(t, *sdk, ctx, accountId, ictWithdrawalId)
	testInstantCashTransfer_TransfersLocateIctReport_LocateIctReport1(t, *sdk, ctx)
}

func testInstantCashTransfer_TransfersCreateIctDeposit_CreateIctDeposit1(t *testing.T,
	s *ascendsdk.SDK, ctx context.Context, account *components.Account) string {
	request := createIctDepositRequest(account)
	resp, err := s.InstantCashTransferICT.CreateIctDeposit(ctx, *account.AccountID, request)
	require.NoError(t, err)

	assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)

	ictDepositName := resp.IctDeposit.Name
	ictDepositNameList := strings.Split(*ictDepositName, "/")
	return ictDepositNameList[3]
}

func createIctDepositRequest(account *components.Account) components.IctDepositCreate {
	address := account.Parties[0].MailingAddress
	person := account.Parties[0].LegalNaturalPerson

	party := &components.TravelRulePartyCreate{
		Address: components.PostalAddressCreate{
			AddressLines:       address.AddressLines,
			RegionCode:         address.RegionCode,
			PostalCode:         address.PostalCode,
			AdministrativeArea: address.AdministrativeArea,
			Locality:           address.Locality,
		},
		FamilyName: *person.FamilyName,
		GivenNames: []string{*person.GivenName},
	}

	return components.IctDepositCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("1000.00"),
		},
		ClientTransferID: GenerateGuid(),
		Program:          components.ProgramBrokerPartner,
		TravelRule: components.IctDepositTravelRuleCreate{
			IndividualOriginatingParty: party,
			IndividualRecipientParty:   party,
			OriginatingInstitution: components.InstitutionCreate{
				AccountID: "09673049",
				Title:     *ascendsdk.String("Default Bank"),
			},
		},
	}
}

func testInstantCashTransfer_TransfersGetIctDeposit_GetIctDeposit1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context, enrolledAccountId string,
	createIctDepositId string) {

	assert.NotNil(t, s)
	res, err := s.InstantCashTransferICT.GetIctDeposit(ctx,
		enrolledAccountId, createIctDepositId)
	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func CreateIctWithdrawal(t *testing.T, s *ascendsdk.SDK, ctx context.Context, accountId string) string {
	res, err := s.InstantCashTransferICT.CreateIctWithdrawal(ctx, accountId, createIctWithdrawalRequest(accountId))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)

	ictWithdrawalName := res.IctWithdrawal.Name
	ictWithdrawalNameList := strings.Split(*ictWithdrawalName, "/")
	return ictWithdrawalNameList[3]
}

func createIctWithdrawalRequest(accountId string) components.IctWithdrawalCreate {
	return components.IctWithdrawalCreate{
		ClientTransferID: GenerateGuid(),
		Program:          components.IctWithdrawalCreateProgramBrokerPartner,
		FullDisbursement: ascendsdk.Bool(true),
		TravelRule: components.IctWithdrawalTravelRuleCreate{
			RecipientInstitution: components.InstitutionCreate{
				AccountID: accountId,
				Title:     *ascendsdk.String("Default Bank"),
			},
		},
	}
}

func testInstantCashTransfer_TransfersCancelIctDeposit_CancelIctDeposit1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context, enrolledAccountId string,
	createIctDepositId string) {

	assert.NotNil(t, s)
	cancelIctDepositRequest := components.CancelIctDepositRequestCreate{
		Name:   "accounts/" + enrolledAccountId + "ictDeposits/" + createIctDepositId + ":cancel",
		Reason: ascendsdk.String("User requested"),
	}
	res, err := s.InstantCashTransferICT.CancelIctDeposit(ctx,
		enrolledAccountId, createIctDepositId, cancelIctDepositRequest)
	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testInstantCashTransfer_TransfersGetIctWithdrawal_GetIctWithdrawal1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context, enrolledAccountId string,
	createIctWithdrawalId string) {

	assert.NotNil(t, s)
	res, err := s.InstantCashTransferICT.GetIctWithdrawal(ctx,
		enrolledAccountId, createIctWithdrawalId)
	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testInstantCashTransfer_TransfersCancelIctWithdrawal_CancelIctWithdrawal1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context, enrolledAccountId string,
	createIctWithdrawalId string) {

	assert.NotNil(t, s)
	cancelIctWithdrawalRequest := components.CancelIctWithdrawalRequestCreate{
		Name:   "accounts/" + enrolledAccountId + "ictWithdrawals/" + createIctWithdrawalId + ":cancel",
		Reason: ascendsdk.String("User requested"),
	}
	res, err := s.InstantCashTransferICT.CancelIctWithdrawal(ctx,
		enrolledAccountId, createIctWithdrawalId, cancelIctWithdrawalRequest)
	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func GenerateGuid() string {
	GUID := uuid.New()
	return GUID.String()
}

func testInstantCashTransfer_TransfersLocateIctReport_LocateIctReport1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context) {

	assert.NotNil(t, s)

	request := operations.IctReconReportsLocateIctReportRequest{
		CorrespondentID:                   os.Getenv("CORRESPONDENT_ID"),
		ProgramDateFilterProgram:          operations.ProgramDateFilterProgramBrokerPartner.ToPointer(),
		ProgramDateFilterProcessDateYear:  ascendsdk.Int(2025),
		ProgramDateFilterProcessDateMonth: ascendsdk.Int(5),
		ProgramDateFilterProcessDateDay:   ascendsdk.Int(28),
	}

	res, err := s.InstantCashTransferICT.LocateIctReport(
		ctx,
		request,
	)

	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
