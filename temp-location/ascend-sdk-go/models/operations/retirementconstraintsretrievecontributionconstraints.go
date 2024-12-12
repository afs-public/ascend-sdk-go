// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type RetirementConstraintsRetrieveContributionConstraintsRequest struct {
	// The account id.
	AccountID                                    string                                                  `pathParam:"style=simple,explode=false,name=account_id"`
	RetrieveContributionConstraintsRequestCreate components.RetrieveContributionConstraintsRequestCreate `request:"mediaType=application/json"`
}

func (o *RetirementConstraintsRetrieveContributionConstraintsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetirementConstraintsRetrieveContributionConstraintsRequest) GetRetrieveContributionConstraintsRequestCreate() components.RetrieveContributionConstraintsRequestCreate {
	if o == nil {
		return components.RetrieveContributionConstraintsRequestCreate{}
	}
	return o.RetrieveContributionConstraintsRequestCreate
}

type RetirementConstraintsRetrieveContributionConstraintsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ContributionConstraints *components.ContributionConstraints
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *RetirementConstraintsRetrieveContributionConstraintsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetirementConstraintsRetrieveContributionConstraintsResponse) GetContributionConstraints() *components.ContributionConstraints {
	if o == nil {
		return nil
	}
	return o.ContributionConstraints
}

func (o *RetirementConstraintsRetrieveContributionConstraintsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
