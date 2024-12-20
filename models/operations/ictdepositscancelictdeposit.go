// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type IctDepositsCancelIctDepositRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The ictDeposit id.
	IctDepositID                  string                                   `pathParam:"style=simple,explode=false,name=ictDeposit_id"`
	CancelIctDepositRequestCreate components.CancelIctDepositRequestCreate `request:"mediaType=application/json"`
}

func (o *IctDepositsCancelIctDepositRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *IctDepositsCancelIctDepositRequest) GetIctDepositID() string {
	if o == nil {
		return ""
	}
	return o.IctDepositID
}

func (o *IctDepositsCancelIctDepositRequest) GetCancelIctDepositRequestCreate() components.CancelIctDepositRequestCreate {
	if o == nil {
		return components.CancelIctDepositRequestCreate{}
	}
	return o.CancelIctDepositRequestCreate
}

type IctDepositsCancelIctDepositResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	IctDeposit *components.IctDeposit
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *IctDepositsCancelIctDepositResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IctDepositsCancelIctDepositResponse) GetIctDeposit() *components.IctDeposit {
	if o == nil {
		return nil
	}
	return o.IctDeposit
}

func (o *IctDepositsCancelIctDepositResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
