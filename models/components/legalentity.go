// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LegalEntityBusinessIndustrialClassification string

const (
	LegalEntityBusinessIndustrialClassificationBusinessIndustrialClassificationUnspecified                LegalEntityBusinessIndustrialClassification = "BUSINESS_INDUSTRIAL_CLASSIFICATION_UNSPECIFIED"
	LegalEntityBusinessIndustrialClassificationAgricultureForestryAndFishing                              LegalEntityBusinessIndustrialClassification = "AGRICULTURE_FORESTRY_AND_FISHING"
	LegalEntityBusinessIndustrialClassificationMining                                                     LegalEntityBusinessIndustrialClassification = "MINING"
	LegalEntityBusinessIndustrialClassificationConstruction                                               LegalEntityBusinessIndustrialClassification = "CONSTRUCTION"
	LegalEntityBusinessIndustrialClassificationManufacturing                                              LegalEntityBusinessIndustrialClassification = "MANUFACTURING"
	LegalEntityBusinessIndustrialClassificationTransportationCommunicationsElectricGasAndSanitaryServices LegalEntityBusinessIndustrialClassification = "TRANSPORTATION_COMMUNICATIONS_ELECTRIC_GAS_AND_SANITARY_SERVICES"
	LegalEntityBusinessIndustrialClassificationWholesaleTrade                                             LegalEntityBusinessIndustrialClassification = "WHOLESALE_TRADE"
	LegalEntityBusinessIndustrialClassificationRetailTrade                                                LegalEntityBusinessIndustrialClassification = "RETAIL_TRADE"
	LegalEntityBusinessIndustrialClassificationFinanceInsuranceAndRealEstate                              LegalEntityBusinessIndustrialClassification = "FINANCE_INSURANCE_AND_REAL_ESTATE"
	LegalEntityBusinessIndustrialClassificationServices                                                   LegalEntityBusinessIndustrialClassification = "SERVICES"
	LegalEntityBusinessIndustrialClassificationPublicAdministration                                       LegalEntityBusinessIndustrialClassification = "PUBLIC_ADMINISTRATION"
)

func (e LegalEntityBusinessIndustrialClassification) ToPointer() *LegalEntityBusinessIndustrialClassification {
	return &e
}

// LegalEntityCorporateStructure - Corporate structure of the entity.
type LegalEntityCorporateStructure string

const (
	LegalEntityCorporateStructureEntityCorporateStructureUnspecified LegalEntityCorporateStructure = "ENTITY_CORPORATE_STRUCTURE_UNSPECIFIED"
	LegalEntityCorporateStructureCorporationCCorp                    LegalEntityCorporateStructure = "CORPORATION_C_CORP"
	LegalEntityCorporateStructureCorporationSCorp                    LegalEntityCorporateStructure = "CORPORATION_S_CORP"
	LegalEntityCorporateStructureCorporationBCorp                    LegalEntityCorporateStructure = "CORPORATION_B_CORP"
	LegalEntityCorporateStructureCorporationNonprofit                LegalEntityCorporateStructure = "CORPORATION_NONPROFIT"
)

func (e LegalEntityCorporateStructure) ToPointer() *LegalEntityCorporateStructure {
	return &e
}

// LegalEntityNegativeNews - Information about any negative news against related parties and entities
type LegalEntityNegativeNews struct {
	// Indicates whether there is negative news against related parties
	NegativeNewsAgainstRelatedParties *bool `json:"negative_news_against_related_parties,omitempty"`
	// Description of the negative news against related parties
	NegativeNewsAgainstRelatedPartiesDescription *string `json:"negative_news_against_related_parties_description,omitempty"`
}

func (o *LegalEntityNegativeNews) GetNegativeNewsAgainstRelatedParties() *bool {
	if o == nil {
		return nil
	}
	return o.NegativeNewsAgainstRelatedParties
}

func (o *LegalEntityNegativeNews) GetNegativeNewsAgainstRelatedPartiesDescription() *string {
	if o == nil {
		return nil
	}
	return o.NegativeNewsAgainstRelatedPartiesDescription
}

