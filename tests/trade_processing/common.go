package trade_processing

import (
	"context"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/google/uuid"
)

type Fixtures struct {
	t                           *testing.T
	sdk                         *ascendsdk.SDK
	ctx                         context.Context
	enrolledWithdrawalAccountId string
	deceasedAccountId           string
}

func CreateBooking(sdk *ascendsdk.SDK, ctx context.Context, accountId string) ([]string, error) {
	// Create a new booking request
	now := time.Now()
	assetType := components.TradeCreateAssetTypeEquity
	bookingRequest := components.TradeCreate{
		AccountID:      accountId,
		BrokerCapacity: components.TradeCreateBrokerCapacityPrincipal,
		ClientOrderID:  uuid.New().String(),
		Executions: []components.ExecutionCreate{
			components.ExecutionCreate{
				ExecutionTime: &now,
				ExternalID:    uuid.New().String(),
				Price:         components.DecimalCreate{Value: ascendsdk.String("5")},
				Quantity:      components.DecimalCreate{Value: ascendsdk.String("1")},
			},
		},
		Identifier:        "SBUX",
		IdentifierType:    components.TradeCreateIdentifierTypeSymbol,
		RouteType:         components.RouteTypeQuik,
		Side:              components.TradeCreateSideBuy,
		SourceApplication: "Trading-App",
		AssetType:         assetType,
		Open:              ascendsdk.Bool(true),
	}

	// Send the booking request
	response, err := sdk.TradeBooking.CreateTrade(ctx, accountId, bookingRequest)
	if err != nil {
		return nil, err
	}

	// Return the trade ID and client order ID
	return []string{*response.BookingTrade.TradeID, *response.BookingTrade.ClientOrderID}, nil
}

func CreateExecution(sdk *ascendsdk.SDK, ctx context.Context, accountId string, tradeID string) (string, error) {
	// Create a new execution request
	now := time.Now()
	executionRequest := components.ExecutionCreate{
		ExecutionTime: &now,
		ExternalID:    uuid.New().String(),
		Price:         components.DecimalCreate{Value: ascendsdk.String("5")},
		Quantity:      components.DecimalCreate{Value: ascendsdk.String("1")},
	}

	// Send the execution request
	response, err := sdk.TradeBooking.CreateExecution(ctx, accountId, tradeID, executionRequest)
	if err != nil {
		return "", err
	}

	// Return the execution ID
	return *response.Execution.ExecutionID, nil
}

func CreateTradeAllocation(sdk *ascendsdk.SDK, ctx context.Context, accountId string, deceasedAccountID string) (string, error) {
	// Create a new trade allocation request
	now := time.Now()
	assetType := components.TradeAllocationCreateAssetTypeEquity
	allocationRequest := components.TradeAllocationCreate{
		BrokerCapacity:    components.TradeAllocationCreateBrokerCapacityPrincipal,
		ExecutionTime:     &now,
		FromAccountID:     deceasedAccountID,
		Identifier:        "SBUX",
		IdentifierType:    components.TradeAllocationCreateIdentifierTypeSymbol,
		AssetType:         assetType,
		Price:             components.DecimalCreate{Value: ascendsdk.String("5")},
		Quantity:          components.DecimalCreate{Value: ascendsdk.String("1")},
		SourceApplication: "Trading-App",
		ToAccountID:       accountId,
		ToSide:            components.ToSideBuy,
	}
	request_ID := uuid.New().String()

	// Send the trade allocation request
	response, err := sdk.TradeAllocation.CreateTradeAllocation(ctx, accountId, allocationRequest, &request_ID)
	if err != nil {
		return "", err
	}
	// Return the trade allocation ID
	return *response.TradeAllocation.TradeAllocationID, nil
}
