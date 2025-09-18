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

func Test_Trade_Booking(t *testing.T) {
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

	bookingIds, err := CreateBooking(sdk, ctx, fixture.enrolledWithdrawalAccountId)
	require.NoError(t, err)

	executionID, err := CreateExecution(sdk, ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0])
	require.NoError(t, err)

	rebookExecutionID, err := CreateExecution(sdk, ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0])
	require.NoError(t, err)

	t.Run("test Trade Booking Trade Processing Create Trade Create Trade1", func(t *testing.T) {
		require.Greater(t, len(bookingIds), 0)
		require.NotEmpty(t, bookingIds[0])
		require.NotEmpty(t, bookingIds[1])
	})

	t.Run("test Trade Booking Trade Processing Get Trade Get Trade1", func(t *testing.T) {
		result, err := sdk.TradeBooking.GetTrade(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0])
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Get Trade Get Trade1", func(t *testing.T) {
		result, err := sdk.TradeBooking.GetTrade(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0])
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Create Execution Create Execution1", func(t *testing.T) {
		require.Greater(t, len(executionID), 0)
		require.NotEmpty(t, executionID)
	})

	t.Run("test Trade Booking Trade Processing Get Execution Get Execution1", func(t *testing.T) {
		result, err := sdk.TradeBooking.GetExecution(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], executionID)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Complete Trade Complete Trade1", func(t *testing.T) {

		request := components.CompleteTradeRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/trades/" + bookingIds[0],
		}
		result, err := sdk.TradeBooking.CompleteTrade(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Rebook Execution Rebook Execution1", func(t *testing.T) {

		now := time.Now()
		request := components.RebookExecutionRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/trades/" + bookingIds[0] + "/executions/" + rebookExecutionID,
			Execution: components.ExecutionCreate{
				ExecutionTime: &now,
				ExternalID:    uuid.New().String(),
				Price:         components.DecimalCreate{Value: ascendsdk.String("5")},
				Quantity:      components.DecimalCreate{Value: ascendsdk.String("1")},
			},
		}
		result, err := sdk.TradeBooking.RebookExecution(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], rebookExecutionID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Cancel Execution Cancel Execution1", func(t *testing.T) {

		request := components.CancelExecutionRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/trades/" + bookingIds[0] + "/executions/" + executionID,
		}
		result, err := sdk.TradeBooking.CancelExecution(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], executionID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Rebook Trade Rebook Trade1", func(t *testing.T) {

		now := time.Now()
		assetType := components.TradeCreateAssetTypeEquity
		request := components.RebookTradeRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/trades/" + bookingIds[0],
			Trade: components.TradeCreate{
				AccountID:      fixture.enrolledWithdrawalAccountId,
				BrokerCapacity: components.TradeCreateBrokerCapacityPrincipal,
				ClientOrderID:  bookingIds[1],
				Executions: []components.ExecutionCreate{
					components.ExecutionCreate{
						ExecutionTime: &now,
						ExternalID:    uuid.New().String(),
						Price:         components.DecimalCreate{Value: ascendsdk.String("5")},
						Quantity:      components.DecimalCreate{Value: ascendsdk.String("1")},
					},
				},
				Identifier:        "SBUX",
				IdentifierType:    components.TradeCreateIdentifierTypeSymbol,
				RouteType:         components.RouteTypeQuik,
				Side:              components.TradeCreateSideBuy,
				SourceApplication: "Trading-App",
				AssetType:         assetType,
			},
		}
		result, err := sdk.TradeBooking.RebookTrade(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("test Trade Booking Trade Processing Cancel Trade Cancel Trade1", func(t *testing.T) {

		request := components.CancelTradeRequestCreate{
			Name: "accounts/" + fixture.enrolledWithdrawalAccountId + "/trades/" + bookingIds[0],
		}
		result, err := sdk.TradeBooking.CancelTrade(ctx, fixture.enrolledWithdrawalAccountId, bookingIds[0], request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
