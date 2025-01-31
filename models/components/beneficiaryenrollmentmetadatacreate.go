// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// BeneficiaryEnrollmentMetadataCreate - Enrollment metadata for the BENEFICIARY_DESIGNATION enrollment type.
type BeneficiaryEnrollmentMetadataCreate struct {
	// Contingent Beneficiary list is optional, with a maximum of five contingent beneficiaries.
	ContingentBeneficiaries []BeneficiaryCreate `json:"contingent_beneficiaries,omitempty"`
	// At least one primary beneficiary must be provided, with a maximum of five primary beneficiaries.
	PrimaryBeneficiaries []BeneficiaryCreate `json:"primary_beneficiaries"`
}

func (o *BeneficiaryEnrollmentMetadataCreate) GetContingentBeneficiaries() []BeneficiaryCreate {
	if o == nil {
		return nil
	}
	return o.ContingentBeneficiaries
}

func (o *BeneficiaryEnrollmentMetadataCreate) GetPrimaryBeneficiaries() []BeneficiaryCreate {
	if o == nil {
		return []BeneficiaryCreate{}
	}
	return o.PrimaryBeneficiaries
}
