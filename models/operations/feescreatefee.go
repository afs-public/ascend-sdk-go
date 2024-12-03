// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type FeesCreateFeeRequest struct {
	// The account id.
	AccountID          string                        `pathParam:"style=simple,explode=false,name=account_id"`
	TransfersFeeCreate components.TransfersFeeCreate `request:"mediaType=application/json"`
}

func (o *FeesCreateFeeRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *FeesCreateFeeRequest) GetTransfersFeeCreate() components.TransfersFeeCreate {
	if o == nil {
		return components.TransfersFeeCreate{}
	}
	return o.TransfersFeeCreate
}

type FeesCreateFeeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	TransfersFee *components.TransfersFee
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *FeesCreateFeeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FeesCreateFeeResponse) GetTransfersFee() *components.TransfersFee {
	if o == nil {
		return nil
	}
	return o.TransfersFee
}

func (o *FeesCreateFeeResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
