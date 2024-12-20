// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type IctDepositsForceApproveIctDepositRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The ictDeposit id.
	IctDepositID                        string                                         `pathParam:"style=simple,explode=false,name=ictDeposit_id"`
	ForceApproveIctDepositRequestCreate components.ForceApproveIctDepositRequestCreate `request:"mediaType=application/json"`
}

func (o *IctDepositsForceApproveIctDepositRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *IctDepositsForceApproveIctDepositRequest) GetIctDepositID() string {
	if o == nil {
		return ""
	}
	return o.IctDepositID
}

func (o *IctDepositsForceApproveIctDepositRequest) GetForceApproveIctDepositRequestCreate() components.ForceApproveIctDepositRequestCreate {
	if o == nil {
		return components.ForceApproveIctDepositRequestCreate{}
	}
	return o.ForceApproveIctDepositRequestCreate
}

type IctDepositsForceApproveIctDepositResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	IctDeposit *components.IctDeposit
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *IctDepositsForceApproveIctDepositResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IctDepositsForceApproveIctDepositResponse) GetIctDeposit() *components.IctDeposit {
	if o == nil {
		return nil
	}
	return o.IctDeposit
}

func (o *IctDepositsForceApproveIctDepositResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
