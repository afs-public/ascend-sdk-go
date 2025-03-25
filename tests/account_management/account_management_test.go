package account_management1

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountManagement(t *testing.T) {
	ctx := context.Background()
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	legalNaturalPersonId, err := helpers.CreateLegalNaturalPersonId(sdk, ctx)
	require.NoError(t, err)

	accountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	agreements, err := helpers.EnrollAccountIds(sdk, ctx, *accountId)
	require.NoError(t, err)

	err = helpers.AffirmAgreements(sdk, ctx, *accountId, agreements)
	require.NoError(t, err)

	time.Sleep(5 * time.Second)

	listAccountId := accountId
	response, err := sdk.AccountCreation.GetAccount(ctx, *listAccountId, operations.QueryParamViewFull.ToPointer())
	require.NoError(t, err)

	ownerPartyId := response.Account.Parties[0].PartyID

	testAccountManagementListAccounts(t, sdk, ctx)
	testAccountManagementUpdateAccount(t, sdk, ctx, *listAccountId)
	partyId := testAccountManagementAddParty(t, sdk, ctx, *listAccountId, *ownerPartyId, *legalNaturalPersonId)
	testAccountManagementUpdateParty(t, sdk, ctx, *listAccountId, *partyId)
	replacedPartyId := testAccountManagementReplaceParty(t, sdk, ctx, *listAccountId, *partyId, *legalNaturalPersonId, *ownerPartyId)
	testAccountManagementRemoveParty(t, sdk, ctx, *listAccountId, *replacedPartyId, *ownerPartyId)
	trustedContactId := testAccountManagementCreateTrustedContact(t, sdk, ctx, *listAccountId)
	testAccountManagementUpdateTrustedContact(t, sdk, ctx, *listAccountId, *trustedContactId)
	testAccountManagementDeleteTrustedContact(t, sdk, ctx, *listAccountId, *trustedContactId)
	interestedPartyId := testAccountManagementCreateInterestedParty(t, sdk, ctx, *listAccountId)
	testAccountManagementUpdateInterestedParty(t, sdk, ctx, *listAccountId, *interestedPartyId)
	testAccountManagementDeleteInterestedParty(t, sdk, ctx, *listAccountId, *interestedPartyId)
	testAccountManagementListAvailableRestrictions(t, sdk, ctx, *listAccountId)
	restrictionId := testAccountManagementCreateRestriction(t, sdk, ctx, *listAccountId)
	testAccountManagementEndRestriction(t, sdk, ctx, *listAccountId, *restrictionId)
	testAccountManagementCloseAccount(t, sdk, ctx, *listAccountId)
}

func testAccountManagementListAccounts(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context) (*string, *string) {

	res, err := sdk.AccountManagement.ListAccounts(ctx, operations.AccountsListAccountsRequest{
		View: operations.ViewFull.ToPointer(),
	})
	require.NoError(t, err)
	assert.NotNil(t, res.ListAccountsResponse)
	for _, account := range res.ListAccountsResponse.Accounts {
		if len(account.ActiveRestrictions) == 0 {
			for _, party := range account.Parties {
				if party.RelationType != nil && *party.RelationType == components.PartyRelationTypePrimaryOwner {
					assert.NotNil(t, account.AccountID)
					assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
					return account.AccountID, party.PartyID
				}
			}
		}
	}
	return nil, nil
}

