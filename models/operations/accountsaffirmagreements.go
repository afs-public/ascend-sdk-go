// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsAffirmAgreementsRequest struct {
	// The account id.
	AccountID                     string                                   `pathParam:"style=simple,explode=false,name=account_id"`
	AffirmAgreementsRequestCreate components.AffirmAgreementsRequestCreate `request:"mediaType=application/json"`
}

func (o *AccountsAffirmAgreementsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountsAffirmAgreementsRequest) GetAffirmAgreementsRequestCreate() components.AffirmAgreementsRequestCreate {
	if o == nil {
		return components.AffirmAgreementsRequestCreate{}
	}
	return o.AffirmAgreementsRequestCreate
}

type AccountsAffirmAgreementsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	AffirmAgreementsResponse *components.AffirmAgreementsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsAffirmAgreementsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsAffirmAgreementsResponse) GetAffirmAgreementsResponse() *components.AffirmAgreementsResponse {
	if o == nil {
		return nil
	}
	return o.AffirmAgreementsResponse
}

func (o *AccountsAffirmAgreementsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
