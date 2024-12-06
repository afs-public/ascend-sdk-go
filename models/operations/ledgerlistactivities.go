// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type LedgerListActivitiesRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The maximum number of activities to return. The service may return fewer than this value Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListActivity` call. Provide this to retrieve the subsequent page When paginating, all other parameters provided to `ListActivity` must match the call that provided the page token in order to maintain a stable result set
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *LedgerListActivitiesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *LedgerListActivitiesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *LedgerListActivitiesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *LedgerListActivitiesRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type LedgerListActivitiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListActivitiesResponse *components.ListActivitiesResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *LedgerListActivitiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *LedgerListActivitiesResponse) GetListActivitiesResponse() *components.ListActivitiesResponse {
	if o == nil {
		return nil
	}
	return o.ListActivitiesResponse
}

func (o *LedgerListActivitiesResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
