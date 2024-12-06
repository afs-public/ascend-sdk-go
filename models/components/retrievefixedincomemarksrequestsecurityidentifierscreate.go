// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType - The identifier type of the asset being sought
type RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType string

const (
	RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeAssetID RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType = "ASSET_ID"
	RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeCusip   RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType = "CUSIP"
	RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierTypeIsin    RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType = "ISIN"
)

func (e RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType) ToPointer() *RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType {
	return &e
}
func (e *RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASSET_ID":
		fallthrough
	case "CUSIP":
		fallthrough
	case "ISIN":
		*e = RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType: %v", v)
	}
}

// RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate - The identifier for the asset.
type RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate struct {
	// Identifier of the asset (of the type specified in `identifier_type`).
	Identifier string `json:"identifier"`
	// The identifier type of the asset being sought
	IdentifierType RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType `json:"identifier_type"`
}

func (o *RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreate) GetIdentifierType() RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType {
	if o == nil {
		return RetrieveFixedIncomeMarksRequestSecurityIdentifiersCreateIdentifierType("")
	}
	return o.IdentifierType
}