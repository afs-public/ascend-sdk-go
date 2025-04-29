package transfers

import (
	"context"
	"strings"
	"testing"
	"time"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/google/uuid"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/sdkerrors"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (fixture *Fixtures) AchDeposit(accountID string, bankRelationship string) string {
	achDepositId, err := getAchDeposit(accountID, bankRelationship, *fixture)

	require.NoError(fixture.t, err)

	fixture.pendingDepositAchAccountId = &achDepositId

	return achDepositId
}

func (fixture *Fixtures) AchWithdrawal(accountID string, bankRelationship string) string {
	achWithdrawalId, err := getAchWithdrawal(accountID, bankRelationship, *fixture)

	require.NoError(fixture.t, err)

	fixture.pendingWithdrawalAchAccountId = &achWithdrawalId

	return achWithdrawalId
}

func (fixture *Fixtures) pendingIctWithdrawal() string {
	pendingIctWithdrawalId, err := getPendingIctWithdrawal(*fixture)

	require.NoError(fixture.t, err)

	fixture.pendingIctWithdrawalId = &pendingIctWithdrawalId

	return pendingIctWithdrawalId
}

func (fixture *Fixtures) pendingIctDeposit() string {
	if fixture.pendingIctDepositId != nil {
		return *fixture.pendingIctDepositId
	}
	pendingIctDepositId, err := getPendingIctDeposit(*fixture)
	require.NoError(fixture.t, err)
	fixture.pendingIctDepositId = &pendingIctDepositId
	return pendingIctDepositId
}

func (fixture *Fixtures) completedWithdrawalId() string {
	completedWithdrawalAccountId, err := getCompletedWithdrawalId(*fixture)
	require.NoError(fixture.t, err)
	fixture.completedWithdrawalAccountId = completedWithdrawalAccountId
	return *completedWithdrawalAccountId
}

func (fixture *Fixtures) pendingWireWithdrawal(accountId string) string {
	if fixture.wireId != nil {
		return *fixture.wireId
	}
	pendingWireWithdrawalId, err := getWireWithdrawalId(accountId, *fixture)
	require.NoError(fixture.t, err)
	fixture.wireId = pendingWireWithdrawalId
	return *pendingWireWithdrawalId
}

func (fixture *Fixtures) pendingCashJournal(accountId string) string {
	if fixture.cashJournalId != nil {
		return *fixture.cashJournalId
	}
	pendingCashJournalId, err := createCashJournal(fixture.ctx, *fixture.sdk, &accountId)
	require.NoError(fixture.t, err)
	fixture.cashJournalId = &pendingCashJournalId
	return pendingCashJournalId
}

func Test_Test_Simulation(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	accountIdPtr, err := helpers.CreateEnrolledAccount(sdk, ctx, t)
	require.NoError(t, err)
	helpers.Wait()
	accountId := *accountIdPtr
	bankRelationShipPtr, err := helpers.CreateBankRelationship(sdk, ctx, accountId)
	require.NoError(t, err)
	helpers.Wait()
	bankRelationShip := *bankRelationShipPtr
	amounts, err := helpers.GetCorrectMicroDeposits(sdk, ctx, accountId, bankRelationShip)
	require.NoError(t, err)
	helpers.Wait()
	err = helpers.VerifyMicroDeposits(sdk, ctx, accountId, bankRelationShip, amounts)
	require.NoError(t, err)

	fixture := &Fixtures{
		t:                           t,
		sdk:                         sdk,
		ctx:                         ctx,
		deceasedAccountId:           "01JHK07CRQ9X8P5XE9JWG4PFSP",
		deceasedBankRelationshipId:  "6786a8e8ea916b424a53cc24",
		enrolledDepositAccountId:    accountId,
		bankRelationshipDepositId:   bankRelationShip,
		enrolledWithdrawalAccountId: "01JHGTEPC6ZTAHCFRH2MD3VJJT",
	}

	t.Run("test Test Simulation Transfers Force Approve Ach Deposit Force Approve Ach Deposit1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveAchDepositForceApproveAchDeposit1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Noc Ach Deposit Force Noc Ach Deposit1", func(t *testing.T) {
		testTestSimulationTransfersForceNocAchDepositForceNocAchDeposit1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Ach Deposit Force Reject Ach Deposit1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectAchDepositForceRejectAchDeposit1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Ach Deposit Return Force Ach Deposit Return1", func(t *testing.T) {
		testTestSimulationTransfersForceAchDepositReturnForceAchDepositReturn1(t, *fixture)
	})

	t.Run("Test Test Simulation Transfers Force Approve Ach Withdrawal Force Approve Ach Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveAchWithdrawalForceApproveAchWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Noc Ach Withdrawal Force Noc Ach Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceNocAchWithdrawalForceNocAchWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Ach Withdrawal Force Reject Ach Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectAchWithdrawalForceRejectAchWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Ach Withdrawal Return Force Ach Withdrawal Return1", func(t *testing.T) {
		testTestSimulationTransfersForceAchWithdrawalReturnForceAchWithdrawalReturn1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Approve Ict Deposit Force Approve Ict Deposit1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveIctDepositForceApproveIctDeposit1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Ict Deposit Force Reject Ict Deposit1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectIctDepositForceRejectIctDeposit1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Approve Ict Withdrawal Force Approve Ict Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveIctWithdrawalForceApproveIctWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Ict Withdrawal Force Reject Ict Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectIctWithdrawalForceRejectIctWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Approve Wire Withdrawal Force Approve Wire Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveWireWithdrawalForceApproveWireWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Wire Withdrawal Force Reject Wire Withdrawal1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectWireWithdrawalForceRejectWireWithdrawal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Approve Cash Journal Force Approve Cash Journal1", func(t *testing.T) {
		testTestSimulationTransfersForceApproveCashJournalForceApproveCashJournal1(t, *fixture)
	})

	t.Run("test Test Simulation Transfers Force Reject Cash Journal Force Reject Cash Journal1", func(t *testing.T) {
		testTestSimulationTransfersForceRejectCashJournalForceRejectCashJournal1(t, *fixture)
	})
}

func testTestSimulationTransfersForceNocAchDepositForceNocAchDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Approval")
	}
	assert.NotNil(t, fixture.AchDeposit(fixture.enrolledDepositAccountId, fixture.bankRelationshipDepositId))
	time.Sleep(5 * time.Second)
	request := components.ForceNocAchDepositRequestCreate{
		NachaNoc: components.NachaNocCreate{
			Code:                   components.CodeC05,
			UpdatedBankAccountType: components.UpdatedBankAccountTypeChecking.ToPointer(),
		},
		Name: "accounts/" + fixture.enrolledDepositAccountId + "/achDeposits/" + *fixture.pendingDepositAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceNocAchDeposit(fixture.ctx, fixture.enrolledDepositAccountId, *fixture.pendingDepositAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceRejectAchDepositForceRejectAchDeposit1(t *testing.T,
	fixture Fixtures) {
	if isNotDuringTradingHours() {
		fixture.t.Skip("Skipping Endpoint Test: ACH Deposit Approval")
	}
	assert.NotNil(t, fixture.AchDeposit(fixture.deceasedAccountId, fixture.deceasedBankRelationshipId))
	time.Sleep(5 * time.Second)
	request := components.ForceRejectAchDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achDeposits/" + *fixture.pendingDepositAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectAchDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingDepositAchAccountId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceAchDepositReturnForceAchDepositReturn1(t *testing.T,
	fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Return")
	}
	assert.NotNil(t, fixture.AchDeposit(fixture.enrolledDepositAccountId, fixture.bankRelationshipDepositId))
	time.Sleep(5 * time.Second)
	request := components.ForceReturnAchDepositRequestCreate{
		NachaReturn: components.NachaReturnCreate{
			Code: components.NachaReturnCreateCodeR16,
		},
		Name: "accounts/" + fixture.enrolledDepositAccountId + "/achDeposits/" + *fixture.pendingDepositAchAccountId,
	}
	_, err := fixture.sdk.TestSimulation.ForceReturnAchDeposit(fixture.ctx, fixture.enrolledDepositAccountId,
		*fixture.pendingDepositAchAccountId, request)
	statusErr, ok := err.(*sdkerrors.Status)
	require.True(t, ok)
	assert.Equal(t, 3, *statusErr.Code)
	assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "current state"), statusErr.Code)
}

