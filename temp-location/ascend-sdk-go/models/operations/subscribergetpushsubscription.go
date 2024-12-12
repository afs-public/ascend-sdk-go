// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type SubscriberGetPushSubscriptionRequest struct {
	// The subscription id.
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
}

func (o *SubscriberGetPushSubscriptionRequest) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

type SubscriberGetPushSubscriptionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	PushSubscription *components.PushSubscription
	// INVALID_ARGUMENT: The request was not well formed.
	Status *components.Status
}

func (o *SubscriberGetPushSubscriptionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SubscriberGetPushSubscriptionResponse) GetPushSubscription() *components.PushSubscription {
	if o == nil {
		return nil
	}
	return o.PushSubscription
}

func (o *SubscriberGetPushSubscriptionResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
