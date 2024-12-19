// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type BusinessIndustrialClassification string

const (
	BusinessIndustrialClassificationBusinessIndustrialClassificationUnspecified                BusinessIndustrialClassification = "BUSINESS_INDUSTRIAL_CLASSIFICATION_UNSPECIFIED"
	BusinessIndustrialClassificationAgricultureForestryAndFishing                              BusinessIndustrialClassification = "AGRICULTURE_FORESTRY_AND_FISHING"
	BusinessIndustrialClassificationMining                                                     BusinessIndustrialClassification = "MINING"
	BusinessIndustrialClassificationConstruction                                               BusinessIndustrialClassification = "CONSTRUCTION"
	BusinessIndustrialClassificationManufacturing                                              BusinessIndustrialClassification = "MANUFACTURING"
	BusinessIndustrialClassificationTransportationCommunicationsElectricGasAndSanitaryServices BusinessIndustrialClassification = "TRANSPORTATION_COMMUNICATIONS_ELECTRIC_GAS_AND_SANITARY_SERVICES"
	BusinessIndustrialClassificationWholesaleTrade                                             BusinessIndustrialClassification = "WHOLESALE_TRADE"
	BusinessIndustrialClassificationRetailTrade                                                BusinessIndustrialClassification = "RETAIL_TRADE"
	BusinessIndustrialClassificationFinanceInsuranceAndRealEstate                              BusinessIndustrialClassification = "FINANCE_INSURANCE_AND_REAL_ESTATE"
	BusinessIndustrialClassificationServices                                                   BusinessIndustrialClassification = "SERVICES"
	BusinessIndustrialClassificationPublicAdministration                                       BusinessIndustrialClassification = "PUBLIC_ADMINISTRATION"
)

func (e BusinessIndustrialClassification) ToPointer() *BusinessIndustrialClassification {
	return &e
}

// CorporateStructure - Corporate structure of the entity.
type CorporateStructure string

const (
	CorporateStructureEntityCorporateStructureUnspecified CorporateStructure = "ENTITY_CORPORATE_STRUCTURE_UNSPECIFIED"
	CorporateStructureCorporationCCorp                    CorporateStructure = "CORPORATION_C_CORP"
	CorporateStructureCorporationSCorp                    CorporateStructure = "CORPORATION_S_CORP"
	CorporateStructureCorporationBCorp                    CorporateStructure = "CORPORATION_B_CORP"
	CorporateStructureCorporationNonprofit                CorporateStructure = "CORPORATION_NONPROFIT"
)

func (e CorporateStructure) ToPointer() *CorporateStructure {
	return &e
}

// EntityType - The entity type.
type EntityType string

const (
	EntityTypeEntityTypeUnspecified               EntityType = "ENTITY_TYPE_UNSPECIFIED"
	EntityTypeCorporation                         EntityType = "CORPORATION"
	EntityTypeLimitedLiabilityCompany             EntityType = "LIMITED_LIABILITY_COMPANY"
	EntityTypePartnership                         EntityType = "PARTNERSHIP"
	EntityTypeSoleProprietorshipOrSingleMemberLlc EntityType = "SOLE_PROPRIETORSHIP_OR_SINGLE_MEMBER_LLC"
	EntityTypeTrust                               EntityType = "TRUST"
	EntityTypeEstate                              EntityType = "ESTATE"
)

func (e EntityType) ToPointer() *EntityType {
	return &e
}

// ExemptCustomerReason - The reason the customer is exempt from verifying beneficial owners, if applicable.
type ExemptCustomerReason string

