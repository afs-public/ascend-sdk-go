// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"os"
	"testing"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountManagement_AccountsListAccounts_ListAccounts1(t *testing.T) {
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
	res, err := s.AccountManagement.ListAccounts(ctx, operations.AccountsListAccountsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
