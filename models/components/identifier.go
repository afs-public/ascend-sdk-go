// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// IdentifierType1 - The type of identifier
type IdentifierType1 string

const (
	IdentifierType1IdentifierTypeUnspecified IdentifierType1 = "IDENTIFIER_TYPE_UNSPECIFIED"
	IdentifierType1OriginatingAccountID      IdentifierType1 = "ORIGINATING_ACCOUNT_ID"
	IdentifierType1OriginatingFdid           IdentifierType1 = "ORIGINATING_FDID"
	IdentifierType1OriginatingCatReporterCrd IdentifierType1 = "ORIGINATING_CAT_REPORTER_CRD"
	IdentifierType1ClientAccountID           IdentifierType1 = "CLIENT_ACCOUNT_ID"
)

func (e IdentifierType1) ToPointer() *IdentifierType1 {
	return &e
}

// Identifier - An identifier associated with an account
type Identifier struct {
	// The type of identifier
	Type *IdentifierType1 `json:"type,omitempty"`
	// The value of the identifier
	Value *string `json:"value,omitempty"`
}

func (o *Identifier) GetType() *IdentifierType1 {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Identifier) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
