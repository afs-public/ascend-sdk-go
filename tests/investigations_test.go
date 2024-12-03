// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvestigations_InvestigationServiceGetInvestigation_GetInvestigation1(t *testing.T) {
	s := ascendsdk.New(
		ascendsdk.WithServerURL("https://uat.apexapis.com"),
		ascendsdk.WithSecurity(components.Security{
			APIKey: ascendsdk.String(os.Getenv("API_KEY")),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY"),
				Name:         os.Getenv("SERVICE_ACCOUNT_CREDS_NAME"),
				Organization: os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION"),
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Investigations.GetInvestigation(ctx, "b90a7bb5-de15-49ee-8af8-0f30a29e8c72")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestInvestigations_InvestigationServiceListInvestigations_ListInvestigations1(t *testing.T) {
	s := ascendsdk.New(
		ascendsdk.WithServerURL("https://uat.apexapis.com"),
		ascendsdk.WithSecurity(components.Security{
			APIKey: ascendsdk.String(os.Getenv("API_KEY")),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY"),
				Name:         os.Getenv("SERVICE_ACCOUNT_CREDS_NAME"),
				Organization: os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION"),
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Investigations.ListInvestigations(ctx, ascendsdk.Int(100), nil, ascendsdk.String("investigation_subject.person_investigation.given_name == 'Jane' && investigation_subject.person_investigation.family_name == 'Dough'"))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestInvestigations_WatchlistServiceGetWatchlistItem_GetWatchlistItem1(t *testing.T) {
	s := ascendsdk.New(
		ascendsdk.WithServerURL("https://uat.apexapis.com"),
		ascendsdk.WithSecurity(components.Security{
			APIKey: ascendsdk.String(os.Getenv("API_KEY")),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY"),
				Name:         os.Getenv("SERVICE_ACCOUNT_CREDS_NAME"),
				Organization: os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION"),
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Investigations.GetWatchlistItem(ctx, "DOWJONES", "123456")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
