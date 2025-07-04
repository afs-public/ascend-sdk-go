// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ReuseBankRelationshipRequestCreate - Request to reuse a bank relationship.
type ReuseBankRelationshipRequestCreate struct {
	// The account to create the reused bank relationship under. The account must be related to the parent account of the `source_bank_relationship`. Format: `accounts/{account}`
	Parent string `json:"parent"`
	// The source bank relationship to reuse. The bank relationship must be in an approved state. Format: `accounts/{account}/bankRelationships/{bank_relationship}`
	SourceBankRelationship string `json:"source_bank_relationship"`
}

func (o *ReuseBankRelationshipRequestCreate) GetParent() string {
	if o == nil {
		return ""
	}
	return o.Parent
}

func (o *ReuseBankRelationshipRequestCreate) GetSourceBankRelationship() string {
	if o == nil {
		return ""
	}
	return o.SourceBankRelationship
}
