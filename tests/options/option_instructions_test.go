package options

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// isOutsideExerciseSubmissionWindow reports whether it's currently outside the
// window exercise instructions (DO_NOT_EXERCISE, etc.) are accepted in:
// 3:00:00 PM - 4:19:59 PM Central Time, the real-world options exercise
// cutoff near market close. Outside that window every attempt 400s with
// "outside allowed time window" regardless of which contract or expiration
// date is used.
func isOutsideExerciseSubmissionWindow() bool {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return true
	}
	now := time.Now().In(loc)
	currentMinutes := now.Hour()*60 + now.Minute()
	windowStart := 15 * 60      // 3:00 PM
	windowEnd := 16*60 + 19 + 1 // 4:19:59 PM, rounded up to the minute
	return !(windowStart <= currentMinutes && currentMinutes < windowEnd)
}

// findOptionExpiringToday finds a usable equity option contract expiring
// today. DO_NOT_EXERCISE instructions can only be submitted on an option's
// expiration date, so a fixed/hardcoded asset_id only works on the one
// calendar day it happens to expire. Looking one up dynamically each run
// stays valid indefinitely instead of breaking again the day after whatever
// asset was hardcoded. Restricted to EQUITY options -- 0DTE index options
// (XSP, APXSIM, etc.) expire daily but don't support exercise instructions at
// all ("exercise instructions are not supported for index options"). Equity
// options only expire on specific days (weekly/monthly), so none may be
// expiring today; the caller should skip in that case rather than treat it
// as a failure.
func findOptionExpiringToday(sdk *ascendsdk.SDK, ctx context.Context) (string, error) {
	// "Today" must match the America/Chicago submission window this test is
	// gated on -- system-local time diverges from it on runners east of UTC.
	chicago, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return "", err
	}
	now := time.Now().In(chicago)
	// Filter server-side: without the expiration/usable constraints this
	// walks the entire option universe page by page on no-match days.
	filter := fmt.Sprintf(
		`type == "OPTION" && usable && option.expiration_date == date("%04d-%02d-%02d")`,
		now.Year(), int(now.Month()), now.Day())
	pageSize := 200
	var pageToken *string

	for {
		res, err := sdk.Assets.ListAssets(ctx, nil, &pageSize, pageToken, &filter)
		if err != nil {
			return "", err
		}
		if res.ListAssetsResponse == nil {
			return "", nil
		}
		for _, asset := range res.ListAssetsResponse.Assets {
			if asset.AssetID == nil || asset.Option == nil {
				continue
			}
			option := asset.Option
			if option.OptionType == nil || *option.OptionType != components.OptionTypeEquity {
				continue
			}
			return *asset.AssetID, nil
		}
		if res.ListAssetsResponse.NextPageToken == nil || *res.ListAssetsResponse.NextPageToken == "" {
			return "", nil
		}
		pageToken = res.ListAssetsResponse.NextPageToken
	}
}

type Fixtures struct {
	t             *testing.T
	sdk           *ascendsdk.SDK
	ctx           context.Context
	accountId     string
	assetId       string
	instructionId *string
}

func (f *Fixtures) InstructionId(t *testing.T) *string {
	if f.instructionId != nil {
		return f.instructionId
	}

	instructionId, err := CreateOptionInstruction(t, f.sdk, f.ctx, f.accountId, f.assetId)

	fmt.Println("instructionId", instructionId)
	require.NoError(f.t, err)

	f.instructionId = &instructionId

	return &instructionId
}

func CreateOptionInstruction(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, assetId string) (string, error) {
	create := components.OptionInstructionCreate{
		AccountID:      accountId,
		Identifier:     assetId,
		IdentifierType: components.OptionInstructionCreateIdentifierTypeAssetID,
		Quantity:       components.DecimalCreate{Value: ascendsdk.String("1")},
		Type:           components.OptionInstructionCreateTypeDoNotExercise,
	}

	fmt.Printf("OptionInstructionCreate: AccountID=%s, Identifier=%s, IdentifierType=%v, Quantity=%v, Type=%v\n",
		create.AccountID, create.Identifier, create.IdentifierType, *create.Quantity.Value, create.Type)

	res, err := sdk.OptionInstructions.CreateOptionInstruction(ctx, accountId, assetId, create)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		return *res.OptionInstruction.InstructionID, nil
	}
	return "", errors.New("Error creating option instruction")
}

func TestOptionInstructionService(t *testing.T) {
	ctx := context.Background()

	if isOutsideExerciseSubmissionWindow() {
		t.Skip("Exercise instructions are only accepted 3:00-4:19:59 PM Central Time")
	}

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	assetId, err := findOptionExpiringToday(sdk, ctx)
	if err != nil {
		t.Fatalf("Error finding an option expiring today: %v", err)
	}
	if assetId == "" {
		t.Skip("No equity option contract expiring today was found")
	}

	fixtures := &Fixtures{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	accountId, err := helpers.CreateAccountId(fixtures.sdk, fixtures.ctx)
	if err != nil {
		t.Fatalf("Error creating account: %v", err)
	}
	fmt.Println("accountId", *accountId)
	fixtures.accountId = *accountId
	fixtures.assetId = assetId

	agreements, enrollErr := helpers.EnrollAccountIds(sdk, ctx, *accountId)
	if enrollErr != nil {
		t.Fatalf("Error enrolling account: %v", enrollErr)
	}

	if err := helpers.AffirmAgreements(sdk, ctx, *accountId, agreements); err != nil {
		t.Fatalf("Error affirming agreements: %v", err)
	}

	t.Run("CreateOptionInstruction", func(t *testing.T) {
		fmt.Printf("CreateOptionInstruction fixtures: accountId=%s, assetId=%s, instructionId=%v\n", fixtures.accountId, fixtures.assetId, fixtures.instructionId)
		assert.NotNil(t, fixtures.InstructionId(t))
	})

	t.Run("GetOptionInstruction", func(t *testing.T) {
		fmt.Printf("GetOptionInstruction fixtures: accountId=%s, assetId=%s, instructionId=%v\n", fixtures.accountId, fixtures.assetId, fixtures.instructionId)
		res, err := sdk.OptionInstructions.GetOptionInstruction(ctx, fixtures.accountId, fixtures.assetId, *fixtures.InstructionId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("ListOptionInstructions", func(t *testing.T) {
		fmt.Printf("ListOptionInstructions fixtures: accountId=%s, assetId=%s, instructionId=%v\n", fixtures.accountId, fixtures.assetId, fixtures.instructionId)
		request := operations.ExerciseServiceListOptionInstructionsRequest{
			AccountID: fixtures.accountId,
			AssetID:   fixtures.assetId,
		}
		res, err := sdk.OptionInstructions.ListOptionInstructions(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res.ListOptionInstructionsResponse)
	})

	t.Run("CancelOptionInstruction", func(t *testing.T) {
		fmt.Printf("CancelOptionInstruction fixtures: accountId=%s, assetId=%s, instructionId=%v\n", fixtures.accountId, fixtures.assetId, fixtures.instructionId)
		request := components.CancelOptionInstructionRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/assets/" + fixtures.assetId + "/instructions/" + *fixtures.InstructionId(t),
		}
		res, err := sdk.OptionInstructions.CancelOptionInstruction(ctx, fixtures.accountId, fixtures.assetId, *fixtures.InstructionId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res.OptionInstruction.InstructionID)
	})
}