func testTestSimulationTransfersForceApproveAchWithdrawalForceApproveAchWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal Return")
	}

	assert.NotNil(t, fixture.AchWithdrawal(fixture.deceasedAccountId, fixture.deceasedBankRelationshipId))
	time.Sleep(5 * time.Second)

	request := components.ForceApproveAchWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achWithdrawals/" + *fixture.pendingWithdrawalAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveAchWithdrawal(fixture.ctx, fixture.deceasedAccountId,
		*fixture.pendingWithdrawalAchAccountId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceNocAchWithdrawalForceNocAchWithdrawal1(
	t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal NOC")
	}
	assert.NotNil(t, fixture.completedWithdrawalId())
	request := components.ForceNocAchWithdrawalRequestCreate{
		NachaNoc: components.NachaNocCreate{
			Code:                   components.CodeC05,
			UpdatedBankAccountType: components.UpdatedBankAccountTypeChecking.ToPointer(),
		},
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/achWithdrawals/" + *fixture.completedWithdrawalAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceNocAchWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, *fixture.completedWithdrawalAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceRejectAchWithdrawalForceRejectAchWithdrawal1(
	t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal Rejection")
	}
	assert.NotNil(t, fixture.AchWithdrawal(fixture.deceasedAccountId, fixture.deceasedBankRelationshipId))
	time.Sleep(5 * time.Second)
	request := components.ForceRejectAchWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achWithdrawals/" + *fixture.pendingWithdrawalAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectAchWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingWithdrawalAchAccountId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceAchWithdrawalReturnForceAchWithdrawalReturn1(
	t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal Return")
	}
	assert.NotNil(t, fixture.completedWithdrawalId())
	time.Sleep(5 * time.Second)
	request := components.ForceReturnAchWithdrawalRequestCreate{
		NachaReturn: components.NachaReturnCreate{
			Code: components.NachaReturnCreateCodeR16,
		},
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/achWithdrawals/" + *fixture.completedWithdrawalAccountId,
	}
	_, err := fixture.sdk.TestSimulation.ForceReturnAchWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, *fixture.completedWithdrawalAccountId, request)
	statusErr, ok := err.(*sdkerrors.Status)
	require.True(t, ok)
	assert.Equal(t, 3, *statusErr.Code)
	assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "current state"), statusErr.Code)
}

