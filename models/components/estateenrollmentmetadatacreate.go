// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EstateEnrollmentMetadataCreateDividendReinvestmentPlan - Option to auto-enroll in Dividend Reinvestment; defaults to true
type EstateEnrollmentMetadataCreateDividendReinvestmentPlan string

const (
	EstateEnrollmentMetadataCreateDividendReinvestmentPlanAutoEnrollDividendReinvestmentUnspecified EstateEnrollmentMetadataCreateDividendReinvestmentPlan = "AUTO_ENROLL_DIVIDEND_REINVESTMENT_UNSPECIFIED"
	EstateEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentEnroll                EstateEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_ENROLL"
	EstateEnrollmentMetadataCreateDividendReinvestmentPlanDividendReinvestmentDecline               EstateEnrollmentMetadataCreateDividendReinvestmentPlan = "DIVIDEND_REINVESTMENT_DECLINE"
)

func (e EstateEnrollmentMetadataCreateDividendReinvestmentPlan) ToPointer() *EstateEnrollmentMetadataCreateDividendReinvestmentPlan {
	return &e
}

// EstateEnrollmentMetadataCreate - Enrollment metadata for estate enrollments
type EstateEnrollmentMetadataCreate struct {
	// The document id for the certificate of appointment
	CertificateOfAppointmentDocumentID *string `json:"certificate_of_appointment_document_id,omitempty"`
	// Option to auto-enroll in Dividend Reinvestment; defaults to true
	DividendReinvestmentPlan *EstateEnrollmentMetadataCreateDividendReinvestmentPlan `json:"dividend_reinvestment_plan,omitempty"`
}

func (o *EstateEnrollmentMetadataCreate) GetCertificateOfAppointmentDocumentID() *string {
	if o == nil {
		return nil
	}
	return o.CertificateOfAppointmentDocumentID
}

func (o *EstateEnrollmentMetadataCreate) GetDividendReinvestmentPlan() *EstateEnrollmentMetadataCreateDividendReinvestmentPlan {
	if o == nil {
		return nil
	}
	return o.DividendReinvestmentPlan
}