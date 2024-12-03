// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// VerificationMethod - The verification method of the bank relationship.
type VerificationMethod string

const (
	VerificationMethodVerificationMethodUnspecified VerificationMethod = "VERIFICATION_METHOD_UNSPECIFIED"
	VerificationMethodMicroDeposit                  VerificationMethod = "MICRO_DEPOSIT"
	VerificationMethodYodlee                        VerificationMethod = "YODLEE"
	VerificationMethodQuovo                         VerificationMethod = "QUOVO"
	VerificationMethodGiact                         VerificationMethod = "GIACT"
	VerificationMethodSynapse                       VerificationMethod = "SYNAPSE"
	VerificationMethodSophtron                      VerificationMethod = "SOPHTRON"
	VerificationMethodUseExisting                   VerificationMethod = "USE_EXISTING"
	VerificationMethodInternalBank                  VerificationMethod = "INTERNAL_BANK"
	VerificationMethodMx                            VerificationMethod = "MX"
	VerificationMethodFiserv                        VerificationMethod = "FISERV"
	VerificationMethodPlaidToken                    VerificationMethod = "PLAID_TOKEN"
)

func (e VerificationMethod) ToPointer() *VerificationMethod {
	return &e
}
func (e *VerificationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VERIFICATION_METHOD_UNSPECIFIED":
		fallthrough
	case "MICRO_DEPOSIT":
		fallthrough
	case "YODLEE":
		fallthrough
	case "QUOVO":
		fallthrough
	case "GIACT":
		fallthrough
	case "SYNAPSE":
		fallthrough
	case "SOPHTRON":
		fallthrough
	case "USE_EXISTING":
		fallthrough
	case "INTERNAL_BANK":
		fallthrough
	case "MX":
		fallthrough
	case "FISERV":
		fallthrough
	case "PLAID_TOKEN":
		*e = VerificationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VerificationMethod: %v", v)
	}
}

// BankRelationshipCreate - A relationship between a bank account and an Apex account.
type BankRelationshipCreate struct {
	// A representation of a bank account.
	BankAccount *BankAccountCreate `json:"bank_account,omitempty"`
	// The nickname of the bank relationship.
	Nickname string `json:"nickname"`
	// A processor token from Plaid (vendor). Required if using `PLAID_TOKEN` verification method.
	PlaidProcessorToken *string `json:"plaid_processor_token,omitempty"`
	// The verification method of the bank relationship.
	VerificationMethod VerificationMethod `json:"verification_method"`
}

func (o *BankRelationshipCreate) GetBankAccount() *BankAccountCreate {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *BankRelationshipCreate) GetNickname() string {
	if o == nil {
		return ""
	}
	return o.Nickname
}

func (o *BankRelationshipCreate) GetPlaidProcessorToken() *string {
	if o == nil {
		return nil
	}
	return o.PlaidProcessorToken
}

func (o *BankRelationshipCreate) GetVerificationMethod() VerificationMethod {
	if o == nil {
		return VerificationMethod("")
	}
	return o.VerificationMethod
}
