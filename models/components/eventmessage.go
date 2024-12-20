// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/utils"
)

// EventMessage - Represents an envelope and the data of an event
type EventMessage struct {
	// The account ID related to the event (if applicable)
	AccountID *string `json:"account_id,omitempty"`
	// The client ID related to the event
	ClientID *string `json:"client_id,omitempty"`
	// The correspondent ID related to the event (if applicable)
	CorrespondentID *string `json:"correspondent_id,omitempty"`
	// A data payload containing the fields specific to the type of event being sent
	Data map[string]any `json:"data,omitempty"`
	// Specifies the action that triggered the event as well as what resource changed
	EventType *string `json:"event_type,omitempty"`
	// A unique identifier for the event
	MessageID *string `json:"message_id,omitempty"`
	// The resource name of the event; Format: messages/{message}
	Name *string `json:"name,omitempty"`
	// A value, if present, is used to group related events together. Events with the same partition key are guaranteed to be sent to the consumer in the same order they were published.
	PartitionKey *string `json:"partition_key,omitempty"`
	// The date and time of the event publication (not necessarily the time the event occurred)
	PublishTime *time.Time `json:"publish_time,omitempty"`
}

func (e EventMessage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventMessage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EventMessage) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *EventMessage) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *EventMessage) GetCorrespondentID() *string {
	if o == nil {
		return nil
	}
	return o.CorrespondentID
}

func (o *EventMessage) GetData() map[string]any {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *EventMessage) GetEventType() *string {
	if o == nil {
		return nil
	}
	return o.EventType
}

func (o *EventMessage) GetMessageID() *string {
	if o == nil {
		return nil
	}
	return o.MessageID
}

func (o *EventMessage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EventMessage) GetPartitionKey() *string {
	if o == nil {
		return nil
	}
	return o.PartitionKey
}

func (o *EventMessage) GetPublishTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.PublishTime
}
