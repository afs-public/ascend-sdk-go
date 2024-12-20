// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EntityIdentificationType - Tax id type for entities (e.g. ein, lei, etc.))
type EntityIdentificationType string

const (
	EntityIdentificationTypeIDEntityTypeUnspecified EntityIdentificationType = "ID_ENTITY_TYPE_UNSPECIFIED"
	EntityIdentificationTypeEin                     EntityIdentificationType = "EIN"
	EntityIdentificationTypeLei                     EntityIdentificationType = "LEI"
	EntityIdentificationTypeDuns                    EntityIdentificationType = "DUNS"
)

func (e EntityIdentificationType) ToPointer() *EntityIdentificationType {
	return &e
}

// EntityIdentification - stores various Entity identification types
type EntityIdentification struct {
	// Administrative area that issued the identification For example, this can be a state, a province, an oblast, or a prefecture.
	AdministrativeArea *string `json:"administrative_area,omitempty"`
	// One or more UUIDs from the documents api of the image(s) of the document that relates to the identification for the person investigation.
	DocumentReferenceIds []string `json:"document_reference_ids,omitempty"`
	// Country that issued identification Two character region code, complies with https://cldr.unicode.org/index
	RegionCode *string `json:"region_code,omitempty"`
	// Tax id type for entities (e.g. ein, lei, etc.))
	Type *EntityIdentificationType `json:"type,omitempty"`
	// Tax id value
	Value *string `json:"value,omitempty"`
}

func (o *EntityIdentification) GetAdministrativeArea() *string {
	if o == nil {
		return nil
	}
	return o.AdministrativeArea
}

func (o *EntityIdentification) GetDocumentReferenceIds() []string {
	if o == nil {
		return nil
	}
	return o.DocumentReferenceIds
}

func (o *EntityIdentification) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}

func (o *EntityIdentification) GetType() *EntityIdentificationType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *EntityIdentification) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
