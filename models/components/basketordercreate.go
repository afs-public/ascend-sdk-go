// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/utils"
)

// BasketOrderCreateAssetType - The type of the asset in this order
type BasketOrderCreateAssetType string

const (
	BasketOrderCreateAssetTypeAssetTypeUnspecified BasketOrderCreateAssetType = "ASSET_TYPE_UNSPECIFIED"
	BasketOrderCreateAssetTypeEquity               BasketOrderCreateAssetType = "EQUITY"
	BasketOrderCreateAssetTypeMutualFund           BasketOrderCreateAssetType = "MUTUAL_FUND"
)

func (e BasketOrderCreateAssetType) ToPointer() *BasketOrderCreateAssetType {
	return &e
}

// BasketOrderCreateIdentifierType - The identifier type of the asset being ordered. For Equities: only SYMBOL is supported For Mutual Funds: only SYMBOL and CUSIP are supported
type BasketOrderCreateIdentifierType string

const (
	BasketOrderCreateIdentifierTypeSymbol BasketOrderCreateIdentifierType = "SYMBOL"
	BasketOrderCreateIdentifierTypeCusip  BasketOrderCreateIdentifierType = "CUSIP"
	BasketOrderCreateIdentifierTypeIsin   BasketOrderCreateIdentifierType = "ISIN"
)

func (e BasketOrderCreateIdentifierType) ToPointer() *BasketOrderCreateIdentifierType {
	return &e
}
func (e *BasketOrderCreateIdentifierType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SYMBOL":
		fallthrough
	case "CUSIP":
		fallthrough
	case "ISIN":
		*e = BasketOrderCreateIdentifierType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BasketOrderCreateIdentifierType: %v", v)
	}
}

// BasketOrderCreateOrderType - The execution type of this order.
type BasketOrderCreateOrderType string

const (
	BasketOrderCreateOrderTypeOrderTypeUnspecified BasketOrderCreateOrderType = "ORDER_TYPE_UNSPECIFIED"
	BasketOrderCreateOrderTypeMarket               BasketOrderCreateOrderType = "MARKET"
)

func (e BasketOrderCreateOrderType) ToPointer() *BasketOrderCreateOrderType {
	return &e
}

// BasketOrderCreateSide - The side of this order.
type BasketOrderCreateSide string

const (
	BasketOrderCreateSideSideUnspecified BasketOrderCreateSide = "SIDE_UNSPECIFIED"
	BasketOrderCreateSideBuy             BasketOrderCreateSide = "BUY"
	BasketOrderCreateSideSell            BasketOrderCreateSide = "SELL"
)

func (e BasketOrderCreateSide) ToPointer() *BasketOrderCreateSide {
	return &e
}

// BasketOrderCreateTimeInForce - Must be the value "DAY". Regulatory requirements dictate that the system capture the intended time_in_force, which is why this a mandatory field.
type BasketOrderCreateTimeInForce string

const (
	BasketOrderCreateTimeInForceTimeInForceUnspecified BasketOrderCreateTimeInForce = "TIME_IN_FORCE_UNSPECIFIED"
	BasketOrderCreateTimeInForceDay                    BasketOrderCreateTimeInForce = "DAY"
)

func (e BasketOrderCreateTimeInForce) ToPointer() *BasketOrderCreateTimeInForce {
	return &e
}

// BasketOrderCreate - The message describing an order that has been added to a basket
type BasketOrderCreate struct {
	// The identifier of the account transacting this order
	AccountID string `json:"account_id"`
	// The type of the asset in this order
	AssetType BasketOrderCreateAssetType `json:"asset_type"`
	// User-supplied unique order ID. Cannot be more than 40 characters long.
	ClientOrderID string `json:"client_order_id"`
	// Time the order request was received by the client
	ClientOrderReceivedTime *time.Time `json:"client_order_received_time,omitempty"`
	// Defaults to "USD". Only "USD" is supported. Full list of currency codes is defined at: https://en.wikipedia.org/wiki/ISO_4217
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Fees that will be applied to this order.
	Fees []FeeCreate `json:"fees,omitempty"`
	// Identifier of the asset (of the type specified in `identifier_type`).
	Identifier string `json:"identifier"`
	// The identifier type of the asset being ordered. For Equities: only SYMBOL is supported For Mutual Funds: only SYMBOL and CUSIP are supported
	IdentifierType BasketOrderCreateIdentifierType `json:"identifier_type"`
	// Letter of Intent (LOI). An LOI allows investors to receive sales charge discounts based on a commitment to buy a specified monetary amount of shares over a period of time, usually 13 months.
	LetterOfIntent *LetterOfIntentCreate `json:"letter_of_intent,omitempty"`
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	NotionalValue *DecimalCreate `json:"notional_value,omitempty"`
	// The execution type of this order.
	OrderType BasketOrderCreateOrderType `json:"order_type"`
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Quantity *DecimalCreate `json:"quantity,omitempty"`
	// Rights of Accumulation (ROA). An ROA allows an investor to aggregate their own fund shares with the holdings of certain related parties toward achieving the investment thresholds at which sales charge discounts become available.
	RightsOfAccumulation *RightsOfAccumulationCreate `json:"rights_of_accumulation,omitempty"`
	// The side of this order.
	Side BasketOrderCreateSide `json:"side"`
	// Must be the value "DAY". Regulatory requirements dictate that the system capture the intended time_in_force, which is why this a mandatory field.
	TimeInForce BasketOrderCreateTimeInForce `json:"time_in_force"`
}

func (b BasketOrderCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BasketOrderCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BasketOrderCreate) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BasketOrderCreate) GetAssetType() BasketOrderCreateAssetType {
	if o == nil {
		return BasketOrderCreateAssetType("")
	}
	return o.AssetType
}

func (o *BasketOrderCreate) GetClientOrderID() string {
	if o == nil {
		return ""
	}
	return o.ClientOrderID
}

func (o *BasketOrderCreate) GetClientOrderReceivedTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClientOrderReceivedTime
}

func (o *BasketOrderCreate) GetCurrencyCode() *string {
	if o == nil {
		return nil
	}
	return o.CurrencyCode
}

func (o *BasketOrderCreate) GetFees() []FeeCreate {
	if o == nil {
		return nil
	}
	return o.Fees
}

func (o *BasketOrderCreate) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *BasketOrderCreate) GetIdentifierType() BasketOrderCreateIdentifierType {
	if o == nil {
		return BasketOrderCreateIdentifierType("")
	}
	return o.IdentifierType
}

func (o *BasketOrderCreate) GetLetterOfIntent() *LetterOfIntentCreate {
	if o == nil {
		return nil
	}
	return o.LetterOfIntent
}

func (o *BasketOrderCreate) GetNotionalValue() *DecimalCreate {
	if o == nil {
		return nil
	}
	return o.NotionalValue
}

func (o *BasketOrderCreate) GetOrderType() BasketOrderCreateOrderType {
	if o == nil {
		return BasketOrderCreateOrderType("")
	}
	return o.OrderType
}

func (o *BasketOrderCreate) GetQuantity() *DecimalCreate {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *BasketOrderCreate) GetRightsOfAccumulation() *RightsOfAccumulationCreate {
	if o == nil {
		return nil
	}
	return o.RightsOfAccumulation
}

func (o *BasketOrderCreate) GetSide() BasketOrderCreateSide {
	if o == nil {
		return BasketOrderCreateSide("")
	}
	return o.Side
}

func (o *BasketOrderCreate) GetTimeInForce() BasketOrderCreateTimeInForce {
	if o == nil {
		return BasketOrderCreateTimeInForce("")
	}
	return o.TimeInForce
}
