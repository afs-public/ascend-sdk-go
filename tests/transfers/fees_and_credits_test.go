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

func (f *Fixtures) FeeId(t *testing.T) *string {
	if f.feeId != nil {
		return f.feeId
	}
	feeId, err := CreateFee(t, f.sdk, f.ctx, f.accountId)
	fmt.Println("feeId:", feeId)
	require.NoError(f.t, err)
	f.feeId = &feeId
	return &feeId
}

func (f *Fixtures) CreditId(t *testing.T) *string {
	if f.creditId != nil {
		return f.creditId
	}
	creditId, err := CreateCredit(t, f.sdk, f.ctx, f.accountId)
	fmt.Println("creditId: ", creditId)
	require.NoError(f.t, err)
	f.creditId = &creditId
	return &creditId
}

func CreateFee(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, enrolledAccountId string) (string, error) {
	feeCreate := components.TransfersFeeCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("10.00"),
		},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Fee charged"),
		Type:             components.TransfersFeeCreateType("MANAGEMENT"),
	}
	res, err := sdk.FeesAndCredits.CreateFee(ctx, enrolledAccountId, feeCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	feeId := strings.Split(*res.TransfersFee.Name, "/")[3]
	if res.HTTPMeta.Response.StatusCode == 200 {
		return feeId, nil
	} else {
		return "", errors.New("Error creating fee")
	}
}

func CreateCredit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, enrolledAccountId string) (string, error) {
	creditCreate := components.TransfersCreditCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("10.00"),
		},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Credit awarded"),
		Type:             components.TransfersCreditCreateTypePromotional,
	}
	res, err := sdk.FeesAndCredits.CreateCredit(ctx, enrolledAccountId, creditCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	creditId := strings.Split(*res.TransfersCredit.Name, "/")[3]
	if res.HTTPMeta.Response.StatusCode == 200 {
		return creditId, nil
	} else {
		return "", errors.New("Error creating credit")
	}
}

func TestFeesAndCredits(t *testing.T) {
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
	print("accountId:", fixtures.accountId)

	t.Run("CreateFee", func(t *testing.T) {
		assert.NotNil(t, fixtures.FeeId(t))
	})
	t.Run("GetFee", func(t *testing.T) {
		res, err := sdk.FeesAndCredits.GetFee(ctx, fixtures.accountId, *fixtures.FeeId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
	t.Run("CancelFee", func(t *testing.T) {
		request := components.CancelFeeRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/feesAndCredits/" + *fixtures.FeeId(t),
		}
		res, err := sdk.FeesAndCredits.CancelFee(ctx, fixtures.accountId, *fixtures.FeeId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res)
	})

	t.Run("CreateCredit", func(t *testing.T) {
		assert.NotNil(t, fixtures.CreditId(t))
	})
	t.Run("GetCredit", func(t *testing.T) {
		res, err := sdk.FeesAndCredits.GetCredit(ctx, fixtures.accountId, *fixtures.CreditId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
	t.Run("CancelCredit", func(t *testing.T) {
		request := components.CancelCreditRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/feesAndCredits/" + *fixtures.CreditId(t),
		}
		res, err := sdk.FeesAndCredits.CancelCredit(ctx, fixtures.accountId, *fixtures.CreditId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res)
	})
}