const (
	ExemptCustomerReasonExemptReasonUnspecified                        ExemptCustomerReason = "EXEMPT_REASON_UNSPECIFIED"
	ExemptCustomerReasonRegulatedFinancialInstitution                  ExemptCustomerReason = "REGULATED_FINANCIAL_INSTITUTION"
	ExemptCustomerReasonDepartmentOrAgencyOfFederalStateOrSubdivision  ExemptCustomerReason = "DEPARTMENT_OR_AGENCY_OF_FEDERAL_STATE_OR_SUBDIVISION"
	ExemptCustomerReasonNonBankListedEntity                            ExemptCustomerReason = "NON_BANK_LISTED_ENTITY"
	ExemptCustomerReasonSection12SecuritiesExchangeAct1934Or15D        ExemptCustomerReason = "SECTION_12_SECURITIES_EXCHANGE_ACT_1934_OR_15D"
	ExemptCustomerReasonSection3InvestmentCompanyAct1940               ExemptCustomerReason = "SECTION_3_INVESTMENT_COMPANY_ACT_1940"
	ExemptCustomerReasonSection202AInvestmentAdvisorsAct1940           ExemptCustomerReason = "SECTION_202A_INVESTMENT_ADVISORS_ACT_1940"
	ExemptCustomerReasonSection3SecuritiesExchangeAct1934Section6Or17A ExemptCustomerReason = "SECTION_3_SECURITIES_EXCHANGE_ACT_1934_SECTION_6_OR_17A"
	ExemptCustomerReasonAnyOtherSecuritiesExchangeAct1934              ExemptCustomerReason = "ANY_OTHER_SECURITIES_EXCHANGE_ACT_1934"
	ExemptCustomerReasonCommodityFuturesTradingCommissionRegistered    ExemptCustomerReason = "COMMODITY_FUTURES_TRADING_COMMISSION_REGISTERED"
	ExemptCustomerReasonPublicAccountingFirmSection102SarbanesOxley    ExemptCustomerReason = "PUBLIC_ACCOUNTING_FIRM_SECTION_102_SARBANES_OXLEY"
	ExemptCustomerReasonStateRegulatedInsuranceCompany                 ExemptCustomerReason = "STATE_REGULATED_INSURANCE_COMPANY"
)

func (e ExemptCustomerReason) ToPointer() *ExemptCustomerReason {
	return &e
}

// LegalEntityCreateTaxIDType - The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN.
type LegalEntityCreateTaxIDType string

const (
	LegalEntityCreateTaxIDTypeTaxIDTypeUnspecified LegalEntityCreateTaxIDType = "TAX_ID_TYPE_UNSPECIFIED"
	LegalEntityCreateTaxIDTypeTaxIDTypeSsn         LegalEntityCreateTaxIDType = "TAX_ID_TYPE_SSN"
	LegalEntityCreateTaxIDTypeTaxIDTypeItin        LegalEntityCreateTaxIDType = "TAX_ID_TYPE_ITIN"
	LegalEntityCreateTaxIDTypeTaxIDTypeEin         LegalEntityCreateTaxIDType = "TAX_ID_TYPE_EIN"
)

func (e LegalEntityCreateTaxIDType) ToPointer() *LegalEntityCreateTaxIDType {
	return &e
}

