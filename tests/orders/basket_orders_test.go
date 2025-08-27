package orders

import (
	"context"
	"fmt"
	"os"
	"testing"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createBasketOrder(ctx context.Context, sdk *ascendsdk.SDK, t *testing.T) *string {
	request := components.BasketCreate{
		ClientBasketID:  uuid.New().String(),
		CorrespondentID: os.Getenv("CORRESPONDENT_ID"),
	}

	result, err := sdk.BasketOrders.CreateBasket(ctx, os.Getenv("CORRESPONDENT_ID"), request)
	require.NoError(t, err)

	return result.Basket.BasketID
}

func TestBasketOrders(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	correspondentId := os.Getenv("CORRESPONDENT_ID")
	basketOrderToRemove := uuid.NewString()
	basketOrderId := createBasketOrder(ctx, sdk, t)
	fmt.Println("basketOrderID", basketOrderId)

	basketOrderName := fmt.Sprintf("correspondents/%s/baskets/%s", correspondentId, *basketOrderId)

	t.Run("Basket Orders Orders Create Basket Create Basket1", func(t *testing.T) {
		assert.Greater(t, len(*basketOrderId), 0)
	})

	t.Run("Basket Orders Orders Add Orders Add Orders1", func(t *testing.T) {
		onestr := "1"

		request := components.AddOrdersRequestCreate{
			BasketOrders: []components.BasketOrderCreate{
				{
					AccountID:      "01JHGTEPC6ZTAHCFRH2MD3VJJT",
					AssetType:      components.BasketOrderCreateAssetTypeEquity,
					ClientOrderID:  uuid.NewString(),
					Identifier:     "SBUX",
					IdentifierType: components.BasketOrderCreateIdentifierTypeSymbol,
					OrderType:      components.BasketOrderCreateOrderTypeMarket,
					Quantity:       &components.DecimalCreate{Value: &onestr},
					Side:           components.BasketOrderCreateSideBuy,
					TimeInForce:    components.BasketOrderCreateTimeInForceDay,
				},
				{
					AccountID:      "01JHGTEPC6ZTAHCFRH2MD3VJJT",
					AssetType:      components.BasketOrderCreateAssetTypeEquity,
					ClientOrderID:  basketOrderToRemove,
					Identifier:     "SBUX",
					IdentifierType: components.BasketOrderCreateIdentifierTypeSymbol,
					OrderType:      components.BasketOrderCreateOrderTypeMarket,
					Quantity:       &components.DecimalCreate{Value: &onestr},
					Side:           components.BasketOrderCreateSideBuy,
					TimeInForce:    components.BasketOrderCreateTimeInForceDay,
				},
			},
			Name: basketOrderName,
		}

		result, err := sdk.BasketOrders.AddOrders(ctx, correspondentId, *basketOrderId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Basket Orders Orders Get Basket Get Basket1", func(t *testing.T) {
		result, err := sdk.BasketOrders.GetBasket(ctx, correspondentId, *basketOrderId)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Basket Orders Orders Remove Orders Remove Orders1", func(t *testing.T) {
		request := components.RemoveOrdersRequestCreate{
			ClientOrderIds: []string{basketOrderToRemove},
			Name:           basketOrderName,
		}

		result, err := sdk.BasketOrders.RemoveOrders(ctx, correspondentId, *basketOrderId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Basket Orders Orders Submit Basket Submit Basket1", func(t *testing.T) {
		request := components.SubmitBasketRequestCreate{
			Name: basketOrderName,
		}

		result, err := sdk.BasketOrders.SubmitBasket(ctx, correspondentId, *basketOrderId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Basket Orders List Basket Orders List Basket Orders1", func(t *testing.T) {
		pageSize := 10
		pageToken := ""
		request := operations.BasketOrdersServiceListBasketOrdersRequest{
			CorrespondentID: correspondentId,
			BasketID:        *basketOrderId,
			PageSize:        &pageSize,
			PageToken:       &pageToken,
		}

		result, err := sdk.BasketOrders.ListBasketOrders(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Basket Orders Orders List Compressed Orders List Compressed Orders1", func(t *testing.T) {
		pageSize := 10
		pageToken := ""
		result, err := sdk.BasketOrders.ListCompressedOrders(ctx, correspondentId, *basketOrderId, &pageSize, &pageToken)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
