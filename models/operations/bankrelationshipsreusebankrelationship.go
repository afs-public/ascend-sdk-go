// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type BankRelationshipsReuseBankRelationshipRequest struct {
	// The account id.
	AccountID                          string                                        `pathParam:"style=simple,explode=false,name=account_id"`
	ReuseBankRelationshipRequestCreate components.ReuseBankRelationshipRequestCreate `request:"mediaType=application/json"`
}

func (o *BankRelationshipsReuseBankRelationshipRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BankRelationshipsReuseBankRelationshipRequest) GetReuseBankRelationshipRequestCreate() components.ReuseBankRelationshipRequestCreate {
	if o == nil {
		return components.ReuseBankRelationshipRequestCreate{}
	}
	return o.ReuseBankRelationshipRequestCreate
}

type BankRelationshipsReuseBankRelationshipResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	BankRelationship *components.BankRelationship
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *BankRelationshipsReuseBankRelationshipResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BankRelationshipsReuseBankRelationshipResponse) GetBankRelationship() *components.BankRelationship {
	if o == nil {
		return nil
	}
	return o.BankRelationship
}

func (o *BankRelationshipsReuseBankRelationshipResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
