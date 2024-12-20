// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AchWithdrawalsCreateAchWithdrawalRequest struct {
	// The account id.
	AccountID           string                         `pathParam:"style=simple,explode=false,name=account_id"`
	AchWithdrawalCreate components.AchWithdrawalCreate `request:"mediaType=application/json"`
}

func (o *AchWithdrawalsCreateAchWithdrawalRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchWithdrawalsCreateAchWithdrawalRequest) GetAchWithdrawalCreate() components.AchWithdrawalCreate {
	if o == nil {
		return components.AchWithdrawalCreate{}
	}
	return o.AchWithdrawalCreate
}

type AchWithdrawalsCreateAchWithdrawalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchWithdrawal *components.AchWithdrawal
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchWithdrawalsCreateAchWithdrawalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchWithdrawalsCreateAchWithdrawalResponse) GetAchWithdrawal() *components.AchWithdrawal {
	if o == nil {
		return nil
	}
	return o.AchWithdrawal
}

func (o *AchWithdrawalsCreateAchWithdrawalResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
