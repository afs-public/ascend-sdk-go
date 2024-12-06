// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RelatedPep - Detail around the related politically exposed Person
type RelatedPep struct {
	// Information about the immediate family members of the related politically exposed person
	ImmediateFamilyMembers []ImmediateFamilyMember `json:"immediate_family_members,omitempty"`
	// The organization a politically exposed person is associated with causing them to be considered politically exposed
	Organization *string `json:"organization,omitempty"`
	// The name of the related politically exposed person
	PepName *string `json:"pep_name,omitempty"`
	// The title of the related politically exposed person
	Title *string `json:"title,omitempty"`
}

func (o *RelatedPep) GetImmediateFamilyMembers() []ImmediateFamilyMember {
	if o == nil {
		return nil
	}
	return o.ImmediateFamilyMembers
}

func (o *RelatedPep) GetOrganization() *string {
	if o == nil {
		return nil
	}
	return o.Organization
}

func (o *RelatedPep) GetPepName() *string {
	if o == nil {
		return nil
	}
	return o.PepName
}

func (o *RelatedPep) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}