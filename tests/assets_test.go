// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"os"
	"testing"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssets_AssetsListAssets1_AssetsListAssets1(t *testing.T) {
	s := ascendsdkgo.New(
		ascendsdkgo.WithServerURL("https://uat.apexapis.com"),
		ascendsdkgo.WithSecurity(components.Security{
			APIKey: ascendsdkgo.String(os.Getenv("API_KEY")),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY"),
				Name:         os.Getenv("SERVICE_ACCOUNT_CREDS_NAME"),
				Organization: os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION"),
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Assets.ListAssets(ctx, ascendsdkgo.String("correspondents/1234"), ascendsdkgo.Int(100), ascendsdkgo.String("Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA=="), ascendsdkgo.String("(symbol == 'IBM' && usable) || symbol == 'USD'"))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestAssets_AssetsGetAsset_AssetsGetAsset1(t *testing.T) {
	s := ascendsdkgo.New(
		ascendsdkgo.WithServerURL("https://uat.apexapis.com"),
		ascendsdkgo.WithSecurity(components.Security{
			APIKey: ascendsdkgo.String(os.Getenv("API_KEY")),
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
