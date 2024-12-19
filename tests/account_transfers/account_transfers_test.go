package account_transfers

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

type Fixture struct {
	t                 *testing.T
	sdk               *ascendsdk.SDK
	ctx               context.Context
	accountId         *string
	accountTransferId *string
}

func (f *Fixture) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount()
	return f.accountId
}

func (f *Fixture) AccountTransferId() *string {
	return f.accountTransferId
}

func (f *Fixture) createAndEnrollAccount() *string {
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

func TestAccoutTransfers_getAccountTransfer(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	ctx := context.Background()

	require.NoError(t, err)

	fixtures := &Fixture{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateAccountTransfer", func(t *testing.T) {
		request := components.TransferCreate{
			Deliverer: components.TransferAccountCreate{
				ExternalAccount: &components.ExternalAccountCreate{
					AccountNumber:     "1234567890",
					ParticipantNumber: "987",
				},
			},
		}

		res, err := sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(), request, nil)
		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)

		name := res.AcatsTransfer.Name
		parts := strings.Split(*name, "/")
		accountTransferId := &parts[len(parts)-1]

		fixtures.accountTransferId = accountTransferId
	})

	t.Run("GetAccountTransfer", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountTransferId(), "accountTransferId is required to get account transfer")

		res, err := sdk.AccountTransfers.GetTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(), *fixtures.AccountTransferId())

		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)
	})
}
