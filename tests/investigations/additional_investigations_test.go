package tests

import (
	"context"
	"os"
	"strings"
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

func TestIdentityLookupServiceCreateIdentityLookup(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	correspondentId := os.Getenv("CORRESPONDENT_ID")
	require.NotEmpty(t, correspondentId, "CORRESPONDENT_ID environment variable must be set")

	identityLookupId := testIdentityLookupServiceCreateIdentityLookup(t, sdk, ctx, correspondentId)
	assert.NotNil(t, identityLookupId)
	assert.Greater(t, len(identityLookupId), 0)
}

func testIdentityLookupServiceCreateIdentityLookup(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, correspondentId string) string {
	request := components.IdentityLookupCreate{
		DeviceMetadata: components.DeviceMetadataCreate{
			IPAddress: "203.0.113.42",
		},
		Identification: components.IdentificationCreate{
			RegionCode: "US",
			Type:       components.IdentificationCreateTypeSsn,
			Value:      "123-45-6789",
		},
		PhoneNumber: components.PhoneNumberCreate{
			E164Number: ascendsdk.String("+15035550123"),
			Extension:  ascendsdk.String("123"),
		},
		UserConsent: true,
	}

	res, err := sdk.Investigations.CreateIdentityLookup(ctx, correspondentId, request)
	require.NoError(t, err)
	require.NotNil(t, res.IdentityLookup)
	require.NotNil(t, res.IdentityLookup.Name)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)

	// Extract identity lookup ID from name
	nameParts := strings.Split(*res.IdentityLookup.Name, "/")
	lookupId := nameParts[len(nameParts)-1]
	return lookupId
}

func TestIdentityLookupServiceVerifyIdentityLookup(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	correspondentId := os.Getenv("CORRESPONDENT_ID")
	require.NotEmpty(t, correspondentId, "CORRESPONDENT_ID environment variable must be set")

	// First create an identity lookup
	lookupId := testIdentityLookupServiceCreateIdentityLookup(t, sdk, ctx, correspondentId)
	require.NotNil(t, lookupId)

	testIdentityLookupServiceVerifyIdentityLookup(t, sdk, ctx, correspondentId, lookupId)
}

func testIdentityLookupServiceVerifyIdentityLookup(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, correspondentId string, identityLookupId string) {
	verifyRequest := components.VerifyIdentityLookupRequestCreate{
		Name:             "correspondents/" + correspondentId + "/identityLookups/" + identityLookupId,
		VerificationCode: "123456", // This is a test verification code
	}

	res, err := sdk.Investigations.VerifyIdentityLookup(ctx, correspondentId, identityLookupId, verifyRequest)

	// The verification may fail with invalid code, which is expected in test environment
	// We're just testing that the endpoint is callable
	if err != nil {
		// Verify it's the expected error about invalid verification code
		assert.Contains(t, err.Error(), "verification")
	} else {
		require.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	}
}
