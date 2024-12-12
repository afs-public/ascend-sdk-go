// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AchDepositScheduleUpdate - A deposit transfer schedule using the ACH mechanism
type AchDepositScheduleUpdate struct {
	// Details of deposit schedule transfers
	ScheduleDetails *DepositScheduleDetailsUpdate `json:"schedule_details,omitempty"`
}

func (o *AchDepositScheduleUpdate) GetScheduleDetails() *DepositScheduleDetailsUpdate {
	if o == nil {
		return nil
	}
	return o.ScheduleDetails
}
