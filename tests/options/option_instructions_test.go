package options

import (
	"context"
	"errors"
	"fmt"
	"testing"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const optionAssetID = "9098656"

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

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

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
	fixtures.assetId = optionAssetID

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
