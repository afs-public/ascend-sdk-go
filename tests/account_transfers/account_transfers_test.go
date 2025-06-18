package account_transfers

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/afs-public/ascend-sdk-go/models/operations"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

type Fixture struct {
	t                 *testing.T
	sdk               *ascendsdk.SDK
	ctx               context.Context
	accountId         *string
	accountNumber     *string
	accountTransferId *string
}

func (f *Fixture) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount()
	return f.accountId
}

func (f *Fixture) AccountNumber() *string {
	if f.accountNumber != nil {
		return f.accountNumber
	}

	accountId := f.AccountId()
	require.NotNil(f.t, accountId, "Account ID should not be nil")

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(f.t, err)
	ctx := context.Background()
	account, _ := sdk.AccountCreation.GetAccount(ctx, *accountId, nil)
	require.NotNil(f.t, account, "Account should not be nil")
	f.accountNumber = account.GetAccount().AccountNumber
	return f.accountNumber
}

func (f *Fixture) AccountTransferId() *string {
	if f.accountTransferId != nil {
		return f.accountTransferId
	}
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(f.t, err)
	ctx := context.Background()
	request := components.TransferCreate{
		Assets: []components.AssetCreate{
			{
				Identifier: "USD",
				Position: components.PositionCreate{
					Quantity: components.DecimalCreate{Value: ascendsdk.String("1")},
				},
				Type: components.AssetCreateTypeCurrencyCode,
			},
		},
		Deliverer: components.TransferAccountCreate{
			ExternalAccount: &components.ExternalAccountCreate{
				AccountNumber:     *f.AccountNumber(),
				ParticipantNumber: "158",
			},
		},
	}

	res, err := sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), helpers.WITHDRAWAL_ACCOUNT_ID, request, nil)
	require.NoError(f.t, err)

	name := res.AcatsTransfer.Name
	parts := strings.Split(*name, "/")
	accountTransferId := &parts[len(parts)-1]

	f.accountTransferId = accountTransferId

	return accountTransferId
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

func TestAccountTransfers(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	ctx := context.Background()

	require.NoError(t, err)

	fixtures := &Fixture{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateAccountTransfer", func(t *testing.T) {
		// Fund Account
		creditCreate := components.TransfersCreditCreate{
			Amount: components.DecimalCreate{
				Value: ascendsdk.String("1000.00"),
			},
			ClientTransferID: uuid.New().String(),
			Description:      ascendsdk.String("Credit awarded"),
			Type:             components.TransfersCreditCreateTypePromotional,
		}

		_, err := sdk.FeesAndCredits.CreateCredit(ctx, *fixtures.AccountId(), creditCreate)
		require.NoError(t, err)

		transferID := fixtures.AccountTransferId()
		assert.NotNil(t, transferID, "Account transfer ID should not be nil")
	})

	t.Run("ListAccountTransfers", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountId(), "accountId is required to list account transfers")

		request := operations.AccountTransfersListTransfersRequest{
			CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
			AccountID:       *fixtures.AccountId(),
		}

		res, err := sdk.AccountTransfers.ListTransfers(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, res.ListTransfersResponse)
	})

	t.Run("RejectTransfer", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountTransferId(), "accountTransferId is required to reject account transfer")

		request := components.RejectTransferRequestCreate{
			Name: "correspondents/" + os.Getenv("CORRESPONDENT_ID") + "/accounts/" + *fixtures.AccountId() + "/transfers/" + *fixtures.AccountTransferId(),
		}
		res, err := sdk.AccountTransfers.RejectTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(), *fixtures.AccountTransferId(), request)

		require.NoError(t, err)
		assert.NotNil(t, res.RejectTransferResponse)
	})

	t.Run("AcceptTransfer", func(t *testing.T) {
		request := components.TransferCreate{
			Assets: []components.AssetCreate{
				{
					Identifier: "USD",
					Position: components.PositionCreate{
						Quantity: components.DecimalCreate{Value: ascendsdk.String("1")},
					},
					Type: components.AssetCreateTypeCurrencyCode,
				},
			},
			Deliverer: components.TransferAccountCreate{
				ExternalAccount: &components.ExternalAccountCreate{
					AccountNumber:     *fixtures.AccountNumber(),
					ParticipantNumber: "158",
				},
			},
		}

		res, err := sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), helpers.WITHDRAWAL_ACCOUNT_ID, request, nil)
		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)

		transferID := res.AcatsTransfer.Name
		parts := strings.Split(*transferID, "/")
		accountTransferId := &parts[len(parts)-1]

		require.NotNil(t, accountTransferId, "accountTransferId should not be nil")

		acceptRequest := components.AcceptTransferRequestCreate{
			Name: "correspondents/" + os.Getenv("CORRESPONDENT_ID") + "/accounts/" + *fixtures.AccountId() + "/transfers/" + *accountTransferId,
		}

		acceptRes, err := sdk.AccountTransfers.AcceptTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(), *accountTransferId, acceptRequest)
		require.NoError(t, err)
		assert.NotNil(t, acceptRes.AcceptTransferResponse)
	})

	t.Run("GetAccountTransfer", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountTransferId(), "accountTransferId is required to get account transfer")

		res, err := sdk.AccountTransfers.GetTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(), *fixtures.AccountTransferId())

		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)
	})
}
