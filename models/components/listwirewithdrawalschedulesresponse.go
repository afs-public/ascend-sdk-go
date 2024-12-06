// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListWireWithdrawalSchedulesResponse - A paged response containing a list of Wire withdrawal transfer schedules
type ListWireWithdrawalSchedulesResponse struct {
	// The next page token
	NextPageToken *string `json:"next_page_token,omitempty"`
	// The list of transfer schedules
	WireWithdrawalSchedules []WireWithdrawalSchedule `json:"wire_withdrawal_schedules,omitempty"`
}

func (o *ListWireWithdrawalSchedulesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *ListWireWithdrawalSchedulesResponse) GetWireWithdrawalSchedules() []WireWithdrawalSchedule {
	if o == nil {
		return nil
	}
	return o.WireWithdrawalSchedules
}