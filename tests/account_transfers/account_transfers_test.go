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
	sdk               *ascendsdk.SDK
	ctx               context.Context
	accountId         *string
	accountNumber     *string
	accountTransferId *string
}

// Fixture methods take the *testing.T of whichever subtest is calling them --
// require.NoError ultimately calls t.FailNow(), which must run on the
// goroutine executing that specific t.Run subtest. Using a T stored at
// Fixture construction produces "subtest may have called FailNow on a parent
// test" panics whenever setup genuinely fails, masking the real error.
func (f *Fixture) AccountId(t *testing.T) *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId = f.createAndEnrollAccount(t)
	return f.accountId
}

func (f *Fixture) AccountNumber(t *testing.T) *string {
	if f.accountNumber != nil {
		return f.accountNumber
	}

	accountId := f.AccountId(t)
	require.NotNil(t, accountId, "Account ID should not be nil")

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	ctx := context.Background()
	account, _ := sdk.AccountCreation.GetAccount(ctx, *accountId, nil)
	require.NotNil(t, account, "Account should not be nil")
	f.accountNumber = account.GetAccount().AccountNumber
	return f.accountNumber
}

func (f *Fixture) AccountTransferId(t *testing.T) *string {
	if f.accountTransferId != nil {
		return f.accountTransferId
	}
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
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
				AccountNumber:     *f.AccountNumber(t),
				ParticipantNumber: "158",
			},
		},
	}

	// The funding credit created just before this posts asynchronously;
	// until it lands the API rejects the transfer for insufficient cash.
	var res *operations.AccountTransfersCreateTransferResponse
	err = helpers.RetryOnTransientError(func() error {
		var opErr error
		res, opErr = sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), helpers.WITHDRAWAL_ACCOUNT_ID, request, nil)
		return opErr
	})
	require.NoError(t, err)

	name := res.AcatsTransfer.Name
	parts := strings.Split(*name, "/")
	accountTransferId := &parts[len(parts)-1]

	f.accountTransferId = accountTransferId

	return accountTransferId
}

func (f *Fixture) createAndEnrollAccount(t *testing.T) *string {
	accountId, err := helpers.CreateEnrolledAccount(f.sdk, f.ctx, t)
	require.NoError(t, err)

	return accountId
}

func TestAccountTransfers(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	ctx := context.Background()

	require.NoError(t, err)

	fixtures := &Fixture{
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

		_, err := sdk.FeesAndCredits.CreateCredit(ctx, *fixtures.AccountId(t), creditCreate)
		require.NoError(t, err)

		transferID := fixtures.AccountTransferId(t)
		assert.NotNil(t, transferID, "Account transfer ID should not be nil")
	})

	t.Run("ListAccountTransfers", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountId(t), "accountId is required to list account transfers")

		request := operations.AccountTransfersListTransfersRequest{
			CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
			AccountID:       *fixtures.AccountId(t),
		}

		res, err := sdk.AccountTransfers.ListTransfers(ctx, request)

		require.NoError(t, err)
		assert.NotNil(t, res.ListTransfersResponse)
	})

	t.Run("RejectTransfer", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountTransferId(t), "accountTransferId is required to reject account transfer")

		request := components.RejectTransferRequestCreate{
			Name: "correspondents/" + os.Getenv("CORRESPONDENT_ID") + "/accounts/" + *fixtures.AccountId(t) + "/transfers/" + *fixtures.AccountTransferId(t),
		}
		res, err := sdk.AccountTransfers.RejectTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(t), *fixtures.AccountTransferId(t), request)

		require.NoError(t, err)
		assert.NotNil(t, res.RejectTransferResponse)
	})

	t.Run("AcceptTransfer", func(t *testing.T) {
		// Use a dedicated account: rejecting the earlier transfer restricts
		// its deliverer account (ACAT_PARTIAL_OUTBOUND entitlement) for an
		// unbounded window, so a second transfer on the same account is
		// rejected as "Account not entitled".
		acceptAccountId := fixtures.createAndEnrollAccount(t)
		account, _ := sdk.AccountCreation.GetAccount(ctx, *acceptAccountId, nil)
		require.NotNil(t, account, "Account should not be nil")
		acceptAccountNumber := account.GetAccount().AccountNumber

		creditCreate := components.TransfersCreditCreate{
			Amount: components.DecimalCreate{
				Value: ascendsdk.String("1000.00"),
			},
			ClientTransferID: uuid.New().String(),
			Description:      ascendsdk.String("Credit awarded"),
			Type:             components.TransfersCreditCreateTypePromotional,
		}
		_, err := sdk.FeesAndCredits.CreateCredit(ctx, *acceptAccountId, creditCreate)
		require.NoError(t, err)

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
					AccountNumber:     *acceptAccountNumber,
					ParticipantNumber: "158",
				},
			},
		}

		var res *operations.AccountTransfersCreateTransferResponse
		err = helpers.RetryOnTransientError(func() error {
			var opErr error
			res, opErr = sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), helpers.WITHDRAWAL_ACCOUNT_ID, request, nil)
			return opErr
		})
		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)

		transferID := res.AcatsTransfer.Name
		parts := strings.Split(*transferID, "/")
		accountTransferId := &parts[len(parts)-1]

		require.NotNil(t, accountTransferId, "accountTransferId should not be nil")

		acceptRequest := components.AcceptTransferRequestCreate{
			Name: "correspondents/" + os.Getenv("CORRESPONDENT_ID") + "/accounts/" + *acceptAccountId + "/transfers/" + *accountTransferId,
		}

		acceptRes, err := sdk.AccountTransfers.AcceptTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *acceptAccountId, *accountTransferId, acceptRequest)
		require.NoError(t, err)
		assert.NotNil(t, acceptRes.AcceptTransferResponse)
	})

	t.Run("GetAccountTransfer", func(t *testing.T) {
		require.NotNil(t, fixtures.AccountTransferId(t), "accountTransferId is required to get account transfer")

		res, err := sdk.AccountTransfers.GetTransfer(ctx, os.Getenv("CORRESPONDENT_ID"), *fixtures.AccountId(t), *fixtures.AccountTransferId(t))

		require.NoError(t, err)
		assert.NotNil(t, res.AcatsTransfer)
	})
}
