// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ForeignIdentificationUpdateType - Identification type
type ForeignIdentificationUpdateType string

const (
	ForeignIdentificationUpdateTypeIdentificationTypeUnspecified ForeignIdentificationUpdateType = "IDENTIFICATION_TYPE_UNSPECIFIED"
	ForeignIdentificationUpdateTypePassport                      ForeignIdentificationUpdateType = "PASSPORT"
	ForeignIdentificationUpdateTypeNationalID                    ForeignIdentificationUpdateType = "NATIONAL_ID"
	ForeignIdentificationUpdateTypeDriversLicense                ForeignIdentificationUpdateType = "DRIVERS_LICENSE"
)

func (e ForeignIdentificationUpdateType) ToPointer() *ForeignIdentificationUpdateType {
	return &e
}

// ForeignIdentificationUpdate - Foreign identification
type ForeignIdentificationUpdate struct {
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	ExpirationDate *DateUpdate `json:"expiration_date,omitempty"`
	// Denotes if the identification is a tax id or other
	Ftin *bool `json:"ftin,omitempty"`
	// Identification number
	IdentificationNumber *string `json:"identification_number,omitempty"`
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	IssueDate *DateUpdate `json:"issue_date,omitempty"`
	// Region of issuance must be provided as a two-character CLDR country code
	IssuingRegionCode *string `json:"issuing_region_code,omitempty"`
	// Identification type
	Type *ForeignIdentificationUpdateType `json:"type,omitempty"`
}

func (o *ForeignIdentificationUpdate) GetExpirationDate() *DateUpdate {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

func (o *ForeignIdentificationUpdate) GetFtin() *bool {
	if o == nil {
		return nil
	}
	return o.Ftin
}

func (o *ForeignIdentificationUpdate) GetIdentificationNumber() *string {
	if o == nil {
		return nil
	}
	return o.IdentificationNumber
}

func (o *ForeignIdentificationUpdate) GetIssueDate() *DateUpdate {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *ForeignIdentificationUpdate) GetIssuingRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.IssuingRegionCode
}

func (o *ForeignIdentificationUpdate) GetType() *ForeignIdentificationUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
