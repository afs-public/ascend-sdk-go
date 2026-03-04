package transfers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CashJournalScheduleFixture struct {
	t                     *testing.T
	sdk                   *ascendsdk.SDK
	ctx                   context.Context
	sourceAccountId       string
	destinationAccountId  *string
	cashJournalScheduleId *string
}

func (f *CashJournalScheduleFixture) DestinationAccountId() *string {
	if f.destinationAccountId != nil {
		return f.destinationAccountId
	}
	f.destinationAccountId = f.createAndEnrollAccount()
	return f.destinationAccountId
}

func (f *CashJournalScheduleFixture) createAndEnrollAccount() *string {
	accountId, err := helpers.CreateAccountId(f.sdk, f.ctx)
	require.NoError(f.t, err)

	agg, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *accountId)
	require.NoError(f.t, err)

	err = helpers.AffirmAgreements(f.sdk, f.ctx, *accountId, agg)
	require.NoError(f.t, err)

	return accountId
}

func (f *CashJournalScheduleFixture) CashJournalScheduleId(t *testing.T) *string {
	if f.cashJournalScheduleId != nil {
		return f.cashJournalScheduleId
	}

	scheduleId, err := CreateCashJournalSchedule(t, f.sdk, f.ctx, f.sourceAccountId, *f.DestinationAccountId())
	fmt.Println("cashJournalScheduleId:", scheduleId)
	require.NoError(f.t, err)

	f.cashJournalScheduleId = &scheduleId
	return &scheduleId
}

func CreateCashJournalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, sourceAccountId string, destinationAccountId string) (string, error) {
	clientScheduleID := uuid.New().String()
	now := time.Now()

	scheduleCreate := components.CashJournalScheduleCreate{
		SourceAccount:      sourceAccountId,
		DestinationAccount: destinationAccountId,
		ScheduleDetails: components.WithdrawalScheduleDetailsCreate{
			Amount:           &components.DecimalCreate{Value: ascendsdk.String("10.00")},
			ClientScheduleID: uuid.New().String(),
			ScheduleProperties: components.SchedulePropertiesCreate{
				Occurrences: ascendsdk.Int(10),
				StartDate: components.DateCreate{
					Day:   ascendsdk.Int(now.Day()),
					Month: ascendsdk.Int(int(now.Month())),
					Year:  ascendsdk.Int(now.Year()),
				},
				TimeUnit:       components.TimeUnitMonth,
				UnitMultiplier: 1,
			},
		},
	}

	fmt.Println("=== CreateCashJournalSchedule Request ===")
	fmt.Println("sourceAccount:", scheduleCreate.SourceAccount)
	fmt.Println("destinationAccount:", scheduleCreate.DestinationAccount)
	fmt.Println("amount:", *scheduleCreate.ScheduleDetails.Amount.Value)
	fmt.Println("clientScheduleID:", clientScheduleID)
	fmt.Println("occurrences:", *scheduleCreate.ScheduleDetails.ScheduleProperties.Occurrences)
	fmt.Printf("startDate: %d-%02d-%02d\n", now.Year(), now.Month(), now.Day())
	fmt.Println("timeUnit:", scheduleCreate.ScheduleDetails.ScheduleProperties.TimeUnit)
	fmt.Println("unitMultiplier:", scheduleCreate.ScheduleDetails.ScheduleProperties.UnitMultiplier)

	res, err := sdk.ScheduleTransfers.CreateCashJournalSchedule(ctx, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		parts := strings.Split(*res.CashJournalSchedule.Name, "/")
		scheduleId := parts[len(parts)-1]
		return scheduleId, nil
	}
	return "", errors.New("Error creating cash journal schedule")
}

func TestCashJournalSchedules(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &CashJournalScheduleFixture{
		t:               t,
		sdk:             sdk,
		ctx:             ctx,
		sourceAccountId: withdrawal_account_id,
	}

	t.Run("CreateCashJournalSchedule", func(t *testing.T) {
		assert.NotNil(t, fixtures.CashJournalScheduleId(t))
	})

	t.Run("GetCashJournalSchedule", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.GetCashJournalSchedule(ctx, *fixtures.CashJournalScheduleId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("SearchCashJournalSchedules", func(t *testing.T) {
		sourceAccount := fixtures.sourceAccountId
		searchReq := operations.CashJournalSchedulesSearchCashJournalSchedulesRequest{
			SourceAccount: &sourceAccount,
		}
		res, err := sdk.ScheduleTransfers.SearchCashJournalSchedules(ctx, searchReq)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("UpdateCashJournalSchedule", func(t *testing.T) {
		scheduleUpdate := components.CashJournalScheduleUpdate{
			ScheduleDetails: &components.WithdrawalScheduleDetailsUpdate{
				Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
			},
		}

		res, err := sdk.ScheduleTransfers.UpdateCashJournalSchedule(ctx, *fixtures.CashJournalScheduleId(t), scheduleUpdate, ascendsdk.String("schedule_details.amount"))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelCashJournalSchedule", func(t *testing.T) {
		req := components.CancelCashJournalScheduleRequestCreate{
			Name:    *fixtures.CashJournalScheduleId(t),
			Comment: ascendsdk.String("canceled due to test"),
		}
		res, err := sdk.ScheduleTransfers.CancelCashJournalSchedule(ctx, *fixtures.CashJournalScheduleId(t), req)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
