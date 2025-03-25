package ledger

import (
	"context"
	"testing"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLedger(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	listEntries, err := sdk.Ledger.ListEntries(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, ascendsdk.Pointer(10), ascendsdk.String(""), ascendsdk.String(""))
	require.NoError(t, err)
	assert.Equal(t, 200, listEntries.HTTPMeta.Response.StatusCode)
	entryId := listEntries.ListEntriesResponse.Entries[0]

	listActivities, err := sdk.Ledger.ListActivities(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, ascendsdk.Pointer(10), ascendsdk.String(""), ascendsdk.String(""))
	require.NoError(t, err)
	assert.Equal(t, 200, listActivities.HTTPMeta.Response.StatusCode)
	activityId := listActivities.ListActivitiesResponse.Activities[0]
	assert.NotNil(t, activityId)

	t.Run("Ledger Ledger List Entries List Entries1", func(t *testing.T) {
		assert.NotNil(t, entryId)
	})

	t.Run("Ledger Ledger List Activities List Activities1", func(t *testing.T) {
		assert.NotNil(t, activityId)
	})

	t.Run("Ledger Ledger List Positions List Positions1", func(t *testing.T) {
		result, err := sdk.Ledger.ListPositions(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Ledger Ledger Get Activity Get Activity1", func(t *testing.T) {
		result, err := sdk.Ledger.GetActivity(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, *activityId.ActivityID)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Ledger Ledger Get Entry Get Entry1", func(t *testing.T) {
		result, err := sdk.Ledger.GetEntry(ctx, helpers.WITHDRAWAL_ACCOUNT_ID, *entryId.EntryID)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
