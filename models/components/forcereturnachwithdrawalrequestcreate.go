// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ForceReturnAchWithdrawalRequestCreate - Request to force a Nacha return on a completed ACH withdrawal. FOR TESTING ONLY!
type ForceReturnAchWithdrawalRequestCreate struct {
	// A return on an ACH transfer from the Nacha network.
	NachaReturn NachaReturnCreate `json:"nacha_return"`
	// The name of the ACH withdrawal to force a Nacha return on.
	Name string `json:"name"`
}

func (o *ForceReturnAchWithdrawalRequestCreate) GetNachaReturn() NachaReturnCreate {
	if o == nil {
		return NachaReturnCreate{}
	}
	return o.NachaReturn
}

func (o *ForceReturnAchWithdrawalRequestCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
