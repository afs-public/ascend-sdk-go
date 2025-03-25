package tests

import (
	"context"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const APEX_INVESTIGATION_ID = "01JP8EHZ3CJKCTMHKTT4FZ51HC"

func TestInvestigationServiceUpdate(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	testInvestigationServiceUpdate(t, sdk, ctx, "01JHGRJG62CZ0TV805CSWYHJ31")
}

func testInvestigationServiceUpdate(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, investigationID string) {
	newComment := "updated comment"
	update := components.InvestigationUpdate{Comment: &newComment}

	response, err := sdk.Investigations.UpdateInvestigation(ctx, investigationID, update, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, response.HTTPMeta.Response.StatusCode)
}

func TestInvestigationServiceLinkDocuments(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	testInvestigationServiceLinkDocuments(t, sdk, ctx, "01JHGRJG62CZ0TV805CSWYHJ31")
}

func testInvestigationServiceLinkDocuments(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, investigationID string) {
	request := components.LinkDocumentsRequestCreate{IdentityVerificationDocumentIds: []string{uuid.NewString()}}

	response, err := sdk.Investigations.LinkDocuments(ctx, investigationID, request)

	require.NoError(t, err)
	assert.Equal(t, 200, response.HTTPMeta.Response.StatusCode)
}

func TestInvestigationServiceGetIdentityVerification(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	testInvestigationsGetIdentityVerification(t, sdk, ctx, APEX_INVESTIGATION_ID)
}
func testInvestigationsGetIdentityVerification(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, investigationID string) {
	investigation, err := sdk.Investigations.GetInvestigation(ctx, investigationID)
	require.NoError(t, err)

	results := investigation.Investigation.IdentityVerificationResults

	assert.NotNil(t, results)

	if results != nil {
		assert.NotEmpty(t, results)
	}

	firstIdResult := results[0]

	if firstIdResult.CustomerIdentificationID != nil {
		assert.NotEmpty(t, *firstIdResult.CustomerIdentificationID)
	}

	response, err := sdk.Investigations.GetCustomerIdentification(ctx, os.Getenv("CORRESPONDENT_ID"), *firstIdResult.CustomerIdentificationID, nil)

	require.NoError(t, err)
	assert.Equal(t, 200, response.HTTPMeta.Response.StatusCode)
}
