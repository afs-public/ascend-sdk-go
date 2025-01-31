// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type WireWithdrawalsForceApproveWireWithdrawalRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The wireWithdrawal id.
	WireWithdrawalID                        string                                             `pathParam:"style=simple,explode=false,name=wireWithdrawal_id"`
	ForceApproveWireWithdrawalRequestCreate components.ForceApproveWireWithdrawalRequestCreate `request:"mediaType=application/json"`
}

func (o *WireWithdrawalsForceApproveWireWithdrawalRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *WireWithdrawalsForceApproveWireWithdrawalRequest) GetWireWithdrawalID() string {
	if o == nil {
		return ""
	}
	return o.WireWithdrawalID
}

func (o *WireWithdrawalsForceApproveWireWithdrawalRequest) GetForceApproveWireWithdrawalRequestCreate() components.ForceApproveWireWithdrawalRequestCreate {
	if o == nil {
		return components.ForceApproveWireWithdrawalRequestCreate{}
	}
	return o.ForceApproveWireWithdrawalRequestCreate
}

type WireWithdrawalsForceApproveWireWithdrawalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	WireWithdrawal *components.WireWithdrawal
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *WireWithdrawalsForceApproveWireWithdrawalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WireWithdrawalsForceApproveWireWithdrawalResponse) GetWireWithdrawal() *components.WireWithdrawal {
	if o == nil {
		return nil
	}
	return o.WireWithdrawal
}

func (o *WireWithdrawalsForceApproveWireWithdrawalResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
