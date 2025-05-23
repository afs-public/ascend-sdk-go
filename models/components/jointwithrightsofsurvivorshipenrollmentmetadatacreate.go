// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
type JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan string

const (
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll                JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentDecline               JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan) ToPointer() *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan {
	return &e
}

// JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep - Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
type JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep string

const (
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweepAutoEnrollFdicCashSweepUnspecified JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep = "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED"
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll                JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_ENROLL"
	JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweepFdicCashSweepDecline               JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_DECLINE"
)

func (e JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep) ToPointer() *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep {
	return &e
}

// JointWithRightsOfSurvivorshipEnrollmentMetadataCreate - Enrollment metadata for the With Right of Survivorship enrollment type
type JointWithRightsOfSurvivorshipEnrollmentMetadataCreate struct {
	// Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
	DividendReinvestmentPlan *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
	// Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
	FdicCashSweep *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep `json:"fdic_cash_sweep,omitempty"`
}

func (o *JointWithRightsOfSurvivorshipEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateDividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}

func (o *JointWithRightsOfSurvivorshipEnrollmentMetadataCreate) GetFdicCashSweep() *JointWithRightsOfSurvivorshipEnrollmentMetadataCreateFdicCashSweep {
	if o == nil {
		return nil
	}
	return o.FdicCashSweep
}
