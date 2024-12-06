// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// InterestedPartyUpdateStatementDeliveryPreference - Delivery method instruction for account statements for a given Interested Party; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
type InterestedPartyUpdateStatementDeliveryPreference string

const (
	InterestedPartyUpdateStatementDeliveryPreferenceDeliveryPreferenceUnspecified InterestedPartyUpdateStatementDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	InterestedPartyUpdateStatementDeliveryPreferenceDigital                       InterestedPartyUpdateStatementDeliveryPreference = "DIGITAL"
	InterestedPartyUpdateStatementDeliveryPreferencePhysical                      InterestedPartyUpdateStatementDeliveryPreference = "PHYSICAL"
	InterestedPartyUpdateStatementDeliveryPreferenceSuppress                      InterestedPartyUpdateStatementDeliveryPreference = "SUPPRESS"
)

func (e InterestedPartyUpdateStatementDeliveryPreference) ToPointer() *InterestedPartyUpdateStatementDeliveryPreference {
	return &e
}

// InterestedPartyUpdateTradeConfirmationDeliveryPreference - Delivery method instruction for trade confirmations for a given Interested Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
type InterestedPartyUpdateTradeConfirmationDeliveryPreference string

const (
	InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDeliveryPreferenceUnspecified InterestedPartyUpdateTradeConfirmationDeliveryPreference = "DELIVERY_PREFERENCE_UNSPECIFIED"
	InterestedPartyUpdateTradeConfirmationDeliveryPreferenceDigital                       InterestedPartyUpdateTradeConfirmationDeliveryPreference = "DIGITAL"
	InterestedPartyUpdateTradeConfirmationDeliveryPreferencePhysical                      InterestedPartyUpdateTradeConfirmationDeliveryPreference = "PHYSICAL"
	InterestedPartyUpdateTradeConfirmationDeliveryPreferenceSuppress                      InterestedPartyUpdateTradeConfirmationDeliveryPreference = "SUPPRESS"
)

func (e InterestedPartyUpdateTradeConfirmationDeliveryPreference) ToPointer() *InterestedPartyUpdateTradeConfirmationDeliveryPreference {
	return &e
}

// InterestedPartyUpdate - An interested party.
type InterestedPartyUpdate struct {
	// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).
	//
	//  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.
	//
	//  Advice on address input / editing: - Use an i18n-ready address widget such as  https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of  fields outside countries where that field is used.
	//
	//  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
	MailingAddress *PostalAddressUpdate `json:"mailing_address,omitempty"`
	// The sending address name for mailings to Interested Parties The name of an Interested Party; Used for envelope/communication addressing
	Recipient *string `json:"recipient,omitempty"`
	// Delivery method instruction for account statements for a given Interested Party; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
	StatementDeliveryPreference *InterestedPartyUpdateStatementDeliveryPreference `json:"statement_delivery_preference,omitempty"`
	// Delivery method instruction for trade confirmations for a given Interested Party record; Can be `DIGITAL`, `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
	TradeConfirmationDeliveryPreference *InterestedPartyUpdateTradeConfirmationDeliveryPreference `json:"trade_confirmation_delivery_preference,omitempty"`
}

func (o *InterestedPartyUpdate) GetMailingAddress() *PostalAddressUpdate {
	if o == nil {
		return nil
	}
	return o.MailingAddress
}

func (o *InterestedPartyUpdate) GetRecipient() *string {
	if o == nil {
		return nil
	}
	return o.Recipient
}

func (o *InterestedPartyUpdate) GetStatementDeliveryPreference() *InterestedPartyUpdateStatementDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.StatementDeliveryPreference
}

func (o *InterestedPartyUpdate) GetTradeConfirmationDeliveryPreference() *InterestedPartyUpdateTradeConfirmationDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TradeConfirmationDeliveryPreference
}