// EntityDueDiligence - Due Diligence for Legal Entities
type EntityDueDiligence struct {
	// Indicates whether the entity issues bearer shares
	EntityIssuesBearerShares *bool `json:"entity_issues_bearer_shares,omitempty"`
	// Information about any negative news against related parties and entities
	NegativeNews *LegalEntityNegativeNews `json:"negative_news,omitempty"`
}

func (o *EntityDueDiligence) GetEntityIssuesBearerShares() *bool {
	if o == nil {
		return nil
	}
	return o.EntityIssuesBearerShares
}

func (o *EntityDueDiligence) GetNegativeNews() *LegalEntityNegativeNews {
	if o == nil {
		return nil
	}
	return o.NegativeNews
}

// LegalEntityEntityType - The entity type.
type LegalEntityEntityType string

const (
	LegalEntityEntityTypeEntityTypeUnspecified   LegalEntityEntityType = "ENTITY_TYPE_UNSPECIFIED"
	LegalEntityEntityTypeCorporation             LegalEntityEntityType = "CORPORATION"
	LegalEntityEntityTypeLimitedLiabilityCompany LegalEntityEntityType = "LIMITED_LIABILITY_COMPANY"
	LegalEntityEntityTypePartnership             LegalEntityEntityType = "PARTNERSHIP"
	LegalEntityEntityTypeTrust                   LegalEntityEntityType = "TRUST"
	LegalEntityEntityTypeEstate                  LegalEntityEntityType = "ESTATE"
)

func (e LegalEntityEntityType) ToPointer() *LegalEntityEntityType {
	return &e
}

// LegalEntityExemptCustomerReason - **Field Dependencies:**
//
// Exempt entities must set `exempt_verifying_beneficial_owners` to `true` and provide an `exempt_customer_reason` on the owner record.
//
// Required if `exempt_verifying_beneficial_owners` is `true`.
//
// Otherwise, must be empty.
type LegalEntityExemptCustomerReason string

const (
	LegalEntityExemptCustomerReasonExemptReasonUnspecified                        LegalEntityExemptCustomerReason = "EXEMPT_REASON_UNSPECIFIED"
	LegalEntityExemptCustomerReasonRegulatedFinancialInstitution                  LegalEntityExemptCustomerReason = "REGULATED_FINANCIAL_INSTITUTION"
	LegalEntityExemptCustomerReasonDepartmentOrAgencyOfFederalStateOrSubdivision  LegalEntityExemptCustomerReason = "DEPARTMENT_OR_AGENCY_OF_FEDERAL_STATE_OR_SUBDIVISION"
	LegalEntityExemptCustomerReasonNonBankListedEntity                            LegalEntityExemptCustomerReason = "NON_BANK_LISTED_ENTITY"
	LegalEntityExemptCustomerReasonSection12SecuritiesExchangeAct1934Or15D        LegalEntityExemptCustomerReason = "SECTION_12_SECURITIES_EXCHANGE_ACT_1934_OR_15D"
	LegalEntityExemptCustomerReasonSection3InvestmentCompanyAct1940               LegalEntityExemptCustomerReason = "SECTION_3_INVESTMENT_COMPANY_ACT_1940"
	LegalEntityExemptCustomerReasonSection202AInvestmentAdvisorsAct1940           LegalEntityExemptCustomerReason = "SECTION_202A_INVESTMENT_ADVISORS_ACT_1940"
	LegalEntityExemptCustomerReasonSection3SecuritiesExchangeAct1934Section6Or17A LegalEntityExemptCustomerReason = "SECTION_3_SECURITIES_EXCHANGE_ACT_1934_SECTION_6_OR_17A"
	LegalEntityExemptCustomerReasonAnyOtherSecuritiesExchangeAct1934              LegalEntityExemptCustomerReason = "ANY_OTHER_SECURITIES_EXCHANGE_ACT_1934"
	LegalEntityExemptCustomerReasonCommodityFuturesTradingCommissionRegistered    LegalEntityExemptCustomerReason = "COMMODITY_FUTURES_TRADING_COMMISSION_REGISTERED"
	LegalEntityExemptCustomerReasonPublicAccountingFirmSection102SarbanesOxley    LegalEntityExemptCustomerReason = "PUBLIC_ACCOUNTING_FIRM_SECTION_102_SARBANES_OXLEY"
	LegalEntityExemptCustomerReasonStateRegulatedInsuranceCompany                 LegalEntityExemptCustomerReason = "STATE_REGULATED_INSURANCE_COMPANY"
)

