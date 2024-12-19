// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TrustEnrollmentMetadataCreateDividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to true
type TrustEnrollmentMetadataCreateDividendReinvestmentPlan string

const (
	TrustEnrollmentMetadataCreateDividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified TrustEnrollmentMetadataCreateDividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	TrustEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll                TrustEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	TrustEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentDecline               TrustEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e TrustEnrollmentMetadataCreateDividendReinvestmentPlan) ToPointer() *TrustEnrollmentMetadataCreateDividendReinvestmentPlan {
	return &e
}

// TrustEnrollmentMetadataCreateFdicCashSweep - Option to auto-enroll in FDIC cash sweep; defaults to true
type TrustEnrollmentMetadataCreateFdicCashSweep string

const (
	TrustEnrollmentMetadataCreateFdicCashSweepAutoEnrollFdicCashSweepUnspecified TrustEnrollmentMetadataCreateFdicCashSweep = "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED"
	TrustEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll                TrustEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_ENROLL"
	TrustEnrollmentMetadataCreateFdicCashSweepFdicCashSweepDecline               TrustEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_DECLINE"
)

func (e TrustEnrollmentMetadataCreateFdicCashSweep) ToPointer() *TrustEnrollmentMetadataCreateFdicCashSweep {
	return &e
}

// OpenedOnBehalfOf - Trust account is opened on behalf of
type OpenedOnBehalfOf string

const (
	OpenedOnBehalfOfOpenedOnBehalfOfUnspecified OpenedOnBehalfOf = "OPENED_ON_BEHALF_OF_UNSPECIFIED"
	OpenedOnBehalfOfPersonalTrust               OpenedOnBehalfOf = "PERSONAL_TRUST"
	OpenedOnBehalfOfBusinessTrust               OpenedOnBehalfOf = "BUSINESS_TRUST"
	OpenedOnBehalfOfThirdPartyAdministrator     OpenedOnBehalfOf = "THIRD_PARTY_ADMINISTRATOR"
)

func (e OpenedOnBehalfOf) ToPointer() *OpenedOnBehalfOf {
	return &e
}

type TrustEnrollmentMetadataCreate struct {
	// Option to auto-enroll in Dividend Reinvestment; defaults to true
	DividendReinvestmentPlan *TrustEnrollmentMetadataCreateDividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
	// Option to auto-enroll in FDIC cash sweep; defaults to true
	FdicCashSweep *TrustEnrollmentMetadataCreateFdicCashSweep `json:"fdic_cash_sweep,omitempty"`
	// Trust account is opened on behalf of
	OpenedOnBehalfOf OpenedOnBehalfOf `json:"opened_on_behalf_of"`
}

func (o *TrustEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *TrustEnrollmentMetadataCreateDividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}

func (o *TrustEnrollmentMetadataCreate) GetFdicCashSweep() *TrustEnrollmentMetadataCreateFdicCashSweep {
	if o == nil {
		return nil
	}
	return o.FdicCashSweep
}

func (o *TrustEnrollmentMetadataCreate) GetOpenedOnBehalfOf() OpenedOnBehalfOf {
	if o == nil {
		return OpenedOnBehalfOf("")
	}
	return o.OpenedOnBehalfOf
}
