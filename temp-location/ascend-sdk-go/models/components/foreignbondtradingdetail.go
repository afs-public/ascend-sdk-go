// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Percentage - The percentage of the account's trades which will involve foreign bond
type Percentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Percentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// ForeignBondTradingDetail - Foreign bond trading detail
type ForeignBondTradingDetail struct {
	// The percentage of the account's trades which will involve foreign bond
	Percentage *Percentage `json:"percentage,omitempty"`
	// The region where the foreign bond trading activity is taking place. Must be a two-character CLDR code.
	RegionCode *string `json:"region_code,omitempty"`
}

func (o *ForeignBondTradingDetail) GetPercentage() *Percentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

func (o *ForeignBondTradingDetail) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
