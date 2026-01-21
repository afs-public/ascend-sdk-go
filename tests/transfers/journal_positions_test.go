package transfers

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"

	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/sdkerrors"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPositionJournals(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	sourceAccountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	destinationAccountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	var positionJournalId string

	t.Run("Position Journals Create Position Journal", func(t *testing.T) {
		request := components.PositionJournalCreate{
			ClientTransferID:   uuid.New().String(),
			DestinationAccount: "accounts/" + *destinationAccountId,
			Identifier:         "AAPL",
			IdentifierType:     components.IdentifierTypeSymbol,
			Quantity: components.DecimalCreate{
				Value: ascendsdkgo.String("1.00"),
			},
			SourceAccount: "accounts/" + *sourceAccountId,
			Type:          components.PositionJournalCreateTypeReward,
			FairMarketValue: &components.DecimalCreate{
				Value: ascendsdkgo.String("150.00"),
			},
		}

		result, err := sdk.PositionJournals.CreatePositionJournal(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, result.PositionJournal)

		if result.PositionJournal != nil && result.PositionJournal.Name != nil {
			parts := strings.Split(*result.PositionJournal.Name, "/")
			positionJournalId = parts[len(parts)-1]
		}
	})

	t.Run("Position Journals Get Position Journal", func(t *testing.T) {
		if positionJournalId == "" {
			t.Skip("Skipping: No position journal ID available from create test")
		}

		result, err := sdk.PositionJournals.GetPositionJournal(ctx, positionJournalId)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, result.PositionJournal)
	})

	t.Run("Position Journals Cancel Position Journal", func(t *testing.T) {
		if positionJournalId == "" {
			t.Skip("Skipping: No position journal ID available from create test")
		}

		request := components.CancelPositionJournalRequestCreate{
			Name:   "positionJournals/" + positionJournalId,
			Reason: ascendsdkgo.String("Test cancellation"),
		}

		result, err := sdk.PositionJournals.CancelPositionJournal(ctx, positionJournalId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, result.PositionJournal)
	})
}

func TestPositionJournalsTestSimulation(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	sourceAccountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	destinationAccountId, err := helpers.CreateAccountId(sdk, ctx)
	require.NoError(t, err)

	t.Run("Test Simulation Force Approve Position Journal", func(t *testing.T) {
		if isNotJournalHours() {
			t.Skip("Skipping Endpoint Test: Force Approve Position Journal")
		}

		request := components.PositionJournalCreate{
			ClientTransferID:   uuid.New().String(),
			DestinationAccount: "accounts/" + *destinationAccountId,
			Identifier:         "AAPL",
			IdentifierType:     components.IdentifierTypeSymbol,
			Quantity: components.DecimalCreate{
				Value: ascendsdkgo.String("1.00"),
			},
			SourceAccount: "accounts/" + *sourceAccountId,
			Type:          components.PositionJournalCreateTypeReward,
			FairMarketValue: &components.DecimalCreate{
				Value: ascendsdkgo.String("150.00"),
			},
			Description: ascendsdkgo.String("Stock reward for testing"),
		}

		createResult, err := sdk.PositionJournals.CreatePositionJournal(ctx, request)
		require.NoError(t, err)
		require.NotNil(t, createResult.PositionJournal)
		require.NotNil(t, createResult.PositionJournal.Name)

		parts := strings.Split(*createResult.PositionJournal.Name, "/")
		positionJournalId := parts[len(parts)-1]

		forceApproveRequest := components.ForceApprovePositionJournalRequestCreate{
			Name: "positionJournals/" + positionJournalId,
		}

		result, err := sdk.TestSimulation.ForceApprovePositionJournal(ctx, positionJournalId, forceApproveRequest)
		if err != nil {
			statusErr, ok := err.(*sdkerrors.Status)
			require.True(t, ok)
			assert.Equal(t, 3, *statusErr.Code)
			assert.True(t, strings.Contains(strings.ToLower(*statusErr.Message), "that does not need review"), *statusErr.Message)
		} else {
			assert.NotNil(t, result)
			assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
			assert.NotNil(t, result.PositionJournal)
		}
	})

	t.Run("Test Simulation Force Reject Position Journal", func(t *testing.T) {
		if isNotJournalHours() {
			t.Skip("Skipping Endpoint Test: Force Reject Position Journal")
		}

		request := components.PositionJournalCreate{
			ClientTransferID:   uuid.New().String(),
			DestinationAccount: "accounts/" + *destinationAccountId,
			Identifier:         "AAPL",
			IdentifierType:     components.IdentifierTypeSymbol,
			Quantity: components.DecimalCreate{
				Value: ascendsdkgo.String("1.00"),
			},
			SourceAccount: "accounts/" + *sourceAccountId,
			Type:          components.PositionJournalCreateTypeReward,
			FairMarketValue: &components.DecimalCreate{
				Value: ascendsdkgo.String("150.00"),
			},
			Description: ascendsdkgo.String("Stock reward for testing"),
		}

		createResult, err := sdk.PositionJournals.CreatePositionJournal(ctx, request)
		require.NoError(t, err)
		require.NotNil(t, createResult.PositionJournal)
		require.NotNil(t, createResult.PositionJournal.Name)

		parts := strings.Split(*createResult.PositionJournal.Name, "/")
		positionJournalId := parts[len(parts)-1]

		forceRejectRequest := components.ForceRejectPositionJournalRequestCreate{
			Name:   "positionJournals/" + positionJournalId,
			Reason: ascendsdkgo.String("Force reject for testing"),
		}

		result, err := sdk.TestSimulation.ForceRejectPositionJournal(ctx, positionJournalId, forceRejectRequest)
		if err != nil {
			statusErr, ok := err.(*sdkerrors.Status)
			require.True(t, ok)
			assert.Equal(t, 3, *statusErr.Code)
		} else {
			assert.NotNil(t, result)
			assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
			assert.NotNil(t, result.PositionJournal)
		}
	})
}
