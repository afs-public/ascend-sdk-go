// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type IctWithdrawalsCancelIctWithdrawalRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The ictWithdrawal id.
	IctWithdrawalID                  string                                      `pathParam:"style=simple,explode=false,name=ictWithdrawal_id"`
	CancelIctWithdrawalRequestCreate components.CancelIctWithdrawalRequestCreate `request:"mediaType=application/json"`
}

func (o *IctWithdrawalsCancelIctWithdrawalRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *IctWithdrawalsCancelIctWithdrawalRequest) GetIctWithdrawalID() string {
	if o == nil {
		return ""
	}
	return o.IctWithdrawalID
}

func (o *IctWithdrawalsCancelIctWithdrawalRequest) GetCancelIctWithdrawalRequestCreate() components.CancelIctWithdrawalRequestCreate {
	if o == nil {
		return components.CancelIctWithdrawalRequestCreate{}
	}
	return o.CancelIctWithdrawalRequestCreate
}

type IctWithdrawalsCancelIctWithdrawalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	IctWithdrawal *components.IctWithdrawal
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *IctWithdrawalsCancelIctWithdrawalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IctWithdrawalsCancelIctWithdrawalResponse) GetIctWithdrawal() *components.IctWithdrawal {
	if o == nil {
		return nil
	}
	return o.IctWithdrawal
}

func (o *IctWithdrawalsCancelIctWithdrawalResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}