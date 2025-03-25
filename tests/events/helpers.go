package events

import (
	"context"
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
	res, err := s.Subscriber.ListPushSubscriptionDeliveries(ctx, *subscriberID, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return res.ListPushSubscriptionDeliveriesResponse.PushSubscriptionDeliveries[0].DeliveryID, nil
}
