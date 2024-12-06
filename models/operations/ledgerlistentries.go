// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type LedgerListEntriesRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The maximum number of entries to return. The service may return fewer than this value Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListEntries` call. Provide this to retrieve the subsequent page When paginating, all other parameters provided to `ListEntries` must match the call that provided the page token in order to maintain a stable result set
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *LedgerListEntriesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *LedgerListEntriesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *LedgerListEntriesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *LedgerListEntriesRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type LedgerListEntriesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListEntriesResponse *components.ListEntriesResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *LedgerListEntriesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *LedgerListEntriesResponse) GetListEntriesResponse() *components.ListEntriesResponse {
	if o == nil {
		return nil
	}
	return o.ListEntriesResponse
}

func (o *LedgerListEntriesResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
