package tests

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/models/components"
	"fmt"
	"os"
)

func SetupAscendSDK() (*ascendsdk.SDK, error) {
	privateKey := os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY")
	apiKey := os.Getenv("API_KEY")
	name := os.Getenv("SERVICE_ACCOUNT_CREDS_NAME")
	organization := os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION")

	url := "https://uat.apexapis.com"
	accountType := "serviceAccount"

	if privateKey == "" || apiKey == "" || name == "" || organization == "" {
		return nil, fmt.Errorf("environment variables are not set properly")
	}

	return ascendsdk.New(
		ascendsdk.WithServerURL(url),
		ascendsdk.WithSecurity(components.Security{
			APIKey: ascendsdk.String(apiKey),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   privateKey,
				Name:         name,
				Organization: organization,
				Type:         accountType,
			},
		})), nil
}