func (e LegalEntityExemptCustomerReason) ToPointer() *LegalEntityExemptCustomerReason {
	return &e
}

// FormationDate - If the legal entity is a trust, the formation date is required.
type FormationDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *FormationDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *FormationDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *FormationDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LegalEntityEffectiveDate - The date on which the trader meets or exceeds the large trader reporting threshold, which is defined by the U.S. Securities and Exchange Commission (SEC) as trades of 2 million shares or $20 million in a single day or 20 million shares or $200 million during a calendar month
type LegalEntityEffectiveDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *LegalEntityEffectiveDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *LegalEntityEffectiveDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *LegalEntityEffectiveDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LegalEntityLargeTrader - Large trader for the legal entity.
type LegalEntityLargeTrader struct {
	// The date on which the trader meets or exceeds the large trader reporting threshold, which is defined by the U.S. Securities and Exchange Commission (SEC) as trades of 2 million shares or $20 million in a single day or 20 million shares or $200 million during a calendar month
	EffectiveDate *LegalEntityEffectiveDate `json:"effective_date,omitempty"`
	// SEC-issued ID signifying the person/entity as a large trader; Required for CAIS regulatory reporting.
	LargeTraderID *string `json:"large_trader_id,omitempty"`
}

func (o *LegalEntityLargeTrader) GetEffectiveDate() *LegalEntityEffectiveDate {
	if o == nil {
		return nil
	}
	return o.EffectiveDate
}

func (o *LegalEntityLargeTrader) GetLargeTraderID() *string {
	if o == nil {
		return nil
	}
	return o.LargeTraderID
}

