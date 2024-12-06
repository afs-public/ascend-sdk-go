// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type ReaderListEventMessagesRequest struct {
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; If left empty, all events the user has permission to view are returned; Filter options include:
	//  `name`
	//  `message_id`
	//  `event_type`
	//  `publish_time`
	//  `partition_key`
	//  `client_id`
	//  `correspondent_id`
	//  `account_id`
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
	// The number of entries to return in a single page; Default = 100; Maximum = 1000
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// Page token used for pagination; Supplying a page token returns the next page of results
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *ReaderListEventMessagesRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ReaderListEventMessagesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ReaderListEventMessagesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type ReaderListEventMessagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListEventMessagesResponse *components.ListEventMessagesResponse
	// INVALID_ARGUMENT: The request was not well formed.
	Status *components.Status
}

func (o *ReaderListEventMessagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReaderListEventMessagesResponse) GetListEventMessagesResponse() *components.ListEventMessagesResponse {
	if o == nil {
		return nil
	}
	return o.ListEventMessagesResponse
}

func (o *ReaderListEventMessagesResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}