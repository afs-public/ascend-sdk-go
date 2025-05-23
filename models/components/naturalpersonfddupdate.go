// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// NaturalPersonFddUpdate - Foreign Due Diligence for Legal Natural Persons required when a Legal Natural Person is the Primary Owner on a non-resident/non-citizen Account.
type NaturalPersonFddUpdate struct {
	// Customer Non-referral Source
	CustomerNonReferralSource *string `json:"customer_non_referral_source,omitempty"`
	// Customer Referral Source
	CustomerReferralSource *CustomerReferralSourceUpdate `json:"customer_referral_source,omitempty"`
	// The description of the applicant's source of wealth
	EmploymentAndEmployerDescription *string `json:"employment_and_employer_description,omitempty"`
	// Negative News detail.
	NegativeNews *NegativeNewsUpdate `json:"negative_news,omitempty"`
	// Applicant's other source of wealth
	OtherSourcesOfWealth *OtherSourcesOfWealthUpdate `json:"other_sources_of_wealth,omitempty"`
}

func (o *NaturalPersonFddUpdate) GetCustomerNonReferralSource() *string {
	if o == nil {
		return nil
	}
	return o.CustomerNonReferralSource
}

func (o *NaturalPersonFddUpdate) GetCustomerReferralSource() *CustomerReferralSourceUpdate {
	if o == nil {
		return nil
	}
	return o.CustomerReferralSource
}

func (o *NaturalPersonFddUpdate) GetEmploymentAndEmployerDescription() *string {
	if o == nil {
		return nil
	}
	return o.EmploymentAndEmployerDescription
}

func (o *NaturalPersonFddUpdate) GetNegativeNews() *NegativeNewsUpdate {
	if o == nil {
		return nil
	}
	return o.NegativeNews
}

func (o *NaturalPersonFddUpdate) GetOtherSourcesOfWealth() *OtherSourcesOfWealthUpdate {
	if o == nil {
		return nil
	}
	return o.OtherSourcesOfWealth
}
