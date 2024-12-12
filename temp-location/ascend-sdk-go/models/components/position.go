// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AdjustedSettled - `settled` + any as of settled amounts for the date
type AdjustedSettled struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AdjustedSettled) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// AdjustedTrade - `trade` + any as of trade amounts for the date
type AdjustedTrade struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AdjustedTrade) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Date - The date for which the positions were calculated
type Date struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *Date) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *Date) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *Date) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// PositionFpsl - Quantity of asset in use by the FPSL program. Should not be used by currency assets
type PositionFpsl struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PositionFpsl) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Free - Quantity of asset available for allocation for use by the FPSL program. Raw bucket values. These denote that a position is allocated to this purpose. Values may be negative
type Free struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Free) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// LastAdjustedDate - The most recent date an adjustment occurred
type LastAdjustedDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *LastAdjustedDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *LastAdjustedDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *LastAdjustedDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// PendingDrip - Quantity of currency from a dividend being reserved for reinvestment. should not be used by non-currency assets
type PendingDrip struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PendingDrip) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PendingOutgoingAcat - Quantity/ amount of asset restricted due to an outgoing acat request
type PendingOutgoingAcat struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PendingOutgoingAcat) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PendingWithdrawal - Quantity of currency being reserved for withdrawal. should not be used by non-currency assets
type PendingWithdrawal struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PendingWithdrawal) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Settled - Computed fieldsOriginal Settled Position before and as-of changesComputed based on the bucket values to represet the total settled position in an account  Currently defined as `free` + `fpsl` + `pending_outgoing_acat` + `drip` + `pending_withdrawal`, but if/when new buckets are added this value will need to change to reflect them
type Settled struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Settled) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PositionTrade - original trade position
type PositionTrade struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *PositionTrade) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Unrestricted - Computed based on the bucket values to represent the total unrestricted position in an account. Will always be less than or equal to `settled`  settled - (pending_outgoing_acat + pending_drip + pending_withdrawal) ; however, if/when the API adds new buckets, Apex may adjust this to either incorporate the new value or not
type Unrestricted struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Unrestricted) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// Position
type Position struct {
	// A globally unique identifier referencing a single account; this is the main identifier for an account used for machine-to-machine interactions
	AccountID *string `json:"account_id,omitempty"`
	// `settled` + any as of settled amounts for the date
	AdjustedSettled *AdjustedSettled `json:"adjusted_settled,omitempty"`
	// `trade` + any as of trade amounts for the date
	AdjustedTrade *AdjustedTrade `json:"adjusted_trade,omitempty"`
	// An Apex-provided, global identifier created on a per asset bases which provides connectivity across all areas
	AssetID *string `json:"asset_id,omitempty"`
	// The correspondent id associated with the account for the position
	CorrespondentID *string `json:"correspondent_id,omitempty"`
	// The date for which the positions were calculated
	Date *Date `json:"date,omitempty"`
	// Quantity of asset in use by the FPSL program. Should not be used by currency assets
	Fpsl *PositionFpsl `json:"fpsl,omitempty"`
	// Quantity of asset available for allocation for use by the FPSL program. Raw bucket values. These denote that a position is allocated to this purpose. Values may be negative
	Free *Free `json:"free,omitempty"`
	// The most recent date an adjustment occurred
	LastAdjustedDate *LastAdjustedDate `json:"last_adjusted_date,omitempty"`
	// accounts/{account_id}/positions/{position_id}
	Name *string `json:"name,omitempty"`
	// Quantity of currency from a dividend being reserved for reinvestment. should not be used by non-currency assets
	PendingDrip *PendingDrip `json:"pending_drip,omitempty"`
	// Quantity/ amount of asset restricted due to an outgoing acat request
	PendingOutgoingAcat *PendingOutgoingAcat `json:"pending_outgoing_acat,omitempty"`
	// Quantity of currency being reserved for withdrawal. should not be used by non-currency assets
	PendingWithdrawal *PendingWithdrawal `json:"pending_withdrawal,omitempty"`
	// Computed fieldsOriginal Settled Position before and as-of changesComputed based on the bucket values to represet the total settled position in an account  Currently defined as `free` + `fpsl` + `pending_outgoing_acat` + `drip` + `pending_withdrawal`, but if/when new buckets are added this value will need to change to reflect them
	Settled *Settled `json:"settled,omitempty"`
	// original trade position
	Trade *PositionTrade `json:"trade,omitempty"`
	// Computed based on the bucket values to represent the total unrestricted position in an account. Will always be less than or equal to `settled`  settled - (pending_outgoing_acat + pending_drip + pending_withdrawal) ; however, if/when the API adds new buckets, Apex may adjust this to either incorporate the new value or not
	Unrestricted *Unrestricted `json:"unrestricted,omitempty"`
}

func (o *Position) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *Position) GetAdjustedSettled() *AdjustedSettled {
	if o == nil {
		return nil
	}
	return o.AdjustedSettled
}

func (o *Position) GetAdjustedTrade() *AdjustedTrade {
	if o == nil {
		return nil
	}
	return o.AdjustedTrade
}

func (o *Position) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *Position) GetCorrespondentID() *string {
	if o == nil {
		return nil
	}
	return o.CorrespondentID
}

func (o *Position) GetDate() *Date {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Position) GetFpsl() *PositionFpsl {
	if o == nil {
		return nil
	}
	return o.Fpsl
}

func (o *Position) GetFree() *Free {
	if o == nil {
		return nil
	}
	return o.Free
}

func (o *Position) GetLastAdjustedDate() *LastAdjustedDate {
	if o == nil {
		return nil
	}
	return o.LastAdjustedDate
}

func (o *Position) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Position) GetPendingDrip() *PendingDrip {
	if o == nil {
		return nil
	}
	return o.PendingDrip
}

func (o *Position) GetPendingOutgoingAcat() *PendingOutgoingAcat {
	if o == nil {
		return nil
	}
	return o.PendingOutgoingAcat
}

func (o *Position) GetPendingWithdrawal() *PendingWithdrawal {
	if o == nil {
		return nil
	}
	return o.PendingWithdrawal
}

func (o *Position) GetSettled() *Settled {
	if o == nil {
		return nil
	}
	return o.Settled
}

func (o *Position) GetTrade() *PositionTrade {
	if o == nil {
		return nil
	}
	return o.Trade
}

func (o *Position) GetUnrestricted() *Unrestricted {
	if o == nil {
		return nil
	}
	return o.Unrestricted
}
