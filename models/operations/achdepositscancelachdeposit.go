// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AchDepositsCancelAchDepositRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The achDeposit id.
	AchDepositID                  string                                   `pathParam:"style=simple,explode=false,name=achDeposit_id"`
	CancelAchDepositRequestCreate components.CancelAchDepositRequestCreate `request:"mediaType=application/json"`
}

func (o *AchDepositsCancelAchDepositRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchDepositsCancelAchDepositRequest) GetAchDepositID() string {
	if o == nil {
		return ""
	}
	return o.AchDepositID
}

func (o *AchDepositsCancelAchDepositRequest) GetCancelAchDepositRequestCreate() components.CancelAchDepositRequestCreate {
	if o == nil {
		return components.CancelAchDepositRequestCreate{}
	}
	return o.CancelAchDepositRequestCreate
}

type AchDepositsCancelAchDepositResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchDeposit *components.AchDeposit
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchDepositsCancelAchDepositResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchDepositsCancelAchDepositResponse) GetAchDeposit() *components.AchDeposit {
	if o == nil {
		return nil
	}
	return o.AchDeposit
}

func (o *AchDepositsCancelAchDepositResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
