// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListAchDepositSchedulesResponse - A paged response containing a list of ACH deposit transfer schedules
type ListAchDepositSchedulesResponse struct {
	// The list of transfer schedules
	AchDepositSchedules []AchDepositSchedule `json:"ach_deposit_schedules,omitempty"`
	// The next page token
	NextPageToken *string `json:"next_page_token,omitempty"`
}

func (o *ListAchDepositSchedulesResponse) GetAchDepositSchedules() []AchDepositSchedule {
	if o == nil {
		return nil
	}
	return o.AchDepositSchedules
}

func (o *ListAchDepositSchedulesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
