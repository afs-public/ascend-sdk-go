// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CustomerReferralSourceUpdate - Customer Referral Source
type CustomerReferralSourceUpdate struct {
	// The name of the referrer
	Name *string `json:"name,omitempty"`
	// The relationship of the referrer to the applicant
	RelationshipToApplicant *string `json:"relationship_to_applicant,omitempty"`
	// The years the referrer has known the applicant If the referrer has known the applicant for less than a year, they must specify 1
	RelationshipYearsWithApplicant *int `json:"relationship_years_with_applicant,omitempty"`
	// The years the referrer has known the broker If the referrer has known the broker for less than a year, they must specify 1
	RelationshipYearsWithBroker *int `json:"relationship_years_with_broker,omitempty"`
}

func (o *CustomerReferralSourceUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomerReferralSourceUpdate) GetRelationshipToApplicant() *string {
	if o == nil {
		return nil
	}
	return o.RelationshipToApplicant
}

func (o *CustomerReferralSourceUpdate) GetRelationshipYearsWithApplicant() *int {
	if o == nil {
		return nil
	}
	return o.RelationshipYearsWithApplicant
}

func (o *CustomerReferralSourceUpdate) GetRelationshipYearsWithBroker() *int {
	if o == nil {
		return nil
	}
	return o.RelationshipYearsWithBroker
}
