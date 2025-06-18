package trade_processing

import (
	"context"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Trade_Allocation(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixture := &Fixtures{
		t:                           t,
		sdk:                         sdk,
		ctx:                         ctx,
		deceasedAccountId:           "01JHK07CRQ9X8P5XE9JWG4PFSP",
		enrolledWithdrawalAccountId: "01JHGTEPC6ZTAHCFRH2MD3VJJT",
	}

	tradeAllocationID, err := CreateTradeAllocation(sdk, ctx, fixture.enrolledWithdrawalAccountId, fixture.deceasedAccountId)
	require.NoError(t, err)

	rebookTradeAllocationID, err := CreateTradeAllocation(sdk, ctx, fixture.enrolledWithdrawalAccountId, fixture.deceasedAccountId)
	require.NoError(t, err)

	t.Run("test Trade Allocation Trade Processing Create Trade Allocation Create TradeAllocation1", func(t *testing.T) {
		require.Greater(t, len(tradeAllocationID), 0)
		require.NotEmpty(t, tradeAllocationID)
	})

	t.Run("test Trade Allocation Trade Processing Get Trade Allocation Get TradeAllocation1", func(t *testing.T) {
		result, err := sdk.TradeAllocation.GetTradeAllocation(ctx, fixture.enrolledWithdrawalAccountId, tradeAllocationID)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Allocation Trade Processing Rebook Trade Allocation Rebook TradeAllocation1", func(t *testing.T) {
		now := time.Now()
		assetType := components.TradeAllocationCreateAssetTypeEquity
		request := components.RebookTradeAllocationRequestCreate{
			Name:      "accounts/" + fixture.enrolledWithdrawalAccountId + "/tradeAllocations/" + rebookTradeAllocationID,
			RequestID: uuid.New().String(),
			TradeAllocation: components.TradeAllocationCreate{
				BrokerCapacity:    components.TradeAllocationCreateBrokerCapacityPrincipal,
				ExecutionTime:     &now,
				FromAccountID:     fixture.deceasedAccountId,
				Identifier:        "SBUX",
				IdentifierType:    components.TradeAllocationCreateIdentifierTypeSymbol,
				AssetType:         &assetType,
				Price:             components.DecimalCreate{Value: ascendsdk.String("5")},
				Quantity:          components.DecimalCreate{Value: ascendsdk.String("1")},
				SourceApplication: "Trading-App",
				ToAccountID:       fixture.enrolledWithdrawalAccountId,
				ToSide:            components.ToSideBuy,
			},
		}
		result, err := sdk.TradeAllocation.RebookTradeAllocation(ctx, fixture.enrolledWithdrawalAccountId, rebookTradeAllocationID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Allocation Trade Processing Cancel Trade Allocation Cancel TradeAllocation1", func(t *testing.T) {
		request := components.CancelTradeAllocationRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/tradeAllocations/" + tradeAllocationID,
		}
		result, err := sdk.TradeAllocation.CancelTradeAllocation(ctx, fixture.enrolledWithdrawalAccountId, tradeAllocationID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
