// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AddressDetail - Address detail used for Dow Jones Profile details
type AddressDetail struct {
	// Dow Jones persons address city
	AddressCity *string `json:"address_city,omitempty"`
	// Dow Jones persons address line
	AddressLine *string `json:"address_line,omitempty"`
	// Dow Jones persons address administrative area
	AdministrativeArea *string `json:"administrative_area,omitempty"`
	// Dow Jones persons address postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// Two character region code, complies with https://cldr.unicode.org/index Example values: "US", "CA"
	RegionCode *string `json:"region_code,omitempty"`
}

func (o *AddressDetail) GetAddressCity() *string {
	if o == nil {
		return nil
	}
	return o.AddressCity
}

func (o *AddressDetail) GetAddressLine() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine
}

func (o *AddressDetail) GetAdministrativeArea() *string {
	if o == nil {
		return nil
	}
	return o.AdministrativeArea
}

func (o *AddressDetail) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AddressDetail) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
