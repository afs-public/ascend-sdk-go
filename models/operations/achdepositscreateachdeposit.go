// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type AchDepositsCreateAchDepositRequest struct {
	// The account id.
	AccountID        string                      `pathParam:"style=simple,explode=false,name=account_id"`
	AchDepositCreate components.AchDepositCreate `request:"mediaType=application/json"`
}

func (o *AchDepositsCreateAchDepositRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchDepositsCreateAchDepositRequest) GetAchDepositCreate() components.AchDepositCreate {
	if o == nil {
		return components.AchDepositCreate{}
	}
	return o.AchDepositCreate
}

type AchDepositsCreateAchDepositResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchDeposit *components.AchDeposit
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchDepositsCreateAchDepositResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchDepositsCreateAchDepositResponse) GetAchDeposit() *components.AchDeposit {
	if o == nil {
		return nil
	}
	return o.AchDeposit
}

func (o *AchDepositsCreateAchDepositResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
