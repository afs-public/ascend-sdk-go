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

func TestPersonManagement_AccountsListLegalNaturalPersons_ListLegalNaturalPersons1(t *testing.T) {
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
	res, err := s.PersonManagement.ListLegalNaturalPersons(ctx, ascendsdk.Int(25), ascendsdk.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), nil, ascendsdk.String("legal_natural_person_id == \"e6716139-da77-46d1-9f15-13599161db0b\""))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPersonManagement_AccountsListLegalEntities_ListLegalEntities1(t *testing.T) {
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
	res, err := s.PersonManagement.ListLegalEntities(ctx, ascendsdk.Int(25), ascendsdk.String("AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h"), nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
