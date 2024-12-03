// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EntityDueDiligenceUpdate - Due Diligence for Legal Entities required when a Legal Entity is the Primary Owner on an Account.
type EntityDueDiligenceUpdate struct {
	// Indicates whether the entity issues bearer shares
	EntityIssuesBearerShares *bool `json:"entity_issues_bearer_shares,omitempty"`
	// Negative News detail.
	NegativeNews *NegativeNewsUpdate `json:"negative_news,omitempty"`
}

func (o *EntityDueDiligenceUpdate) GetEntityIssuesBearerShares() *bool {
	if o == nil {
		return nil
	}
	return o.EntityIssuesBearerShares
}

func (o *EntityDueDiligenceUpdate) GetNegativeNews() *NegativeNewsUpdate {
	if o == nil {
		return nil
	}
	return o.NegativeNews
}