// LegalAddress - The mailing address of the legal entity. Required fields within the `legal_address` object include:
//   - `administrative_area`
//   - `region_code` - 2 character CLDR Code
//   - `postal_code`
//   - `locality`
//   - `address_lines` - max 5 lines
type LegalAddress struct {
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

func (o *LegalAddress) GetAddressLines() []string {
	if o == nil {
		return nil
	}
	return o.AddressLines
}

func (o *LegalAddress) GetAdministrativeArea() *string {
	if o == nil {
		return nil
	}
	return o.AdministrativeArea
}

func (o *LegalAddress) GetLanguageCode() *string {
	if o == nil {
		return nil
	}
	return o.LanguageCode
}

func (o *LegalAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *LegalAddress) GetOrganization() *string {
	if o == nil {
		return nil
	}
	return o.Organization
}

func (o *LegalAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *LegalAddress) GetRecipients() []string {
	if o == nil {
		return nil
	}
	return o.Recipients
}

func (o *LegalAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}

func (o *LegalAddress) GetRevision() *int {
	if o == nil {
		return nil
	}
	return o.Revision
}

func (o *LegalAddress) GetSortingCode() *string {
	if o == nil {
		return nil
	}
	return o.SortingCode
}

func (o *LegalAddress) GetSublocality() *string {
	if o == nil {
		return nil
	}
	return o.Sublocality
}

// LegalEntityTaxIDType - The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN.
type LegalEntityTaxIDType string

const (
	LegalEntityTaxIDTypeTaxIDTypeUnspecified LegalEntityTaxIDType = "TAX_ID_TYPE_UNSPECIFIED"
	LegalEntityTaxIDTypeTaxIDTypeSsn         LegalEntityTaxIDType = "TAX_ID_TYPE_SSN"
	LegalEntityTaxIDTypeTaxIDTypeItin        LegalEntityTaxIDType = "TAX_ID_TYPE_ITIN"
	LegalEntityTaxIDTypeTaxIDTypeEin         LegalEntityTaxIDType = "TAX_ID_TYPE_EIN"
)

func (e LegalEntityTaxIDType) ToPointer() *LegalEntityTaxIDType {
	return &e
}

// LegalEntityCNoticeDate - C Notice date.
type LegalEntityCNoticeDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *LegalEntityCNoticeDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *LegalEntityCNoticeDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *LegalEntityCNoticeDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LegalEntityFederalTaxClassification - Federal tax classification.
type LegalEntityFederalTaxClassification string

const (
	LegalEntityFederalTaxClassificationFederalTaxClassificationUnspecified LegalEntityFederalTaxClassification = "FEDERAL_TAX_CLASSIFICATION_UNSPECIFIED"
	LegalEntityFederalTaxClassificationIndivSolepropOrSinglememberllc      LegalEntityFederalTaxClassification = "INDIV_SOLEPROP_OR_SINGLEMEMBERLLC"
	LegalEntityFederalTaxClassificationPartnership                         LegalEntityFederalTaxClassification = "PARTNERSHIP"
	LegalEntityFederalTaxClassificationCCorporation                        LegalEntityFederalTaxClassification = "C_CORPORATION"
	LegalEntityFederalTaxClassificationSCorporation                        LegalEntityFederalTaxClassification = "S_CORPORATION"
	LegalEntityFederalTaxClassificationTrustEstate                         LegalEntityFederalTaxClassification = "TRUST_ESTATE"
	LegalEntityFederalTaxClassificationLlcTaxedAsCCorp                     LegalEntityFederalTaxClassification = "LLC_TAXED_AS_C_CORP"
	LegalEntityFederalTaxClassificationLlcTaxedAsSCorp                     LegalEntityFederalTaxClassification = "LLC_TAXED_AS_S_CORP"
	LegalEntityFederalTaxClassificationLlcTaxedAsPartnership               LegalEntityFederalTaxClassification = "LLC_TAXED_AS_PARTNERSHIP"
	LegalEntityFederalTaxClassificationOther                               LegalEntityFederalTaxClassification = "OTHER"
)

func (e LegalEntityFederalTaxClassification) ToPointer() *LegalEntityFederalTaxClassification {
	return &e
}

// LegalEntityFirstBNoticeDate - Initial B Notice date.
type LegalEntityFirstBNoticeDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *LegalEntityFirstBNoticeDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *LegalEntityFirstBNoticeDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *LegalEntityFirstBNoticeDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LegalEntityIrsFormType - IRS form type.
type LegalEntityIrsFormType string

const (
	LegalEntityIrsFormTypeIrsFormTypeUnspecified LegalEntityIrsFormType = "IRS_FORM_TYPE_UNSPECIFIED"
	LegalEntityIrsFormTypeW9                     LegalEntityIrsFormType = "W_9"
	LegalEntityIrsFormTypeW8Ben                  LegalEntityIrsFormType = "W_8BEN"
)

func (e LegalEntityIrsFormType) ToPointer() *LegalEntityIrsFormType {
	return &e
}

// LegalEntityReportingEligibility - Tax reporting eligibility.
type LegalEntityReportingEligibility string

const (
	LegalEntityReportingEligibilityTaxReportingEligibilityUnspecified LegalEntityReportingEligibility = "TAX_REPORTING_ELIGIBILITY_UNSPECIFIED"
	LegalEntityReportingEligibilityEligible                           LegalEntityReportingEligibility = "ELIGIBLE"
	LegalEntityReportingEligibilityIneligible                         LegalEntityReportingEligibility = "INELIGIBLE"
)

func (e LegalEntityReportingEligibility) ToPointer() *LegalEntityReportingEligibility {
	return &e
}

// LegalEntityTaxCertificationDate - Tax Certification date.
type LegalEntityTaxCertificationDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *LegalEntityTaxCertificationDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *LegalEntityTaxCertificationDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *LegalEntityTaxCertificationDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// LegalEntityTaxpayerCertificationState - Taxpayer certification status.
type LegalEntityTaxpayerCertificationState string

const (
	LegalEntityTaxpayerCertificationStateTaxpayerCertificationStateUnspecified LegalEntityTaxpayerCertificationState = "TAXPAYER_CERTIFICATION_STATE_UNSPECIFIED"
	LegalEntityTaxpayerCertificationStateCertified                             LegalEntityTaxpayerCertificationState = "CERTIFIED"
	LegalEntityTaxpayerCertificationStateUncertified                           LegalEntityTaxpayerCertificationState = "UNCERTIFIED"
)

func (e LegalEntityTaxpayerCertificationState) ToPointer() *LegalEntityTaxpayerCertificationState {
	return &e
}

// LegalEntityUsTinStatus - United States Individual Taxpayer Identification Number (ITIN) status.
type LegalEntityUsTinStatus string

const (
	LegalEntityUsTinStatusUsTinStatusUnspecified LegalEntityUsTinStatus = "US_TIN_STATUS_UNSPECIFIED"
	LegalEntityUsTinStatusPassing                LegalEntityUsTinStatus = "PASSING"
	LegalEntityUsTinStatusFailing                LegalEntityUsTinStatus = "FAILING"
)

func (e LegalEntityUsTinStatus) ToPointer() *LegalEntityUsTinStatus {
	return &e
}

// LegalEntityWithholdingState - B/C Notice status.
type LegalEntityWithholdingState string

const (
	LegalEntityWithholdingStateWithholdingStateUnspecified LegalEntityWithholdingState = "WITHHOLDING_STATE_UNSPECIFIED"
	LegalEntityWithholdingStateFirstBNoticeReceived        LegalEntityWithholdingState = "FIRST_B_NOTICE_RECEIVED"
	LegalEntityWithholdingStateSecondBNoticeReceived       LegalEntityWithholdingState = "SECOND_B_NOTICE_RECEIVED"
	LegalEntityWithholdingStateCNoticeReceived             LegalEntityWithholdingState = "C_NOTICE_RECEIVED"
	LegalEntityWithholdingStateCNoticeIndicatedByCustomer  LegalEntityWithholdingState = "C_NOTICE_INDICATED_BY_CUSTOMER"
)

func (e LegalEntityWithholdingState) ToPointer() *LegalEntityWithholdingState {
	return &e
}

// LegalEntityTaxProfile - The tax profile for the legal entity.
type LegalEntityTaxProfile struct {
	// C Notice date.
	CNoticeDate *LegalEntityCNoticeDate `json:"c_notice_date,omitempty"`
	// Federal tax classification.
	FederalTaxClassification *LegalEntityFederalTaxClassification `json:"federal_tax_classification,omitempty"`
	// Initial B Notice date.
	FirstBNoticeDate *LegalEntityFirstBNoticeDate `json:"first_b_notice_date,omitempty"`
	// IRS form type.
	IrsFormType *LegalEntityIrsFormType `json:"irs_form_type,omitempty"`
	// Legal tax region must be "US" if provided W-9, otherwise must be a non-US country.
	LegalTaxRegionCode *string `json:"legal_tax_region_code,omitempty"`
	// Tax reporting eligibility.
	ReportingEligibility *LegalEntityReportingEligibility `json:"reporting_eligibility,omitempty"`
	// Tax Certification date.
	TaxCertificationDate *LegalEntityTaxCertificationDate `json:"tax_certification_date,omitempty"`
	// Taxpayer certification status.
	TaxpayerCertificationState *LegalEntityTaxpayerCertificationState `json:"taxpayer_certification_state,omitempty"`
	// United States Individual Taxpayer Identification Number (ITIN) status.
	UsTinStatus *LegalEntityUsTinStatus `json:"us_tin_status,omitempty"`
	// B/C Notice status.
	WithholdingState *LegalEntityWithholdingState `json:"withholding_state,omitempty"`
}

func (o *LegalEntityTaxProfile) GetCNoticeDate() *LegalEntityCNoticeDate {
	if o == nil {
		return nil
	}
	return o.CNoticeDate
}

func (o *LegalEntityTaxProfile) GetFederalTaxClassification() *LegalEntityFederalTaxClassification {
	if o == nil {
		return nil
	}
	return o.FederalTaxClassification
}

func (o *LegalEntityTaxProfile) GetFirstBNoticeDate() *LegalEntityFirstBNoticeDate {
	if o == nil {
		return nil
	}
	return o.FirstBNoticeDate
}

func (o *LegalEntityTaxProfile) GetIrsFormType() *LegalEntityIrsFormType {
	if o == nil {
		return nil
	}
	return o.IrsFormType
}

func (o *LegalEntityTaxProfile) GetLegalTaxRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.LegalTaxRegionCode
}

