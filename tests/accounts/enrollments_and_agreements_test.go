package accounts

import (
	"context"
	"os"
	"testing"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Fixtures struct {
	t                  *testing.T
	sdk                *ascendsdk.SDK
	ctx                context.Context
	accountId          *string
	enrolledAgreements []components.Agreement
}

func (f *Fixtures) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}

	accountId, err := helpers.CreateAccountId(f.sdk, f.ctx)
	require.NoError(f.t, err)

	f.accountId = accountId

	return accountId
}

func (f *Fixtures) EnrollAccountIds() []components.Agreement {
	if len(f.enrolledAgreements) > 0 {
		return f.enrolledAgreements
	}

	agreements, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *f.AccountId())
	require.NoError(f.t, err)

	f.enrolledAgreements = agreements

	return agreements
}

func TestEnrollmentsAndAgreements(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	ctx := context.Background()

	fixtures := &Fixtures{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("TestEnrollAccountIds", func(t *testing.T) {
		assert.NotNil(t, fixtures.EnrollAccountIds())
	})

	t.Run("TestListAvailableEnrollments", func(t *testing.T) {
		res, err := sdk.EnrollmentsAndAgreements.ListAvailableEnrollments(ctx, *fixtures.AccountId(), nil, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestListEnrollments", func(t *testing.T) {
		res, err := sdk.EnrollmentsAndAgreements.ListEnrollments(ctx, *fixtures.AccountId(), nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestAffirmAgreement", func(t *testing.T) {
		var agreementIds []string

		for _, agreement := range fixtures.EnrollAccountIds() {
			agreementIds = append(agreementIds, *agreement.AgreementID)
		}

		affirmAgreementRequest := components.AffirmAgreementsRequestCreate{AccountAgreementIds: agreementIds}

		res, err := sdk.EnrollmentsAndAgreements.AffirmAgreements(ctx, *fixtures.AccountId(), affirmAgreementRequest)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestListAgreements", func(t *testing.T) {
		res, err := sdk.EnrollmentsAndAgreements.ListAgreements(ctx, *fixtures.AccountId(), nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestListEntitlements", func(t *testing.T) {
		res, err := sdk.EnrollmentsAndAgreements.ListEntitlements(ctx, *fixtures.AccountId(), nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestDeactivateEnrollment", func(t *testing.T) {
		deactivate_enrollment_id, err := GetEnrollmentToDeactivate(sdk, ctx, *fixtures.AccountId())
		require.NoError(t, err)
		deactivateEnrollmentRequest := components.DeactivateEnrollmentRequestCreate{
			EnrollmentID: deactivate_enrollment_id,
		}

		res, err := sdk.EnrollmentsAndAgreements.DeactivateEnrollment(ctx, *fixtures.AccountId(), deactivateEnrollmentRequest)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("TestAccounts_ListAvailableEnrollmentsByAccountGroup", func(t *testing.T) {
		res, err := sdk.EnrollmentsAndAgreements.AccountsListAvailableEnrollmentsByAccountGroup(ctx, os.Getenv("ACCOUNT_GROUP_ID"), nil, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
