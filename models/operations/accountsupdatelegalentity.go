// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsUpdateLegalEntityRequest struct {
	// The legalEntity id.
	LegalEntityID     string                       `pathParam:"style=simple,explode=false,name=legalEntity_id"`
	LegalEntityUpdate components.LegalEntityUpdate `request:"mediaType=application/json"`
}

func (o *AccountsUpdateLegalEntityRequest) GetLegalEntityID() string {
	if o == nil {
		return ""
	}
	return o.LegalEntityID
}

func (o *AccountsUpdateLegalEntityRequest) GetLegalEntityUpdate() components.LegalEntityUpdate {
	if o == nil {
		return components.LegalEntityUpdate{}
	}
	return o.LegalEntityUpdate
}

type AccountsUpdateLegalEntityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LegalEntity *components.LegalEntity
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsUpdateLegalEntityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsUpdateLegalEntityResponse) GetLegalEntity() *components.LegalEntity {
	if o == nil {
		return nil
	}
	return o.LegalEntity
}

func (o *AccountsUpdateLegalEntityResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
