// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ProvidedIdentityVerificationUpdate - Clients must supply data and confirmation attesting to identity verification.
type ProvidedIdentityVerificationUpdate struct {
	// Indicates whether the identity's address was verified
	AddressVerified *bool `json:"address_verified,omitempty"`
	// Indicates whether the identity's date of birth was verified
	BirthDateVerified *bool `json:"birth_date_verified,omitempty"`
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	ExecutionDate *DateUpdate `json:"execution_date,omitempty"`
	// Client-generated identifier associated with the KYC results for the appropriate case
	ExternalCaseID *string `json:"external_case_id,omitempty"`
	// A collection of unique identifiers provided by the documents api that correspond to any number of identity verification documents used in support of the external vendor to verify the identity, such as a driver's license, passport, etc
	IdentityVerificationDocumentIds []string `json:"identity_verification_document_ids,omitempty"`
	// Indicates whether the identity's name was verified
	NameVerified *bool `json:"name_verified,omitempty"`
	// Id of this identity verification record
	ProvidedIdentityVerificationID *string `json:"provided_identity_verification_id,omitempty"`
	// A unique identifier provided from the documents api that corresponds to an identity verification result
	RawVendorDataDocumentID *string `json:"raw_vendor_data_document_id,omitempty"`
	// Indicates whether the identity's tax id was verified
	TaxIDVerified *bool `json:"tax_id_verified,omitempty"`
	// Name of the vendor that performed identity verification
	Vendor *string `json:"vendor,omitempty"`
}

func (o *ProvidedIdentityVerificationUpdate) GetAddressVerified() *bool {
	if o == nil {
		return nil
	}
	return o.AddressVerified
}

func (o *ProvidedIdentityVerificationUpdate) GetBirthDateVerified() *bool {
	if o == nil {
		return nil
	}
	return o.BirthDateVerified
}

func (o *ProvidedIdentityVerificationUpdate) GetExecutionDate() *DateUpdate {
	if o == nil {
		return nil
	}
	return o.ExecutionDate
}

func (o *ProvidedIdentityVerificationUpdate) GetExternalCaseID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalCaseID
}

func (o *ProvidedIdentityVerificationUpdate) GetIdentityVerificationDocumentIds() []string {
	if o == nil {
		return nil
	}
	return o.IdentityVerificationDocumentIds
}

func (o *ProvidedIdentityVerificationUpdate) GetNameVerified() *bool {
	if o == nil {
		return nil
	}
	return o.NameVerified
}

func (o *ProvidedIdentityVerificationUpdate) GetProvidedIdentityVerificationID() *string {
	if o == nil {
		return nil
	}
	return o.ProvidedIdentityVerificationID
}

func (o *ProvidedIdentityVerificationUpdate) GetRawVendorDataDocumentID() *string {
	if o == nil {
		return nil
	}
	return o.RawVendorDataDocumentID
}

func (o *ProvidedIdentityVerificationUpdate) GetTaxIDVerified() *bool {
	if o == nil {
		return nil
	}
	return o.TaxIDVerified
}

func (o *ProvidedIdentityVerificationUpdate) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}