// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Mechanism - Cash transfer mechanism to search constraints for
type Mechanism string

const (
	MechanismAch Mechanism = "ACH"
	MechanismIct Mechanism = "ICT"
)

func (e Mechanism) ToPointer() *Mechanism {
	return &e
}
func (e *Mechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACH":
		fallthrough
	case "ICT":
		*e = Mechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mechanism: %v", v)
	}
}

// RetrieveContributionConstraintsRequestCreate - Request to retrieve retirement contribution constraints
type RetrieveContributionConstraintsRequestCreate struct {
	// Cash transfer mechanism to search constraints for
	Mechanism Mechanism `json:"mechanism"`
	// Name of the account being queried, for retirement contribution constraints Format: accounts/{account}
	Name string `json:"name"`
}

func (o *RetrieveContributionConstraintsRequestCreate) GetMechanism() Mechanism {
	if o == nil {
		return Mechanism("")
	}
	return o.Mechanism
}

func (o *RetrieveContributionConstraintsRequestCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}