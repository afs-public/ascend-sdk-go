// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// NetAmount - The net amount of the trade in USD. This value is always positive.
type NetAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *NetAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CalculateCashBalanceResponseTradeSummary - A summary of a trade.
type CalculateCashBalanceResponseTradeSummary struct {
	// The ledger activity for the trade.
	Activity *string `json:"activity,omitempty"`
	// The asset that was traded.
	Asset *string `json:"asset,omitempty"`
	// The net amount of the trade in USD. This value is always positive.
	NetAmount *NetAmount `json:"net_amount,omitempty"`
}

func (o *CalculateCashBalanceResponseTradeSummary) GetActivity() *string {
	if o == nil {
		return nil
	}
	return o.Activity
}

func (o *CalculateCashBalanceResponseTradeSummary) GetAsset() *string {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *CalculateCashBalanceResponseTradeSummary) GetNetAmount() *NetAmount {
	if o == nil {
		return nil
	}
	return o.NetAmount
}
