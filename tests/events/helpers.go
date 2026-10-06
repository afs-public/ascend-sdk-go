package events

import (
	"context"
	"fmt"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
)

func subscriberId(s *ascendsdk.SDK, ctx context.Context, correspondentId *string) (*string, error) {
	now := time.Now()

	httpCallback := components.HTTPPushCallbackCreate{
		URL:            "https://brokercheck.finra.org/",
		ClientSecret:   "mysecretkey1",
		TimeoutSeconds: ascendsdk.Int(30),
	}

	request := components.PushSubscriptionCreate{
		CorrespondentID: correspondentId,
		DisplayName:     now.Format(time.RFC1123),
		EventTypes:      []string{"position.v1.updated"},
		HTTPCallback:    &httpCallback,
	}

	res, err := s.Subscriber.CreatePushSubscription(ctx, request)

	if err != nil {
		return nil, err
	}

	return res.PushSubscription.SubscriptionID, nil
}

func deliveryID(s *ascendsdk.SDK, ctx context.Context, subscriberID *string) (*string, error) {
	if subscriberID == nil {
		return nil, fmt.Errorf("subscriberID is nil")
	}

	res, err := s.Subscriber.ListPushSubscriptionDeliveries(ctx, *subscriberID, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	deliveries := res.ListPushSubscriptionDeliveriesResponse.PushSubscriptionDeliveries
	if len(deliveries) == 0 {
		return nil, fmt.Errorf("subscription %s has no deliveries yet", *subscriberID)
	}

	return deliveries[0].DeliveryID, nil
}

// subscriptionDelivery atomically picks a subscription that has deliveries
// together with its first delivery id. Callers that read a delivery should
// re-pick through this on failure: a concurrently running suite can delete
// the picked subscription between the pick and the read.
func subscriptionDelivery(s *ascendsdk.SDK, ctx context.Context) (*string, *string, error) {
	res, err := s.Subscriber.ListPushSubscriptions(ctx, nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	for _, subscription := range res.ListPushSubscriptionsResponse.PushSubscriptions {
		if subscription.SubscriptionID == nil {
			continue
		}
		deliveryId, err := deliveryID(s, ctx, subscription.SubscriptionID)
		if err != nil {
			continue
		}
		return subscription.SubscriptionID, deliveryId, nil
	}
	return nil, nil, fmt.Errorf("no subscription with deliveries found")
}
