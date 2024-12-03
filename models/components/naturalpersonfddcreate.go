// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// NaturalPersonFddCreate - Foreign Due Diligence for Legal Natural Persons required when a Legal Natural Person is the Primary Owner on a non-resident/non-citizen Account.
type NaturalPersonFddCreate struct {
	// Customer Non-referral Source
	CustomerNonReferralSource *string `json:"customer_non_referral_source,omitempty"`
	// Customer Referral Source
	CustomerReferralSource *CustomerReferralSourceCreate `json:"customer_referral_source,omitempty"`
	// The description of the applicant's source of wealth
	EmploymentAndEmployerDescription string `json:"employment_and_employer_description"`
	// Negative News detail.
	NegativeNews NegativeNewsCreate `json:"negative_news"`
	// Applicant's other source of wealth
	OtherSourcesOfWealth OtherSourcesOfWealthCreate `json:"other_sources_of_wealth"`
}

func (o *NaturalPersonFddCreate) GetCustomerNonReferralSource() *string {
	if o == nil {
		return nil
	}
	return o.CustomerNonReferralSource
}

func (o *NaturalPersonFddCreate) GetCustomerReferralSource() *CustomerReferralSourceCreate {
	if o == nil {
		return nil
	}
	return o.CustomerReferralSource
}

func (o *NaturalPersonFddCreate) GetEmploymentAndEmployerDescription() string {
	if o == nil {
		return ""
	}
	return o.EmploymentAndEmployerDescription
}

func (o *NaturalPersonFddCreate) GetNegativeNews() NegativeNewsCreate {
	if o == nil {
		return NegativeNewsCreate{}
	}
	return o.NegativeNews
}

func (o *NaturalPersonFddCreate) GetOtherSourcesOfWealth() OtherSourcesOfWealthCreate {
	if o == nil {
		return OtherSourcesOfWealthCreate{}
	}
	return o.OtherSourcesOfWealth
}
