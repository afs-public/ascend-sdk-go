// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type AchDepositSchedulesCancelAchDepositScheduleRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The achDepositSchedule id.
	AchDepositScheduleID                  string                                           `pathParam:"style=simple,explode=false,name=achDepositSchedule_id"`
	CancelAchDepositScheduleRequestCreate components.CancelAchDepositScheduleRequestCreate `request:"mediaType=application/json"`
}

func (o *AchDepositSchedulesCancelAchDepositScheduleRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchDepositSchedulesCancelAchDepositScheduleRequest) GetAchDepositScheduleID() string {
	if o == nil {
		return ""
	}
	return o.AchDepositScheduleID
}

func (o *AchDepositSchedulesCancelAchDepositScheduleRequest) GetCancelAchDepositScheduleRequestCreate() components.CancelAchDepositScheduleRequestCreate {
	if o == nil {
		return components.CancelAchDepositScheduleRequestCreate{}
	}
	return o.CancelAchDepositScheduleRequestCreate
}

type AchDepositSchedulesCancelAchDepositScheduleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchDepositSchedule *components.AchDepositSchedule
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchDepositSchedulesCancelAchDepositScheduleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchDepositSchedulesCancelAchDepositScheduleResponse) GetAchDepositSchedule() *components.AchDepositSchedule {
	if o == nil {
		return nil
	}
	return o.AchDepositSchedule
}

func (o *AchDepositSchedulesCancelAchDepositScheduleResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
