// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IctWithdrawalTravelRuleCreate - The travel rules associated with an ICT withdrawal
type IctWithdrawalTravelRuleCreate struct {
	// Institution representing originator or recipient of funds from an Instant Cash Transfer
	RecipientInstitution InstitutionCreate `json:"recipient_institution"`
}

func (o *IctWithdrawalTravelRuleCreate) GetRecipientInstitution() InstitutionCreate {
	if o == nil {
		return InstitutionCreate{}
	}
	return o.RecipientInstitution
}
