package events

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/tests/helpers"
	"context"
	"testing"

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
	filter, pageSize, pageToken := getListEventParams()
	res, err := sdk.Reader.ListEventMessages(ctx, &filter, &pageSize, &pageToken)

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

func getListEventParams() (filter string, pageSize int, pageToken string) {
	return "", 1, ""
}