func (o *LegalEntityTaxProfile) GetReportingEligibility() *LegalEntityReportingEligibility {
	if o == nil {
		return nil
	}
	return o.ReportingEligibility
}

func (o *LegalEntityTaxProfile) GetTaxCertificationDate() *LegalEntityTaxCertificationDate {
	if o == nil {
		return nil
	}
	return o.TaxCertificationDate
}

func (o *LegalEntityTaxProfile) GetTaxpayerCertificationState() *LegalEntityTaxpayerCertificationState {
	if o == nil {
		return nil
	}
	return o.TaxpayerCertificationState
}

func (o *LegalEntityTaxProfile) GetUsTinStatus() *LegalEntityUsTinStatus {
	if o == nil {
		return nil
	}
	return o.UsTinStatus
}

func (o *LegalEntityTaxProfile) GetWithholdingState() *LegalEntityWithholdingState {
	if o == nil {
		return nil
	}
	return o.WithholdingState
}

// LegalEntity - A legal entity. Legal entities are organizations, such as companies, that participate in financial transactions
type LegalEntity struct {
	// Indicates whether the entity is an accredited investor. By default, this is set to `false`.
	AccreditedInvestor *bool `json:"accredited_investor,omitempty"`
	// Indicates whether the entity is an adviser. By default, this is set to `false`.
	Adviser *bool `json:"adviser,omitempty"`
	// Indicates whether the entity is a broker dealer. By default, this is set to `false`.
	BrokerDealer                     *bool                                        `json:"broker_dealer,omitempty"`
	BusinessIndustrialClassification *LegalEntityBusinessIndustrialClassification `json:"business_industrial_classification,omitempty"`
	// Corporate structure of the entity.
	CorporateStructure *LegalEntityCorporateStructure `json:"corporate_structure,omitempty"`
	// The correspondent id associated with the legal entity.
	CorrespondentID *string `json:"correspondent_id,omitempty"`
	// DBA (Doing Business As) names. Can list up to 5 associated with the Legal Entity
	DoingBusinessAs []string `json:"doing_business_as,omitempty"`
	// Due Diligence for Legal Entities
	EntityDueDiligence *EntityDueDiligence `json:"entity_due_diligence,omitempty"`
	// The legal entity name.
	EntityName *string `json:"entity_name,omitempty"`
	// The entity type.
	EntityType *LegalEntityEntityType `json:"entity_type,omitempty"`
	// **Field Dependencies:**
	//
	// Exempt entities must set `exempt_verifying_beneficial_owners` to `true` and provide an `exempt_customer_reason` on the owner record.
	//
	// Required if `exempt_verifying_beneficial_owners` is `true`.
	//
	// Otherwise, must be empty.
	ExemptCustomerReason *LegalEntityExemptCustomerReason `json:"exempt_customer_reason,omitempty"`
	// Indicates whether the entity is exempt from verifying beneficial owners and Enhanced Due Diligence. By default, this is set to `false`
	ExemptVerifyingBeneficialOwners *bool `json:"exempt_verifying_beneficial_owners,omitempty"`
	// If the legal entity is a trust, they may set this field to convey ownership and value to a trustee.
	ForTheBenefitOf *string `json:"for_the_benefit_of,omitempty"`
	// Indicates whether the entity is a foreign entity. By default, this is set to `false`.
	ForeignEntity *bool `json:"foreign_entity,omitempty"`
	// Indicates whether the entity is a foreign financial institution. By default, this is set to `false`.
	ForeignFinancialInstitution *bool `json:"foreign_financial_institution,omitempty"`
	// If the legal entity is a trust, the formation date is required.
	FormationDate *FormationDate `json:"formation_date,omitempty"`
	// Globally Unique identifier for a legal natural person
	GlobalPersonID *string `json:"global_person_id,omitempty"`
	// Indicates whether the entity is an institutional customer
	InstitutionalCustomer *bool `json:"institutional_customer,omitempty"`
	// Investigation id relating to the Customer Identification Program (CIP) and Customer Due Diligence (CDD).
	InvestigationID *string `json:"investigation_id,omitempty"`
	// Large trader for the legal entity.
	LargeTrader *LegalEntityLargeTrader `json:"large_trader,omitempty"`
	// The mailing address of the legal entity. Required fields within the `legal_address` object include:
	//  - `administrative_area`
	//  - `region_code` - 2 character CLDR Code
	//  - `postal_code`
	//  - `locality`
	//  - `address_lines` - max 5 lines
	LegalAddress *LegalAddress `json:"legal_address,omitempty"`
	// A system-generated unique identifier referencing a single juridical (non-natural) person (e.g., a corporation); Used to access the record after creation
	LegalEntityID *string `json:"legal_entity_id,omitempty"`
	// The Legal Entity Identifier (LEI) is the financial industry term for a unique global identifier for legal entities participating in financial transactions
	LeiCode *string `json:"lei_code,omitempty"`
	// The name field Format: legalEntities/{legalEntity}
	Name *string `json:"name,omitempty"`
	// The operational footprint of an entity. Operating regions encompass all countries and regions where a company has a significant business presence This includes locations with physical offices, manufacturing plants, service centers, and sales and marketing activities Regions must be provided as two-character CLDR country codes
	OperatingRegions []string `json:"operating_regions,omitempty"`
	// The legal home of an entity. A region of registration, in the context of a corporation, refers to the specific geographic area where the corporation is legally registered and incorporated Defines the legal jurisdiction and framework under which the corporation operates, including legal regulations, tax obligations, and compliance requirements Region must be provided as a two-character CLDR country code
	RegistrationRegion *string `json:"registration_region,omitempty"`
	// Indicates whether the entity is a regulated investment company. By default, this is set to `false`.
	RegulatedInvestmentCompany *bool `json:"regulated_investment_company,omitempty"`
	// Document ids related to the legal entity. At least one is required for RIA correspondents when creating Estate or Trust accounts.
	RelatedDocumentIds []string `json:"related_document_ids,omitempty"`
	// Indicates whether the trust is a revocable trust. By default, this is set to `false`.
	RevocableTrust *bool `json:"revocable_trust,omitempty"`
	// The full U.S. tax ID for a related entity; Must be provided with `EIN` tax ID type
	TaxID *string `json:"tax_id,omitempty"`
	// The last four characters of the related person's tax identifier; Masked/truncated to "last four" in most usage contexts to preserve data privacy.
	TaxIDLastFour *string `json:"tax_id_last_four,omitempty"`
	// The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN.
	TaxIDType *LegalEntityTaxIDType `json:"tax_id_type,omitempty"`
	// The tax profile for the legal entity.
	TaxProfile *LegalEntityTaxProfile `json:"tax_profile,omitempty"`
}

