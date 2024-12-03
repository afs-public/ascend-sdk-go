// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"ascend-sdk/internal/utils"
	"time"
)

// AccruedInterestAmount - The amount of accrued interest exchanged in this execution. Will only be present for orders of Fixed Income assets.
type AccruedInterestAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AccruedInterestAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// GrossCreditAmount - The net currency amount exchanged in this transaction, in the order currency. Will only be present for orders of Fixed Income assets.
type GrossCreditAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *GrossCreditAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PrevailingMarketPrice - The prevailing market price of the asset, without fees or commissions. Will only be present for orders of Fixed Income assets.
type PrevailingMarketPrice struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PrevailingMarketPrice) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Quantity - The quantity of the order. For Equities: measured in shares. For Fixed Income assets: measured in the face value of the currency of the order.
type Quantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Quantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Executions - Details of an individual execution within an order
type Executions struct {
	// The amount of accrued interest exchanged in this execution. Will only be present for orders of Fixed Income assets.
	AccruedInterestAmount *AccruedInterestAmount `json:"accrued_interest_amount,omitempty"`
	// The prices at which the order was executed. For Equities: there will be one price measured in PRICE_PER_UNIT (using the order currency). For Fixed Income assets: there will always be an entry measured in the PERCENTAGE_OF_PAR (100 X cost / total par value), and there may be additional entries measured in yield.
	ExecutedPrices []ExecutedPrice `json:"executed_prices,omitempty"`
	// The timestamp that this fill was transacted at the market
	ExecutedTime *time.Time `json:"executed_time,omitempty"`
	// The net currency amount exchanged in this transaction, in the order currency. Will only be present for orders of Fixed Income assets.
	GrossCreditAmount *GrossCreditAmount `json:"gross_credit_amount,omitempty"`
	// The prevailing market price of the asset, without fees or commissions. Will only be present for orders of Fixed Income assets.
	PrevailingMarketPrice *PrevailingMarketPrice `json:"prevailing_market_price,omitempty"`
	// The quantity of the order. For Equities: measured in shares. For Fixed Income assets: measured in the face value of the currency of the order.
	Quantity *Quantity `json:"quantity,omitempty"`
}

func (e Executions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Executions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Executions) GetAccruedInterestAmount() *AccruedInterestAmount {
	if o == nil {
		return nil
	}
	return o.AccruedInterestAmount
}

func (o *Executions) GetExecutedPrices() []ExecutedPrice {
	if o == nil {
		return nil
	}
	return o.ExecutedPrices
}

func (o *Executions) GetExecutedTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExecutedTime
}

func (o *Executions) GetGrossCreditAmount() *GrossCreditAmount {
	if o == nil {
		return nil
	}
	return o.GrossCreditAmount
}

func (o *Executions) GetPrevailingMarketPrice() *PrevailingMarketPrice {
	if o == nil {
		return nil
	}
	return o.PrevailingMarketPrice
}

func (o *Executions) GetQuantity() *Quantity {
	if o == nil {
		return nil
	}
	return o.Quantity
}
