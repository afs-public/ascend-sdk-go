// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type WireWithdrawalSchedulesGetWireWithdrawalScheduleRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The wireWithdrawalSchedule id.
	WireWithdrawalScheduleID string `pathParam:"style=simple,explode=false,name=wireWithdrawalSchedule_id"`
}

func (o *WireWithdrawalSchedulesGetWireWithdrawalScheduleRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *WireWithdrawalSchedulesGetWireWithdrawalScheduleRequest) GetWireWithdrawalScheduleID() string {
	if o == nil {
		return ""
	}
	return o.WireWithdrawalScheduleID
}

type WireWithdrawalSchedulesGetWireWithdrawalScheduleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	WireWithdrawalSchedule *components.WireWithdrawalSchedule
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *WireWithdrawalSchedulesGetWireWithdrawalScheduleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WireWithdrawalSchedulesGetWireWithdrawalScheduleResponse) GetWireWithdrawalSchedule() *components.WireWithdrawalSchedule {
	if o == nil {
		return nil
	}
	return o.WireWithdrawalSchedule
}

func (o *WireWithdrawalSchedulesGetWireWithdrawalScheduleResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
