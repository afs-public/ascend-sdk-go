// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// InterestedPartyMailingAddress - The object containing data for the purpose of delivery physical mailings to a party; Typically used for statements, account updates, tax documents, and other postal mailings; May also be used as an alternative identity verification address to personalAddress. Required fields within the `mailing_address` object include:
//   - `administrative_area`
//   - `region_code` - 2 character CLDR Code
//   - `postal_code`
//   - `locality`
//   - `address_lines` - max 5 lines
type InterestedPartyMailingAddress struct {
	// Unstructured address lines describing the lower levels of an address.
	//
	//  Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way, the most specific line of an address can be selected based on the language.
	//
	//  The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved.
	//
	//  Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
	AddressLines []string `json:"address_lines,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
	AdministrativeArea *string `json:"administrative_area,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations.
	//
	//  If this value is not known, it should be omitted (rather than specifying a possibly incorrect default).
	//
	//  Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode *string `json:"language_code,omitempty"`
	// Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
	Locality *string `json:"locality,omitempty"`
	// Optional. The name of the organization at the address.
	Organization *string `json:"organization,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
	PostalCode *string `json:"postal_code,omitempty"`
	// Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain "care of" information.
	Recipients []string `json:"recipients,omitempty"`
	// Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See http://cldr.unicode.org/ and http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
	RegionCode *string `json:"region_code,omitempty"`
	// The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision.
	//
	//  All new revisions **must** be backward compatible with old revisions.
	Revision *int `json:"revision,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number alone, representing the "sector code" (Jamaica), "delivery area indicator" (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode *string `json:"sorting_code,omitempty"`
	// Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
	Sublocality *string `json:"sublocality,omitempty"`
}

func (o *InterestedPartyMailingAddress) GetAddressLines() []string {
	if o == nil {
		return nil
	}
	return o.AddressLines
}

func (o *InterestedPartyMailingAddress) GetAdministrativeArea() *string {
	if o == nil {
		return nil
	}
	return o.AdministrativeArea
}

func (o *InterestedPartyMailingAddress) GetLanguageCode() *string {
	if o == nil {
		return nil
	}
	return o.LanguageCode
}

func (o *InterestedPartyMailingAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *InterestedPartyMailingAddress) GetOrganization() *string {
	if o == nil {
		return nil
	}
	return o.Organization
}

func (o *InterestedPartyMailingAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *InterestedPartyMailingAddress) GetRecipients() []string {
	if o == nil {
		return nil
	}
	return o.Recipients
}

func (o *InterestedPartyMailingAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}

func (o *InterestedPartyMailingAddress) GetRevision() *int {
	if o == nil {
		return nil
	}
	return o.Revision
}

func (o *InterestedPartyMailingAddress) GetSortingCode() *string {
	if o == nil {
		return nil
	}
	return o.SortingCode
}

func (o *InterestedPartyMailingAddress) GetSublocality() *string {
	if o == nil {
		return nil
	}
	return o.Sublocality
}

// InterestedPartyStatementDeliveryPreference - Delivery method instruction for account statements for a given Interested Party; Can be `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
type InterestedPartyStatementDeliveryPreference string

const (
	InterestedPartyStatementDeliveryPreferencePhysical InterestedPartyStatementDeliveryPreference = "PHYSICAL"
	InterestedPartyStatementDeliveryPreferenceSuppress InterestedPartyStatementDeliveryPreference = "SUPPRESS"
)

func (e InterestedPartyStatementDeliveryPreference) ToPointer() *InterestedPartyStatementDeliveryPreference {
	return &e
}

// InterestedPartyTradeConfirmationDeliveryPreference - Delivery method instruction for trade confirmations for a given Interested Party record; Can be `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
type InterestedPartyTradeConfirmationDeliveryPreference string

const (
	InterestedPartyTradeConfirmationDeliveryPreferencePhysical InterestedPartyTradeConfirmationDeliveryPreference = "PHYSICAL"
	InterestedPartyTradeConfirmationDeliveryPreferenceSuppress InterestedPartyTradeConfirmationDeliveryPreference = "SUPPRESS"
)

func (e InterestedPartyTradeConfirmationDeliveryPreference) ToPointer() *InterestedPartyTradeConfirmationDeliveryPreference {
	return &e
}

// InterestedParty - An interested party.
type InterestedParty struct {
	// A system-generated unique identifier for an Interested Party on an account; Used to access the record after creation
	InterestedPartyID *string `json:"interested_party_id,omitempty"`
	// The object containing data for the purpose of delivery physical mailings to a party; Typically used for statements, account updates, tax documents, and other postal mailings; May also be used as an alternative identity verification address to personalAddress. Required fields within the `mailing_address` object include:
	//  - `administrative_area`
	//  - `region_code` - 2 character CLDR Code
	//  - `postal_code`
	//  - `locality`
	//  - `address_lines` - max 5 lines
	MailingAddress *InterestedPartyMailingAddress `json:"mailing_address,omitempty"`
	// The name field Format: accounts/[{account}/interestedParties/{interestedParty}
	Name *string `json:"name,omitempty"`
	// The sending address name for mailings to Interested Parties The name of an Interested Party; Used for envelope/communication addressing
	Recipient *string `json:"recipient,omitempty"`
	// Delivery method instruction for account statements for a given Interested Party; Can be `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
	StatementDeliveryPreference *InterestedPartyStatementDeliveryPreference `json:"statement_delivery_preference,omitempty"`
	// Delivery method instruction for trade confirmations for a given Interested Party record; Can be `PHYSICAL`, `SUPPRESS`; Defaults to `PHYSICAL` on party creation
	TradeConfirmationDeliveryPreference *InterestedPartyTradeConfirmationDeliveryPreference `json:"trade_confirmation_delivery_preference,omitempty"`
}

func (o *InterestedParty) GetInterestedPartyID() *string {
	if o == nil {
		return nil
	}
	return o.InterestedPartyID
}

func (o *InterestedParty) GetMailingAddress() *InterestedPartyMailingAddress {
	if o == nil {
		return nil
	}
	return o.MailingAddress
}

func (o *InterestedParty) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *InterestedParty) GetRecipient() *string {
	if o == nil {
		return nil
	}
	return o.Recipient
}

func (o *InterestedParty) GetStatementDeliveryPreference() *InterestedPartyStatementDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.StatementDeliveryPreference
}

func (o *InterestedParty) GetTradeConfirmationDeliveryPreference() *InterestedPartyTradeConfirmationDeliveryPreference {
	if o == nil {
		return nil
	}
	return o.TradeConfirmationDeliveryPreference
}
