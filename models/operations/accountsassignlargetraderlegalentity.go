// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsAssignLargeTraderLegalEntityRequest struct {
	// The legalEntity id.
	LegalEntityID                  string                                    `pathParam:"style=simple,explode=false,name=legalEntity_id"`
	AssignLargeTraderRequestCreate components.AssignLargeTraderRequestCreate `request:"mediaType=application/json"`
}

func (o *AccountsAssignLargeTraderLegalEntityRequest) GetLegalEntityID() string {
	if o == nil {
		return ""
	}
	return o.LegalEntityID
}

func (o *AccountsAssignLargeTraderLegalEntityRequest) GetAssignLargeTraderRequestCreate() components.AssignLargeTraderRequestCreate {
	if o == nil {
		return components.AssignLargeTraderRequestCreate{}
	}
	return o.AssignLargeTraderRequestCreate
}

type AccountsAssignLargeTraderLegalEntityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LargeTrader *components.LargeTrader
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsAssignLargeTraderLegalEntityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsAssignLargeTraderLegalEntityResponse) GetLargeTrader() *components.LargeTrader {
	if o == nil {
		return nil
	}
	return o.LargeTrader
}

func (o *AccountsAssignLargeTraderLegalEntityResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}