package ach_transfers

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"ascend-sdk/tests"
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAchTransfers(t *testing.T) {
	accountId := "01J5T7W3T7PYCYCH4YQADAQ7VD"
	bankRelationshipId := "66c5c6bd8603496314c4678f"
	ctx := context.Background()
	sdk, err := tests.SetupAscendSDK()
	require.NoError(t, err)
	achDepositId := testAchTransfers_TransfersCreateAchDeposit_CreateAchDeposit1(t, *sdk, ctx, accountId, bankRelationshipId)
	testAchTransfers_TransfersGetAchDeposit_GetAchDeposit1(t, *sdk, ctx, accountId, achDepositId)
	testAchTransfers_TransfersCancelAchDeposit_CancelAchDeposit1(t, *sdk, ctx, accountId, achDepositId)
	withdrawalId := testAchTransfers_TransfersCreateAchWithdrawal_CreateAchWithdrawal1(t, *sdk, ctx, accountId, bankRelationshipId)
	testAchTransfers_TransfersGetAchWithdrawal_GetAchWithdrawal1(t, *sdk, ctx, accountId, withdrawalId)
	testAchTransfers_TransfersCancelAchWithdrawal_CancelAchWithdrawal1(t, *sdk, ctx, accountId, withdrawalId)
}

func testAchTransfers_TransfersCreateAchDeposit_CreateAchDeposit1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string, verifiedBankRelationshipId string) string {

	createAchDepositRequest := components.AchDepositCreate{
		Amount:           components.DecimalCreate{Value: ascendsdk.String("1000.00")},
		BankRelationship: "accounts/" + enrolledAccountId + "/bankRelationship/" + verifiedBankRelationshipId,
		ClientTransferID: uuid.New().String(),
	}
	resp, err := sdk.ACHTransfers.CreateAchDeposit(ctx, enrolledAccountId, createAchDepositRequest)
	require.NoError(t, err)
	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())

	name := resp.AchDeposit.Name
	depositId := strings.Split(*name, "/")[3]
	return depositId
}

func testAchTransfers_TransfersCreateAchWithdrawal_CreateAchWithdrawal1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string, verifiedBankRelationshipId string) string {

	createAchWithdrawalRequest := components.AchWithdrawalCreate{
		Amount:           &components.DecimalCreate{Value: ascendsdk.String("0.01")},
		BankRelationship: "accounts/" + enrolledAccountId + "/bankRelationship/" + verifiedBankRelationshipId,
		ClientTransferID: uuid.New().String(),
	}
	resp, err := sdk.ACHTransfers.CreateAchWithdrawal(ctx, enrolledAccountId, createAchWithdrawalRequest)
	require.NoError(t, err)
	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())

	name := resp.AchWithdrawal.Name
	splitName := strings.Split(*name, "/")
	return splitName[3]
}

func testAchTransfers_TransfersGetAchDeposit_GetAchDeposit1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string, createAchDepositId string) {

	resp, err := sdk.ACHTransfers.GetAchDeposit(ctx, enrolledAccountId, createAchDepositId)
	assert.Nil(t, err)

	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())
}

func testAchTransfers_TransfersCancelAchDeposit_CancelAchDeposit1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string,
	createAchDepositId string) {

	name := "accounts/" + enrolledAccountId + "/achTransfers/" + createAchDepositId
	cancelRequest := components.CancelAchDepositRequestCreate{
		Name: name,
	}
	resp, err := sdk.ACHTransfers.CancelAchDeposit(ctx, enrolledAccountId, createAchDepositId,
		cancelRequest)
	assert.Nil(t, err)

	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())
}

func testAchTransfers_TransfersGetAchWithdrawal_GetAchWithdrawal1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string, createAchWithdrawalId string) {

	resp, err := sdk.ACHTransfers.GetAchWithdrawal(ctx, enrolledAccountId, createAchWithdrawalId)
	assert.Nil(t, err)

	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())
}

func testAchTransfers_TransfersCancelAchWithdrawal_CancelAchWithdrawal1(t *testing.T,
	sdk ascendsdk.SDK, ctx context.Context, enrolledAccountId string, createAchWithdrawalId string) {

	name := "accounts/" + enrolledAccountId + "/achTransfers/" + createAchWithdrawalId
	cancelRequest := components.CancelAchWithdrawalRequestCreate{
		Name: name,
	}
	resp, err := sdk.ACHTransfers.CancelAchWithdrawal(ctx, enrolledAccountId, createAchWithdrawalId,
		cancelRequest)
	assert.Nil(t, err)

	assert.NotNil(t, resp.HTTPMeta)
	assert.NotNil(t, resp.HTTPMeta.Response)
	assert.Equal(t, 200, resp.GetStatus().GetCode())
}
