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

type Fixture struct {
	t                        *testing.T
	sdk                      *ascendsdk.SDK
	ctx                      context.Context
	accountId                *string
	bankRelationshipId       *string
	scheduleId               *string
	withdrawalScheduleId     *string
	wireWithdrawalScheduleId *string
}

func (f *Fixture) WireScheduleId(t *testing.T) *string {
	if f.wireWithdrawalScheduleId != nil {
		return f.wireWithdrawalScheduleId
	}

	scheduleId, err := createScheduledWireWithdrawal(f.ctx, *f.AccountId(), *f.sdk)
	require.NoError(t, err)
	return &scheduleId
}

func (f *Fixture) ScheduleId(t *testing.T) *string {
	if f.scheduleId != nil {
		return f.scheduleId
	}

	scheduleId, err := CreateDepositSchedule(t, f.sdk, f.ctx, *f.AccountId(), *f.BankRelationshipId())
	fmt.Println("schedule Id:", scheduleId)
	require.NoError(f.t, err)

	f.scheduleId = &scheduleId
	time.Sleep(5 * time.Second)
	return &scheduleId
}

func (f *Fixture) WithdrawalScheduleId(t *testing.T) *string {
	if f.withdrawalScheduleId != nil {
		return f.withdrawalScheduleId
	}

	withdrawalScheduleId, err := CreateWithdrawalSchedule(t, f.sdk, f.ctx, *f.AccountId(), *f.BankRelationshipId())
	fmt.Println("withdraw schedule Id:", withdrawalScheduleId)
	require.NoError(f.t, err)

	f.withdrawalScheduleId = &withdrawalScheduleId
	time.Sleep(5 * time.Second)
	return &withdrawalScheduleId
}

func (f *Fixture) AccountId() *string {
	if f.accountId != nil {
		return f.accountId
	}
	f.accountId = f.createAndEnrollAccount()
	return f.accountId
}

func (f *Fixture) BankRelationshipId() *string {
	if f.bankRelationshipId != nil {
		return f.bankRelationshipId
	}
	f.bankRelationshipId = f.setupBankRelationship(*f.AccountId())
	return f.bankRelationshipId
}

func (f *Fixture) createAndEnrollAccount() *string {
	accountId, err := helpers.CreateAccountId(f.sdk, f.ctx)
	require.NoError(f.t, err)
	helpers.Wait()

	agg, err := helpers.EnrollAccountIds(f.sdk, f.ctx, *accountId)
	require.NoError(f.t, err)
	helpers.Wait()

	err = helpers.AffirmAgreements(f.sdk, f.ctx, *accountId, agg)
	require.NoError(f.t, err)
	helpers.Wait()

	return accountId
}

func (f *Fixture) setupBankRelationship(accountID string) *string {
	bankRelationshipId, err := helpers.CreateBankRelationship(f.sdk, f.ctx, accountID)
	require.NoError(f.t, err)

	correctMicroDeposits, err := helpers.GetCorrectMicroDeposits(f.sdk, f.ctx, accountID, *bankRelationshipId)
	require.NoError(f.t, err)
	helpers.Wait()

	err = helpers.VerifyMicroDeposits(f.sdk, f.ctx, accountID, *bankRelationshipId, correctMicroDeposits)
	require.NoError(f.t, err)

	return bankRelationshipId
}

func CreateDepositSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) (string, error) {
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
	time.Sleep(3 * time.Second)
	res, err := sdk.ScheduleTransfers.CreateAchDepositSchedule(ctx, accountId, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		scheduleId := strings.Split(*res.AchDepositSchedule.Name, "/")[3]
		return scheduleId, nil
	} else {
		return "", errors.New("Error creating deposit schedule.")
	}
}

