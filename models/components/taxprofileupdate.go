// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

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
func (e *TaxProfileUpdateFederalTaxClassification) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FEDERAL_TAX_CLASSIFICATION_UNSPECIFIED":
		fallthrough
	case "INDIV_SOLEPROP_OR_SINGLEMEMBERLLC":
		fallthrough
	case "PARTNERSHIP":
		fallthrough
	case "C_CORPORATION":
		fallthrough
	case "S_CORPORATION":
		fallthrough
	case "TRUST_ESTATE":
		fallthrough
	case "LLC_TAXED_AS_C_CORP":
		fallthrough
	case "LLC_TAXED_AS_S_CORP":
		fallthrough
	case "LLC_TAXED_AS_PARTNERSHIP":
		fallthrough
	case "OTHER":
		*e = TaxProfileUpdateFederalTaxClassification(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxProfileUpdateFederalTaxClassification: %v", v)
	}
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
func (e *TaxProfileUpdateIrsFormType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IRS_FORM_TYPE_UNSPECIFIED":
		fallthrough
	case "W_9":
		*e = TaxProfileUpdateIrsFormType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxProfileUpdateIrsFormType: %v", v)
	}
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
func (e *TaxProfileUpdateUsTinStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US_TIN_STATUS_UNSPECIFIED":
		fallthrough
	case "PASSING":
		fallthrough
	case "FAILING":
		*e = TaxProfileUpdateUsTinStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaxProfileUpdateUsTinStatus: %v", v)
	}
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
