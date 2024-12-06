// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListAssetsResponse is the response message for ListAssets method.
type ListAssetsResponse struct {
	// The assets returned in the response.
	Assets []Asset `json:"assets,omitempty"`
	// The next_page_token value to include in a subsequent ListAssets request to retrieve the next page of results.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

func (o *ListAssetsResponse) GetAssets() []Asset {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *ListAssetsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}