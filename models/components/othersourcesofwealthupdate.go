// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// OtherSourcesOfWealthUpdate - Applicant's other source of wealth
type OtherSourcesOfWealthUpdate struct {
	// Indicates whether the applicant has other sources of wealth.
	ApplicantHasOtherSourcesOfWealth *bool `json:"applicant_has_other_sources_of_wealth,omitempty"`
	// The applicant's other source of wealth description. If the applicant has no other sources of wealth, they must specify "N/A."
	OtherSourcesOfWealth *string `json:"other_sources_of_wealth,omitempty"`
	// The applicant's other source of wealth verification. If the applicant has no other sources of wealth, they must specify "N/A."
	OtherSourcesOfWealthVerification *string `json:"other_sources_of_wealth_verification,omitempty"`
}

func (o *OtherSourcesOfWealthUpdate) GetApplicantHasOtherSourcesOfWealth() *bool {
	if o == nil {
		return nil
	}
	return o.ApplicantHasOtherSourcesOfWealth
}

func (o *OtherSourcesOfWealthUpdate) GetOtherSourcesOfWealth() *string {
	if o == nil {
		return nil
	}
	return o.OtherSourcesOfWealth
}

func (o *OtherSourcesOfWealthUpdate) GetOtherSourcesOfWealthVerification() *string {
	if o == nil {
		return nil
	}
	return o.OtherSourcesOfWealthVerification
}
