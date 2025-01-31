// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// HTTPCallback - The information about an HTTP target callback
type HTTPCallback struct {
	// The maximum amount of time, in seconds, the service will wait for an acknowledgement of a delivery. If a value of 0 or no value is specified, the timeout will default to 10 seconds.
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
	// The URL address of the client HTTP server that will receive the events via POST; URLs must be in the form of https://{domain}[/{path}]
	URL *string `json:"url,omitempty"`
}

func (o *HTTPCallback) GetTimeoutSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutSeconds
}

func (o *HTTPCallback) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// State - The current status of the subscription
type State string

const (
	StatePushSubscriptionStateUnspecified State = "PUSH_SUBSCRIPTION_STATE_UNSPECIFIED"
	StateCreating                         State = "CREATING"
	StateActive                           State = "ACTIVE"
	StateUpdating                         State = "UPDATING"
	StateDeleting                         State = "DELETING"
)

func (e State) ToPointer() *State {
	return &e
}

// PushSubscription - Configuration information about a push subscription
type PushSubscription struct {
	// The client that owns the subscription. A client subscription will receive events for it and all of its correspondents. This can only be set at creation time and is mutually exclusive with correspondent_id.
	ClientID *string `json:"client_id,omitempty"`
	// The correspondent that owns the subscription. A correspondent subscription will receive events only for itself. This can only be set at creation time and is mutually exclusive with client_id.
	CorrespondentID *string `json:"correspondent_id,omitempty"`
	// The user-defined name for the subscription
	DisplayName *string `json:"display_name,omitempty"`
	// Filter for event types; ["\*"] matches all values; Suffix wildcards using "\*" (e.g. ["account.\*"]) are supported
	EventTypes []string `json:"event_types,omitempty"`
	// The information about an HTTP target callback
	HTTPCallback *HTTPCallback `json:"http_callback,omitempty"`
	// The resource name of the subscription; Format: subscriptions/{subscription}
	Name *string `json:"name,omitempty"`
	// The current status of the subscription
	State *State `json:"state,omitempty"`
	// The unique identifier for the subscription
	SubscriptionID *string `json:"subscription_id,omitempty"`
}

func (o *PushSubscription) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *PushSubscription) GetCorrespondentID() *string {
	if o == nil {
		return nil
	}
	return o.CorrespondentID
}

func (o *PushSubscription) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *PushSubscription) GetEventTypes() []string {
	if o == nil {
		return nil
	}
	return o.EventTypes
}

func (o *PushSubscription) GetHTTPCallback() *HTTPCallback {
	if o == nil {
		return nil
	}
	return o.HTTPCallback
}

func (o *PushSubscription) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PushSubscription) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *PushSubscription) GetSubscriptionID() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionID
}