func (o *LegalEntity) GetAccreditedInvestor() *bool {
	if o == nil {
		return nil
	}
	return o.AccreditedInvestor
}

func (o *LegalEntity) GetAdviser() *bool {
	if o == nil {
		return nil
	}
	return o.Adviser
}

func (o *LegalEntity) GetBrokerDealer() *bool {
	if o == nil {
		return nil
	}
	return o.BrokerDealer
}

func (o *LegalEntity) GetBusinessIndustrialClassification() *LegalEntityBusinessIndustrialClassification {
	if o == nil {
		return nil
	}
	return o.BusinessIndustrialClassification
}

func (o *LegalEntity) GetCorporateStructure() *LegalEntityCorporateStructure {
	if o == nil {
		return nil
	}
	return o.CorporateStructure
}

func (o *LegalEntity) GetCorrespondentID() *string {
	if o == nil {
		return nil
	}
	return o.CorrespondentID
}

func (o *LegalEntity) GetDoingBusinessAs() []string {
	if o == nil {
		return nil
	}
	return o.DoingBusinessAs
}

func (o *LegalEntity) GetEntityDueDiligence() *EntityDueDiligence {
	if o == nil {
		return nil
	}
	return o.EntityDueDiligence
}

func (o *LegalEntity) GetEntityName() *string {
	if o == nil {
		return nil
	}
	return o.EntityName
}

