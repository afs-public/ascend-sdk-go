// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ExpectedNotionalCeilingAmount - The notional value the order is not reasonably expected to exceed in USD. This value is always positive.
type ExpectedNotionalCeilingAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *ExpectedNotionalCeilingAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CalculateCashBalanceResponseOpenOrderSummary - A summary of an open order.
type CalculateCashBalanceResponseOpenOrderSummary struct {
	// The asset for the open order.
	Asset *string `json:"asset,omitempty"`
	// The notional value the order is not reasonably expected to exceed in USD. This value is always positive.
	ExpectedNotionalCeilingAmount *ExpectedNotionalCeilingAmount `json:"expected_notional_ceiling_amount,omitempty"`
}

func (o *CalculateCashBalanceResponseOpenOrderSummary) GetAsset() *string {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *CalculateCashBalanceResponseOpenOrderSummary) GetExpectedNotionalCeilingAmount() *ExpectedNotionalCeilingAmount {
	if o == nil {
		return nil
	}
	return o.ExpectedNotionalCeilingAmount
}
