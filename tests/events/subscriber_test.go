package events

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"

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
	return ascendsdk.String("01JJYZ16TVYZM6A6BDJ8RJRMTQ")
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

func (f *Fixtures) DeliveryId() *string {
	if f.deliveryId != nil {
		return f.deliveryId
	}

	deliveryId, err := deliveryID(f.sdk, f.ctx, f.FixedSubscriberId())

	fmt.Println("DeliveryID", deliveryId)
	require.NoError(f.t, err)

	f.deliveryId = deliveryId

	time.Sleep(5 * time.Second)

	return deliveryId
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
	})

	t.Run("UpdatePushSubscription", func(t *testing.T) {
		var updatedEventType = "position.v2.updated"
		pushSubscriptionUpdateCreate := components.PushSubscriptionUpdate{
			EventTypes: []string{
				updatedEventType,
			},
		}

		var status *operations.SubscriberGetPushSubscriptionResponse
		for i := 0; i < 12; i++ {
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
		assert.NotNil(t, fixtures.DeliveryId())
	})

	t.Run("GetPushSubscriptionDelivery", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscriptionDelivery(ctx, *fixtures.FixedSubscriberId(), *fixtures.DeliveryId())
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("DeletePushSubscription", func(t *testing.T) {
		time.Sleep(5 * time.Second)
		res, err := sdk.Subscriber.DeletePushSubscription(ctx, *fixtures.SubscriberId())
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
