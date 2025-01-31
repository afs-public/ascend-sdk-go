// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EntityUpdate - investigation details on an entity
type EntityUpdate struct {
	// Other names the entity is known by (Doing Business As)
	DbaNames []string `json:"dba_names,omitempty"`
	// Email addresses
	EmailAddresses []string `json:"email_addresses,omitempty"`
	// Identification details including id value, and type (e.g. ein, lei)
	Identifications []EntityIdentificationUpdate `json:"identifications,omitempty"`
	// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).
	//
	//  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.
	//
	//  Advice on address input / editing: - Use an i18n-ready address widget such as  https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of  fields outside countries where that field is used.
	//
	//  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
	LegalAddress *PostalAddressUpdate `json:"legal_address,omitempty"`
	// The legal name of the entity
	LegalName *string `json:"legal_name,omitempty"`
	// mailing address
	MailingAddresses []PostalAddressUpdate `json:"mailing_addresses,omitempty"`
	// The countries where an entity does business Two character region code, complies with https://cldr.unicode.org/index Example values: "US", "CA"
	OperatingRegionCodes []string `json:"operating_region_codes,omitempty"`
	// phone numbers
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
	// Region of registration Two character region code, complies with https://cldr.unicode.org/index Example values: "US", "CA"
	RegistrationRegion *string `json:"registration_region,omitempty"`
}

func (o *EntityUpdate) GetDbaNames() []string {
	if o == nil {
		return nil
	}
	return o.DbaNames
}

func (o *EntityUpdate) GetEmailAddresses() []string {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *EntityUpdate) GetIdentifications() []EntityIdentificationUpdate {
	if o == nil {
		return nil
	}
	return o.Identifications
}

func (o *EntityUpdate) GetLegalAddress() *PostalAddressUpdate {
	if o == nil {
		return nil
	}
	return o.LegalAddress
}

func (o *EntityUpdate) GetLegalName() *string {
	if o == nil {
		return nil
	}
	return o.LegalName
}

func (o *EntityUpdate) GetMailingAddresses() []PostalAddressUpdate {
	if o == nil {
		return nil
	}
	return o.MailingAddresses
}

func (o *EntityUpdate) GetOperatingRegionCodes() []string {
	if o == nil {
		return nil
	}
	return o.OperatingRegionCodes
}

func (o *EntityUpdate) GetPhoneNumbers() []string {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *EntityUpdate) GetRegistrationRegion() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationRegion
}
