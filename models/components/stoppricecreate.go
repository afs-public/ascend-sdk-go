// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// StopPriceCreateType - The type of this price, which must PRICE_PER_UNIT for equity orders. (Fixed income and mutual fund assets do not support stop orders.)
type StopPriceCreateType string

const (
	StopPriceCreateTypeStopPriceTypeUnspecified StopPriceCreateType = "STOP_PRICE_TYPE_UNSPECIFIED"
	StopPriceCreateTypePricePerUnit             StopPriceCreateType = "PRICE_PER_UNIT"
)

func (e StopPriceCreateType) ToPointer() *StopPriceCreateType {
	return &e
}
func (e *StopPriceCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STOP_PRICE_TYPE_UNSPECIFIED":
		fallthrough
	case "PRICE_PER_UNIT":
		*e = StopPriceCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StopPriceCreateType: %v", v)
	}
}

// StopPriceCreate - A stop price definition
type StopPriceCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Price DecimalCreate `json:"price"`
	// The type of this price, which must PRICE_PER_UNIT for equity orders. (Fixed income and mutual fund assets do not support stop orders.)
	Type StopPriceCreateType `json:"type"`
}

func (o *StopPriceCreate) GetPrice() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Price
}

func (o *StopPriceCreate) GetType() StopPriceCreateType {
	if o == nil {
		return StopPriceCreateType("")
	}
	return o.Type
}