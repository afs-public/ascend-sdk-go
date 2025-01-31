// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type RetirementConstraintsRetrieveDistributionConstraintsRequest struct {
	// The account id.
	AccountID                                    string                                                  `pathParam:"style=simple,explode=false,name=account_id"`
	RetrieveDistributionConstraintsRequestCreate components.RetrieveDistributionConstraintsRequestCreate `request:"mediaType=application/json"`
}

func (o *RetirementConstraintsRetrieveDistributionConstraintsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *RetirementConstraintsRetrieveDistributionConstraintsRequest) GetRetrieveDistributionConstraintsRequestCreate() components.RetrieveDistributionConstraintsRequestCreate {
	if o == nil {
		return components.RetrieveDistributionConstraintsRequestCreate{}
	}
	return o.RetrieveDistributionConstraintsRequestCreate
}

type RetirementConstraintsRetrieveDistributionConstraintsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	DistributionConstraints *components.DistributionConstraints
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *RetirementConstraintsRetrieveDistributionConstraintsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetirementConstraintsRetrieveDistributionConstraintsResponse) GetDistributionConstraints() *components.DistributionConstraints {
	if o == nil {
		return nil
	}
	return o.DistributionConstraints
}

func (o *RetirementConstraintsRetrieveDistributionConstraintsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
