// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NonCitizenResidencyUpdateResidencyStatus string

const (
	NonCitizenResidencyUpdateResidencyStatusResidencyStatusUnspecified NonCitizenResidencyUpdateResidencyStatus = "RESIDENCY_STATUS_UNSPECIFIED"
	NonCitizenResidencyUpdateResidencyStatusUsPermanentResident        NonCitizenResidencyUpdateResidencyStatus = "US_PERMANENT_RESIDENT"
	NonCitizenResidencyUpdateResidencyStatusUsTemporaryResident        NonCitizenResidencyUpdateResidencyStatus = "US_TEMPORARY_RESIDENT"
	NonCitizenResidencyUpdateResidencyStatusUsNonResident              NonCitizenResidencyUpdateResidencyStatus = "US_NON_RESIDENT"
)

func (e NonCitizenResidencyUpdateResidencyStatus) ToPointer() *NonCitizenResidencyUpdateResidencyStatus {
	return &e
}

// NonCitizenResidencyUpdate - Non Citizenship Residency to facilitate non-Citizen lawful US residents to open domestic accounts.
type NonCitizenResidencyUpdate struct {
	ResidencyStatus *NonCitizenResidencyUpdateResidencyStatus `json:"residency_status,omitempty"`
}

func (o *NonCitizenResidencyUpdate) GetResidencyStatus() *NonCitizenResidencyUpdateResidencyStatus {
	if o == nil {
		return nil
	}
	return o.ResidencyStatus
}
