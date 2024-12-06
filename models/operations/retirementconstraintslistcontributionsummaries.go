// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type RetirementConstraintsListContributionSummariesRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Number of contribution summaries to get (partitioned by tax year) Default = 2 (current year and prior year), maximum = 10
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// When paginating, this is used to retrieve a specific page from the overall response
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *RetirementConstraintsListContributionSummariesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetirementConstraintsListContributionSummariesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *RetirementConstraintsListContributionSummariesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type RetirementConstraintsListContributionSummariesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListContributionSummariesResponse *components.ListContributionSummariesResponse
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *RetirementConstraintsListContributionSummariesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetirementConstraintsListContributionSummariesResponse) GetListContributionSummariesResponse() *components.ListContributionSummariesResponse {
	if o == nil {
		return nil
	}
	return o.ListContributionSummariesResponse
}

func (o *RetirementConstraintsListContributionSummariesResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}