// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListEventMessagesResponse - A response to a list events method
type ListEventMessagesResponse struct {
	// The returned collection of events
	EventMessages []EventMessage `json:"event_messages,omitempty"`
	// Page token used for pagination; Supplying a page token returns the next page of results
	NextPageToken *string `json:"next_page_token,omitempty"`
}

func (o *ListEventMessagesResponse) GetEventMessages() []EventMessage {
	if o == nil {
		return nil
	}
	return o.EventMessages
}

func (o *ListEventMessagesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