// LegalEntityCreate - A legal entity. Legal entities are organizations, such as companies, that participate in financial transactions
type LegalEntityCreate struct {
	// Indicates whether the entity is an accredited investor. By default, this is set to `false`.
	AccreditedInvestor *bool `json:"accredited_investor,omitempty"`
	// Indicates whether the entity is an adviser. By default, this is set to `false`.
	Adviser *bool `json:"adviser,omitempty"`
	// Indicates whether the entity is a broker dealer. By default, this is set to `false`.
	BrokerDealer                     *bool                             `json:"broker_dealer,omitempty"`
	BusinessIndustrialClassification *BusinessIndustrialClassification `json:"business_industrial_classification,omitempty"`
	// Corporate structure of the entity.
	CorporateStructure *CorporateStructure `json:"corporate_structure,omitempty"`
	// The correspondent id associated with the legal entity.
	CorrespondentID string `json:"correspondent_id"`
	// DBA (Doing Business As) names. Can list up to 5 associated with the Legal Entity
	DoingBusinessAs []string `json:"doing_business_as,omitempty"`
	// Due Diligence for Legal Entities required when a Legal Entity is the Primary Owner on an Account.
	EntityDueDiligence *EntityDueDiligenceCreate `json:"entity_due_diligence,omitempty"`
	// The legal entity name.
	EntityName string `json:"entity_name"`
	// The entity type.
	EntityType EntityType `json:"entity_type"`
	// The reason the customer is exempt from verifying beneficial owners, if applicable.
	ExemptCustomerReason *ExemptCustomerReason `json:"exempt_customer_reason,omitempty"`
	// Indicates whether the entity is exempt from verifying beneficial owners. By default, this is set to `false`.
	ExemptVerifyingBeneficialOwners *bool `json:"exempt_verifying_beneficial_owners,omitempty"`
	// If the legal entity is a trust, they may set this field to convey ownership and value to a trustee.
	ForTheBenefitOf *string `json:"for_the_benefit_of,omitempty"`
	// Indicates whether the entity is a foreign financial institution. By default, this is set to `false`.
	ForeignFinancialInstitution *bool `json:"foreign_financial_institution,omitempty"`
	// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following:
	//
	//  * A full date, with non-zero year, month, and day values * A month and day value, with a zero year, such as an anniversary * A year on its own, with zero month and day values * A year and month value, with a zero day, such as a credit card expiration date
	//
	//  Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and `google.protobuf.Timestamp`.
	FormationDate *DateCreate `json:"formation_date,omitempty"`
	// Indicates whether the entity is an institutional customer
	InstitutionalCustomer *bool `json:"institutional_customer,omitempty"`
	// A large trader.
	LargeTrader *LargeTraderCreate `json:"large_trader,omitempty"`
	// Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains).
	//
	//  In typical usage an address would be created via user input or from importing existing data, depending on the type of process.
	//
	//  Advice on address input / editing: - Use an i18n-ready address widget such as  https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of  fields outside countries where that field is used.
	//
	//  For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
	LegalAddress PostalAddressCreate `json:"legal_address"`
	// The Legal Entity Identifier (LEI) is the financial industry term for a unique global identifier for legal entities participating in financial transactions
	LeiCode *string `json:"lei_code,omitempty"`
	// The operational footprint of an entity. Operating regions encompass all countries and regions where a company has a significant business presence This includes locations with physical offices, manufacturing plants, service centers, and sales and marketing activities Regions must be provided as two-character CLDR country codes
	OperatingRegions []string `json:"operating_regions"`
	// The legal home of an entity. A region of registration, in the context of a corporation, refers to the specific geographic area where the corporation is legally registered and incorporated Defines the legal jurisdiction and framework under which the corporation operates, including legal regulations, tax obligations, and compliance requirements Region must be provided as a two-character CLDR country code
	RegistrationRegion string `json:"registration_region"`
	// Indicates whether the entity is a regulated investment company. By default, this is set to `false`.
	RegulatedInvestmentCompany *bool `json:"regulated_investment_company,omitempty"`
	// Document ids related to the legal entity. At least one is required for RIA correspondents when creating Estate or Trust accounts.
	RelatedDocumentIds []string `json:"related_document_ids,omitempty"`
	// Indicates whether the trust is a revocable trust. By default, this is set to `false`.
	RevocableTrust *bool `json:"revocable_trust,omitempty"`
	// Boolean indicator whether the LE is subject to backup withholding
	SubjectToBackupWithholding *bool `json:"subject_to_backup_withholding,omitempty"`
	// The full U.S. tax ID for a related entity; Must be provided with `EIN` tax ID type
	TaxID string `json:"tax_id"`
	// The nature of the U.S. Tax ID indicated in the related tax_id field; Examples include ITIN, SSN, EIN.
	TaxIDType LegalEntityCreateTaxIDType `json:"tax_id_type"`
	// Tax Profile pertaining to the Legal Entity or Natural Person.
	TaxProfile TaxProfileCreate `json:"tax_profile"`
}

func (o *LegalEntityCreate) GetAccreditedInvestor() *bool {
	if o == nil {
		return nil
	}
	return o.AccreditedInvestor
}

func (o *LegalEntityCreate) GetAdviser() *bool {
	if o == nil {
		return nil
	}
	return o.Adviser
}

