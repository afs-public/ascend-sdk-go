// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type InvestorCommunicationServiceListDocumentsRequest struct {
	// The maximum number of items to return; The service may return fewer than this value
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// Token used to get a specific page of results
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// CEL filter to be applied against the documents; Providing a correspondent to search for is required; Only one correspondent can be searched at a time
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *InvestorCommunicationServiceListDocumentsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *InvestorCommunicationServiceListDocumentsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *InvestorCommunicationServiceListDocumentsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type InvestorCommunicationServiceListDocumentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListDocumentsResponse *components.ListDocumentsResponse
	// INVALID_ARGUMENT: The request was not well formed.
	Status *components.Status
}

func (o *InvestorCommunicationServiceListDocumentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InvestorCommunicationServiceListDocumentsResponse) GetListDocumentsResponse() *components.ListDocumentsResponse {
	if o == nil {
		return nil
	}
	return o.ListDocumentsResponse
}

func (o *InvestorCommunicationServiceListDocumentsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}