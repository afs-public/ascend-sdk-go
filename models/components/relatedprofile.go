// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RelatedProfile - Profiles related to dow jones profile
type RelatedProfile struct {
	// Description details
	DescriptionDetails []DescriptionDetail `json:"description_details,omitempty"`
	// Dow Jones person id
	DowJonesPersonID *int `json:"dow_jones_person_id,omitempty"`
	// Given name relating to profile
	GivenName *string `json:"given_name,omitempty"`
	// Middle names relating to profile
	MiddleNames *string `json:"middle_names,omitempty"`
	// Surname relating to profile
	NameSuffix *string `json:"name_suffix,omitempty"`
	// Relationship type
	RelationshipType *string `json:"relationship_type,omitempty"`
	// Sanctions list details
	SanctionsListDetails []SanctionsListDetail `json:"sanctions_list_details,omitempty"`
	// Surname relating to profile
	Surname *string `json:"surname,omitempty"`
	// Dow Jones persons title
	TitleHonorific *string `json:"title_honorific,omitempty"`
}

func (o *RelatedProfile) GetDescriptionDetails() []DescriptionDetail {
	if o == nil {
		return nil
	}
	return o.DescriptionDetails
}

func (o *RelatedProfile) GetDowJonesPersonID() *int {
	if o == nil {
		return nil
	}
	return o.DowJonesPersonID
}

func (o *RelatedProfile) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

func (o *RelatedProfile) GetMiddleNames() *string {
	if o == nil {
		return nil
	}
	return o.MiddleNames
}

func (o *RelatedProfile) GetNameSuffix() *string {
	if o == nil {
		return nil
	}
	return o.NameSuffix
}

func (o *RelatedProfile) GetRelationshipType() *string {
	if o == nil {
		return nil
	}
	return o.RelationshipType
}

func (o *RelatedProfile) GetSanctionsListDetails() []SanctionsListDetail {
	if o == nil {
		return nil
	}
	return o.SanctionsListDetails
}

func (o *RelatedProfile) GetSurname() *string {
	if o == nil {
		return nil
	}
	return o.Surname
}

func (o *RelatedProfile) GetTitleHonorific() *string {
	if o == nil {
		return nil
	}
	return o.TitleHonorific
}