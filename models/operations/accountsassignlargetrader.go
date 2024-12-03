// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type AccountsAssignLargeTraderRequest struct {
	// The legalNaturalPerson id.
	LegalNaturalPersonID           string                                    `pathParam:"style=simple,explode=false,name=legalNaturalPerson_id"`
	AssignLargeTraderRequestCreate components.AssignLargeTraderRequestCreate `request:"mediaType=application/json"`
}

func (o *AccountsAssignLargeTraderRequest) GetLegalNaturalPersonID() string {
	if o == nil {
		return ""
	}
	return o.LegalNaturalPersonID
}

func (o *AccountsAssignLargeTraderRequest) GetAssignLargeTraderRequestCreate() components.AssignLargeTraderRequestCreate {
	if o == nil {
		return components.AssignLargeTraderRequestCreate{}
	}
	return o.AssignLargeTraderRequestCreate
}

type AccountsAssignLargeTraderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LargeTrader *components.LargeTrader
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsAssignLargeTraderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsAssignLargeTraderResponse) GetLargeTrader() *components.LargeTrader {
	if o == nil {
		return nil
	}
	return o.LargeTrader
}

func (o *AccountsAssignLargeTraderResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
