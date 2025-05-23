// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsCreateRestrictionRequest struct {
	// The account id.
	AccountID         string                       `pathParam:"style=simple,explode=false,name=account_id"`
	RestrictionCreate components.RestrictionCreate `request:"mediaType=application/json"`
}

func (o *AccountsCreateRestrictionRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsCreateRestrictionRequest) GetRestrictionCreate() components.RestrictionCreate {
	if o == nil {
		return components.RestrictionCreate{}
	}
	return o.RestrictionCreate
}

type AccountsCreateRestrictionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Restriction *components.Restriction
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsCreateRestrictionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsCreateRestrictionResponse) GetRestriction() *components.Restriction {
	if o == nil {
		return nil
	}
	return o.Restriction
}

func (o *AccountsCreateRestrictionResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
