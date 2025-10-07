package orders

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	ascendsdk "github.com/afs-public/ascend-sdk-go"

	"github.com/afs-public/ascend-sdk-go/tests/helpers"
	"github.com/google/uuid"

	"github.com/afs-public/ascend-sdk-go/models/components"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Fixtures struct {
	t         *testing.T
	sdk       *ascendsdk.SDK
	ctx       context.Context
	accountId string
	orderId   *string
}

func (f *Fixtures) OrderId(t *testing.T) *string {
	if f.orderId != nil {
		return f.orderId
	}

	orderId, err := CreateOrder(t, f.sdk, f.ctx, f.accountId)

	fmt.Println("orderId", orderId)
	require.NoError(f.t, err)

	f.orderId = &orderId

	return &orderId
}

func CreateOrder(t *testing.T, sdk *ascendsdk.SDK, ctx context.Context, enrolledAccountId string) (string, error) {

	s := sdk

	// Fund Account with Credit
	transfersCreditCreate := components.TransfersCreditCreate{
		Amount:           components.DecimalCreate{Value: ascendsdk.String("1000000.00")},
		ClientTransferID: uuid.New().String(),
		Description:      ascendsdk.String("Credit given as promotion"),
		Type:             components.TransfersCreditCreateTypePromotional,
	}

	s.FeesAndCredits.CreateCredit(ctx, enrolledAccountId, transfersCreditCreate)

	// Create Order
	create := components.OrderCreate{
		AssetType:     components.AssetTypeEquity,
		ClientOrderID: uuid.New().String(),
		OrderDate: components.DateCreate{
			Year:  ascendsdk.Int(time.Now().Year()),
			Month: ascendsdk.Int(int(time.Now().Month())),
			Day:   ascendsdk.Int(time.Now().Day()),
		},
		Identifier:     "SBUX",
		IdentifierType: components.IdentifierTypeSymbol,
		Quantity:       &components.DecimalCreate{Value: ascendsdk.String("1")},
		OrderType:      components.OrderTypeLimit,
		LimitPrice: &components.LimitPriceCreate{
			Price: components.DecimalCreate{Value: ascendsdk.String("5.00")},
			Type:  components.LimitPriceCreateTypePricePerUnit,
		},
		Side:        components.SideBuy,
		TimeInForce: components.TimeInForceDay,
	}

	res, err := s.CreateOrder.CreateOrder(ctx, enrolledAccountId, create)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	if res.HTTPMeta.Response.StatusCode == 200 {
		return *res.Order.OrderID, nil
	} else {
		return "", errors.New("Error creating order")
	}
}

func TestOrderService(t *testing.T) {
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

	agreements, enrollErr := helpers.EnrollAccountIds(sdk, ctx, *accountId)
	if enrollErr != nil {
		t.Fatalf("Error enrolling account: %v", enrollErr)
	}

	helpers.AffirmAgreements(sdk, ctx, *accountId, agreements)

	t.Run("CreateOrder", func(t *testing.T) {
		assert.NotNil(t, fixtures.OrderId(t))
	})

	t.Run("GetOrder", func(t *testing.T) {
		res, err := sdk.CreateOrder.GetOrder(ctx, fixtures.accountId, *fixtures.OrderId(t))
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	})

	t.Run("CancelOrder", func(t *testing.T) {
		request := components.CancelOrderRequestCreate{
			Name: "accounts/" + fixtures.accountId + "/orders/" + *fixtures.OrderId(t),
		}
		res, err := sdk.CreateOrder.CancelOrder(ctx, fixtures.accountId, *fixtures.OrderId(t), request)
		require.NoError(t, err)
		assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
		assert.NotNil(t, res.Order.OrderID)
	})

}
