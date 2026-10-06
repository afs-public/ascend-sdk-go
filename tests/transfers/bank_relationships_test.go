package transfers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FixtureBanks struct {
	sdk                *ascendsdk.SDK
	ctx                context.Context
	lnpID              *string
	accountId          *string
	bankRelationshipId *string
	reuseAccountID     *string
}

func (f *FixtureBanks) LNPId() *string {
	if f.lnpID != nil {
		return f.lnpID
	}

	f.lnpID, _ = helpers.CreateLegalNaturalPersonId(f.sdk, f.ctx)
	return f.lnpID
}

// AccountId, ReuseAccountId, and createAndEnrollAccount all take the
// *testing.T of whichever subtest is calling them (not f.t, the parent
// test's T stored at Fixture construction) -- require.NoError ultimately
// calls t.FailNow(), which must run on the goroutine executing that specific
// t.Run subtest. Using the parent's T from inside a subtest's goroutine
// produces "subtest may have called FailNow on a parent test" panics
// whenever setup genuinely fails, masking the real error.
func (f *FixtureBanks) AccountId(t *testing.T) *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount(t)
	return f.accountId
}

func (f *FixtureBanks) ReuseAccountId(t *testing.T) *string {
	if f.reuseAccountID != nil {
		return f.reuseAccountID
	}

	f.reuseAccountID = f.createAndEnrollAccount(t)
	return f.reuseAccountID
}

func (f *FixtureBanks) createAndEnrollAccount(t *testing.T) *string {
	accountId, err := helpers.CreateAccountIdWithLNP(f.sdk, f.ctx, f.LNPId())
	require.NoError(t, err)

	agg, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *accountId)
	require.NoError(t, err)

	err = helpers.AffirmAgreements(f.sdk, f.ctx, *accountId, agg)
	require.NoError(t, err)

	require.NoError(t, helpers.WaitForAccountOpen(f.sdk, f.ctx, *accountId))

	return accountId
}

// BankRelationshipId fails the calling subtest cleanly (instead of panicking
// on a nil dereference) if CreateBankRelationship hasn't run yet or failed
// to set it.
func (f *FixtureBanks) BankRelationshipId(t *testing.T) *string {
	require.NotNil(t, f.bankRelationshipId, "bank relationship was not created")
	return f.bankRelationshipId
}

func TestBankRelationships(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &FixtureBanks{
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateBankRelationship", func(t *testing.T) {
		bankRelationshipCreate := components.BankRelationshipCreate{
			BankAccount: &components.BankAccountCreate{
				AccountNumber: strconv.Itoa(rand.Intn(90000000) + 10000000),
				Owner:         "TEST123",
				RoutingNumber: "112203216",
				Type:          components.BankAccountCreateTypeChecking,
			},
			Nickname:           "TEST ACCOUNT",
			VerificationMethod: components.VerificationMethodMicroDeposit,
		}
		res, err := sdk.BankRelationships.CreateBankRelationship(ctx, *fixtures.AccountId(t), bankRelationshipCreate)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		bankId := strings.Split(*res.BankRelationship.Name, "/")[3]
		fixtures.bankRelationshipId = &bankId
	})

	t.Run("ListBankRelationships", func(t *testing.T) {
		res, err := sdk.BankRelationships.ListBankRelationships(ctx, *fixtures.AccountId(t), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("GetBankRelationship", func(t *testing.T) {
		res, err := sdk.BankRelationships.GetBankRelationship(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStatePending)
	})

	t.Run("UpdateBankRelationship", func(t *testing.T) {
		bankRelationshipUpdate := components.BankRelationshipUpdate{
			Nickname: ascendsdk.String("updated nickname"),
		}
		res, err := sdk.BankRelationships.UpdateBankRelationship(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t), bankRelationshipUpdate, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, "updated nickname", *res.BankRelationship.Nickname)
	})

	microDepositAmounts, err := getMicrodepositAmounts(sdk, ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t))
	require.NoError(t, err)

	t.Run("FailMicrodepositVerification", func(t *testing.T) {
		maxAttempts := 3
		for i := 0; i < maxAttempts; i++ {
			amount1, _ := strconv.ParseFloat(*microDepositAmounts.Amount1.Value, 64)
			amount2, _ := strconv.ParseFloat(*microDepositAmounts.Amount2.Value, 64)
			verifyMicrodepositRequest := components.VerifyMicroDepositsRequestCreate{
				Amounts: components.MicroDepositAmountsCreate{
					Amount1: components.DecimalCreate{Value: ascendsdk.String(fmt.Sprintf("%.2f", amount1+0.01))},
					Amount2: components.DecimalCreate{Value: ascendsdk.String(fmt.Sprintf("%.2f", amount2+0.01))},
				},
				Name: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
			}
			_, err = sdk.BankRelationships.VerifyMicroDeposits(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t), verifyMicrodepositRequest)
			require.Error(t, err)
		}
	})

	t.Run("ReissueMicrodeposit", func(t *testing.T) {
		reissueMicrodepositRequest := components.ReissueMicroDepositsRequestCreate{
			Name: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
		}
		res, err := sdk.BankRelationships.ReissueMicroDeposits(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t), reissueMicrodepositRequest)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	microDepositAmounts, err = getMicrodepositAmounts(sdk, ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t))
	require.NoError(t, err)

	t.Run("VerifyMicrodeposit", func(t *testing.T) {
		verifyMicrodepositRequest := components.VerifyMicroDepositsRequestCreate{
			Amounts: components.MicroDepositAmountsCreate{
				Amount1: components.DecimalCreate{Value: microDepositAmounts.Amount1.Value},
				Amount2: components.DecimalCreate{Value: microDepositAmounts.Amount2.Value},
			},
			Name: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
		}
		res, err := sdk.BankRelationships.VerifyMicroDeposits(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t), verifyMicrodepositRequest)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateApproved)
	})

	t.Run("ReuseBankRelationship", func(t *testing.T) {
		reuseBankRelationShipRequest := components.ReuseBankRelationshipRequestCreate{
			Parent:                 "accounts/" + *fixtures.AccountId(t),
			SourceBankRelationship: "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
		}
		var res *operations.BankRelationshipsReuseBankRelationshipResponse
		err := helpers.RetryOnTransientError(func() error {
			var opErr error
			res, opErr = sdk.BankRelationships.ReuseBankRelationship(ctx, *fixtures.ReuseAccountId(t), reuseBankRelationShipRequest)
			return opErr
		})
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStatePending)
	})

	t.Run("CancelBankRelationship", func(t *testing.T) {
		cancelRelationshipCreate := components.CancelBankRelationshipRequestCreate{
			Comment: string("cancelling bank relationship"),
			Name:    "accounts/" + *fixtures.AccountId(t) + "/bankRelationships/" + *fixtures.BankRelationshipId(t),
		}
		res, err := sdk.BankRelationships.CancelBankRelationship(ctx, *fixtures.AccountId(t), *fixtures.BankRelationshipId(t), cancelRelationshipCreate)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateCanceled)
	})
}
