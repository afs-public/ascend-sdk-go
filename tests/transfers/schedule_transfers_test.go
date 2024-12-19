package transfers

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestScheduleTransfers(t *testing.T) {
	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	accountId := os.Getenv("account_id")
	bankRelationshipId := os.Getenv("bank_relationship_id")

	ctx := context.Background()

	scheduleId := testCreateAchDepositSchedule(t, sdk, ctx, accountId, bankRelationshipId)
	testListAchDepositSchedules(t, sdk, ctx, accountId)
	testGetAchDepositSchedule(t, sdk, ctx, accountId, scheduleId)
	testUpdateAchDepositSchedule(t, sdk, ctx, accountId, scheduleId)
	testCancelAchDepositSchedule(t, sdk, ctx, accountId, scheduleId)

	withdrawalScheduleId := testCreateAchWithdrawalSchedule(t, sdk, ctx, accountId, bankRelationshipId)
	testListAchWithdrawalSchedules(t, sdk, ctx, accountId)
	testGetAchWithdrawalSchedule(t, sdk, ctx, accountId, withdrawalScheduleId)
	testUpdateAchWithdrawalSchedule(t, sdk, ctx, accountId, withdrawalScheduleId)
	testCancelAchWithdrawalSchedule(t, sdk, ctx, accountId, withdrawalScheduleId)
}

func testCreateAchDepositSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) string {
	scheduleCreate := components.AchDepositScheduleCreate{
		BankRelationship: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
		ScheduleDetails: components.DepositScheduleDetailsCreate{
			Amount:           components.DecimalCreate{Value: ascendsdk.String("10.00")},
			ClientScheduleID: uuid.New().String(),
			ScheduleProperties: components.SchedulePropertiesCreate{
				Occurrences: ascendsdk.Int(10),
				StartDate: components.DateCreate{
					Day:   ascendsdk.Int(time.Now().Day()),
					Month: ascendsdk.Int(int(time.Now().Month())),
					Year:  ascendsdk.Int(time.Now().Year()),
				},
				TimeUnit:       components.TimeUnit("MONTH"),
				UnitMultiplier: 1,
			},
		},
	}
	res, err := sdk.ScheduleTransfers.CreateAchDepositSchedule(ctx, accountId, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	scheduleId := strings.Split(*res.AchDepositSchedule.Name, "/")[3]
	return scheduleId
}

func testListAchDepositSchedules(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	res, err := sdk.ScheduleTransfers.ListAchDepositSchedules(ctx, accountId, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testGetAchDepositSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	res, err := sdk.ScheduleTransfers.GetAchDepositSchedule(ctx, accountId, scheduleId)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testUpdateAchDepositSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	scheduleUpdate := components.AchDepositScheduleUpdate{
		ScheduleDetails: &components.DepositScheduleDetailsUpdate{
			Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
		},
	}

	res, err := sdk.ScheduleTransfers.UpdateAchDepositSchedule(ctx, accountId, scheduleId, scheduleUpdate, ascendsdk.String("schedule_details.amount"))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, "20.00", *res.AchDepositSchedule.ScheduleDetails.Amount.Value)
}

func testCancelAchDepositSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	req := components.CancelAchDepositScheduleRequestCreate{
		Name:    "accounts/" + accountId + "/scheduleTransfers/achDepositSchedules/" + scheduleId,
		Comment: ascendsdk.String("canceled due to test"),
	}
	res, err := sdk.ScheduleTransfers.CancelAchDepositSchedule(ctx, accountId, scheduleId, req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testCreateAchWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) string {
	scheduleCreate := components.AchWithdrawalScheduleCreate{
		BankRelationship: "accounts/" + accountId + "/bankRelationships/" + bankRelationshipId,
		ScheduleDetails: components.WithdrawalScheduleDetailsCreate{
			Amount:           &components.DecimalCreate{Value: ascendsdk.String("10.00")},
			ClientScheduleID: uuid.New().String(),
			ScheduleProperties: components.SchedulePropertiesCreate{
				Occurrences: ascendsdk.Int(10),
				StartDate: components.DateCreate{
					Day:   ascendsdk.Int(time.Now().Day()),
					Month: ascendsdk.Int(int(time.Now().Month())),
					Year:  ascendsdk.Int(time.Now().Year()),
				},
				TimeUnit:       components.TimeUnit("MONTH"),
				UnitMultiplier: 1,
			},
		},
	}
	res, err := sdk.ScheduleTransfers.CreateAchWithdrawalSchedule(ctx, accountId, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	scheduleId := strings.Split(*res.AchWithdrawalSchedule.Name, "/")[3]
	return scheduleId
}

func testListAchWithdrawalSchedules(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) {
	res, err := sdk.ScheduleTransfers.ListAchWithdrawalSchedules(ctx, accountId, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testGetAchWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	res, err := sdk.ScheduleTransfers.GetAchWithdrawalSchedule(ctx, accountId, scheduleId)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func testUpdateAchWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	scheduleUpdate := components.AchWithdrawalScheduleUpdate{
		ScheduleDetails: &components.WithdrawalScheduleDetailsUpdate{
			Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
		},
	}

	res, err := sdk.ScheduleTransfers.UpdateAchWithdrawalSchedule(ctx, accountId, scheduleId, scheduleUpdate, ascendsdk.String("schedule_details.amount"))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, "20.00", *res.AchWithdrawalSchedule.ScheduleDetails.Amount.Value)
}

func testCancelAchWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, scheduleId string) {
	if scheduleId == "" {
		t.Fatalf("expected a valid scheduleId, but got an empty string")
	}

	req := components.CancelAchWithdrawalScheduleRequestCreate{
		Name:    "accounts/" + accountId + "/scheduleTransfers/achWithdrawalSchedules/" + scheduleId,
		Comment: ascendsdk.String("canceled due to test"),
	}
	res, err := sdk.ScheduleTransfers.CancelAchWithdrawalSchedule(ctx, accountId, scheduleId, req)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
