// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsDeleteTrustedContactRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The trustedContact id.
	TrustedContactID string `pathParam:"style=simple,explode=false,name=trustedContact_id"`
}

func (o *AccountsDeleteTrustedContactRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsDeleteTrustedContactRequest) GetTrustedContactID() string {
	if o == nil {
		return ""
	}
	return o.TrustedContactID
}

type AccountsDeleteTrustedContactResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsDeleteTrustedContactResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsDeleteTrustedContactResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
