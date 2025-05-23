package events

import (
	"context"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	eventMessageID := testListEventMessages(t, sdk, ctx)
	testGetEventMessage(t, sdk, ctx, eventMessageID)
}

func testListEventMessages(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context) string {
	res, err := sdk.Reader.ListEventMessages(ctx, nil, nil, nil)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)

	return *res.ListEventMessagesResponse.EventMessages[0].MessageID
}

func testGetEventMessage(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, messageID string) {
	if messageID == "" {
		t.Fatalf("expected a valid messageID, but got an empty string")
	}

	res, err := sdk.Reader.GetEventMessage(ctx, messageID)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
