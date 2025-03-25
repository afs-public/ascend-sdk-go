package transfers

import (
	"context"
	"strings"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const withdrawal_account_id = "01JHGTEPC6ZTAHCFRH2MD3VJJT"
const deceased_account_id = "01JHK07CRQ9X8P5XE9JWG4PFSP"

func createCashJournal(ctx context.Context, sdk ascendsdk.SDK, accountId *string) (string, error) {
	oneStr := "1.00"

	request := components.CashJournalCreate{
		ClientTransferID:   uuid.NewString(),
		DestinationAccount: "accounts/" + *accountId,
		Amount: &components.DecimalCreate{
			Value: &oneStr,
		},
		SourceAccount: "accounts/" + withdrawal_account_id,
	}

	result, err := sdk.Journals.CreateCashJournal(ctx, request)

	if err != nil {
		return "", err
	}

	split := strings.Split(*result.CashJournal.Name, `/`)
	journalId := split[len(split)-1]
	return journalId, err
}

func TestCashJournals(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	accountId, err := helpers.CreateAccountId(sdk, ctx)

	time.Sleep(5 * time.Second)

	cashJournalId, err := createCashJournal(ctx, *sdk, accountId)

	require.NoError(t, err)

	time.Sleep(5 * time.Second)

	t.Run("Cash Journals Transfers Create Cash Journal Create Cash Journal1", func(t *testing.T) {
		assert.Greater(t, len(cashJournalId), 0)
	})

	t.Run("Cash Journals Transfers Get Cash Journal Get Cash Journal1", func(t *testing.T) {
		result, err := sdk.Journals.GetCashJournal(ctx, cashJournalId)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Cash Journals Transfers Cancel Cash Journal Cancel Cash Journal1", func(t *testing.T) {
		reason := "Test cancel cash journal"

		request := components.CancelCashJournalRequestCreate{
			Name:   "cashJournals/" + cashJournalId,
			Reason: &reason,
		}

		result, err := sdk.Journals.CancelCashJournal(ctx, cashJournalId, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Cash Journals Transfers Retrieve Cash Journal Constraints Retrieve Cash Journal Constraints1", func(t *testing.T) {
		request := components.RetrieveCashJournalConstraintsRequestCreate{
			DestinationAccount: "accounts/" + deceased_account_id,
			SourceAccount:      "accounts/" + withdrawal_account_id,
		}

		result, err := sdk.Journals.RetrieveCashJournalConstraints(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})

	t.Run("Cash Journals Transfers Retrieve Cash Journal Party Retrieve Cash Journal Party1", func(t *testing.T) {
		request := components.CheckPartyTypeRequestCreate{
			DestinationAccount: "accounts/" + deceased_account_id,
			SourceAccount:      "accounts/" + withdrawal_account_id,
		}

		result, err := sdk.Journals.CheckPartyType(ctx, request)
		require.NoError(t, err)
		assert.Equal(t, 200, result.HTTPMeta.Response.StatusCode)
	})
}
