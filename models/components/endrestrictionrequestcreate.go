// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EndRestrictionRequestCreate - The request for ending a Restriction on an Account.
type EndRestrictionRequestCreate struct {
	// The reason to end the restriction.
	Reason string `json:"reason"`
}

func (o *EndRestrictionRequestCreate) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}
