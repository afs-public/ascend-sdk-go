// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// LetterOfIntentCreate - Letter of Intent (LOI). An LOI allows investors to receive sales charge discounts based on a commitment to buy a specified monetary amount of shares over a period of time, usually 13 months.
type LetterOfIntentCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount DecimalCreate `json:"amount"`
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	PeriodStartDate DateCreate `json:"period_start_date"`
}

func (o *LetterOfIntentCreate) GetAmount() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Amount
}

func (o *LetterOfIntentCreate) GetPeriodStartDate() DateCreate {
	if o == nil {
		return DateCreate{}
	}
	return o.PeriodStartDate
}
