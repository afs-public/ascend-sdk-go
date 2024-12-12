package transfers

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/tests/helpers"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getMicrodepositAmounts(sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) (components.MicroDepositAmounts, error) {
	res, err := sdk.TestSimulation.GetMicroDepositAmounts(ctx, accountId, bankRelationshipId)
	if err != nil {
		return components.MicroDepositAmounts{}, err
	}
	return *res.MicroDepositAmounts, nil
}

func TestBankRelationships(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	accountId := "01JC2TNAZ51HHMT07HSXGW68WM"

	bankRelationshipId := testCreateBankRelationship(t, sdk, ctx, accountId)
	testListBankRelationships(t, sdk, ctx, accountId)
	testGetBankRelationship(t, sdk, ctx, accountId, bankRelationshipId)
	testUpdateBankRelationship(t, sdk, ctx, accountId, bankRelationshipId)
	microdepositAmounts, err := getMicrodepositAmounts(sdk, ctx, accountId, bankRelationshipId)
	require.NoError(t, err)
	testFailMicrodepositVerification(t, sdk, ctx, accountId, bankRelationshipId, microdepositAmounts.Amount1.Value, microdepositAmounts.Amount2.Value)
	testReissueMicrodeposit(t, sdk, ctx, accountId, bankRelationshipId)
	microdepositAmounts, err = getMicrodepositAmounts(sdk, ctx, accountId, bankRelationshipId)
	require.NoError(t, err)
	testVerifyMicrodeposit(t, sdk, ctx, accountId, bankRelationshipId, microdepositAmounts.Amount1.Value, microdepositAmounts.Amount2.Value)
	testCancelBankRelationship(t, sdk, ctx, accountId, bankRelationshipId)

}

func testCreateBankRelationship(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) string {
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
	res, err := sdk.BankRelationships.CreateBankRelationship(ctx, accountId, bankRelationshipCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	bankId := strings.Split(*res.BankRelationship.Name, "/")[3]
	return bankId
}

func testListBankRelationships(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	res, err := sdk.BankRelationships.ListBankRelationships(ctx, accountId, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testGetBankRelationship(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	res, err := sdk.BankRelationships.GetBankRelationship(ctx, accountId, bankRelationshipId)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStatePending)
}

func testUpdateBankRelationship(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	bankRelationshipUpdate := components.BankRelationshipUpdate{
		Nickname: ascendsdk.String("updated nickname"),
	}
	res, err := sdk.BankRelationships.UpdateBankRelationship(ctx, accountId, bankRelationshipId, bankRelationshipUpdate, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, "updated nickname", *res.BankRelationship.Nickname)
}

func testFailMicrodepositVerification(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string, amount1 *string, amount2 *string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	amount1Float, err := strconv.ParseFloat(*amount1, 64)
	require.NoError(t, err)
	amount2Float, err := strconv.ParseFloat(*amount2, 64)
	require.NoError(t, err)

	maxAttempts := 3
	for i := 0; i < maxAttempts; i++ {
		verifyMicrodepositRequest := components.VerifyMicroDepositsRequestCreate{
			Amounts: components.MicroDepositAmountsCreate{
				Amount1: components.DecimalCreate{Value: ascendsdk.String(fmt.Sprintf("%.2f", amount1Float+0.01))},
				Amount2: components.DecimalCreate{Value: ascendsdk.String(fmt.Sprintf("%.2f", amount2Float+0.01))},
			},
			Name: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
		}
		_, err = sdk.BankRelationships.VerifyMicroDeposits(ctx, accountId, bankRelationshipId, verifyMicrodepositRequest)
		require.Error(t, err)
	}
}

func testReissueMicrodeposit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	reissueMicrodepositRequest := components.ReissueMicroDepositsRequestCreate{
		Name: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
	}
	res, err := sdk.BankRelationships.ReissueMicroDeposits(ctx, accountId, bankRelationshipId, reissueMicrodepositRequest)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testVerifyMicrodeposit(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string, amount1 *string, amount2 *string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	verifyMicrodepositRequest := components.VerifyMicroDepositsRequestCreate{
		Amounts: components.MicroDepositAmountsCreate{
			Amount1: components.DecimalCreate{Value: amount1},
			Amount2: components.DecimalCreate{Value: amount2},
		},
		Name: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
	}
	res, err := sdk.BankRelationships.VerifyMicroDeposits(ctx, accountId, bankRelationshipId, verifyMicrodepositRequest)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	time.Sleep(2 * time.Second)
	assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateApproved)
}

func testCancelBankRelationship(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) {
	if accountId == "" {
		t.Fatalf("expected a valid accountId, but got an empty string")
	}
	if bankRelationshipId == "" {
		t.Fatalf("expected a valid bankRelationshipId, but got an empty string")
	}
	cancelRelationshipCreate := components.CancelBankRelationshipRequestCreate{
		Comment: string("cancelling bank relationship"),
		Name:    "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
	}
	res, err := sdk.BankRelationships.CancelBankRelationship(ctx, accountId, bankRelationshipId, cancelRelationshipCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, *res.BankRelationship.State.State, components.BankRelationshipStateStateCanceled)
}
