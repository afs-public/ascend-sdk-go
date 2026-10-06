package ach_transfers

import (
	"context"
	"strings"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"testing"

	"github.com/stretchr/testify/assert"
)

type Fixture struct {
	sdk                *ascendsdk.SDK
	ctx                context.Context
	accountId          *string
	bankRelationshipId *string
	depositId          *string
	withdrawalId       *string
}

// AccountId, BankRelationshipId, createAndEnrollAccount, and
// setupBankRelationship all take the *testing.T of whichever subtest is
// calling them -- require.NoError ultimately calls t.FailNow(), which must
// run on the goroutine executing that specific t.Run subtest. Using a
// different subtest's T (or a T stored at Fixture construction) produces
// "subtest may have called FailNow on a parent test" panics whenever setup
// genuinely fails, masking the real error.
func (f *Fixture) AccountId(t *testing.T) *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount(t)
	return f.accountId
}

func (f *Fixture) BankRelationshipId(t *testing.T) *string {
	if f.bankRelationshipId != nil {
		return f.bankRelationshipId
	}

	f.bankRelationshipId = f.setupBankRelationship(t, *f.AccountId(t))
	return f.bankRelationshipId
}

func (f *Fixture) DepositId() *string {
	return f.depositId
}

func (f *Fixture) WithdrawalId() *string {
	return f.withdrawalId
}

func (f *Fixture) createAndEnrollAccount(t *testing.T) *string {
	accountId, err := helpers.CreateEnrolledAccount(f.sdk, f.ctx, t)
	require.NoError(t, err)

	return accountId
}

func (f *Fixture) setupBankRelationship(t *testing.T, accountID string) *string {
	bankRelationshipId, err := helpers.CreateBankRelationship(f.sdk, f.ctx, accountID)
	require.NoError(t, err)

	correctMicroDeposits, err := helpers.GetCorrectMicroDeposits(f.sdk, f.ctx, accountID, *bankRelationshipId)
	require.NoError(t, err)

	err = helpers.VerifyMicroDeposits(f.sdk, f.ctx, accountID, *bankRelationshipId, correctMicroDeposits)
	require.NoError(t, err)

	return bankRelationshipId
}

func TestAchTransfers(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &Fixture{
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateAchDeposit", func(t *testing.T) {
		createAchDepositRequest := components.AchDepositCreate{
			Amount:           components.DecimalCreate{Value: ascendsdk.String("1000.00")},
			BankRelationship: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
			ClientTransferID: uuid.New().String(),
		}
		resp, err := sdk.ACHTransfers.CreateAchDeposit(ctx, *fixtures.AccountId(t), createAchDepositRequest)
		require.NoError(t, err)
		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)

		name := resp.AchDeposit.Name
		fixtures.depositId = &strings.Split(*name, "/")[3]
	})

	t.Run("GetAchDeposit", func(t *testing.T) {
		require.NotNil(t, fixtures.DepositId(), "DepositId is nil and required for this test")

		resp, err := sdk.ACHTransfers.GetAchDeposit(ctx, *fixtures.AccountId(t), *fixtures.DepositId())
		assert.Nil(t, err)

		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelAchDeposit", func(t *testing.T) {
		require.NotNil(t, fixtures.DepositId(), "DepositId is nil and required for this test")

		name := "accounts/" + *fixtures.AccountId(t) + "/achTransfers/" + *fixtures.DepositId()
		cancelRequest := components.CancelAchDepositRequestCreate{
			Name: name,
		}
		resp, err := sdk.ACHTransfers.CancelAchDeposit(ctx, *fixtures.AccountId(t), *fixtures.DepositId(),
			cancelRequest)
		assert.Nil(t, err)

		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)
	})

	t.Run("CreateAchWithdrawal", func(t *testing.T) {
		createAchWithdrawalRequest := components.AchWithdrawalCreate{
			Amount:           &components.DecimalCreate{Value: ascendsdk.String("0.01")},
			BankRelationship: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
			ClientTransferID: uuid.New().String(),
		}
		resp, err := sdk.ACHTransfers.CreateAchWithdrawal(ctx, *fixtures.AccountId(t), createAchWithdrawalRequest)
		require.NoError(t, err)
		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)

		name := resp.AchWithdrawal.Name
		fixtures.withdrawalId = &strings.Split(*name, "/")[3]
	})

	t.Run("GetAchWithdrawal", func(t *testing.T) {
		require.NotNil(t, fixtures.WithdrawalId(), "WithdrawalId is nil and required for this test")

		resp, err := sdk.ACHTransfers.GetAchWithdrawal(ctx, *fixtures.AccountId(t), *fixtures.WithdrawalId())
		assert.Nil(t, err)

		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelAchWithdrawal", func(t *testing.T) {
		require.NotNil(t, fixtures.WithdrawalId(), "WithdrawalId is nil and required for this test")

		name := "accounts/" + *fixtures.AccountId(t) + "/achTransfers/" + *fixtures.WithdrawalId()
		cancelRequest := components.CancelAchWithdrawalRequestCreate{
			Name: name,
		}
		resp, err := sdk.ACHTransfers.CancelAchWithdrawal(ctx, *fixtures.AccountId(t), *fixtures.WithdrawalId(),
			cancelRequest)
		assert.Nil(t, err)

		assert.NotNil(t, resp.HTTPMeta)
		assert.NotNil(t, resp.HTTPMeta.Response)
		assert.Equal(t, 200, resp.HTTPMeta.Response.StatusCode)
	})
}
