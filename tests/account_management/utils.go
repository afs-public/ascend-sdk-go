package account_management1

import (
	"context"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
)

const LARGE_TRADER_ID = "123456789100"

func AssignLargeTrader(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, lnp_id string) (*components.LargeTrader, error) {
	request := components.AssignLargeTraderRequestCreate{
		LargeTraderID: LARGE_TRADER_ID,
	}

	result, err := sdk.PersonManagement.AssignLargeTrader(ctx, lnp_id, request)
	require.NoError(t, err)
	assert.NotNil(t, result.LargeTrader)
	assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)

	return result.LargeTrader, nil
}

func CreateLegalEntity(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context) (*components.LegalEntity, error) {
	falsePtr := false

	legal_entity_request := components.LegalEntityCreate{
		AccreditedInvestor:              &falsePtr,
		Adviser:                         &falsePtr,
		BrokerDealer:                    &falsePtr,
		CorrespondentID:                 os.Getenv("CORRESPONDENT_ID"),
		EntityName:                      "AcmeInc",
		EntityType:                      components.EntityTypeEstate,
		ExemptVerifyingBeneficialOwners: &falsePtr,
		ForeignFinancialInstitution:     &falsePtr,
		LegalAddress: components.PostalAddressCreate{
			AddressLines:       []string{"19409 Sherilyn Courts"},
			AdministrativeArea: ascendsdk.String("OR"),
			Locality:           ascendsdk.String("Portland"),
			PostalCode:         ascendsdk.String("97035"),
			RegionCode:         ascendsdk.String("US"),
		},
		LeiCode:            ascendsdk.String("12340012345678911000"),
		OperatingRegions:   []string{"US"},
		RegistrationRegion: "US",
		TaxID:              "987-65-4321",
		TaxIDType:          components.LegalEntityCreateTaxIDTypeTaxIDTypeItin,
		TaxProfile: components.TaxProfileCreate{
			FederalTaxClassification: components.FederalTaxClassificationTrustEstate,
			IrsFormType:              components.IrsFormTypeW9,
			LegalTaxRegionCode:       "US",
			UsTinStatus:              components.UsTinStatusPassing,
		},
	}

	result, err := sdk.PersonManagement.CreateLegalEntity(ctx, legal_entity_request)
	require.NoError(t, err)
	assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	return result.LegalEntity, nil
}
