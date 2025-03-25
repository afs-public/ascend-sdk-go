package account_management1

import (
	"context"
	"testing"
	"time"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestPersonManagementAdditional(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	legalEntity, err := CreateLegalEntity(t, sdk, ctx)
	require.NoError(t, err)

	t.Logf("Running additional person management tests...")

	legalNaturalPersonId, err := helpers.CreateLegalNaturalPersonId(sdk, ctx)
	require.NoError(t, err)

	time.Sleep(5 * time.Second)

	t.Run("Person Management Accounts Get Legal Natural Person Get Legal Natural Person1", func(t *testing.T) {
		result, err := sdk.PersonManagement.GetLegalNaturalPerson(ctx, *legalNaturalPersonId)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Person Management Accounts Update Legal Natural Person Update Legal Natural Person1", func(t *testing.T) {
		status := components.LegalNaturalPersonUpdateMaritalStatusMarried
		request := components.LegalNaturalPersonUpdate{
			MaritalStatus: &status,
		}
		result, err := sdk.PersonManagement.UpdateLegalNaturalPerson(ctx, *legalNaturalPersonId, request, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)

	})

	t.Run("Test Person Management Accounts Assign Large Trader Assign Large Trader1", func(t *testing.T) {
		result, err := AssignLargeTrader(t, sdk, ctx, *legalNaturalPersonId)
		require.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Test Person Management Accounts End Large Trader End Large Trader1", func(t *testing.T) {
		endRequest := components.EndLargeTraderRequestCreate{
			EndReason: components.EndReasonEventReasonCreated,
		}

		result, err := sdk.PersonManagement.EndLargeTraderLegalNaturalPerson(ctx, *legalNaturalPersonId, endRequest)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Person Management Accounts Create Legal Entity Create Legal Entity1", func(t *testing.T) {
		assert.NotNil(t, legalEntity)
	})

	t.Run("Test Person Management Accounts Get Legal Entity Get Legal Entity1", func(t *testing.T) {
		result, err := sdk.PersonManagement.GetLegalEntity(ctx, *legalEntity.LegalEntityID)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Person Management Accounts Update Legal Entity Update Legal Entity1", func(t *testing.T) {
		request := components.LegalEntityUpdate{
			EntityName: ascendsdkgo.String("John Doe"),
		}
		result, err := sdk.PersonManagement.UpdateLegalEntity(ctx, *legalEntity.LegalEntityID, request, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Person Management Accounts Assign Large Trader Legal Entity Assign Large Trader Legal Entity1", func(t *testing.T) {
		request := components.AssignLargeTraderRequestCreate{
			LargeTraderID: LARGE_TRADER_ID,
		}

		result, err := sdk.PersonManagement.AssignLargeTraderLegalEntity(ctx, *legalEntity.LegalEntityID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Person Management Accounts End Large Trader Legal Entity End Large Trader Legal Entity1", func(t *testing.T) {
		request := components.EndLargeTraderRequestCreate{
			EndReason: components.EndReasonEventReasonCreated,
		}

		result, err := sdk.PersonManagement.EndLargeTrader(ctx, *legalEntity.LegalEntityID, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
