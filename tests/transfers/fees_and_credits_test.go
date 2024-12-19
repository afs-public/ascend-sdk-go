package transfers

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeesAndCredits(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	accountId := os.Getenv("account_id")

	ctx := context.Background()

	feeId := testCreateFee(t, sdk, ctx, accountId)
	testGetFee(t, sdk, ctx, accountId, feeId)
	testCancelFee(t, sdk, ctx, accountId, feeId)

	creditId := testCreateCredit(t, sdk, ctx, accountId)
	testGetCredit(t, sdk, ctx, accountId, creditId)
	testCancelCredit(t, sdk, ctx, accountId, creditId)
}

func testCreateFee(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) string {
	feeCreate := components.TransfersFeeCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("10.00"),
		},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Fee charged"),
		Type:             components.TransfersFeeCreateType("MANAGEMENT"),
	}

	res, err := sdk.FeesAndCredits.CreateFee(ctx, accountId, feeCreate)
	require.NoError(t, err)
	assert.NotNil(t, res)
	feeId := strings.Split(*res.TransfersFee.Name, "/")[3]
	return feeId
}

func testGetFee(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, feeId string) {
	if feeId == "" {
		t.Fatalf("expected a valid feeId, but got an empty string")
	}

	res, err := sdk.FeesAndCredits.GetFee(ctx, accountId, feeId)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testCancelFee(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, feeId string) {
	if feeId == "" {
		t.Fatalf("expected a valid feeId, but got an empty string")
	}

	req := components.CancelFeeRequestCreate{
		Name: "accounts/" + accountId + "/feesAndCredits/" + feeId,
	}
	res, err := sdk.FeesAndCredits.CancelFee(ctx, accountId, feeId, req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testCreateCredit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) string {
	creditCreate := components.TransfersCreditCreate{
		Amount: components.DecimalCreate{
			Value: ascendsdk.String("10.00"),
		},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Fee charged"),
		Type:             components.TransfersCreditCreateTypePromotional,
	}

	res, err := sdk.FeesAndCredits.CreateCredit(ctx, accountId, creditCreate)
	require.NoError(t, err)
	assert.NotNil(t, res)
	feeId := strings.Split(*res.TransfersCredit.Name, "/")[3]
	return feeId
}

func testGetCredit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, creditId string) {
	if creditId == "" {
		t.Fatalf("expected a valid creditId, but got an empty string")
	}

	res, err := sdk.FeesAndCredits.GetCredit(ctx, accountId, creditId)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testCancelCredit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, creditId string) {
	if creditId == "" {
		t.Fatalf("expected a valid creditId, but got an empty string")
	}

	req := components.CancelCreditRequestCreate{
		Name: "accounts/" + accountId + "/feesAndCredits/" + creditId,
	}
	res, err := sdk.FeesAndCredits.CancelCredit(ctx, accountId, creditId, req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
