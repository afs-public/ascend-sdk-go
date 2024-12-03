// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type AccountsCreateAccountResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Account *components.Account
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsCreateAccountResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsCreateAccountResponse) GetAccount() *components.Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *AccountsCreateAccountResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
