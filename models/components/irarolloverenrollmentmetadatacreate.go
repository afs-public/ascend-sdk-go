// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to true
type IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan string

const (
	IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll                IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentDecline               IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan) ToPointer() *IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan {
	return &e
}
func (e *IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED":
		fallthrough
	case "DIVIDEND_REINVESTMENT_ENROLL":
		fallthrough
	case "DIVIDEND_REINVESTMENT_DECLINE":
		*e = IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan: %v", v)
	}
}

// IRARolloverEnrollmentMetadataCreateFdicCashSweep - Option to auto-enroll in FDIC cash sweep; defaults to true
type IRARolloverEnrollmentMetadataCreateFdicCashSweep string

const (
	IRARolloverEnrollmentMetadataCreateFdicCashSweepAutoEnrollFdicCashSweepUnspecified IRARolloverEnrollmentMetadataCreateFdicCashSweep = "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED"
	IRARolloverEnrollmentMetadataCreateFdicCashSweepFdicCashSweepEnroll                IRARolloverEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_ENROLL"
	IRARolloverEnrollmentMetadataCreateFdicCashSweepFdicCashSweepDecline               IRARolloverEnrollmentMetadataCreateFdicCashSweep = "FDIC_CASH_SWEEP_DECLINE"
)

func (e IRARolloverEnrollmentMetadataCreateFdicCashSweep) ToPointer() *IRARolloverEnrollmentMetadataCreateFdicCashSweep {
	return &e
}
func (e *IRARolloverEnrollmentMetadataCreateFdicCashSweep) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUTO_ENROLL_FDIC_CASH_SWEEP_UNSPECIFIED":
		fallthrough
	case "FDIC_CASH_SWEEP_ENROLL":
		fallthrough
	case "FDIC_CASH_SWEEP_DECLINE":
		*e = IRARolloverEnrollmentMetadataCreateFdicCashSweep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IRARolloverEnrollmentMetadataCreateFdicCashSweep: %v", v)
	}
}

// IRARolloverEnrollmentMetadataCreate - Enrollment metadata for Rollover IRA accounts enrollment type
type IRARolloverEnrollmentMetadataCreate struct {
	// Option to auto-enroll in Dividend Reinvestment; defaults to true
	DividendReinvestmentPlan *IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
	// Option to auto-enroll in FDIC cash sweep; defaults to true
	FdicCashSweep *IRARolloverEnrollmentMetadataCreateFdicCashSweep `json:"fdic_cash_sweep,omitempty"`
}

func (o *IRARolloverEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *IRARolloverEnrollmentMetadataCreateDividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}

func (o *IRARolloverEnrollmentMetadataCreate) GetFdicCashSweep() *IRARolloverEnrollmentMetadataCreateFdicCashSweep {
	if o == nil {
		return nil
	}
	return o.FdicCashSweep
}
