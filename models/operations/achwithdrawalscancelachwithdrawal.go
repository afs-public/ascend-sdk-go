// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AchWithdrawalsCancelAchWithdrawalRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The achWithdrawal id.
	AchWithdrawalID                  string                                      `pathParam:"style=simple,explode=false,name=achWithdrawal_id"`
	CancelAchWithdrawalRequestCreate components.CancelAchWithdrawalRequestCreate `request:"mediaType=application/json"`
}

func (o *AchWithdrawalsCancelAchWithdrawalRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchWithdrawalsCancelAchWithdrawalRequest) GetAchWithdrawalID() string {
	if o == nil {
		return ""
	}
	return o.AchWithdrawalID
}

func (o *AchWithdrawalsCancelAchWithdrawalRequest) GetCancelAchWithdrawalRequestCreate() components.CancelAchWithdrawalRequestCreate {
	if o == nil {
		return components.CancelAchWithdrawalRequestCreate{}
	}
	return o.CancelAchWithdrawalRequestCreate
}

type AchWithdrawalsCancelAchWithdrawalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchWithdrawal *components.AchWithdrawal
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchWithdrawalsCancelAchWithdrawalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchWithdrawalsCancelAchWithdrawalResponse) GetAchWithdrawal() *components.AchWithdrawal {
	if o == nil {
		return nil
	}
	return o.AchWithdrawal
}

func (o *AchWithdrawalsCancelAchWithdrawalResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}