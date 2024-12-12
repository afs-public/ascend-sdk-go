// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListInvestigationsResponse - ListInvestigationStatesResponse
type ListInvestigationsResponse struct {
	// List of investigations matching request search criteria
	Investigations []Investigation `json:"investigations,omitempty"`
	// The next pagination token in the Search response; an empty value means no more results
	NextPageToken *string `json:"next_page_token,omitempty"`
}

func (o *ListInvestigationsResponse) GetInvestigations() []Investigation {
	if o == nil {
		return nil
	}
	return o.Investigations
}

func (o *ListInvestigationsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
