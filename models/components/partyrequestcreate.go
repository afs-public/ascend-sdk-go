// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ProspectusDeliveryPreference - Delivery method instruction for prospectuses for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type ProspectusDeliveryPreference string

const (
	ProspectusDeliveryPreferenceDeliveryPreferenceUnspecified ProspectusDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	ProspectusDeliveryPreferenceDigital                       ProspectusDeliveryPreference = "DIGITAL"
	ProspectusDeliveryPreferencePhysical                      ProspectusDeliveryPreference = "PHYSICAL"
	ProspectusDeliveryPreferenceSuppress                      ProspectusDeliveryPreference = "SUPPRESS"
)

func (e ProspectusDeliveryPreference) ToPointer() *ProspectusDeliveryPreference {
	return &e
}

// ProxyDeliveryPreference - Delivery method instruction for proxy voting for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type ProxyDeliveryPreference string

const (
	ProxyDeliveryPreferenceDeliveryPreferenceUnspecified ProxyDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	ProxyDeliveryPreferenceDigital                       ProxyDeliveryPreference = "DIGITAL"
	ProxyDeliveryPreferencePhysical                      ProxyDeliveryPreference = "PHYSICAL"
	ProxyDeliveryPreferenceSuppress                      ProxyDeliveryPreference = "SUPPRESS"
)

func (e ProxyDeliveryPreference) ToPointer() *ProxyDeliveryPreference {
	return &e
}

// RelationType - Conveys how a person is related to account; Located on each account Party record; Examples are `PRIMARY_OWNER`, `JOINT_OWNER`, `EXECUTOR`, etc.
type RelationType string

const (
	RelationTypePartyRelationTypeUnspecified RelationType = "PARTY_RELATION_TYPE_UNSPECIFIED"
	RelationTypePrimaryOwner                 RelationType = "PRIMARY_OWNER"
	RelationTypeJointOwner                   RelationType = "JOINT_OWNER"
	RelationTypeCustodian                    RelationType = "CUSTODIAN"
	RelationTypeExecutor                     RelationType = "EXECUTOR"
	RelationTypeAuthorizedSigner             RelationType = "AUTHORIZED_SIGNER"
	RelationTypeBeneficialOwner              RelationType = "BENEFICIAL_OWNER"
	RelationTypeControlPerson                RelationType = "CONTROL_PERSON"
	RelationTypeAuthorizedRepresentative     RelationType = "AUTHORIZED_REPRESENTATIVE"
	RelationTypeTrustee                      RelationType = "TRUSTEE"
	RelationTypeAuthTrusteeRep               RelationType = "AUTH_TRUSTEE_REP"
)

func (e RelationType) ToPointer() *RelationType {
	return &e
}

// PartyRequestCreateStatementDeliveryPreference - Delivery method instruction for account statements for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestCreateStatementDeliveryPreference string

const (
	PartyRequestCreateStatementDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestCreateStatementDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestCreateStatementDeliveryPreferenceDigital                       PartyRequestCreateStatementDeliveryPreference = "DIGITAL"
	PartyRequestCreateStatementDeliveryPreferencePhysical                      PartyRequestCreateStatementDeliveryPreference = "PHYSICAL"
	PartyRequestCreateStatementDeliveryPreferenceSuppress                      PartyRequestCreateStatementDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestCreateStatementDeliveryPreference) ToPointer() *PartyRequestCreateStatementDeliveryPreference {
	return &e
}

// TaxDocumentDeliveryPreference - Delivery method instruction for tax documents for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated; Per regulation, selected tax forms will be mailed by regulation regardless of this setting
type TaxDocumentDeliveryPreference string

const (
	TaxDocumentDeliveryPreferenceDeliveryPreferenceUnspecified TaxDocumentDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	TaxDocumentDeliveryPreferenceDigital                       TaxDocumentDeliveryPreference = "DIGITAL"
	TaxDocumentDeliveryPreferencePhysical                      TaxDocumentDeliveryPreference = "PHYSICAL"
	TaxDocumentDeliveryPreferenceSuppress                      TaxDocumentDeliveryPreference = "SUPPRESS"
)

func (e TaxDocumentDeliveryPreference) ToPointer() *TaxDocumentDeliveryPreference {
	return &e
}

// PartyRequestCreateTradeConfirmationDeliveryPreference - Delivery method instruction for trade confirmations for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestCreateTradeConfirmationDeliveryPreference string

