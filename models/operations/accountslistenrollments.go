// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsListEnrollmentsRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The maximum number of enrollments to return.
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListEnrollments` call. Provide this to retrieve the subsequent page.
	//
	//  When paginating, all other parameters provided to `ListEnrollments` must match the call that provided the page token.
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *AccountsListEnrollmentsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsListEnrollmentsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *AccountsListEnrollmentsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type AccountsListEnrollmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListEnrollmentsResponse *components.ListEnrollmentsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsListEnrollmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsListEnrollmentsResponse) GetListEnrollmentsResponse() *components.ListEnrollmentsResponse {
	if o == nil {
		return nil
	}
	return o.ListEnrollmentsResponse
}

func (o *AccountsListEnrollmentsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}