package orders

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type OptionOrderFixtures struct {
	t             *testing.T
	sdk           *ascendsdk.SDK
	ctx           context.Context
	accountId     string
	optionOrderId *string
}

func (f *OptionOrderFixtures) OptionOrderId(t *testing.T) *string {
	if f.optionOrderId != nil {
		return f.optionOrderId
	}

	optionOrderId, err := CreateOptionOrder(t, f.sdk, f.ctx, f.accountId)
	fmt.Println("optionOrderId", optionOrderId)
	require.NoError(t, err)

	f.optionOrderId = &optionOrderId

	return &optionOrderId
}

func CreateOptionOrder(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, enrolledAccountId string) (string, error) {
	s := sdk

	// Fund Account with Credit
	transfersCreditCreate := components.TransfersCreditCreate{
		Amount:           components.DecimalCreate{Value: ascendsdk.String("1000000.00")},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Credit given as promotion"),
		Type:             components.TransfersCreditCreateTypePromotional,
	}

	s.FeesAndCredits.CreateCredit(ctx, enrolledAccountId, transfersCreditCreate)

	// Create Option Order
	create := components.OptionOrderCreate{
		BrokerCapacity: components.OptionOrderCreateBrokerCapacityAgency,
		ClientOrderID:  uuid.New().String(),
		CurrencyCode:   "USD",
		Legs: []components.OptionOrderLegCreate{
			{
				AssetType:      components.OptionOrderLegCreateAssetTypeOption,
				Identifier:     "SBUX250221C00084000",
				IdentifierType: components.OptionOrderLegCreateIdentifierTypeOsi,
				RatioQuantity:  1,
				Side:           components.OptionOrderLegCreateSideBuy,
				SideModifier:   components.SideModifierOpen,
			},
		},
		LimitPrice: components.DecimalCreate{Value: ascendsdk.String("5.00")},
		OrderDate: components.DateCreate{
			Year:  ascendsdk.Int(time.Now().Year()),
			Month: ascendsdk.Int(int(time.Now().Month())),
			Day:   ascendsdk.Int(time.Now().Day()),
		},
		OrderType:      components.OptionOrderCreateOrderTypeLimit,
		PriceDirection: components.PriceDirectionDebit,
		Quantity:       components.DecimalCreate{Value: ascendsdk.String("1")},
		TimeInForce:    components.OptionOrderCreateTimeInForceDay,
	}

	res, err := s.OptionOrders.CreateOptionOrder(ctx, enrolledAccountId, create)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		return *res.OptionOrder.OptionOrderID, nil
	} else {
		return "", errors.New("Error creating option order")
	}
}

func TestOptionOrderService(t *testing.T) {
	ctx := context.Background()

	sdk, err := helpers.SetupAscendSDK()
	require.NoError(t, err)

	fixtures := &OptionOrderFixtures{
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

	agreements, enrollErr := helpers.EnrollAccountIds(sdk, ctx, *accountId)
	if enrollErr != nil {
		t.Fatalf("Error enrolling account: %v", enrollErr)
	}

	helpers.AffirmAgreements(sdk, ctx, *accountId, agreements)

	t.Run("CreateOptionOrder", func(t *testing.T) {
		assert.NotNil(t, fixtures.OptionOrderId(t))
	})

	t.Run("GetOptionOrder", func(t *testing.T) {
		res, err := sdk.OptionOrders.GetOptionOrder(ctx, fixtures.accountId, *fixtures.OptionOrderId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelOptionOrder", func(t *testing.T) {
		// Check if the order is in a cancelable state first
		getRes, getErr := sdk.OptionOrders.GetOptionOrder(ctx, fixtures.accountId, *fixtures.OptionOrderId(t))
		require.NoError(t, getErr)

		status := *getRes.OptionOrder.OrderStatus
		if status == components.OptionOrderOrderStatusRejected ||
			status == components.OptionOrderOrderStatusFilled ||
			status == components.OptionOrderOrderStatusCanceled {
			t.Skipf("Skipping: option order is already in terminal status %s", status)
		}

		request := components.CancelOptionOrderRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/optionOrders/" + *fixtures.OptionOrderId(t),
		}
		res, err := sdk.OptionOrders.CancelOptionOrder(ctx, fixtures.accountId, *fixtures.OptionOrderId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res.OptionOrder.OptionOrderID)
	})
}
