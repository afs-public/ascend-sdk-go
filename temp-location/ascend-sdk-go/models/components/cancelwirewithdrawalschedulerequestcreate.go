// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CancelWireWithdrawalScheduleRequestCreate - Request to cancel a Wire withdrawal transfer schedule
type CancelWireWithdrawalScheduleRequestCreate struct {
	// A comment as to why the Wire withdrawal schedule is being canceled
	Comment *string `json:"comment,omitempty"`
	// The name of the Wire withdrawal transfer schedule to cancel
	Name string `json:"name"`
}

func (o *CancelWireWithdrawalScheduleRequestCreate) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *CancelWireWithdrawalScheduleRequestCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
