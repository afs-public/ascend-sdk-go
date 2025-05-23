// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// BeneficiaryCreateEntityType - The entity type of the beneficiary; Required if the beneficiary is a legal entity.
type BeneficiaryCreateEntityType string

const (
	BeneficiaryCreateEntityTypeEntityTypeUnspecified   BeneficiaryCreateEntityType = "ENTITY_TYPE_UNSPECIFIED"
	BeneficiaryCreateEntityTypeCorporation             BeneficiaryCreateEntityType = "CORPORATION"
	BeneficiaryCreateEntityTypeLimitedLiabilityCompany BeneficiaryCreateEntityType = "LIMITED_LIABILITY_COMPANY"
	BeneficiaryCreateEntityTypePartnership             BeneficiaryCreateEntityType = "PARTNERSHIP"
	BeneficiaryCreateEntityTypeTrust                   BeneficiaryCreateEntityType = "TRUST"
	BeneficiaryCreateEntityTypeEstate                  BeneficiaryCreateEntityType = "ESTATE"
)

func (e BeneficiaryCreateEntityType) ToPointer() *BeneficiaryCreateEntityType {
	return &e
}

// BeneficiaryCreateRelationType - The relationship of the beneficiary to the account owner
type BeneficiaryCreateRelationType string

const (
	BeneficiaryCreateRelationTypeRelationTypeUnspecified BeneficiaryCreateRelationType = "RELATION_TYPE_UNSPECIFIED"
	BeneficiaryCreateRelationTypeSpouse                  BeneficiaryCreateRelationType = "SPOUSE"
	BeneficiaryCreateRelationTypeTrust                   BeneficiaryCreateRelationType = "TRUST"
	BeneficiaryCreateRelationTypeOther                   BeneficiaryCreateRelationType = "OTHER"
)

func (e BeneficiaryCreateRelationType) ToPointer() *BeneficiaryCreateRelationType {
	return &e
}

// BeneficiaryCreateTaxIDType - The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN. Tax id type is required if birth date is not provided.
type BeneficiaryCreateTaxIDType string

const (
	BeneficiaryCreateTaxIDTypeTaxIDTypeUnspecified BeneficiaryCreateTaxIDType = "TAX_ID_TYPE_UNSPECIFIED"
	BeneficiaryCreateTaxIDTypeTaxIDTypeSsn         BeneficiaryCreateTaxIDType = "TAX_ID_TYPE_SSN"
	BeneficiaryCreateTaxIDTypeTaxIDTypeItin        BeneficiaryCreateTaxIDType = "TAX_ID_TYPE_ITIN"
	BeneficiaryCreateTaxIDTypeTaxIDTypeEin         BeneficiaryCreateTaxIDType = "TAX_ID_TYPE_EIN"
)

func (e BeneficiaryCreateTaxIDType) ToPointer() *BeneficiaryCreateTaxIDType {
	return &e
}

// BeneficiaryCreate - The beneficiary for transfer on death.
type BeneficiaryCreate struct {
	// An integer conveying the percentage of interest the related Beneficiary has in the account if the owner(s) become deceased; The sum of all beneficiary percentages must equal "100"
	BeneficiaryPercentage int `json:"beneficiary_percentage"`
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	BirthDate *DateCreate `json:"birth_date,omitempty"`
	// Beneficiaries may provide an email, a mailing_address, or both An email address indicated for account communications
	Email *string `json:"email,omitempty"`
	// The legal entity name; Required if the beneficiary is a legal entity.
	EntityName *string `json:"entity_name,omitempty"`
	// The entity type of the beneficiary; Required if the beneficiary is a legal entity.
	EntityType *BeneficiaryCreateEntityType `json:"entity_type,omitempty"`
	// Family name of a natural person; Required if the beneficiary is a natural person.
	FamilyName *string `json:"family_name,omitempty"`
	// The given name of a natural person; Conventionally known as 'first name' in most English-speaking countries.Required if the beneficiary is a natural person.
	GivenName *string `json:"given_name,omitempty"`
	// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).
	//
	//  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.
	//
	//  Advice on address input / editing: - Use an i18n-ready address widget such as  https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of  fields outside countries where that field is used.
	//
	//  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
	MailingAddress *PostalAddressCreate `json:"mailing_address,omitempty"`
	// Non-primary names representing a natural person; Name attributed to a person other than "Given" and "Family" names.
	MiddleNames *string `json:"middle_names,omitempty"`
	// An object representing a phone number, suitable as an API wire format.
	//
	//  This representation:
	//
	//  - should not be used for locale-specific formatting of a phone number, such  as "+1 (650) 253-0000 ext. 123"
	//
	//  - is not designed for efficient storage - may not be suitable for dialing - specialized libraries (see references)  should be used to parse the number for that purpose
	//
	//  To do something meaningful with this number, such as format it for various use-cases, convert it to an `i18n.phonenumbers.PhoneNumber` object first.
	//
	//  For instance, in Java this would be:
	//
	//   com.google.type.PhoneNumber wireProto =    com.google.type.PhoneNumber.newBuilder().build();  com.google.i18n.phonenumbers.Phonenumber.PhoneNumber phoneNumber =    PhoneNumberUtil.getInstance().parse(wireProto.getE164Number(), "ZZ");  if (!wireProto.getExtension().isEmpty()) {   phoneNumber.setExtension(wireProto.getExtension());  }
	//
	//  Reference(s):
	//   - https://github.com/google/libphonenumber
	PhoneNumber *PhoneNumberCreate `json:"phone_number,omitempty"`
	// The relationship of the beneficiary to the account owner
	RelationType BeneficiaryCreateRelationType `json:"relation_type"`
	// The full U.S. tax ID for a related person; Tax ID is required if birth date is not provided.
	TaxID *string `json:"tax_id,omitempty"`
	// The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN. Tax id type is required if birth date is not provided.
	TaxIDType *BeneficiaryCreateTaxIDType `json:"tax_id_type,omitempty"`
}

func (o *BeneficiaryCreate) GetBeneficiaryPercentage() int {
	if o == nil {
		return 0
	}
	return o.BeneficiaryPercentage
}

func (o *BeneficiaryCreate) GetBirthDate() *DateCreate {
	if o == nil {
		return nil
	}
	return o.BirthDate
}

func (o *BeneficiaryCreate) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *BeneficiaryCreate) GetEntityName() *string {
	if o == nil {
		return nil
	}
	return o.EntityName
}

func (o *BeneficiaryCreate) GetEntityType() *BeneficiaryCreateEntityType {
	if o == nil {
		return nil
	}
	return o.EntityType
}

func (o *BeneficiaryCreate) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *BeneficiaryCreate) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *BeneficiaryCreate) GetMailingAddress() *PostalAddressCreate {
	if o == nil {
		return nil
	}
	return o.MailingAddress
}

func (o *BeneficiaryCreate) GetMiddleNames() *string {
	if o == nil {
		return nil
	}
	return o.MiddleNames
}

func (o *BeneficiaryCreate) GetPhoneNumber() *PhoneNumberCreate {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *BeneficiaryCreate) GetRelationType() BeneficiaryCreateRelationType {
	if o == nil {
		return BeneficiaryCreateRelationType("")
	}
	return o.RelationType
}

func (o *BeneficiaryCreate) GetTaxID() *string {
	if o == nil {
		return nil
	}
	return o.TaxID
}

func (o *BeneficiaryCreate) GetTaxIDType() *BeneficiaryCreateTaxIDType {
	if o == nil {
		return nil
	}
	return o.TaxIDType
}