func testTestSimulationTransfersForceApproveIctWithdrawalForceApproveIctWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringICTHours() {
		t.Skip("Skipping Endpoint Test: Force Approve ICT Withdrawal")
	}
	assert.NotNil(t, fixture.pendingIctWithdrawal())
	time.Sleep(10 * time.Second)
	request := components.ForceApproveIctWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictWithdrawals/" + *fixture.pendingIctWithdrawalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveIctWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctWithdrawalId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceRejectIctWithdrawalForceRejectIctWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringICTHours() {
		t.Skip("Skipping Endpoint Test: Force Reject ICT Withdrawal")
	}
	assert.NotNil(t, fixture.pendingIctWithdrawal())
	time.Sleep(10 * time.Second)
	request := components.ForceRejectIctWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictWithdrawals/" + *fixture.pendingIctWithdrawalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectIctWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctWithdrawalId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceApproveIctDepositForceApproveIctDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringICTHours() {
		t.Skip("Skipping Endpoint Test: Force Approve ICT Deposit")
	}
	assert.NotNil(t, fixture.pendingIctDeposit())
	time.Sleep(10 * time.Second)
	request := components.ForceApproveIctDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictDeposits/" + *fixture.pendingIctDepositId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveIctDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctDepositId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceRejectIctDepositForceRejectIctDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringICTHours() {
		t.Skip("Skipping Endpoint Test: Force Reject ICT Deposit")
	}
	assert.NotNil(t, fixture.pendingIctDeposit())
	time.Sleep(10 * time.Second)
	request := components.ForceRejectIctDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictDeposits/" + *fixture.pendingIctDepositId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectIctDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctDepositId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceApproveAchDepositForceApproveAchDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Approval")
	}
	assert.NotNil(t, fixture.AchDeposit(fixture.deceasedAccountId, fixture.deceasedBankRelationshipId))
	request := components.ForceApproveAchDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achDeposits/" + *fixture.pendingDepositAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveAchDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingDepositAchAccountId, request)
	if err != nil {
		statusErr, ok := err.(*sdkerrors.Status)
		require.True(t, ok)
		assert.Equal(t, 3, *statusErr.Code)
		assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), statusErr.Code)
	} else {
		assert.NotNil(t, res)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}