func CreateWithdrawalSchedule(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, accountId string, bankRelationshipId string) (string, error) {
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
	time.Sleep(3 * time.Second)
	res, err := sdk.ScheduleTransfers.CreateAchWithdrawalSchedule(ctx, accountId, scheduleCreate)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		scheduleId := strings.Split(*res.AchWithdrawalSchedule.Name, "/")[3]
		return scheduleId, nil
	} else {
		return "", errors.New("Error creating withdrawal schedule.")
	}
}

func createScheduledWireWithdrawal(ctx context.Context, enrolledAccount string, sdk ascendsdk.SDK) (string, error) {
	chicagoLocation, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return "", err
	}
	today := time.Now().In(chicagoLocation)
	year := today.Year()
	month := int(today.Month())
	day := today.Day()

	value := "100"
	occurences := 12

	request := components.WireWithdrawalScheduleCreate{
		Beneficiary: components.WireWithdrawalBeneficiaryCreate{Account: enrolledAccount},
		RecipientBank: components.WireWithdrawalRecipientBankCreate{
			BankID: components.RecipientBankBankIDCreate{
				ID:   "ABNANL2AXXX",
				Type: components.RecipientBankBankIDCreateTypeAba,
			},
		},
		ScheduleDetails: components.WithdrawalScheduleDetailsCreate{
			Amount: &components.DecimalCreate{
				Value: &value,
			},
			ClientScheduleID: uuid.NewString(),
			ScheduleProperties: components.SchedulePropertiesCreate{
				Occurrences: &occurences,
				StartDate: components.DateCreate{
					Year:  &year,
					Month: &month,
					Day:   &day,
				},
				TimeUnit:       components.TimeUnitMonth,
				UnitMultiplier: 1,
			},
		},
	}

	result, err := sdk.ScheduleTransfers.CreateWireWithdrawalSchedule(ctx, enrolledAccount, request)
	if err != nil || result.HTTPMeta.Response.StatusCode != 200 {
		return "", err
	}

	split := strings.Split(*result.WireWithdrawalSchedule.Name, `/`)
	return split[len(split)-1], err
}

