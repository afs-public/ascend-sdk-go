// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
type IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan string

const (
	IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll                IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentDecline               IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan) ToPointer() *IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan {
	return &e
}

// IRASimpleEnrollmentMetadataCreateFdicCashSweep - Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
type IRASimpleEnrollmentMetadataCreateFdicCashSweep string

const (
	IRASimpleEnrollmentMetadataCreateFdicCashSweepAutoEnrollFdicCashSweepUnspecified IRASimpleEnrollmentMetadataCreateFdicCashSweep = "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED"
	IRASimpleEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll                IRASimpleEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_ENROLL"
	IRASimpleEnrollmentMetadataCreateFdicCashSweepFdicCashSweepDecline               IRASimpleEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_DECLINE"
)

func (e IRASimpleEnrollmentMetadataCreateFdicCashSweep) ToPointer() *IRASimpleEnrollmentMetadataCreateFdicCashSweep {
	return &e
}

// IRASimpleEnrollmentMetadataCreate - Enrollment metadata for Simple IRA accounts enrollment type
type IRASimpleEnrollmentMetadataCreate struct {
	// Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
	DividendReinvestmentPlan *IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
	// Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
	FdicCashSweep *IRASimpleEnrollmentMetadataCreateFdicCashSweep `json:"fdic_cash_sweep,omitempty"`
}

func (o *IRASimpleEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *IRASimpleEnrollmentMetadataCreateDividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}

func (o *IRASimpleEnrollmentMetadataCreate) GetFdicCashSweep() *IRASimpleEnrollmentMetadataCreateFdicCashSweep {
	if o == nil {
		return nil
	}
	return o.FdicCashSweep
}
