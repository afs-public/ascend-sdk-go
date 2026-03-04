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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CheckWithdrawalScheduleFixture struct {
	t                         *testing.T
	sdk                       *ascendsdk.SDK
	ctx                       context.Context
	accountId                 *string
	checkWithdrawalScheduleId *string
}

func (f *CheckWithdrawalScheduleFixture) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}
	f.accountId = f.createAndEnrollAccount()
	return f.accountId
}

func (f *CheckWithdrawalScheduleFixture) createAndEnrollAccount() *string {
	accountId, err := helpers.CreateAccountId(f.sdk, f.ctx)
	require.NoError(f.t, err)

	agg, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *accountId)
	require.NoError(f.t, err)

	err = helpers.AffirmAgreements(f.sdk, f.ctx, *accountId, agg)
	require.NoError(f.t, err)

	return accountId
}

func (f *CheckWithdrawalScheduleFixture) CheckWithdrawalScheduleId(t *testing.T) *string {
	if f.checkWithdrawalScheduleId != nil {
		return f.checkWithdrawalScheduleId
	}

	scheduleId, err := CreateCheckWithdrawalSchedule(t, f.sdk, f.ctx, *f.AccountId())
	fmt.Println("checkWithdrawalScheduleId:", scheduleId)
	require.NoError(f.t, err)

	f.checkWithdrawalScheduleId = &scheduleId
	return &scheduleId
}

func CreateCheckWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string) (string, error) {
	scheduleCreate := components.CheckWithdrawalScheduleCreate{
		DeliveryMethod: components.DeliveryMethodStandard,
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
				TimeUnit:       components.TimeUnitMonth,
				UnitMultiplier: 1,
			},
		},
	}

	res, err := sdk.ScheduleTransfers.CreateCheckWithdrawalSchedule(ctx, accountId, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		scheduleId := strings.Split(*res.CheckWithdrawalSchedule.Name, "/")[3]
		return scheduleId, nil
	}
	return "", errors.New("Error creating check withdrawal schedule")
}

func TestCheckWithdrawalSchedules(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &CheckWithdrawalScheduleFixture{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateCheckWithdrawalSchedule", func(t *testing.T) {
		assert.NotNil(t, fixtures.CheckWithdrawalScheduleId(t))
	})

	t.Run("ListCheckWithdrawalSchedules", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.ListCheckWithdrawalSchedules(ctx, *fixtures.AccountId(), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("GetCheckWithdrawalSchedule", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.GetCheckWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.CheckWithdrawalScheduleId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("UpdateCheckWithdrawalSchedule", func(t *testing.T) {
		scheduleUpdate := components.CheckWithdrawalScheduleUpdate{
			ScheduleDetails: &components.WithdrawalScheduleDetailsUpdate{
				Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
			},
		}

		res, err := sdk.ScheduleTransfers.UpdateCheckWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.CheckWithdrawalScheduleId(t), scheduleUpdate, ascendsdk.String("schedule_details.amount"))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelCheckWithdrawalSchedule", func(t *testing.T) {
		req := components.CancelCheckWithdrawalScheduleRequestCreate{
			Name:    "accounts/" + *fixtures.AccountId() + "/checkWithdrawalSchedules/" + *fixtures.CheckWithdrawalScheduleId(t),
			Comment: ascendsdk.String("canceled due to test"),
		}
		res, err := sdk.ScheduleTransfers.CancelCheckWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.CheckWithdrawalScheduleId(t), req)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})
}
