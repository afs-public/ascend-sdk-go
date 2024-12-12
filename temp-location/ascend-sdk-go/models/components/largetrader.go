// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EffectiveDate - The date on which the trader meets or exceeds the large trader reporting threshold, which is defined by the U.S. Securities and Exchange Commission (SEC) as trades of 2 million shares or $20 million in a single day or 20 million shares or $200 million during a calendar month
type EffectiveDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *EffectiveDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *EffectiveDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *EffectiveDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LargeTrader - A large trader.
type LargeTrader struct {
	// The date on which the trader meets or exceeds the large trader reporting threshold, which is defined by the U.S. Securities and Exchange Commission (SEC) as trades of 2 million shares or $20 million in a single day or 20 million shares or $200 million during a calendar month
	EffectiveDate *EffectiveDate `json:"effective_date,omitempty"`
	// SEC-issued ID signifying the person/entity as a large trader; Required for CAIS regulatory reporting.
	LargeTraderID *string `json:"large_trader_id,omitempty"`
}

func (o *LargeTrader) GetEffectiveDate() *EffectiveDate {
	if o == nil {
		return nil
	}
	return o.EffectiveDate
}

func (o *LargeTrader) GetLargeTraderID() *string {
	if o == nil {
		return nil
	}
	return o.LargeTraderID
}
