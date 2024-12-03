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

func TestAssets_AssetsListAssets1_AssetsListAssets1(t *testing.T) {
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
	res, err := s.Assets.ListAssets(ctx, ascendsdk.String("correspondents/1234"), ascendsdk.Int(100), ascendsdk.String("Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA=="), ascendsdk.String("(symbol == 'IBM' && usable) || symbol == 'USD'"))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestAssets_AssetsGetAsset_AssetsGetAsset1(t *testing.T) {
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
	res, err := s.Assets.GetAsset(ctx, "8395")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
