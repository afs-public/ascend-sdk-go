// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TaxProfileUpdateFederalTaxClassification - Federal tax classification.
type TaxProfileUpdateFederalTaxClassification string

const (
	TaxProfileUpdateFederalTaxClassificationFederalTaxClassificationUnspecified TaxProfileUpdateFederalTaxClassification = "FEDERAL_TAX_CLASSIFICATION_UNSPECIFIED"
	TaxProfileUpdateFederalTaxClassificationIndivSolepropOrSinglememberllc      TaxProfileUpdateFederalTaxClassification = "INDIV_SOLEPROP_OR_SINGLEMEMBERLLC"
	TaxProfileUpdateFederalTaxClassificationPartnership                         TaxProfileUpdateFederalTaxClassification = "PARTNERSHIP"
	TaxProfileUpdateFederalTaxClassificationCCorporation                        TaxProfileUpdateFederalTaxClassification = "C_CORPORATION"
	TaxProfileUpdateFederalTaxClassificationSCorporation                        TaxProfileUpdateFederalTaxClassification = "S_CORPORATION"
	TaxProfileUpdateFederalTaxClassificationTrustEstate                         TaxProfileUpdateFederalTaxClassification = "TRUST_ESTATE"
	TaxProfileUpdateFederalTaxClassificationLlcTaxedAsCCorp                     TaxProfileUpdateFederalTaxClassification = "LLC_TAXED_AS_C_CORP"
	TaxProfileUpdateFederalTaxClassificationLlcTaxedAsSCorp                     TaxProfileUpdateFederalTaxClassification = "LLC_TAXED_AS_S_CORP"
	TaxProfileUpdateFederalTaxClassificationLlcTaxedAsPartnership               TaxProfileUpdateFederalTaxClassification = "LLC_TAXED_AS_PARTNERSHIP"
	TaxProfileUpdateFederalTaxClassificationOther                               TaxProfileUpdateFederalTaxClassification = "OTHER"
)

func (e TaxProfileUpdateFederalTaxClassification) ToPointer() *TaxProfileUpdateFederalTaxClassification {
	return &e
}

// TaxProfileUpdateIrsFormType - IRS form type.
type TaxProfileUpdateIrsFormType string

const (
	TaxProfileUpdateIrsFormTypeIrsFormTypeUnspecified TaxProfileUpdateIrsFormType = "IRS_FORM_TYPE_UNSPECIFIED"
	TaxProfileUpdateIrsFormTypeW9                     TaxProfileUpdateIrsFormType = "W_9"
)

func (e TaxProfileUpdateIrsFormType) ToPointer() *TaxProfileUpdateIrsFormType {
	return &e
}

// TaxProfileUpdateUsTinStatus - United States Individual Taxpayer Identification Number (ITIN) status.
type TaxProfileUpdateUsTinStatus string

const (
	TaxProfileUpdateUsTinStatusUsTinStatusUnspecified TaxProfileUpdateUsTinStatus = "US_TIN_STATUS_UNSPECIFIED"
	TaxProfileUpdateUsTinStatusPassing                TaxProfileUpdateUsTinStatus = "PASSING"
	TaxProfileUpdateUsTinStatusFailing                TaxProfileUpdateUsTinStatus = "FAILING"
)

func (e TaxProfileUpdateUsTinStatus) ToPointer() *TaxProfileUpdateUsTinStatus {
	return &e
}

// TaxProfileUpdate - Tax Profile pertaining to the Legal Entity or Natural Person.
type TaxProfileUpdate struct {
	// Federal tax classification.
	FederalTaxClassification *TaxProfileUpdateFederalTaxClassification `json:"federal_tax_classification,omitempty"`
	// IRS form type.
	IrsFormType *TaxProfileUpdateIrsFormType `json:"irs_form_type,omitempty"`
	// Legal tax region must be "US" if provided W-9, otherwise must be a non-US country.
	LegalTaxRegionCode *string `json:"legal_tax_region_code,omitempty"`
	// United States Individual Taxpayer Identification Number (ITIN) status.
	UsTinStatus *TaxProfileUpdateUsTinStatus `json:"us_tin_status,omitempty"`
}

func (o *TaxProfileUpdate) GetFederalTaxClassification() *TaxProfileUpdateFederalTaxClassification {
	if o == nil {
		return nil
	}
	return o.FederalTaxClassification
}

func (o *TaxProfileUpdate) GetIrsFormType() *TaxProfileUpdateIrsFormType {
	if o == nil {
		return nil
	}
	return o.IrsFormType
}

func (o *TaxProfileUpdate) GetLegalTaxRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.LegalTaxRegionCode
}

func (o *TaxProfileUpdate) GetUsTinStatus() *TaxProfileUpdateUsTinStatus {
	if o == nil {
		return nil
	}
	return o.UsTinStatus
}
