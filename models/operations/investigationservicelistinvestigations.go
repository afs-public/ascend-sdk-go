// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type InvestigationServiceListInvestigationsRequest struct {
	// The maximum number of records to return. Default is 50 The maximum is 200, values exceeding this will be set to 200
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// The page token used to request a specific page of the search results
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:
	//  ListInvestigationStatesResponse.investigation_states
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *InvestigationServiceListInvestigationsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *InvestigationServiceListInvestigationsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *InvestigationServiceListInvestigationsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type InvestigationServiceListInvestigationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListInvestigationsResponse *components.ListInvestigationsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *InvestigationServiceListInvestigationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *InvestigationServiceListInvestigationsResponse) GetListInvestigationsResponse() *components.ListInvestigationsResponse {
	if o == nil {
		return nil
	}
	return o.ListInvestigationsResponse
}

func (o *InvestigationServiceListInvestigationsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
