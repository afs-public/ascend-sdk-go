package accounts

import (
	"context"
	"testing"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPersonManagement(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("TestLegalEntities2", func(t *testing.T) {
		filterString := "legal_entity_id==\"0157dd4e-fd04-4219-8a4b-695a26547418\""
		var filter = &filterString

		res, err := sdk.PersonManagement.ListLegalEntities(ctx, nil, nil, nil, filter)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, 1, len(res.ListLegalEntitiesResponse.LegalEntities))
		assert.Equal(t, components.LegalEntityIrsFormType("W_8"), *res.ListLegalEntitiesResponse.LegalEntities[0].TaxProfile.IrsFormType)
	})
}
