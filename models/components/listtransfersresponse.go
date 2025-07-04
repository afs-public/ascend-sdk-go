// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListTransfersResponse - Response to list existing transfers.
type ListTransfersResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// transfers resulting from filter query
	Transfers []AcatsTransfer `json:"transfers,omitempty"`
}

func (o *ListTransfersResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *ListTransfersResponse) GetTransfers() []AcatsTransfer {
	if o == nil {
		return nil
	}
	return o.Transfers
}
