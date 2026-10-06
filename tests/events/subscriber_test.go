package events

import (
	"context"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/models/sdkerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Fixtures struct {
	t                   *testing.T
	sdk                 *ascendsdk.SDK
	ctx                 context.Context
	subscriberId        *string
	subscriberIdDeleted bool
	testSubscriberId    *string
	deliveryId          *string
	correspondentId     *string
}

// Fixture methods take the *testing.T of whichever subtest is calling them --
// require failures must FailNow the subtest's own T, not the parent's (see
// the ach_transfers Fixture comment). f.t is kept only for t.Cleanup, which
// must register on the parent so teardown runs after all subtests.
func (f *Fixtures) CorrespondentId(t *testing.T) *string {
	if f.correspondentId != nil {
		return f.correspondentId
	}
	correspondentId := ascendsdk.String(os.Getenv("CORRESPONDENT_ID"))
	require.NotNil(t, correspondentId, "CORRESPONDENT_ID is required and must be set as an environment variable")
	f.correspondentId = correspondentId
	return correspondentId
}

func (f *Fixtures) SubscriberId(t *testing.T) *string {
	if f.subscriberId != nil {
		return f.subscriberId
	}

	subscriberId, err := subscriberId(f.sdk, f.ctx, f.CorrespondentId(t))
	require.NoError(t, err)

	f.subscriberId = subscriberId

	// Safety net: guarantee this subscription is deleted even if a later
	// subtest fails before the explicit DeletePushSubscription step runs.
	// Left over subscriptions accumulate against a fixed per-account quota
	// and eventually make every subsequent run fail at creation time.
	f.t.Cleanup(func() {
		if f.subscriberIdDeleted {
			return
		}
		_ = helpers.RetryOnTransientError(func() error {
			_, opErr := f.sdk.Subscriber.DeletePushSubscription(f.ctx, *subscriberId)
			return opErr
		})
	})

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
	if len(subscriptions) == 0 {
		return nil
	}

	// The first listed subscription can be a freshly created one with no
	// delivery history (e.g. from a concurrently running suite's create
	// test); prefer a subscription that already has deliveries.
	for _, subscription := range subscriptions {
		if _, err := deliveryID(s, ctx, subscription.SubscriptionID); err == nil {
			f.testSubscriberId = subscription.SubscriptionID
			return f.testSubscriberId
		}
	}
	f.testSubscriberId = subscriptions[0].SubscriptionID

	return f.testSubscriberId
}

func (f *Fixtures) DeliveryId(t *testing.T) *string {
	if f.deliveryId != nil {
		return f.deliveryId
	}

	deliveryId, err := deliveryID(f.sdk, f.ctx, f.TestSubscriberId())
	require.NoError(t, err)

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
		assert.NotNil(t, fixtures.SubscriberId(t))
	})

	t.Run("GetPushSubscription", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscription(ctx, *fixtures.SubscriberId(t))
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

		// A fresh subscription is occasionally not yet mutable for a window
		// observed up to ~18s after creation against the real UAT
		// environment; UpdatePushSubscription surfaces that window as a
		// generic message-less internal error, so retry on any error.
		var res *operations.SubscriberUpdatePushSubscriptionResponse
		err := helpers.RetryOnTransientError(func() error {
			var opErr error
			res, opErr = sdk.Subscriber.UpdatePushSubscription(ctx, *fixtures.SubscriberId(t), pushSubscriptionUpdate, nil)
			return opErr
		})
		require.NoError(t, err)
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("ListPushSubscriptionDeliveries", func(t *testing.T) {
		assert.NotNil(t, fixtures.DeliveryId(t))
	})

	t.Run("GetSubscriptionEventDelivery", func(t *testing.T) {
		// Re-pick the subscription/delivery pair on each attempt: a
		// concurrently running suite can delete the picked subscription
		// between the pick and the read.
		var res *operations.SubscriberGetPushSubscriptionDeliveryResponse
		err := helpers.RetryOnTransientError(func() error {
			subscriptionId, deliveryId, pickErr := subscriptionDelivery(sdk, ctx)
			if pickErr != nil {
				return pickErr
			}
			var opErr error
			res, opErr = sdk.Subscriber.GetPushSubscriptionDelivery(ctx, *subscriptionId, *deliveryId)
			return opErr
		})
		require.NoError(t, err)
		assert.NotNil(t, res.HTTPMeta)
		assert.NotNil(t, res.HTTPMeta.Response)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("DeletePushSubscription", func(t *testing.T) {
		// Deletes are not idempotent: if an earlier attempt succeeded
		// server-side but its response was lost, retries see NOT_FOUND.
		// Treat that as success instead of retrying a completed delete
		// into a guaranteed failure.
		var res *operations.SubscriberDeletePushSubscriptionResponse
		err := helpers.RetryOnTransientError(func() error {
			var opErr error
			res, opErr = sdk.Subscriber.DeletePushSubscription(ctx, *fixtures.SubscriberId(t))
			if opErr != nil {
				if statusErr, ok := opErr.(*sdkerrors.Status); ok && statusErr.Code != nil && *statusErr.Code == 5 {
					res = nil
					return nil
				}
				return opErr
			}
			return nil
		})
		require.NoError(t, err)
		if res != nil {
			assert.NotNil(t, res.HTTPMeta)
			assert.NotNil(t, res.HTTPMeta.Response)
			assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		}
		fixtures.subscriberIdDeleted = true
	})
}
