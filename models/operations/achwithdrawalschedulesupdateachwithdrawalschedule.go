// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AchWithdrawalSchedulesUpdateAchWithdrawalScheduleRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The achWithdrawalSchedule id.
	AchWithdrawalScheduleID string `pathParam:"style=simple,explode=false,name=achWithdrawalSchedule_id"`
	// A field mask representing the update. Note: only the 'schedule_details.amount' field of a schedule is currently updatable
	UpdateMask                  *string                                `queryParam:"style=form,explode=true,name=update_mask"`
	AchWithdrawalScheduleUpdate components.AchWithdrawalScheduleUpdate `request:"mediaType=application/json"`
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleRequest) GetAchWithdrawalScheduleID() string {
	if o == nil {
		return ""
	}
	return o.AchWithdrawalScheduleID
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleRequest) GetAchWithdrawalScheduleUpdate() components.AchWithdrawalScheduleUpdate {
	if o == nil {
		return components.AchWithdrawalScheduleUpdate{}
	}
	return o.AchWithdrawalScheduleUpdate
}

type AchWithdrawalSchedulesUpdateAchWithdrawalScheduleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchWithdrawalSchedule *components.AchWithdrawalSchedule
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleResponse) GetAchWithdrawalSchedule() *components.AchWithdrawalSchedule {
	if o == nil {
		return nil
	}
	return o.AchWithdrawalSchedule
}

func (o *AchWithdrawalSchedulesUpdateAchWithdrawalScheduleResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
