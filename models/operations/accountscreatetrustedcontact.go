// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsCreateTrustedContactRequest struct {
	// The account id.
	AccountID            string                          `pathParam:"style=simple,explode=false,name=account_id"`
	TrustedContactCreate components.TrustedContactCreate `request:"mediaType=application/json"`
}

func (o *AccountsCreateTrustedContactRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsCreateTrustedContactRequest) GetTrustedContactCreate() components.TrustedContactCreate {
	if o == nil {
		return components.TrustedContactCreate{}
	}
	return o.TrustedContactCreate
}

type AccountsCreateTrustedContactResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	TrustedContact *components.TrustedContact
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsCreateTrustedContactResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsCreateTrustedContactResponse) GetTrustedContact() *components.TrustedContact {
	if o == nil {
		return nil
	}
	return o.TrustedContact
}

func (o *AccountsCreateTrustedContactResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
