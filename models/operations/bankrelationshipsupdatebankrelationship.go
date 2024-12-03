// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type BankRelationshipsUpdateBankRelationshipRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The bankRelationship id.
	BankRelationshipID string `pathParam:"style=simple,explode=false,name=bankRelationship_id"`
	// The field of the bank relationship to update. Only `nickname` is supported.
	UpdateMask             *string                           `queryParam:"style=form,explode=true,name=update_mask"`
	BankRelationshipUpdate components.BankRelationshipUpdate `request:"mediaType=application/json"`
}

func (o *BankRelationshipsUpdateBankRelationshipRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BankRelationshipsUpdateBankRelationshipRequest) GetBankRelationshipID() string {
	if o == nil {
		return ""
	}
	return o.BankRelationshipID
}

func (o *BankRelationshipsUpdateBankRelationshipRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}

func (o *BankRelationshipsUpdateBankRelationshipRequest) GetBankRelationshipUpdate() components.BankRelationshipUpdate {
	if o == nil {
		return components.BankRelationshipUpdate{}
	}
	return o.BankRelationshipUpdate
}

type BankRelationshipsUpdateBankRelationshipResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	BankRelationship *components.BankRelationship
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *BankRelationshipsUpdateBankRelationshipResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BankRelationshipsUpdateBankRelationshipResponse) GetBankRelationship() *components.BankRelationship {
	if o == nil {
		return nil
	}
	return o.BankRelationship
}

func (o *BankRelationshipsUpdateBankRelationshipResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
