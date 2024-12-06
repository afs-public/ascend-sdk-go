// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CommissionAmount - Monetary amount associated with the commission
type CommissionAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CommissionAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type Commission struct {
	// Monetary amount associated with the commission
	Amount *CommissionAmount `json:"amount,omitempty"`
}

func (o *Commission) GetAmount() *CommissionAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}