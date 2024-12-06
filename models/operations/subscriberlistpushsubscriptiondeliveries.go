// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type SubscriberListPushSubscriptionDeliveriesRequest struct {
	// The subscription id.
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; If left empty, all deliveries the user has permission to view are returned; Filter options include:
	//  `name`
	//  `delivery_id`
	//  `event`
	//  `event_publish_time`
	//  `result`
	//  `last_response`
	//  `last_send_time`
	//  `duration`
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
	// The number of entries to return in a single page; Default = 100; Maximum = 1000
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// Page token used for pagination; Supplying a page token returns the next page of results
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *SubscriberListPushSubscriptionDeliveriesRequest) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

func (o *SubscriberListPushSubscriptionDeliveriesRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *SubscriberListPushSubscriptionDeliveriesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SubscriberListPushSubscriptionDeliveriesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type SubscriberListPushSubscriptionDeliveriesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListPushSubscriptionDeliveriesResponse *components.ListPushSubscriptionDeliveriesResponse
	// INVALID_ARGUMENT: The request was not well formed.
	Status *components.Status
}

func (o *SubscriberListPushSubscriptionDeliveriesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscriberListPushSubscriptionDeliveriesResponse) GetListPushSubscriptionDeliveriesResponse() *components.ListPushSubscriptionDeliveriesResponse {
	if o == nil {
		return nil
	}
	return o.ListPushSubscriptionDeliveriesResponse
}

func (o *SubscriberListPushSubscriptionDeliveriesResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