func TestScheduleTransfers(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &Fixture{
		t:   t,
		sdk: sdk,
		ctx: ctx,
	}

	t.Run("CreateDepositSchedule", func(t *testing.T) {
		assert.NotNil(t, fixtures.ScheduleId(t))
	})

	t.Run("ListAchDepositSchedules", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.ListAchDepositSchedules(ctx, *fixtures.AccountId(), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("GetAchDepositSchedule", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.GetAchDepositSchedule(ctx, *fixtures.AccountId(), *fixtures.ScheduleId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("UpdateAchDepositSchedule", func(t *testing.T) {
		scheduleUpdate := components.AchDepositScheduleUpdate{
			ScheduleDetails: &components.DepositScheduleDetailsUpdate{
				Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
			},
		}

		res, err := sdk.ScheduleTransfers.UpdateAchDepositSchedule(ctx, *fixtures.AccountId(), *fixtures.ScheduleId(t), scheduleUpdate, ascendsdk.String("schedule_details.amount"))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, "20.00", *res.AchDepositSchedule.ScheduleDetails.Amount.Value)
	})

	t.Run("CancelAchDepositSchedule", func(t *testing.T) {
		req := components.CancelAchDepositScheduleRequestCreate{
			Name:    "accounts/" + *fixtures.AccountId() + "/scheduleTransfers/achDepositSchedules/" + *fixtures.ScheduleId(t),
			Comment: ascendsdk.String("canceled due to test"),
		}
		res, err := sdk.ScheduleTransfers.CancelAchDepositSchedule(ctx, *fixtures.AccountId(), *fixtures.ScheduleId(t), req)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("CreateAchWithdrawalSchedule", func(t *testing.T) {
		assert.NotNil(t, fixtures.WithdrawalScheduleId(t))
	})

	t.Run("ListAchWithdrawalSchedules", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.ListAchWithdrawalSchedules(ctx, *fixtures.AccountId(), nil, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("GetAchWithdrawalSchedule", func(t *testing.T) {
		res, err := sdk.ScheduleTransfers.GetAchWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.WithdrawalScheduleId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("UpdateAchWithdrawalSchedule", func(t *testing.T) {
		scheduleUpdate := components.AchWithdrawalScheduleUpdate{
			ScheduleDetails: &components.WithdrawalScheduleDetailsUpdate{
				Amount: &components.DecimalUpdate{Value: ascendsdk.String("20.00")},
			},
		}

		res, err := sdk.ScheduleTransfers.UpdateAchWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.WithdrawalScheduleId(t), scheduleUpdate, ascendsdk.String("schedule_details.amount"))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.Equal(t, "20.00", *res.AchWithdrawalSchedule.ScheduleDetails.Amount.Value)
	})

	t.Run("CancelAchWithdrawalSchedule", func(t *testing.T) {
		req := components.CancelAchWithdrawalScheduleRequestCreate{
			Name:    "accounts/" + *fixtures.AccountId() + "/scheduleTransfers/achWithdrawalSchedules/" + *fixtures.WithdrawalScheduleId(t),
			Comment: ascendsdk.String("canceled due to test"),
		}
		res, err := sdk.ScheduleTransfers.CancelAchWithdrawalSchedule(ctx, *fixtures.AccountId(), *fixtures.WithdrawalScheduleId(t), req)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Schedule Transfers Transfers Create Wire Withdrawal Schedule Create Wire Withdrawal Schedule1", func(t *testing.T) {
		fmt.Println("wire withdrawal schedule id:", *fixtures.WireScheduleId(t))
		assert.Greater(t, len(*fixtures.WireScheduleId(t)), 0)
	})

	t.Run("Test Schedule Transfers Transfers List Wire Withdrawals Schedules List Wire Withdrawals Schedules1", func(t *testing.T) {
		filter := ""
		pageToken := ""
		pageSize := 10
		result, err := sdk.ScheduleTransfers.ListWireWithdrawalSchedules(ctx, *fixtures.AccountId(), &filter, &pageSize, &pageToken)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Schedule Transfers Transfers Get Wire Withdrawal Schedule Get Wire Withdrawal Schedule1", func(t *testing.T) {
		fixtures.WireScheduleId(t)
		time.Sleep(5 * time.Second)
		result, err := sdk.ScheduleTransfers.GetWireWithdrawalSchedule(fixtures.ctx, *fixtures.AccountId(), *fixtures.WireScheduleId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Schedule Transfers Transfers Update Wire Withdrawal Schedule Update Wire Withdrawal Schedule1", func(t *testing.T) {
		value := "100"
		request := components.WireWithdrawalScheduleUpdate{
			ScheduleDetails: &components.WithdrawalScheduleDetailsUpdate{
				Amount: &components.DecimalUpdate{Value: &value},
			},
		}

		updateMask := "schedule_details.amount"

		result, err := sdk.ScheduleTransfers.UpdateWireWithdrawalSchedule(fixtures.ctx, *fixtures.AccountId(), *fixtures.WireScheduleId(t), request, &updateMask)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Schedule Transfers Transfers Cancel Wire Withdrawal Schedule Cancel Wire Withdrawal Schedule1", func(t *testing.T) {
		request := components.CancelWireWithdrawalScheduleRequestCreate{
			Name: "accounts/" + *fixtures.AccountId() + "/scheduleTransfers/" + *fixtures.WithdrawalScheduleId(t),
		}
		result, err := sdk.ScheduleTransfers.CancelWireWithdrawalSchedule(fixtures.ctx, *fixtures.AccountId(), *fixtures.WireScheduleId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Test Schedule Transfers Transfers List Schedule Summaries", func(t *testing.T) {

		filter, pageToken, pageSize := "", "", 10
		result, err := sdk.ScheduleTransfers.ListScheduleSummaries(fixtures.ctx, &filter, &pageSize, &pageToken)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
