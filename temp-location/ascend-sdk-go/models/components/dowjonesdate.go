// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DowJonesDate - Custom date object to present related Dow Jones dates because any value can be null, Month is represented randomly like June, Jun, 6 etc. etc.
type DowJonesDate struct {
	// Day
	Day *string `json:"day,omitempty"`
	// Month
	Month *string `json:"month,omitempty"`
	// Year
	Year *string `json:"year,omitempty"`
}

func (o *DowJonesDate) GetDay() *string {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *DowJonesDate) GetMonth() *string {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *DowJonesDate) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}
