// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type InvestigationServiceLinkDocumentsRequest struct {
	// The investigation id.
	InvestigationID            string                                `pathParam:"style=simple,explode=false,name=investigation_id"`
	LinkDocumentsRequestCreate components.LinkDocumentsRequestCreate `request:"mediaType=application/json"`
}

func (o *InvestigationServiceLinkDocumentsRequest) GetInvestigationID() string {
	if o == nil {
		return ""
	}
	return o.InvestigationID
}

func (o *InvestigationServiceLinkDocumentsRequest) GetLinkDocumentsRequestCreate() components.LinkDocumentsRequestCreate {
	if o == nil {
		return components.LinkDocumentsRequestCreate{}
	}
	return o.LinkDocumentsRequestCreate
}

type InvestigationServiceLinkDocumentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LinkDocumentsResponse *components.LinkDocumentsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *InvestigationServiceLinkDocumentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InvestigationServiceLinkDocumentsResponse) GetLinkDocumentsResponse() *components.LinkDocumentsResponse {
	if o == nil {
		return nil
	}
	return o.LinkDocumentsResponse
}

func (o *InvestigationServiceLinkDocumentsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
