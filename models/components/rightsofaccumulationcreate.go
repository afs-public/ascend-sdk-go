// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RightsOfAccumulationCreate - Rights of Accumulation (ROA). An ROA allows an investor to aggregate their own fund shares with the holdings of certain related parties toward achieving the investment thresholds at which sales charge discounts become available.
type RightsOfAccumulationCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount DecimalCreate `json:"amount"`
}

func (o *RightsOfAccumulationCreate) GetAmount() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Amount
}
