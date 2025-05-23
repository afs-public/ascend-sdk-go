// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsGetLegalEntityRequest struct {
	// The legalEntity id.
	LegalEntityID string `pathParam:"style=simple,explode=false,name=legalEntity_id"`
}

func (o *AccountsGetLegalEntityRequest) GetLegalEntityID() string {
	if o == nil {
		return ""
	}
	return o.LegalEntityID
}

type AccountsGetLegalEntityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LegalEntity *components.LegalEntity
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsGetLegalEntityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsGetLegalEntityResponse) GetLegalEntity() *components.LegalEntity {
	if o == nil {
		return nil
	}
	return o.LegalEntity
}

func (o *AccountsGetLegalEntityResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
