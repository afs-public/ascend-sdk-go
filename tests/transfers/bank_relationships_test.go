package transfers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FixtureBanks struct {
	t                  *testing.T
	sdk                *ascendsdk.SDK
	ctx                context.Context
	accountId          *string
	bankRelationshipId *string
}

func (f *FixtureBanks) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount()
	return f.accountId
}

func (f *FixtureBanks) createAndEnrollAccount() *string {
	accountId, err := helpers.CreateAccountId(f.sdk, f.ctx)
	require.NoError(f.t, err)
	helpers.Wait()

	agg, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *accountId)
	require.NoError(f.t, err)
	helpers.Wait()

	err = helpers.AffirmAgreements(f.sdk, f.ctx, *accountId, agg)
	require.NoError(f.t, err)
	helpers.Wait()

	return accountId
}

func TestBankRelationships(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &FixtureBanks{
		t:   t,
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
		res, err := sdk.BankRelationships.CreateBankRelationship(ctx, *fixtures.AccountId(), bankRelationshipCreate)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		bankId := strings.Split(*res.BankRelationship.Name, "/")[3]
		fixtures.bankRelationshipId = &bankId
	})

	t.Run("ListBankRelationships", func(t *testing.T) {
		res, err := sdk.BankRelationships.ListBankRelationships(ctx, *fixtures.AccountId(), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("GetBankRelationship", func(t *testing.T) {
		res, err := sdk.BankRelationships.GetBankRelationship(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStatePending)
	})

	t.Run("UpdateBankRelationship", func(t *testing.T) {
		bankRelationshipUpdate := components.BankRelationshipUpdate{
			Nickname: ascendsdk.String("updated nickname"),
		}
		res, err := sdk.BankRelationships.UpdateBankRelationship(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId, bankRelationshipUpdate, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, "updated nickname", *res.BankRelationship.Nickname)
	})

	microDepositAmounts, err := getMicrodepositAmounts(sdk, ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId)
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
				Name: "accounts/" + *fixtures.AccountId() + "/bankRelationships/" + *fixtures.bankRelationshipId,
			}
			_, err = sdk.BankRelationships.VerifyMicroDeposits(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId, verifyMicrodepositRequest)
			require.Error(t, err)
		}
	})

	t.Run("ReissueMicrodeposit", func(t *testing.T) {
		reissueMicrodepositRequest := components.ReissueMicroDepositsRequestCreate{
			Name: "accounts/" + *fixtures.AccountId() + "/bankRelationships/" + *fixtures.bankRelationshipId,
		}
		res, err := sdk.BankRelationships.ReissueMicroDeposits(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId, reissueMicrodepositRequest)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	microDepositAmounts, err = getMicrodepositAmounts(sdk, ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId)
	require.NoError(t, err)

	t.Run("VerifyMicrodeposit", func(t *testing.T) {
		verifyMicrodepositRequest := components.VerifyMicroDepositsRequestCreate{
			Amounts: components.MicroDepositAmountsCreate{
				Amount1: components.DecimalCreate{Value: microDepositAmounts.Amount1.Value},
				Amount2: components.DecimalCreate{Value: microDepositAmounts.Amount2.Value},
			},
			Name: "accounts/" + *fixtures.AccountId() + "/bankRelationships/" + *fixtures.bankRelationshipId,
		}
		res, err := sdk.BankRelationships.VerifyMicroDeposits(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId, verifyMicrodepositRequest)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		time.Sleep(2 * time.Second)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateApproved)
	})

	t.Run("CancelBankRelationship", func(t *testing.T) {
		cancelRelationshipCreate := components.CancelBankRelationshipRequestCreate{
			Comment: string("cancelling bank relationship"),
			Name:    "accounts/" + *fixtures.AccountId() + "/bankRelationships/" + *fixtures.bankRelationshipId,
		}
		res, err := sdk.BankRelationships.CancelBankRelationship(ctx, *fixtures.AccountId(), *fixtures.bankRelationshipId, cancelRelationshipCreate)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateCanceled)
	})
}
