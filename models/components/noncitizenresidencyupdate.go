// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

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
func (e *NonCitizenResidencyUpdateResidencyStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RESIDENCY_STATUS_UNSPECIFIED":
		fallthrough
	case "US_PERMANENT_RESIDENT":
		fallthrough
	case "US_TEMPORARY_RESIDENT":
		fallthrough
	case "US_NON_RESIDENT":
		*e = NonCitizenResidencyUpdateResidencyStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NonCitizenResidencyUpdateResidencyStatus: %v", v)
	}
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
