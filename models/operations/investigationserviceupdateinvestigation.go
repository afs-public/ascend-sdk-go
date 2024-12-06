// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type InvestigationServiceUpdateInvestigationRequest struct {
	// The investigation id.
	InvestigationID     string                         `pathParam:"style=simple,explode=false,name=investigation_id"`
	InvestigationUpdate components.InvestigationUpdate `request:"mediaType=application/json"`
}

func (o *InvestigationServiceUpdateInvestigationRequest) GetInvestigationID() string {
	if o == nil {
		return ""
	}
	return o.InvestigationID
}

func (o *InvestigationServiceUpdateInvestigationRequest) GetInvestigationUpdate() components.InvestigationUpdate {
	if o == nil {
		return components.InvestigationUpdate{}
	}
	return o.InvestigationUpdate
}

type InvestigationServiceUpdateInvestigationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Investigation *components.Investigation
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *InvestigationServiceUpdateInvestigationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InvestigationServiceUpdateInvestigationResponse) GetInvestigation() *components.Investigation {
	if o == nil {
		return nil
	}
	return o.Investigation
}

func (o *InvestigationServiceUpdateInvestigationResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}