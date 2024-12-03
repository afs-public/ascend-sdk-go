// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// PartyRequestUpdateProspectusDeliveryPreference - Delivery method instruction for prospectuses for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestUpdateProspectusDeliveryPreference string

const (
	PartyRequestUpdateProspectusDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestUpdateProspectusDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestUpdateProspectusDeliveryPreferenceDigital                       PartyRequestUpdateProspectusDeliveryPreference = "DIGITAL"
	PartyRequestUpdateProspectusDeliveryPreferencePhysical                      PartyRequestUpdateProspectusDeliveryPreference = "PHYSICAL"
	PartyRequestUpdateProspectusDeliveryPreferenceSuppress                      PartyRequestUpdateProspectusDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestUpdateProspectusDeliveryPreference) ToPointer() *PartyRequestUpdateProspectusDeliveryPreference {
	return &e
}
func (e *PartyRequestUpdateProspectusDeliveryPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELIVERY_PREFERENCE_UNSPECIFIED":
		fallthrough
	case "DIGITAL":
		fallthrough
	case "PHYSICAL":
		fallthrough
	case "SUPPRESS":
		*e = PartyRequestUpdateProspectusDeliveryPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateProspectusDeliveryPreference: %v", v)
	}
}

// PartyRequestUpdateProxyDeliveryPreference - Delivery method instruction for proxy voting for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestUpdateProxyDeliveryPreference string

const (
	PartyRequestUpdateProxyDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestUpdateProxyDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestUpdateProxyDeliveryPreferenceDigital                       PartyRequestUpdateProxyDeliveryPreference = "DIGITAL"
	PartyRequestUpdateProxyDeliveryPreferencePhysical                      PartyRequestUpdateProxyDeliveryPreference = "PHYSICAL"
	PartyRequestUpdateProxyDeliveryPreferenceSuppress                      PartyRequestUpdateProxyDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestUpdateProxyDeliveryPreference) ToPointer() *PartyRequestUpdateProxyDeliveryPreference {
	return &e
}
func (e *PartyRequestUpdateProxyDeliveryPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELIVERY_PREFERENCE_UNSPECIFIED":
		fallthrough
	case "DIGITAL":
		fallthrough
	case "PHYSICAL":
		fallthrough
	case "SUPPRESS":
		*e = PartyRequestUpdateProxyDeliveryPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateProxyDeliveryPreference: %v", v)
	}
}

// PartyRequestUpdateRelationType - Conveys how a person is related to account; Located on each account Party record; Examples are `PRIMARY_OWNER`, `JOINT_OWNER`, `EXECUTOR`, etc.
type PartyRequestUpdateRelationType string

const (
	PartyRequestUpdateRelationTypePartyRelationTypeUnspecified PartyRequestUpdateRelationType = "PARTY_RELATION_TYPE_UNSPECIFIED"
	PartyRequestUpdateRelationTypePrimaryOwner                 PartyRequestUpdateRelationType = "PRIMARY_OWNER"
	PartyRequestUpdateRelationTypeJointOwner                   PartyRequestUpdateRelationType = "JOINT_OWNER"
	PartyRequestUpdateRelationTypeCustodian                    PartyRequestUpdateRelationType = "CUSTODIAN"
	PartyRequestUpdateRelationTypeGuardianConservator          PartyRequestUpdateRelationType = "GUARDIAN_CONSERVATOR"
	PartyRequestUpdateRelationTypePowerOfAttorney              PartyRequestUpdateRelationType = "POWER_OF_ATTORNEY"
	PartyRequestUpdateRelationTypeExecutor                     PartyRequestUpdateRelationType = "EXECUTOR"
	PartyRequestUpdateRelationTypeAuthorizedSigner             PartyRequestUpdateRelationType = "AUTHORIZED_SIGNER"
	PartyRequestUpdateRelationTypeBeneficialOwner              PartyRequestUpdateRelationType = "BENEFICIAL_OWNER"
	PartyRequestUpdateRelationTypeControlPerson                PartyRequestUpdateRelationType = "CONTROL_PERSON"
	PartyRequestUpdateRelationTypeAuthorizedRepresentative     PartyRequestUpdateRelationType = "AUTHORIZED_REPRESENTATIVE"
	PartyRequestUpdateRelationTypeTrustee                      PartyRequestUpdateRelationType = "TRUSTEE"
	PartyRequestUpdateRelationTypeAuthTrusteeRep               PartyRequestUpdateRelationType = "AUTH_TRUSTEE_REP"
)