func (o *LegalEntityCreate) GetBrokerDealer() *bool {
	if o == nil {
		return nil
	}
	return o.BrokerDealer
}

func (o *LegalEntityCreate) GetBusinessIndustrialClassification() *BusinessIndustrialClassification {
	if o == nil {
		return nil
	}
	return o.BusinessIndustrialClassification
}

func (o *LegalEntityCreate) GetCorporateStructure() *CorporateStructure {
	if o == nil {
		return nil
	}
	return o.CorporateStructure
}

func (o *LegalEntityCreate) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *LegalEntityCreate) GetDoingBusinessAs() []string {
	if o == nil {
		return nil
	}
	return o.DoingBusinessAs
}

func (o *LegalEntityCreate) GetEntityDueDiligence() *EntityDueDiligenceCreate {
	if o == nil {
		return nil
	}
	return o.EntityDueDiligence
}

func (o *LegalEntityCreate) GetEntityName() string {
	if o == nil {
		return ""
	}
	return o.EntityName
}

func (o *LegalEntityCreate) GetEntityType() EntityType {
	if o == nil {
		return EntityType("")
	}
	return o.EntityType
}

func (o *LegalEntityCreate) GetExemptCustomerReason() *ExemptCustomerReason {
	if o == nil {
		return nil
	}
	return o.ExemptCustomerReason
}

func (o *LegalEntityCreate) GetExemptVerifyingBeneficialOwners() *bool {
	if o == nil {
		return nil
	}
	return o.ExemptVerifyingBeneficialOwners
}

func (o *LegalEntityCreate) GetForTheBenefitOf() *string {
	if o == nil {
		return nil
	}
	return o.ForTheBenefitOf
}

func (o *LegalEntityCreate) GetForeignFinancialInstitution() *bool {
	if o == nil {
		return nil
	}
	return o.ForeignFinancialInstitution
}

func (o *LegalEntityCreate) GetFormationDate() *DateCreate {
	if o == nil {
		return nil
	}
	return o.FormationDate
}

func (o *LegalEntityCreate) GetInstitutionalCustomer() *bool {
	if o == nil {
		return nil
	}
	return o.InstitutionalCustomer
}

func (o *LegalEntityCreate) GetLargeTrader() *LargeTraderCreate {
	if o == nil {
		return nil
	}
	return o.LargeTrader
}

func (o *LegalEntityCreate) GetLegalAddress() PostalAddressCreate {
	if o == nil {
		return PostalAddressCreate{}
	}
	return o.LegalAddress
}

func (o *LegalEntityCreate) GetLeiCode() *string {
	if o == nil {
		return nil
	}
	return o.LeiCode
}

func (o *LegalEntityCreate) GetOperatingRegions() []string {
	if o == nil {
		return []string{}
	}
	return o.OperatingRegions
}

func (o *LegalEntityCreate) GetRegistrationRegion() string {
	if o == nil {
		return ""
	}
	return o.RegistrationRegion
}

func (o *LegalEntityCreate) GetRegulatedInvestmentCompany() *bool {
	if o == nil {
		return nil
	}
	return o.RegulatedInvestmentCompany
}

func (o *LegalEntityCreate) GetRelatedDocumentIds() []string {
	if o == nil {
		return nil
	}
	return o.RelatedDocumentIds
}

func (o *LegalEntityCreate) GetRevocableTrust() *bool {
	if o == nil {
		return nil
	}
	return o.RevocableTrust
}

func (o *LegalEntityCreate) GetSubjectToBackupWithholding() *bool {
	if o == nil {
		return nil
	}
	return o.SubjectToBackupWithholding
}

func (o *LegalEntityCreate) GetTaxID() string {
	if o == nil {
		return ""
	}
	return o.TaxID
}

func (o *LegalEntityCreate) GetTaxIDType() LegalEntityCreateTaxIDType {
	if o == nil {
		return LegalEntityCreateTaxIDType("")
	}
	return o.TaxIDType
}

func (o *LegalEntityCreate) GetTaxProfile() TaxProfileCreate {
	if o == nil {
		return TaxProfileCreate{}
	}
	return o.TaxProfile
}
