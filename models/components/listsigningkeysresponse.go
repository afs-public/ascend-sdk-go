// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ListSigningKeysResponse - A response to a list signing keys method
type ListSigningKeysResponse struct {
	// The returned collection of all currently valid signing keys
	Keys []map[string]any `json:"keys,omitempty"`
	// Page token used for pagination; Supplying a page token returns the next page of results
	NextPageToken *string `json:"next_page_token,omitempty"`
}

func (o *ListSigningKeysResponse) GetKeys() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Keys
}

func (o *ListSigningKeysResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}