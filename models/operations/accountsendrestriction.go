// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsEndRestrictionRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The restriction id.
	RestrictionID               string                                 `pathParam:"style=simple,explode=false,name=restriction_id"`
	EndRestrictionRequestCreate components.EndRestrictionRequestCreate `request:"mediaType=application/json"`
}

func (o *AccountsEndRestrictionRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsEndRestrictionRequest) GetRestrictionID() string {
	if o == nil {
		return ""
	}
	return o.RestrictionID
}

func (o *AccountsEndRestrictionRequest) GetEndRestrictionRequestCreate() components.EndRestrictionRequestCreate {
	if o == nil {
		return components.EndRestrictionRequestCreate{}
	}
	return o.EndRestrictionRequestCreate
}

type AccountsEndRestrictionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsEndRestrictionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsEndRestrictionResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
