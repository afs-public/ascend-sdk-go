// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsListAvailableRestrictionsRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
}

func (o *AccountsListAvailableRestrictionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type AccountsListAvailableRestrictionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListAvailableRestrictionsResponse *components.ListAvailableRestrictionsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsListAvailableRestrictionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsListAvailableRestrictionsResponse) GetListAvailableRestrictionsResponse() *components.ListAvailableRestrictionsResponse {
	if o == nil {
		return nil
	}
	return o.ListAvailableRestrictionsResponse
}

func (o *AccountsListAvailableRestrictionsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
