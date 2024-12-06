// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsDeleteInterestedPartyRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The interestedParty id.
	InterestedPartyID string `pathParam:"style=simple,explode=false,name=interestedParty_id"`
}

func (o *AccountsDeleteInterestedPartyRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsDeleteInterestedPartyRequest) GetInterestedPartyID() string {
	if o == nil {
		return ""
	}
	return o.InterestedPartyID
}

type AccountsDeleteInterestedPartyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsDeleteInterestedPartyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsDeleteInterestedPartyResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}