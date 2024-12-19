// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// LimitPriceCreateType - The type of this price, which must be PRICE_PER_UNIT for equity orders, or PERCENTAGE_OF_PAR for fixed income orders.
type LimitPriceCreateType string

const (
	LimitPriceCreateTypeLimitPriceTypeUnspecified LimitPriceCreateType = "LIMIT_PRICE_TYPE_UNSPECIFIED"
	LimitPriceCreateTypePricePerUnit              LimitPriceCreateType = "PRICE_PER_UNIT"
	LimitPriceCreateTypePercentageOfPar           LimitPriceCreateType = "PERCENTAGE_OF_PAR"
)

func (e LimitPriceCreateType) ToPointer() *LimitPriceCreateType {
	return &e
}

// LimitPriceCreate - A limit price definition
type LimitPriceCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Price DecimalCreate `json:"price"`
	// The type of this price, which must be PRICE_PER_UNIT for equity orders, or PERCENTAGE_OF_PAR for fixed income orders.
	Type LimitPriceCreateType `json:"type"`
}

func (o *LimitPriceCreate) GetPrice() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Price
}

func (o *LimitPriceCreate) GetType() LimitPriceCreateType {
	if o == nil {
		return LimitPriceCreateType("")
	}
	return o.Type
}
