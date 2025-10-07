package events

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

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

func (f *Fixtures) SubscriptionId() *string {
	s := f.sdk
	ctx := f.ctx

	res, err := s.Subscriber.ListPushSubscriptions(ctx, nil, nil, nil)
	if err != nil {
		return nil
	}
	subscriptions := res.ListPushSubscriptionsResponse.PushSubscriptions
	subscription_id := subscriptions[0].SubscriptionID

	return subscription_id
}

func (f *Fixtures) SubscriberId() *string {
	if f.subscriberId != nil {
		return f.subscriberId
	}

	subscriberId, err := subscriberId(f.sdk, f.ctx, f.CorrespondentId())

	fmt.Println("subscriberId", subscriberId)
	require.NoError(f.t, err)

	f.subscriberId = subscriberId

	return subscriberId
}

func (f *Fixtures) DeliveryId() *string {
	if f.deliveryId != nil {
		return f.deliveryId
	}

	deliveryId, err := deliveryID(f.sdk, f.ctx, f.SubscriptionId())

	fmt.Println("DeliveryID", deliveryId)
	require.NoError(f.t, err)

	f.deliveryId = deliveryId

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

		res, err := sdk.Subscriber.UpdatePushSubscription(ctx, *fixtures.SubscriberId(), pushSubscriptionUpdateCreate, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, updatedEventType, res.PushSubscription.EventTypes[0])
	})

	t.Run("ListPushSubscriptionDeliveries", func(t *testing.T) {
		assert.NotNil(t, fixtures.DeliveryId())
	})

	t.Run("GetPushSubscriptionDelivery", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscriptionDelivery(ctx, *fixtures.SubscriptionId(), *fixtures.DeliveryId())
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("DeletePushSubscription", func(t *testing.T) {
		res, err := sdk.Subscriber.DeletePushSubscription(ctx, *fixtures.SubscriberId())
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
