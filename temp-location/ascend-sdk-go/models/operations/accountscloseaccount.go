// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsCloseAccountRequest struct {
	// The account id.
	AccountID                 string                               `pathParam:"style=simple,explode=false,name=account_id"`
	CloseAccountRequestCreate components.CloseAccountRequestCreate `request:"mediaType=application/json"`
}

func (o *AccountsCloseAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsCloseAccountRequest) GetCloseAccountRequestCreate() components.CloseAccountRequestCreate {
	if o == nil {
		return components.CloseAccountRequestCreate{}
	}
	return o.CloseAccountRequestCreate
}

type AccountsCloseAccountResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Account *components.Account
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsCloseAccountResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsCloseAccountResponse) GetAccount() *components.Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *AccountsCloseAccountResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
