// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type BankRelationshipsCancelBankRelationshipRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The bankRelationship id.
	BankRelationshipID                  string                                         `pathParam:"style=simple,explode=false,name=bankRelationship_id"`
	CancelBankRelationshipRequestCreate components.CancelBankRelationshipRequestCreate `request:"mediaType=application/json"`
}

func (o *BankRelationshipsCancelBankRelationshipRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BankRelationshipsCancelBankRelationshipRequest) GetBankRelationshipID() string {
	if o == nil {
		return ""
	}
	return o.BankRelationshipID
}

func (o *BankRelationshipsCancelBankRelationshipRequest) GetCancelBankRelationshipRequestCreate() components.CancelBankRelationshipRequestCreate {
	if o == nil {
		return components.CancelBankRelationshipRequestCreate{}
	}
	return o.CancelBankRelationshipRequestCreate
}

type BankRelationshipsCancelBankRelationshipResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	BankRelationship *components.BankRelationship
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *BankRelationshipsCancelBankRelationshipResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BankRelationshipsCancelBankRelationshipResponse) GetBankRelationship() *components.BankRelationship {
	if o == nil {
		return nil
	}
	return o.BankRelationship
}

func (o *BankRelationshipsCancelBankRelationshipResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
