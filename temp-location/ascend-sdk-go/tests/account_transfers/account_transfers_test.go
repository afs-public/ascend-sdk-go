package account_transfers

import (
	ascendsdk "ascend-sdk"
	"ascend-sdk/tests/helpers"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestAccoutTransfers_getAccountTransfer(t *testing.T) {

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)
	ctx := context.Background()

	accountTransferId := testCreateAccountTransfer(t, *sdk, ctx)

	_, err = sdk.AccountTransfers.GetTransfer(ctx, os.Getenv("correspondent_id"), os.Getenv("account_id"), accountTransferId)

	require.NoError(t, err)

}

func testCreateAccountTransfer(t *testing.T, sdk ascendsdk.SDK, ctx context.Context) string {

	request := components.TransferCreate{
		Deliverer: components.TransferAccountCreate{
			ExternalAccount: &components.ExternalAccountCreate{
				AccountNumber:     "1234567890",
				ParticipantNumber: "987",
			},
		},
	}

	res, err := sdk.AccountTransfers.CreateTransfer(ctx, os.Getenv("correspondent_id"), os.Getenv("account_id"), request, nil)
	require.NoError(t, err)
	assert.NotNil(t, res.AcatsTransfer)

	name := res.AcatsTransfer.Name
	parts := strings.Split(*name, "/")
	accountTransferId := &parts[len(parts)-1]
	return *accountTransferId

}