const (
	PartyRequestCreateTradeConfirmationDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestCreateTradeConfirmationDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestCreateTradeConfirmationDeliveryPreferenceDigital                       PartyRequestCreateTradeConfirmationDeliveryPreference = "DIGITAL"
	PartyRequestCreateTradeConfirmationDeliveryPreferencePhysical                      PartyRequestCreateTradeConfirmationDeliveryPreference = "PHYSICAL"
	PartyRequestCreateTradeConfirmationDeliveryPreferenceSuppress                      PartyRequestCreateTradeConfirmationDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestCreateTradeConfirmationDeliveryPreference) ToPointer() *PartyRequestCreateTradeConfirmationDeliveryPreference {
	return &e
}

// PartyRequestCreate - A single record representing an owner or manager of an Account. Contains fully populated Party Identity object.
type PartyRequestCreate struct {
	// An email address indicated for account communications.
	EmailAddress string `json:"email_address"`
	// Legal entity ID.
	LegalEntityID *string `json:"legal_entity_id,omitempty"`
	// Legal natural person ID.
	LegalNaturalPersonID *string `json:"legal_natural_person_id,omitempty"`
	// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).
	//
	//  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.
	//
	//  Advice on address input / editing: - Use an i18n-ready address widget such as  https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of  fields outside countries where that field is used.
	//
	//  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
	MailingAddress PostalAddressCreate `json:"mailing_address"`
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
	PhoneNumber PhoneNumberCreate `json:"phone_number"`
	// Delivery method instruction for prospectuses for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	ProspectusDeliveryPreference *ProspectusDeliveryPreference `json:"prospectus_delivery_preference,omitempty"`
	// Delivery method instruction for proxy voting for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	ProxyDeliveryPreference *ProxyDeliveryPreference `json:"proxy_delivery_preference,omitempty"`
	// Conveys how a person is related to account; Located on each account Party record; Examples are `PRIMARY_OWNER`, `JOINT_OWNER`, `EXECUTOR`, etc.
	RelationType RelationType `json:"relation_type"`
	// Delivery method instruction for account statements for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	StatementDeliveryPreference *PartyRequestCreateStatementDeliveryPreference `json:"statement_delivery_preference,omitempty"`
	// Delivery method instruction for tax documents for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated; Per regulation, selected tax forms will be mailed by regulation regardless of this setting
	TaxDocumentDeliveryPreference *TaxDocumentDeliveryPreference `json:"tax_document_delivery_preference,omitempty"`
	// Delivery method instruction for trade confirmations for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	TradeConfirmationDeliveryPreference *PartyRequestCreateTradeConfirmationDeliveryPreference `json:"trade_confirmation_delivery_preference,omitempty"`
}

func (o *PartyRequestCreate) GetEmailAddress() string {
	if o == nil {
		return ""
	}
	return o.EmailAddress
}

func (o *PartyRequestCreate) GetLegalEntityID() *string {
	if o == nil {
		return nil
	}
	return o.LegalEntityID
}

func (o *PartyRequestCreate) GetLegalNaturalPersonID() *string {
	if o == nil {
		return nil
	}
	return o.LegalNaturalPersonID
}

func (o *PartyRequestCreate) GetMailingAddress() PostalAddressCreate {
	if o == nil {
		return PostalAddressCreate{}
	}
	return o.MailingAddress
}

func (o *PartyRequestCreate) GetPhoneNumber() PhoneNumberCreate {
	if o == nil {
		return PhoneNumberCreate{}
	}
	return o.PhoneNumber
}

func (o *PartyRequestCreate) GetProspectusDeliveryPreference() *ProspectusDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.ProspectusDeliveryPreference
}

func (o *PartyRequestCreate) GetProxyDeliveryPreference() *ProxyDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.ProxyDeliveryPreference
}

func (o *PartyRequestCreate) GetRelationType() RelationType {
	if o == nil {
		return RelationType("")
	}
	return o.RelationType
}

func (o *PartyRequestCreate) GetStatementDeliveryPreference() *PartyRequestCreateStatementDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.StatementDeliveryPreference
}

func (o *PartyRequestCreate) GetTaxDocumentDeliveryPreference() *TaxDocumentDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TaxDocumentDeliveryPreference
}

func (o *PartyRequestCreate) GetTradeConfirmationDeliveryPreference() *PartyRequestCreateTradeConfirmationDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TradeConfirmationDeliveryPreference
}