func (e PartyRequestUpdateRelationType) ToPointer() *PartyRequestUpdateRelationType {
	return &e
}
func (e *PartyRequestUpdateRelationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARTY_RELATION_TYPE_UNSPECIFIED":
		fallthrough
	case "PRIMARY_OWNER":
		fallthrough
	case "JOINT_OWNER":
		fallthrough
	case "CUSTODIAN":
		fallthrough
	case "GUARDIAN_CONSERVATOR":
		fallthrough
	case "POWER_OF_ATTORNEY":
		fallthrough
	case "EXECUTOR":
		fallthrough
	case "AUTHORIZED_SIGNER":
		fallthrough
	case "BENEFICIAL_OWNER":
		fallthrough
	case "CONTROL_PERSON":
		fallthrough
	case "AUTHORIZED_REPRESENTATIVE":
		fallthrough
	case "TRUSTEE":
		fallthrough
	case "AUTH_TRUSTEE_REP":
		*e = PartyRequestUpdateRelationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateRelationType: %v", v)
	}
}

// PartyRequestUpdateStatementDeliveryPreference - Delivery method instruction for account statements for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestUpdateStatementDeliveryPreference string

const (
	PartyRequestUpdateStatementDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestUpdateStatementDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestUpdateStatementDeliveryPreferenceDigital                       PartyRequestUpdateStatementDeliveryPreference = "DIGITAL"
	PartyRequestUpdateStatementDeliveryPreferencePhysical                      PartyRequestUpdateStatementDeliveryPreference = "PHYSICAL"
	PartyRequestUpdateStatementDeliveryPreferenceSuppress                      PartyRequestUpdateStatementDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestUpdateStatementDeliveryPreference) ToPointer() *PartyRequestUpdateStatementDeliveryPreference {
	return &e
}
func (e *PartyRequestUpdateStatementDeliveryPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELIVERY_PREFERENCE_UNSPECIFIED":
		fallthrough
	case "DIGITAL":
		fallthrough
	case "PHYSICAL":
		fallthrough
	case "SUPPRESS":
		*e = PartyRequestUpdateStatementDeliveryPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateStatementDeliveryPreference: %v", v)
	}
}

// PartyRequestUpdateTaxDocumentDeliveryPreference - Delivery method instruction for tax documents for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated; Per regulation, selected tax forms will be mailed by regulation regardless of this setting
type PartyRequestUpdateTaxDocumentDeliveryPreference string

const (
	PartyRequestUpdateTaxDocumentDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestUpdateTaxDocumentDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestUpdateTaxDocumentDeliveryPreferenceDigital                       PartyRequestUpdateTaxDocumentDeliveryPreference = "DIGITAL"
	PartyRequestUpdateTaxDocumentDeliveryPreferencePhysical                      PartyRequestUpdateTaxDocumentDeliveryPreference = "PHYSICAL"
	PartyRequestUpdateTaxDocumentDeliveryPreferenceSuppress                      PartyRequestUpdateTaxDocumentDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestUpdateTaxDocumentDeliveryPreference) ToPointer() *PartyRequestUpdateTaxDocumentDeliveryPreference {
	return &e
}
func (e *PartyRequestUpdateTaxDocumentDeliveryPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELIVERY_PREFERENCE_UNSPECIFIED":
		fallthrough
	case "DIGITAL":
		fallthrough
	case "PHYSICAL":
		fallthrough
	case "SUPPRESS":
		*e = PartyRequestUpdateTaxDocumentDeliveryPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateTaxDocumentDeliveryPreference: %v", v)
	}
}

// PartyRequestUpdateTradeConfirmationDeliveryPreference - Delivery method instruction for trade confirmations for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
type PartyRequestUpdateTradeConfirmationDeliveryPreference string

const (
	PartyRequestUpdateTradeConfirmationDeliveryPreferenceDeliveryPreferenceUnspecified PartyRequestUpdateTradeConfirmationDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	PartyRequestUpdateTradeConfirmationDeliveryPreferenceDigital                       PartyRequestUpdateTradeConfirmationDeliveryPreference = "DIGITAL"
	PartyRequestUpdateTradeConfirmationDeliveryPreferencePhysical                      PartyRequestUpdateTradeConfirmationDeliveryPreference = "PHYSICAL"
	PartyRequestUpdateTradeConfirmationDeliveryPreferenceSuppress                      PartyRequestUpdateTradeConfirmationDeliveryPreference = "SUPPRESS"
)

