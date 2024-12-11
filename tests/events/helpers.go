package events

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"context"
	"time"
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
