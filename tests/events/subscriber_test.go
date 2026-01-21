package events

import (
	"context"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Fixtures struct {
	t                *testing.T
	sdk              *ascendsdk.SDK
	ctx              context.Context
	subscriberId     *string
	testSubscriberId *string
	deliveryId       *string
	correspondentId  *string
}

func (f *Fixtures) CorrespondentId() *string {
	if f.correspondentId != nil {
		return f.correspondentId
	}
	correspondentId := ascendsdk.String(os.Getenv("CORRESPONDENT_ID"))
	require.NotNil(f.t, correspondentId, "CORRESPONDENT_ID is required and must be set as an environment variable")
	f.correspondentId = correspondentId
	return correspondentId
}

func (f *Fixtures) SubscriberId() *string {
	if f.subscriberId != nil {
		return f.subscriberId
	}

	subscriberId, err := subscriberId(f.sdk, f.ctx, f.CorrespondentId())
	require.NoError(f.t, err)

	f.subscriberId = subscriberId

	return subscriberId
}

func (f *Fixtures) TestSubscriberId() *string {
	if f.testSubscriberId != nil {
		return f.testSubscriberId
	}

	s := f.sdk
	ctx := f.ctx

	res, err := s.Subscriber.ListPushSubscriptions(ctx, nil, nil, nil)
	if err != nil {
		return nil
	}
	subscriptions := res.ListPushSubscriptionsResponse.PushSubscriptions
	if len(subscriptions) > 0 {
		f.testSubscriberId = subscriptions[0].SubscriptionID
	}

	return f.testSubscriberId
}

func (f *Fixtures) DeliveryId() *string {
	if f.deliveryId != nil {
		return f.deliveryId
	}

	deliveryId, err := deliveryID(f.sdk, f.ctx, f.TestSubscriberId())
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
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("UpdatePushSubscription", func(t *testing.T) {
		pushSubscriptionUpdate := components.PushSubscriptionUpdate{
			EventTypes: []string{
				"position.v2.updated",
			},
		}

		res, err := sdk.Subscriber.UpdatePushSubscription(ctx, *fixtures.SubscriberId(), pushSubscriptionUpdate, nil)
		require.NoError(t, err)
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("ListPushSubscriptionDeliveries", func(t *testing.T) {
		assert.NotNil(t, fixtures.DeliveryId())
	})

	t.Run("GetSubscriptionEventDelivery", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscriptionDelivery(ctx, *fixtures.TestSubscriberId(), *fixtures.DeliveryId())
		require.NoError(t, err)
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("DeletePushSubscription", func(t *testing.T) {
		res, err := sdk.Subscriber.DeletePushSubscription(ctx, *fixtures.SubscriberId())
		require.NoError(t, err)
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
