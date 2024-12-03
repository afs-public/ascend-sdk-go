package tests

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"ascend-sdk/tests"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvestigationServiceUpdate(t *testing.T) {
	sdk, err := tests.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	testInvestigationServiceUpdate(t, sdk, ctx, "b90a7bb5-de15-49ee-8af8-0f30a29e8c72")
}

func testInvestigationServiceUpdate(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, investigationID string) {
	newComment := "updated comment"
	update := components.InvestigationUpdate{Comment: &newComment}

	response, err := sdk.Investigations.UpdateInvestigation(ctx, investigationID, update)
	require.NoError(t, err)
	assert.Equal(t, 200, response.HTTPMeta.Response.StatusCode)
}

func TestInvestigationServiceLinkDocuments(t *testing.T) {
	sdk, err := tests.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	testInvestigationServiceLinkDocuments(t, sdk, ctx, "b90a7bb5-de15-49ee-8af8-0f30a29e8c72")
}

func testInvestigationServiceLinkDocuments(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, investigationID string) {
	request := components.LinkDocumentsRequestCreate{IdentityVerificationDocumentIds: []string{uuid.NewString()}}

	response, err := sdk.Investigations.LinkDocuments(ctx, investigationID, request)

	require.NoError(t, err)
	assert.Equal(t, 200, response.HTTPMeta.Response.StatusCode)
}