func (e PartyRequestUpdateTradeConfirmationDeliveryPreference) ToPointer() *PartyRequestUpdateTradeConfirmationDeliveryPreference {
	return &e
}
func (e *PartyRequestUpdateTradeConfirmationDeliveryPreference) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELIVERY_PREFERENCE_UNSPECIFIED":
		fallthrough
	case "DIGITAL":
		fallthrough
	case "PHYSICAL":
		fallthrough
	case "SUPPRESS":
		*e = PartyRequestUpdateTradeConfirmationDeliveryPreference(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyRequestUpdateTradeConfirmationDeliveryPreference: %v", v)
	}
}

// PartyRequestUpdate - A single record representing an owner or manager of an Account. Contains fully populated Party Identity object.
type PartyRequestUpdate struct {
	// An email address indicated for account communications.
	EmailAddress *string `json:"email_address,omitempty"`
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
	MailingAddress *PostalAddressUpdate `json:"mailing_address,omitempty"`
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
	PhoneNumber *PhoneNumberUpdate `json:"phone_number,omitempty"`
	// Delivery method instruction for prospectuses for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	ProspectusDeliveryPreference *PartyRequestUpdateProspectusDeliveryPreference `json:"prospectus_delivery_preference,omitempty"`
	// Delivery method instruction for proxy voting for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	ProxyDeliveryPreference *PartyRequestUpdateProxyDeliveryPreference `json:"proxy_delivery_preference,omitempty"`
	// Conveys how a person is related to account; Located on each account Party record; Examples are `PRIMARY_OWNER`, `JOINT_OWNER`, `EXECUTOR`, etc.
	RelationType *PartyRequestUpdateRelationType `json:"relation_type,omitempty"`
	// Delivery method instruction for account statements for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	StatementDeliveryPreference *PartyRequestUpdateStatementDeliveryPreference `json:"statement_delivery_preference,omitempty"`
	// Delivery method instruction for tax documents for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated; Per regulation, selected tax forms will be mailed by regulation regardless of this setting
	TaxDocumentDeliveryPreference *PartyRequestUpdateTaxDocumentDeliveryPreference `json:"tax_document_delivery_preference,omitempty"`
	// Delivery method instruction for trade confirmations for a given Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `DIGITAL` on account creation but may be updated
	TradeConfirmationDeliveryPreference *PartyRequestUpdateTradeConfirmationDeliveryPreference `json:"trade_confirmation_delivery_preference,omitempty"`
}

func (o *PartyRequestUpdate) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *PartyRequestUpdate) GetLegalEntityID() *string {
	if o == nil {
		return nil
	}
	return o.LegalEntityID
}

func (o *PartyRequestUpdate) GetLegalNaturalPersonID() *string {
	if o == nil {
		return nil
	}
	return o.LegalNaturalPersonID
}

func (o *PartyRequestUpdate) GetMailingAddress() *PostalAddressUpdate {
	if o == nil {
		return nil
	}
	return o.MailingAddress
}

func (o *PartyRequestUpdate) GetPhoneNumber() *PhoneNumberUpdate {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *PartyRequestUpdate) GetProspectusDeliveryPreference() *PartyRequestUpdateProspectusDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.ProspectusDeliveryPreference
}

func (o *PartyRequestUpdate) GetProxyDeliveryPreference() *PartyRequestUpdateProxyDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.ProxyDeliveryPreference
}

func (o *PartyRequestUpdate) GetRelationType() *PartyRequestUpdateRelationType {
	if o == nil {
		return nil
	}
	return o.RelationType
}

func (o *PartyRequestUpdate) GetStatementDeliveryPreference() *PartyRequestUpdateStatementDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.StatementDeliveryPreference
}

func (o *PartyRequestUpdate) GetTaxDocumentDeliveryPreference() *PartyRequestUpdateTaxDocumentDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TaxDocumentDeliveryPreference
}

func (o *PartyRequestUpdate) GetTradeConfirmationDeliveryPreference() *PartyRequestUpdateTradeConfirmationDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TradeConfirmationDeliveryPreference
}