func (o *LegalEntity) GetEntityType() *LegalEntityEntityType {
	if o == nil {
		return nil
	}
	return o.EntityType
}

func (o *LegalEntity) GetExemptCustomerReason() *LegalEntityExemptCustomerReason {
	if o == nil {
		return nil
	}
	return o.ExemptCustomerReason
}

func (o *LegalEntity) GetExemptVerifyingBeneficialOwners() *bool {
	if o == nil {
		return nil
	}
	return o.ExemptVerifyingBeneficialOwners
}

func (o *LegalEntity) GetForTheBenefitOf() *string {
	if o == nil {
		return nil
	}
	return o.ForTheBenefitOf
}

func (o *LegalEntity) GetForeignEntity() *bool {
	if o == nil {
		return nil
	}
	return o.ForeignEntity
}

func (o *LegalEntity) GetForeignFinancialInstitution() *bool {
	if o == nil {
		return nil
	}
	return o.ForeignFinancialInstitution
}

func (o *LegalEntity) GetFormationDate() *FormationDate {
	if o == nil {
		return nil
	}
	return o.FormationDate
}

func (o *LegalEntity) GetGlobalPersonID() *string {
	if o == nil {
		return nil
	}
	return o.GlobalPersonID
}

func (o *LegalEntity) GetInstitutionalCustomer() *bool {
	if o == nil {
		return nil
	}
	return o.InstitutionalCustomer
}

