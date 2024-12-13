package transfers

import (
	"context"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (fixture *Fixtures) pendingAchDeposit() string {
	if fixture.pendingDepositAchAccountId != nil {
		return *fixture.pendingDepositAchAccountId
	}
	achDepositId, err := getPendingAchDeposit(*fixture)

	require.NoError(fixture.t, err)

	fixture.pendingDepositAchAccountId = &achDepositId

	return achDepositId
}

func (fixture *Fixtures) pendingAchWithdrawal() string {

	if fixture.pendingWithdrawalAchAccountId != nil {
		return *fixture.pendingWithdrawalAchAccountId
	}
	achWithdrawalId, err := getPendingAchWithdrawal(*fixture)

	require.NoError(fixture.t, err)

	fixture.pendingWithdrawalAchAccountId = &achWithdrawalId

	return achWithdrawalId
}

func (fixture *Fixtures) pendingIctWithdrawal() string {
	if fixture.pendingIctWithdrawalId != nil {
		return *fixture.pendingIctWithdrawalId
	}
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
	if fixture.completedWithdrawalAccountId != nil {
		return *fixture.completedWithdrawalAccountId
	}
	completedWithdrawalAccountId, err := getCompletedWithdrawalId(*fixture)
	require.NoError(fixture.t, err)
	fixture.completedWithdrawalAccountId = completedWithdrawalAccountId
	return *completedWithdrawalAccountId
}

func Test_Test_Simulation(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	fixture := &Fixtures{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}
	testTestSimulationTransfersForceApproveAchDepositForceApproveAchDeposit1(t, *fixture)
	testTestSimulationTransfersForceRejectAchDepositForceRejectAchDeposit1(t, *fixture)
	testTestSimulationTransfersForceAchDepositReturnForceAchDepositReturn1(t, *fixture)
	testTestSimulationTransfersForceApproveAchWithdrawalForceApproveAchWithdrawal1(t, *fixture)
	testTestSimulationTransfersForceNocAchWithdrawalForceNocAchWithdrawal1(t, *fixture)
	testTestSimulationTransfersForceRejectAchWithdrawalForceRejectAchWithdrawal1(t, *fixture)
	testTestSimulationTransfersForceAchWithdrawalReturnForceAchWithdrawalReturn1(t, *fixture)
	testTestSimulationTransfersForceNocAchDepositForceNocAchDeposit1(t, *fixture)
	testTestSimulationTransfersForceRejectIctWithdrawalForceRejectIctWithdrawal1(t, *fixture)
	testTestSimulationTransfersForceApproveAchWithdrawalForceApproveAchWithdrawal1(t, *fixture)
	testTestSimulationTransfersForceApproveIctWithdrawalForceApproveIctWithdrawal1(t, *fixture)

}

func testTestSimulationTransfersForceNocAchDepositForceNocAchDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Approval")
	}

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
	assert.NotNil(t, fixture.pendingAchDeposit())
	time.Sleep(5 * time.Second)
	request := components.ForceRejectAchDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achDeposits/" + fixture.deceasedBankRelationshipId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectAchDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingDepositAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceAchDepositReturnForceAchDepositReturn1(t *testing.T,
	fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Return")
	}
	assert.NotNil(t, fixture.pendingAchDeposit())
	time.Sleep(5 * time.Second)
	request := components.ForceReturnAchDepositRequestCreate{
		NachaReturn: components.NachaReturnCreate{
			Code: components.NachaReturnCreateCodeR16,
		},
		Name: "accounts/" + fixture.enrolledDepositAccountId + "/achDeposits/" + *fixture.pendingDepositAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceReturnAchDeposit(fixture.ctx, fixture.enrolledDepositAccountId,
		*fixture.pendingDepositAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceApproveAchWithdrawalForceApproveAchWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal Return")
	}

	assert.NotNil(t, fixture.pendingAchWithdrawal())
	time.Sleep(5 * time.Second)

	request := components.ForceApproveAchWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achWithdrawals/" + *fixture.pendingWithdrawalAchAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveAchWithdrawal(fixture.ctx, fixture.deceasedAccountId,
		*fixture.pendingWithdrawalAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
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
	assert.NotNil(t, fixture.pendingAchWithdrawal())
	time.Sleep(5 * time.Second)
	request := components.ForceRejectAchWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achWithdrawals/" + fixture.deceasedBankRelationshipId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectAchWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingWithdrawalAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceAchWithdrawalReturnForceAchWithdrawalReturn1(
	t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Withdrawal Return")
	}
	request := components.ForceReturnAchWithdrawalRequestCreate{
		NachaReturn: components.NachaReturnCreate{
			Code: components.NachaReturnCreateCodeR16,
		},
		Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/achWithdrawals/" + *fixture.completedWithdrawalAccountId,
	}
	res, err := fixture.sdk.TestSimulation.ForceReturnAchWithdrawal(fixture.ctx, fixture.enrolledWithdrawalAccountId, *fixture.completedWithdrawalAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceApproveIctWithdrawalForceApproveIctWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: Force Approve ICT Withdrawal")
	}
	assert.NotNil(t, fixture.pendingIctWithdrawal())
	request := components.ForceApproveIctWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictWithdrawals/" + *fixture.pendingIctWithdrawalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveIctWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctWithdrawalId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceRejectIctWithdrawalForceRejectIctWithdrawal1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: Force Reject ICT Withdrawal")
	}
	assert.NotNil(t, fixture.pendingIctWithdrawal())
	request := components.ForceRejectIctWithdrawalRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/ictWithdrawals/" + *fixture.pendingIctWithdrawalId,
	}
	res, err := fixture.sdk.TestSimulation.ForceRejectIctWithdrawal(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingIctWithdrawalId, request)
	require.NoError(t, err)
	assert.NotNil(t, &res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testTestSimulationTransfersForceApproveAchDepositForceApproveAchDeposit1(t *testing.T, fixture Fixtures) {
	if isNotDuringTradingHours() {
		t.Skip("Skipping Endpoint Test: ACH Deposit Approval")
	}
	assert.NotNil(t, fixture.pendingAchDeposit())
	time.Sleep(5 * time.Second)
	request := components.ForceApproveAchDepositRequestCreate{
		Name: "accounts/" + fixture.deceasedAccountId + "/achDeposits/" + fixture.deceasedBankRelationshipId,
	}
	res, err := fixture.sdk.TestSimulation.ForceApproveAchDeposit(fixture.ctx, fixture.deceasedAccountId, *fixture.pendingDepositAchAccountId, request)
	require.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
