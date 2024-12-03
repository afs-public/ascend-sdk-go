// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RecipientBankBankIDCreateType - The type of bank identifier specified
type RecipientBankBankIDCreateType string

const (
	RecipientBankBankIDCreateTypeTypeUnspecified RecipientBankBankIDCreateType = "TYPE_UNSPECIFIED"
	RecipientBankBankIDCreateTypeAba             RecipientBankBankIDCreateType = "ABA"
	RecipientBankBankIDCreateTypeBic             RecipientBankBankIDCreateType = "BIC"
)

func (e RecipientBankBankIDCreateType) ToPointer() *RecipientBankBankIDCreateType {
	return &e
}
func (e *RecipientBankBankIDCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TYPE_UNSPECIFIED":
		fallthrough
	case "ABA":
		fallthrough
	case "BIC":
		*e = RecipientBankBankIDCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RecipientBankBankIDCreateType: %v", v)
	}
}

// RecipientBankBankIDCreate - A bank identifier
type RecipientBankBankIDCreate struct {
	// The bank identifier
	ID string `json:"id"`
	// The type of bank identifier specified
	Type RecipientBankBankIDCreateType `json:"type"`
}

func (o *RecipientBankBankIDCreate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RecipientBankBankIDCreate) GetType() RecipientBankBankIDCreateType {
	if o == nil {
		return RecipientBankBankIDCreateType("")
	}
	return o.Type
}
