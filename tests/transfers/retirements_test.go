package transfers

import (
	"context"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRetirements(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	accountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	time.Sleep(5 * time.Second)

	t.Run("Retirements Transfers List Contribution Summaries List Contribution Summaries1", func(t *testing.T) {
		pageSize := 10
		pageToken := ""
		result, err := sdk.Retirements.ListContributionSummaries(ctx, *accountId, &pageSize, &pageToken)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Retirements Transfers Retrieve Contribution Constraints Retrieve Contribution Constraints1", func(t *testing.T) {
		request := components.RetrieveContributionConstraintsRequestCreate{
			Mechanism: components.MechanismAch,
			Name:      "accounts/" + *accountId,
		}

		result, err := sdk.Retirements.RetrieveContributionConstraints(ctx, *accountId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Retirements Transfers Retrieve Distribution Constraints Retrieve Distribution Constraints", func(t *testing.T) {
		request := components.RetrieveDistributionConstraintsRequestCreate{
			Mechanism: components.RetrieveDistributionConstraintsRequestCreateMechanismAch,
			Name:      "accounts/" + *accountId,
		}

		result, err := sdk.Retirements.RetrieveDistributionConstraints(ctx, *accountId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Retirements Transfers List Distribution Summaries List Distribution Summaries1", func(t *testing.T) {
		result, err := sdk.Retirements.ListDistributionSummaries(ctx, *accountId, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
