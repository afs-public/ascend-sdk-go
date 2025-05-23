// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AchDepositsGetAchDepositRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The achDeposit id.
	AchDepositID string `pathParam:"style=simple,explode=false,name=achDeposit_id"`
}

func (o *AchDepositsGetAchDepositRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AchDepositsGetAchDepositRequest) GetAchDepositID() string {
	if o == nil {
		return ""
	}
	return o.AchDepositID
}

type AchDepositsGetAchDepositResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AchDeposit *components.AchDeposit
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AchDepositsGetAchDepositResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AchDepositsGetAchDepositResponse) GetAchDeposit() *components.AchDeposit {
	if o == nil {
		return nil
	}
	return o.AchDeposit
}

func (o *AchDepositsGetAchDepositResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
