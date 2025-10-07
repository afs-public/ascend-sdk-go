package transfers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (f *Fixtures) WireId(t *testing.T) *string {
	if f.wireId != nil {
		return f.wireId
	}
	wireId, err := CreateWireWithdrawal(t, f.sdk, f.ctx, f.accountId)
	fmt.Println("Wire Withdrawal Id:", wireId)
	require.NoError(f.t, err)
	f.wireId = &wireId
	return &wireId
}

func CreateWireWithdrawal(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, enrolledAccountId string) (string, error) {
	wireWithdrawalCreate := components.WireWithdrawalCreate{
		Amount:           &components.DecimalCreate{Value: ascendsdk.String("100")},
		ClientTransferID: uuid.New().String(),
		Beneficiary: components.WireWithdrawalBeneficiaryCreate{
			Account: enrolledAccountId,
		},
		RecipientBank: components.WireWithdrawalRecipientBankCreate{
			BankID: components.RecipientBankBankIDCreate{
				ID:   *ascendsdk.String("ABNANL2AXXX"),
				Type: components.RecipientBankBankIDCreateType("ABA"),
			},
		},
	}
	res, err := sdk.Wires.CreateWireWithdrawal(ctx, enrolledAccountId, wireWithdrawalCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	wireId := strings.Split(*res.WireWithdrawal.Name, "/")[3]
	if res.HTTPMeta.Response.StatusCode == 200 {
		return wireId, nil
	} else {
		return "", errors.New("Error creating fee")
	}
}

func TestWireWithdrawals(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	fixtures := &Fixtures{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}
	accountId, err := helpers.CreateAccountId(fixtures.sdk, fixtures.ctx)
	if err != nil {
		t.Fatalf("Error creating account: %v", err)
	}
	fixtures.accountId = *accountId

	t.Run("CreateWireWithdrawal", func(t *testing.T) {
		assert.NotNil(t, fixtures.WireId(t))
	})
	t.Run("GetWireWithdrawal", func(t *testing.T) {
		res, err := sdk.Wires.GetWireWithdrawal(ctx, fixtures.accountId, *fixtures.WireId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
	t.Run("CancelWireWithdrawal", func(t *testing.T) {
		request := components.CancelWireWithdrawalRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/wireWithdrawals/" + *fixtures.WireId(t),
		}
		res, err := sdk.Wires.CancelWireWithdrawal(ctx, fixtures.accountId, *fixtures.WireId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res)
	})

	t.Run("GetWireDeposit", func(t *testing.T) {
		res, err := sdk.Wires.GetWireDeposit(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, Wire_Deposit_ID)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
