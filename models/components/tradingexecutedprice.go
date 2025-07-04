// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Price - The limit price which must be greater than zero if provided. For equity orders in the USD currency, up to 2 decimal places are allowed for prices above $1 and up to 4 decimal places for prices at or below $1. For fixed income orders this is expressed as a percentage of par, which allows up to 5 decimal places in the USD currency.
type Price struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Price) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// TradingExecutedPriceType - The type of this price, which must be PRICE_PER_UNIT for equity orders, or PERCENTAGE_OF_PAR for fixed income orders.
type TradingExecutedPriceType string

const (
	TradingExecutedPriceTypePricePerUnit    TradingExecutedPriceType = "PRICE_PER_UNIT"
	TradingExecutedPriceTypePercentageOfPar TradingExecutedPriceType = "PERCENTAGE_OF_PAR"
	TradingExecutedPriceTypeYieldToWorst    TradingExecutedPriceType = "YIELD_TO_WORST"
	TradingExecutedPriceTypeYieldToMaturity TradingExecutedPriceType = "YIELD_TO_MATURITY"
)

func (e TradingExecutedPriceType) ToPointer() *TradingExecutedPriceType {
	return &e
}

// TradingExecutedPrice - An average price definition
type TradingExecutedPrice struct {
	// The limit price which must be greater than zero if provided. For equity orders in the USD currency, up to 2 decimal places are allowed for prices above $1 and up to 4 decimal places for prices at or below $1. For fixed income orders this is expressed as a percentage of par, which allows up to 5 decimal places in the USD currency.
	Price *Price `json:"price,omitempty"`
	// The type of this price, which must be PRICE_PER_UNIT for equity orders, or PERCENTAGE_OF_PAR for fixed income orders.
	Type *TradingExecutedPriceType `json:"type,omitempty"`
}

func (o *TradingExecutedPrice) GetPrice() *Price {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *TradingExecutedPrice) GetType() *TradingExecutedPriceType {
	if o == nil {
		return nil
	}
	return o.Type
}
