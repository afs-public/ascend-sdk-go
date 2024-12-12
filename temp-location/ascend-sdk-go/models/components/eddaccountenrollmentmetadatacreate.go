// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// DeterminedAccountRiskRating - The client determined account risk rating of the entity customer
type DeterminedAccountRiskRating string

const (
	DeterminedAccountRiskRatingDeterminedAccountRiskRatingUnspecified DeterminedAccountRiskRating = "DETERMINED_ACCOUNT_RISK_RATING_UNSPECIFIED"
	DeterminedAccountRiskRatingLow                                    DeterminedAccountRiskRating = "LOW"
	DeterminedAccountRiskRatingMedium                                 DeterminedAccountRiskRating = "MEDIUM"
	DeterminedAccountRiskRatingHigh                                   DeterminedAccountRiskRating = "HIGH"
)

func (e DeterminedAccountRiskRating) ToPointer() *DeterminedAccountRiskRating {
	return &e
}
func (e *DeterminedAccountRiskRating) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DETERMINED_ACCOUNT_RISK_RATING_UNSPECIFIED":
		fallthrough
	case "LOW":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		*e = DeterminedAccountRiskRating(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeterminedAccountRiskRating: %v", v)
	}
}

// EddAccountEnrollmentMetadataCreate - Enrollment metadata for Entity Accounts
type EddAccountEnrollmentMetadataCreate struct {
	// The initial amount of money placed into the account by the entity upon or after the account's establishment.
	DepositedFunds DepositedFundsCreate `json:"deposited_funds"`
	// The client determined account risk rating of the entity customer
	DeterminedAccountRiskRating DeterminedAccountRiskRating `json:"determined_account_risk_rating"`
	// Disclosure of the entity account owner's financial relationships and source of brokerage funds; facilitates the creation of the overall customer risk profile
	FinancialProfile FinancialProfileCreate `json:"financial_profile"`
	// Details the customer's intended trading and banking-related activities at the time of account application; informs risk checks and forms a baseline for anomalous activity detection
	PlannedActivity PlannedActivityCreate `json:"planned_activity"`
	// Details surrounding the related politically exposed persons
	RelatedPepDetails RelatedPepDetailsCreate `json:"related_pep_details"`
	// The scope of the business for the entity customer
	ScopeOfBusiness string `json:"scope_of_business"`
}

func (o *EddAccountEnrollmentMetadataCreate) GetDepositedFunds() DepositedFundsCreate {
	if o == nil {
		return DepositedFundsCreate{}
	}
	return o.DepositedFunds
}

func (o *EddAccountEnrollmentMetadataCreate) GetDeterminedAccountRiskRating() DeterminedAccountRiskRating {
	if o == nil {
		return DeterminedAccountRiskRating("")
	}
	return o.DeterminedAccountRiskRating
}

func (o *EddAccountEnrollmentMetadataCreate) GetFinancialProfile() FinancialProfileCreate {
	if o == nil {
		return FinancialProfileCreate{}
	}
	return o.FinancialProfile
}

func (o *EddAccountEnrollmentMetadataCreate) GetPlannedActivity() PlannedActivityCreate {
	if o == nil {
		return PlannedActivityCreate{}
	}
	return o.PlannedActivity
}

func (o *EddAccountEnrollmentMetadataCreate) GetRelatedPepDetails() RelatedPepDetailsCreate {
	if o == nil {
		return RelatedPepDetailsCreate{}
	}
	return o.RelatedPepDetails
}

func (o *EddAccountEnrollmentMetadataCreate) GetScopeOfBusiness() string {
	if o == nil {
		return ""
	}
	return o.ScopeOfBusiness
}
