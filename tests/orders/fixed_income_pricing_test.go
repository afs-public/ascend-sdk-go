package orders

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFixedIncomePricing(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()
	accountId, accountIdError := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, accountIdError)

	fmt.Println("accountId", *accountId)

	agreements, enrollErr := helpers.EnrollAccountIds(sdk, ctx, *accountId)
	if enrollErr != nil {
		t.Fatalf("Error enrolling account: %v", enrollErr)
	}
	helpers.AffirmAgreements(sdk, ctx, *accountId, agreements)

	time.Sleep(5 * time.Second)
	testPreviewOrderCost(t, sdk, ctx, *accountId)
	testRetrieveQuote(t, sdk, ctx, *accountId)
	testRetrieveFixedIncomeMarks(t, sdk, ctx)
}

func testPreviewOrderCost(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {

	previewCreate := components.OrderCostPreviewRequestCreate{
		AssetType:      components.OrderCostPreviewRequestCreateAssetTypeFixedIncome,
		Identifier:     "912810SX7",
		IdentifierType: components.OrderCostPreviewRequestCreateIdentifierTypeCusip,
		LimitPrice: components.LimitPriceCreate{
			Price: components.DecimalCreate{Value: ascendsdk.String("2")},
			Type:  components.LimitPriceCreateTypePercentageOfPar,
		},
		Parent:   "accounts/" + accountId,
		Quantity: components.DecimalCreate{Value: ascendsdk.String("5")},
	}
	res, err := sdk.FixedIncomePricing.PreviewOrderCost(ctx, accountId, previewCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testRetrieveQuote(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	retrieveQuoteCreate := components.RetrieveQuoteRequestCreate{
		AssetType:      components.RetrieveQuoteRequestCreateAssetTypeFixedIncome,
		Identifier:     "912810SX7",
		IdentifierType: components.RetrieveQuoteRequestCreateIdentifierTypeCusip,
		Parent:         "accounts/" + accountId,
	}
	res, err := sdk.FixedIncomePricing.RetrieveQuote(ctx, accountId, retrieveQuoteCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testRetrieveFixedIncomeMarks(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context) {
	retrieveMarksCreate := components.RetrieveFixedIncomeMarksRequestCreate{
		Parent: "correspondents/" + os.Getenv("CORRESPONDENT_ID"),
		SecurityIdentifiers: []components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate{
			{
				Identifier:     "912810SX7",
				IdentifierType: components.RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeCusip,
			},
		},
	}
	res, err := sdk.FixedIncomePricing.RetrieveFixedIncomeMarks(ctx, os.Getenv("CORRESPONDENT_ID"), retrieveMarksCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
