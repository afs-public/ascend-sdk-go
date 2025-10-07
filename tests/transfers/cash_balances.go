package transfers

import (
	"context"
	"fmt"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FixtureCashBalances struct {
	t         *testing.T
	sdk       *ascendsdk.SDK
	ctx       context.Context
	accountId *string
}

func (f *FixtureCashBalances) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}

	f.accountId, _ = helpers.CreateAccountId(f.sdk, f.ctx)
	return f.accountId
}

func TestCashBalances(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &FixtureCashBalances{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	fmt.Println("Cash Balance:", *fixtures.AccountId())

	t.Run("GetCashBalance", func(t *testing.T) {
		res, err := sdk.CashBalances.CalculateCashBalance(ctx, *fixtures.AccountId(), nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
