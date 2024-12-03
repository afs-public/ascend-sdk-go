// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WireWithdrawalScheduleCreate - A withdrawal transfer schedule using the Wire mechanism
type WireWithdrawalScheduleCreate struct {
	// The person or entity taking receipt of the wired funds
	Beneficiary WireWithdrawalBeneficiaryCreate `json:"beneficiary"`
	// An intermediary party
	Intermediary *WireWithdrawalIntermediaryCreate `json:"intermediary,omitempty"`
	// A recipient bank / financial institution
	RecipientBank WireWithdrawalRecipientBankCreate `json:"recipient_bank"`
	// A distribution from a retirement account.
	RetirementDistribution *RetirementDistributionCreate `json:"retirement_distribution,omitempty"`
	// Details of withdrawal schedule transfers
	ScheduleDetails WithdrawalScheduleDetailsCreate `json:"schedule_details"`
}

func (o *WireWithdrawalScheduleCreate) GetBeneficiary() WireWithdrawalBeneficiaryCreate {
	if o == nil {
		return WireWithdrawalBeneficiaryCreate{}
	}
	return o.Beneficiary
}

func (o *WireWithdrawalScheduleCreate) GetIntermediary() *WireWithdrawalIntermediaryCreate {
	if o == nil {
		return nil
	}
	return o.Intermediary
}

func (o *WireWithdrawalScheduleCreate) GetRecipientBank() WireWithdrawalRecipientBankCreate {
	if o == nil {
		return WireWithdrawalRecipientBankCreate{}
	}
	return o.RecipientBank
}

func (o *WireWithdrawalScheduleCreate) GetRetirementDistribution() *RetirementDistributionCreate {
	if o == nil {
		return nil
	}
	return o.RetirementDistribution
}

func (o *WireWithdrawalScheduleCreate) GetScheduleDetails() WithdrawalScheduleDetailsCreate {
	if o == nil {
		return WithdrawalScheduleDetailsCreate{}
	}
	return o.ScheduleDetails
}
