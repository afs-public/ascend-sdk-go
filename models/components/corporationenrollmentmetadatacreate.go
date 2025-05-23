// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
type DividendReinvestmentPlan string

const (
	DividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified DividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	DividendReinvestmentPlanDividendReinvestmentEnroll                DividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	DividendReinvestmentPlanDividendReinvestmentDecline               DividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e DividendReinvestmentPlan) ToPointer() *DividendReinvestmentPlan {
	return &e
}

// FdicCashSweep - Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
type FdicCashSweep string

const (
	FdicCashSweepAutoEnrollFdicCashSweepUnspecified FdicCashSweep = "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED"
	FdicCashSweepFdicCashSweepEnroll                FdicCashSweep = "FDIC_CASH_SWEEP_ENROLL"
	FdicCashSweepFdicCashSweepDecline               FdicCashSweep = "FDIC_CASH_SWEEP_DECLINE"
)

func (e FdicCashSweep) ToPointer() *FdicCashSweep {
	return &e
}

type CorporationEnrollmentMetadataCreate struct {
	// Option to auto-enroll in Dividend Reinvestment; defaults to DIVIDEND_REINVESTMENT_ENROLL
	DividendReinvestmentPlan *DividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
	// Enrollment metadata for Entity Accounts
	EddAccountEnrollmentMetadata *EddAccountEnrollmentMetadataCreate `json:"edd_account_enrollment_metadata,omitempty"`
	// Option to auto-enroll in FDIC cash sweep; defaults to FDIC_CASH_SWEEP_ENROLL
	FdicCashSweep *FdicCashSweep `json:"fdic_cash_sweep,omitempty"`
}

func (o *CorporationEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *DividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}

func (o *CorporationEnrollmentMetadataCreate) GetEddAccountEnrollmentMetadata() *EddAccountEnrollmentMetadataCreate {
	if o == nil {
		return nil
	}
	return o.EddAccountEnrollmentMetadata
}

func (o *CorporationEnrollmentMetadataCreate) GetFdicCashSweep() *FdicCashSweep {
	if o == nil {
		return nil
	}
	return o.FdicCashSweep
}
