package events

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"ascend-sdk/tests"
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func subscriberId(s *ascendsdk.SDK, ctx context.Context, t *testing.T) string {
	now := time.Now()

	httpCallback := components.HTTPPushCallbackCreate{
		URL:            "https://brokercheck.finra.org/",
		ClientSecret:   "my-secret",
		TimeoutSeconds: ascendsdk.Int(30),
	}

	request := components.PushSubscriptionCreate{
		CorrespondentID: ascendsdk.String(os.Getenv("correspondent_id")),
		DisplayName:     now.String(),
		EventTypes:      []string{"position.v1.updated"},
		HTTPCallback:    &httpCallback,
	}

	res, err := s.Subscriber.CreatePushSubscription(ctx, request)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)

	return *res.PushSubscription.SubscriptionID
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()

	sdk, err := tests.SetupAscendSDK()
	require.NoError(t, err)

	subscriberId := subscriberId(sdk, context.Background(), t)
	require.NotEmpty(t, subscriberId)

	t.Run("GetPushSubscription", func(t *testing.T) {
		res, err := sdk.Subscriber.GetPushSubscription(ctx, subscriberId)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
