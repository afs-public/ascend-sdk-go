package events

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"ascend-sdk/models/operations"
	"ascend-sdk/tests/helpers"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Fixtures struct {
	t                 *testing.T
	sdk               *ascendsdk.SDK
	ctx               context.Context
	subscriberId      *string
	fixedSubscriberId *string
	listPushSubId     *string
	deliveryId        *string
	correspondentId   *string
}

func (f *Fixtures) CorrespondentId() (correspondentId *string) {
	correspondentId = ascendsdk.String(os.Getenv("CORRESPONDENT_ID"))
	require.NotNil(f.t, correspondentId, "CORRESPONDENT_ID is required and must be set as an environment variable")
	return
}

func (f *Fixtures) FixedSubscriberId() *string {
	return ascendsdk.String("01JBF2P1DGVY9XGHF7780AFCPT")
}

func (f *Fixtures) SubscriberId() *string {
	if f.subscriberId != nil {
		return f.subscriberId
	}

	subscriberId, err := subscriberId(f.sdk, f.ctx, f.CorrespondentId())

	fmt.Println("subscriberId", subscriberId)
	require.NoError(f.t, err)

	f.subscriberId = subscriberId

	time.Sleep(5 * time.Second)

	return subscriberId
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &Fixtures{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreatePushSubscription", func(t *testing.T) {
		assert.NotNil(t, fixtures.SubscriberId())
	})

	t.Run("GetPushSubscription", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscription(ctx, *fixtures.SubscriberId())
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("ListPushSubscriptions", func(t *testing.T) {
		res, err := sdk.Subscriber.ListPushSubscriptions(ctx, nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res.ListPushSubscriptionsResponse.PushSubscriptions)

		fixtures.listPushSubId = res.ListPushSubscriptionsResponse.PushSubscriptions[0].SubscriptionID
	})

	t.Run("UpdatePushSubscription", func(t *testing.T) {
		var updatedEventType = "position.v2.updated"
		pushSubscriptionUpdateCreate := components.PushSubscriptionUpdate{
			EventTypes: []string{
				updatedEventType,
			},
		}

		var status *operations.SubscriberGetPushSubscriptionResponse
		for i := 0; i < 10; i++ {
			status, err = sdk.Subscriber.GetPushSubscription(ctx, *fixtures.SubscriberId())
			require.NoError(t, err)

			if *status.PushSubscription.State == components.StateUpdating || *status.PushSubscription.State == components.StateActive {
				break
			}

			time.Sleep(2 * time.Second)
		}

		require.Contains(t,
			[]components.State{components.StateUpdating, components.StateActive},
			*status.PushSubscription.State,
			"Push subscription must be in the state of UPDATED or ACTIVE to receive an update, otherwise we will get an error",
		)

		res, err := sdk.Subscriber.UpdatePushSubscription(ctx, *fixtures.SubscriberId(), pushSubscriptionUpdateCreate, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, updatedEventType, res.PushSubscription.EventTypes[0])
	})

	t.Run("ListPushSubscriptionDeliveries", func(t *testing.T) {
		res, err := sdk.Subscriber.ListPushSubscriptionDeliveries(ctx, *fixtures.FixedSubscriberId(), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotEmpty(t, res.ListPushSubscriptionDeliveriesResponse.PushSubscriptionDeliveries)

		fixtures.deliveryId = res.ListPushSubscriptionDeliveriesResponse.PushSubscriptionDeliveries[0].DeliveryID
	})

	t.Run("GetPushSubscriptionDelivery", func(t *testing.T) {
		require.NotNil(t, fixtures.deliveryId, "deliveryId is nil and is required for GetPushSubscriptionDelivery")

		res, err := sdk.Subscriber.GetPushSubscriptionDelivery(ctx, *fixtures.FixedSubscriberId(), *fixtures.deliveryId)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("DeletePushSubscription", func(t *testing.T) {
		require.NotNil(t, fixtures.listPushSubId, "listPushSubId is nil and is required for DeletePushSubscription")

		res, err := sdk.Subscriber.DeletePushSubscription(ctx, *fixtures.listPushSubId)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