func testAccountManagementUpdateAccount(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {

	accountRequestUpdateCreate := components.AccountRequestUpdate{
		CatAccountHolderType: components.AccountRequestUpdateCatAccountHolderTypeIIndividual.ToPointer(),
	}

	res, err := sdk.AccountManagement.UpdateAccount(ctx, accountId, accountRequestUpdateCreate, nil)
	require.NoError(t, err)
	assert.NotNil(t, res.Account)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementAddParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, ownerPartyId string, legalNaturalPersonId string) *string {

	partyRequestCreate := components.AddPartyRequestCreate{
		AuthorizedByPartyIds: []string{
			ownerPartyId,
		},
		Parent: fmt.Sprintf("accounts/%s", accountId),
		Party: components.PartyRequestCreate{
			EmailAddress: "example@domain.com",
			MailingAddress: components.PostalAddressCreate{
				AddressLines:       []string{"123 Main Street"},
				AdministrativeArea: ascendsdk.String("CA"),
				Locality:           ascendsdk.String("San Francisco"),
				PostalCode:         ascendsdk.String("12345"),
				RegionCode:         ascendsdk.String("US"),
			},
			PhoneNumber: components.PhoneNumberCreate{
				E164Number: ascendsdk.String("+14155552671"),
			},
			ProspectusDeliveryPreference:        components.ProspectusDeliveryPreferenceDigital.ToPointer(),
			ProxyDeliveryPreference:             components.ProxyDeliveryPreferenceDigital.ToPointer(),
			LegalNaturalPersonID:                ascendsdk.String(legalNaturalPersonId),
			RelationType:                        components.RelationTypeAuthorizedRepresentative,
			StatementDeliveryPreference:         components.PartyRequestCreateStatementDeliveryPreferenceDigital.ToPointer(),
			TaxDocumentDeliveryPreference:       components.TaxDocumentDeliveryPreferenceDigital.ToPointer(),
			TradeConfirmationDeliveryPreference: components.PartyRequestCreateTradeConfirmationDeliveryPreferenceDigital.ToPointer(),
		},
	}

	res, err := sdk.AccountManagement.AddParty(ctx, accountId, partyRequestCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.Party)
	assert.NotNil(t, res.Party.PartyID)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	return res.Party.PartyID
}

func testAccountManagementUpdateParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, partyId string) {
	partyRequestUpdateCreate := components.PartyRequestUpdate{
		EmailAddress: ascendsdk.String("updatedexample@domain.com"),
	}
	res, err := sdk.AccountManagement.UpdateParty(ctx, accountId, partyId, partyRequestUpdateCreate, nil)
	require.NoError(t, err)
	assert.NotNil(t, res.Party)
	assert.Equal(t, "updatedexample@domain.com", *res.Party.EmailAddress)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementReplaceParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, partyId string, legalNaturalPersonId2 string, ownerPartyId string) *string {

	replacePartyRequestCreate := components.ReplacePartyRequestCreate{
		AuthorizedByPartyIds: []string{
			ownerPartyId,
		},
		Name: fmt.Sprintf("accounts/%s/parties/%s", accountId, partyId),
		Party: components.PartyRequestCreate{
			EmailAddress:         "example@myemaildomain.com",
			LegalNaturalPersonID: ascendsdk.String(legalNaturalPersonId2),
			MailingAddress: components.PostalAddressCreate{
				AddressLines:       []string{"123 Main Street"},
				AdministrativeArea: ascendsdk.String("CA"),
				Locality:           ascendsdk.String("San Francisco"),
				PostalCode:         ascendsdk.String("12345"),
				RegionCode:         ascendsdk.String("US"),
			},
			PhoneNumber: components.PhoneNumberCreate{
				E164Number: ascendsdk.String("+14155552671"),
			},
			RelationType: components.RelationTypeAuthorizedRepresentative,
		},
	}
	res, err := sdk.AccountManagement.ReplaceParty(ctx, accountId, partyId, replacePartyRequestCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.Party)
	assert.NotEqual(t, partyId, *res.Party.PartyID)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	return res.Party.PartyID
}

func testAccountManagementRemoveParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, partyId string, ownerPartyId string) {
	removePartyRequestCreate := components.RemovePartyRequestCreate{
		AuthorizedByPartyIds: []string{
			ownerPartyId,
		},
		Name: fmt.Sprintf("accounts/%s/parties/%s", accountId, partyId),
	}
	res, err := sdk.AccountManagement.RemoveParty(ctx, accountId, partyId, removePartyRequestCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementCreateTrustedContact(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) *string {
	trustedContactCreate := components.TrustedContactCreate{
		EmailAddress: ascendsdk.String("example@email.com"),
		FamilyName:   "Doe",
		GivenName:    "John",
		MiddleNames:  ascendsdk.String("Larry"),
	}
	res, err := sdk.AccountManagement.CreateTrustedContact(ctx, accountId, trustedContactCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.TrustedContact)
	assert.NotNil(t, res.TrustedContact.TrustedContactID)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	return res.TrustedContact.TrustedContactID
}

func testAccountManagementUpdateTrustedContact(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, trustedContactId string) {
	trustedContactUpdate := components.TrustedContactUpdate{
		EmailAddress: ascendsdk.String("updatedexample@email.com"),
	}
	res, err := sdk.AccountManagement.UpdateTrustedContact(ctx, accountId, trustedContactId, trustedContactUpdate, nil)
	require.NoError(t, err)
	assert.NotNil(t, res.TrustedContact)
	assert.Equal(t, "updatedexample@email.com", *res.TrustedContact.EmailAddress)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementDeleteTrustedContact(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, trustedContactId string) {
	res, err := sdk.AccountManagement.DeleteTrustedContact(ctx, accountId, trustedContactId)
	require.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementCreateInterestedParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) *string {
	interestedPartyCreate := components.InterestedPartyCreate{
		MailingAddress: components.PostalAddressCreate{
			AddressLines:       []string{"123 Main Street"},
			AdministrativeArea: ascendsdk.String("CA"),
			Locality:           ascendsdk.String("San Francisco"),
			PostalCode:         ascendsdk.String("12345"),
			RegionCode:         ascendsdk.String("US"),
		},
		Recipient: "Bobby Jacobs",
	}
	res, err := sdk.AccountManagement.CreateInterestedParty(ctx, accountId, interestedPartyCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.InterestedParty)
	assert.NotNil(t, res.InterestedParty.InterestedPartyID)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	return res.InterestedParty.InterestedPartyID
}

func testAccountManagementUpdateInterestedParty(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, interestedPartyId string) {
	interestedPartyUpdate := components.InterestedPartyUpdate{
		Recipient: ascendsdk.String("John Doe"),
	}
	res, err := sdk.AccountManagement.UpdateInterestedParty(ctx, accountId, interestedPartyId, interestedPartyUpdate, nil)
	require.NoError(t, err)
	assert.NotNil(t, res.InterestedParty)
	assert.Equal(t, "John Doe", *res.InterestedParty.Recipient)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementDeleteInterestedParty(t *testing.T, skd *ascendsdk.SDK, ctx context.Context, accountId string, interestedPartyId string) {
	res, err := skd.AccountManagement.DeleteInterestedParty(ctx, accountId, interestedPartyId)
	require.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementListAvailableRestrictions(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	res, err := sdk.AccountManagement.ListAvailableRestrictions(ctx, accountId)
	require.NoError(t, err)
	assert.NotNil(t, res.ListAvailableRestrictionsResponse.AvailableRestrictions[0].RestrictionCode)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementCreateRestriction(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) *string {
	restrictionCreate := components.RestrictionCreate{
		CreateReason:    "Some reason for adding",
		EndedReason:     ascendsdk.String("Some reason for removing"),
		RestrictionCode: "TRADING_LIQUIDATION_ONLY_BY_CORRESPONDENT",
	}
	res, err := sdk.AccountManagement.CreateRestriction(ctx, accountId, restrictionCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.Restriction.RestrictionCode)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	return res.Restriction.RestrictionCode
}

func testAccountManagementEndRestriction(t *testing.T, ask *ascendsdk.SDK, ctx context.Context, accountId, restrictionId string) {
	EndRestrictionRequestCreate := components.EndRestrictionRequestCreate{
		Reason: "Some reason for removing",
	}
	res, err := ask.AccountManagement.EndRestriction(ctx, accountId, restrictionId, EndRestrictionRequestCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testAccountManagementCloseAccount(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	closeAccountRequestCreate := components.CloseAccountRequestCreate{
		ReplacedFdid: ascendsdk.String(""),
	}
	res, err := sdk.AccountManagement.CloseAccount(ctx, accountId, closeAccountRequestCreate)
	require.NoError(t, err)
	assert.NotNil(t, res.HTTPMeta.Response)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
