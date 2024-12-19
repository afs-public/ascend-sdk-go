// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CommissionCreateType - The type of commission value being specified. Only the type of "AMOUNT" is supported.
type CommissionCreateType string

const (
	CommissionCreateTypeCommissionTypeUnspecified CommissionCreateType = "COMMISSION_TYPE_UNSPECIFIED"
	CommissionCreateTypeAmount                    CommissionCreateType = "AMOUNT"
)

func (e CommissionCreateType) ToPointer() *CommissionCreateType {
	return &e
}

// CommissionCreate - A custom commission applied to an order
type CommissionCreate struct {
	// The type of commission value being specified. Only the type of "AMOUNT" is supported.
	Type CommissionCreateType `json:"type"`
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Value DecimalCreate `json:"value"`
}

func (o *CommissionCreate) GetType() CommissionCreateType {
	if o == nil {
		return CommissionCreateType("")
	}
	return o.Type
}

func (o *CommissionCreate) GetValue() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Value
}