func testTestSimulationTransfersForceApproveWireWithdrawalForceApproveWireWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotWireHours() {
		t.Skip("Skipping Endpoint Test: Force Approve Wire Withdrawal")
	}
	assert.NotNil(t, fixture.pendingWireWithdrawal(fixture.enrolledWithdrawalAccountId))
	assert.NotNil(t, fixture.wireId)
	request := components.ForceApproveWireWithdrawalRequestCreate{
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/wireWithdrawals/" + *fixture.wireId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveWireWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, *fixture.wireId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceRejectWireWithdrawalForceRejectWireWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotWireHours() {
		t.Skip("Skipping Endpoint Test: Force Reject Wire Withdrawal")
	}
	assert.NotNil(t, fixture.pendingWireWithdrawal(fixture.enrolledWithdrawalAccountId))
	assert.NotNil(t, fixture.wireId)
	request := components.ForceRejectWireWithdrawalRequestCreate{
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/wireWithdrawals/" + *fixture.wireId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectWireWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, *fixture.wireId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceApproveCashJournalForceApproveCashJournal1(t *testing.T, fixture Fixtures) {
	if isNotJournalHours() {
		t.Skip("Skipping Endpoint Test: Force Approve Cash Journal")
	}

	// Counter the amount of money the cash journal is taking
	credit := components.TransfersCreditCreate{
		Amount:           components.DecimalCreate{Value: ascendsdkgo.String("5000000.00")},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdkgo.String("Credit given as promotion"),
		Type:             components.TransfersCreditCreateTypePromotional,
	}
	_, _ = fixture.sdk.FeesAndCredits.CreateCredit(fixture.ctx, fixture.enrolledWithdrawalAccountId, credit)

	assert.NotNil(t, fixture.pendingCashJournal(fixture.deceasedAccountId))

	time.Sleep(5 * time.Second)

	request := components.ForceApproveCashJournalRequestCreate{
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/cashJournals/" + *fixture.cashJournalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveCashJournal(fixture.ctx, *fixture.cashJournalId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceRejectCashJournalForceRejectCashJournal1(t *testing.T, fixture Fixtures) {
	if isNotJournalHours() {
		t.Skip("Skipping Endpoint Test: Force Reject Cash Journal")
	}
	assert.NotNil(t, fixture.pendingCashJournal(fixture.deceasedAccountId))

	time.Sleep(5 * time.Second)

	request := components.ForceRejectCashJournalRequestCreate{
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/cashJournals/" + *fixture.cashJournalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectCashJournal(fixture.ctx, *fixture.cashJournalId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