func (o *LegalEntity) GetInvestigationID() *string {
	if o == nil {
		return nil
	}
	return o.InvestigationID
}

func (o *LegalEntity) GetLargeTrader() *LegalEntityLargeTrader {
	if o == nil {
		return nil
	}
	return o.LargeTrader
}

func (o *LegalEntity) GetLegalAddress() *LegalAddress {
	if o == nil {
		return nil
	}
	return o.LegalAddress
}

func (o *LegalEntity) GetLegalEntityID() *string {
	if o == nil {
		return nil
	}
	return o.LegalEntityID
}

func (o *LegalEntity) GetLeiCode() *string {
	if o == nil {
		return nil
	}
	return o.LeiCode
}

func (o *LegalEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LegalEntity) GetOperatingRegions() []string {
	if o == nil {
		return nil
	}
	return o.OperatingRegions
}

func (o *LegalEntity) GetRegistrationRegion() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationRegion
}

func (o *LegalEntity) GetRegulatedInvestmentCompany() *bool {
	if o == nil {
		return nil
	}
	return o.RegulatedInvestmentCompany
}

func (o *LegalEntity) GetRelatedDocumentIds() []string {
	if o == nil {
		return nil
	}
	return o.RelatedDocumentIds
}

func (o *LegalEntity) GetRevocableTrust() *bool {
	if o == nil {
		return nil
	}
	return o.RevocableTrust
}

func (o *LegalEntity) GetTaxID() *string {
	if o == nil {
		return nil
	}
	return o.TaxID
}

func (o *LegalEntity) GetTaxIDLastFour() *string {
	if o == nil {
		return nil
	}
	return o.TaxIDLastFour
}

func (o *LegalEntity) GetTaxIDType() *LegalEntityTaxIDType {
	if o == nil {
		return nil
	}
	return o.TaxIDType
}

func (o *LegalEntity) GetTaxProfile() *LegalEntityTaxProfile {
	if o == nil {
		return nil
	}
	return o.TaxProfile
}
