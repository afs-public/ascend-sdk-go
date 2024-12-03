package instant_cash_transfers

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"ascend-sdk/models/operations"
	"ascend-sdk/tests"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInstantCashTransfer(t *testing.T) {
	sdk, err := tests.SetupAscendSDK()
	require.NoError(t, err)
	ctx := context.Background()
	accountId := "01J7XASQ2EGHNVENARVENT2HTG"
	ictDepositId := testInstantCashTransfer_TransfersCreateIctDeposit_CreateIctDeposit1(t, sdk, ctx, accountId)
	testInstantCashTransfer_TransfersGetIctDeposit_GetIctDeposit1(t, *sdk, ctx, accountId, ictDepositId)
	testInstantCashTransfer_TransfersCancelIctDeposit_CancelIctDeposit1(t, *sdk, ctx, accountId, ictDepositId)
	ictWithdrawalId := CreateIctWithdrawal(t, sdk, ctx, accountId)
	testInstantCashTransfer_TransfersGetIctWithdrawal_GetIctWithdrawal1(t, *sdk, ctx, accountId, ictWithdrawalId)
	testInstantCashTransfer_TransfersCancelIctWithdrawal_CancelIctWithdrawal1(t, *sdk, ctx, accountId, ictWithdrawalId)
	testInstantCashTransfer_TransfersLocateIctReport_LocateIctReport1(t, *sdk, ctx)
}

func testInstantCashTransfer_TransfersCreateIctDeposit_CreateIctDeposit1(t *testing.T,
	s *ascendsdk.SDK, ctx context.Context, accountId string) string {
	resp, err := s.InstantCashTransferICT.CreateIctDeposit(ctx, accountId, createIctDepositRequest(accountId))
	require.NoError(t, err)

	assert.Equal(t, 200, resp.GetStatus().GetCode())
	ictDepositName := resp.IctDeposit.Name
	ictDepositNameList := strings.Split(*ictDepositName, "/")
	return ictDepositNameList[3]
}

func createIctDepositRequest(accountId string) components.IctDepositCreate {
	return components.IctDepositCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("1000.00"),
		},
		ClientTransferID: GenerateGuid(),
		Program:          components.ProgramBrokerPartner,
		TravelRule: components.IctDepositTravelRuleCreate{
			IndividualRecipientParty: &components.TravelRulePartyCreate{
				Address: components.PostalAddressCreate{
					AddressLines:       []string{"1 Main Street"},
					RegionCode:         ascendsdk.String("US"),
					PostalCode:         ascendsdk.String("12345"),
					AdministrativeArea: ascendsdk.String("NY"),
					Locality:           ascendsdk.String("New York"),
				},
				FamilyName: "Dough",
				GivenNames: []string{"Jane"},
			},
			OriginatingInstitution: components.InstitutionCreate{
				AccountID: accountId,
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

var clientBatchSourceID = GenerateGuid()

func testInstantCashTransfer_TransfersLocateIctReport_LocateIctReport1(
	t *testing.T, s ascendsdk.SDK, ctx context.Context) {

	assert.NotNil(t, s)
	res, err := s.InstantCashTransferICT.LocateIctReport(
		ctx,
		os.Getenv("CORRESPONDENT_ID"),
		&clientBatchSourceID,
		operations.ProgramDateFilterProgramBrokerPartner.ToPointer(),
		&components.DateCreate{Day: ascendsdk.Int(1), Month: ascendsdk.Int(1), Year: ascendsdk.Int(2021)})
	assert.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
