// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type OrderPriceServiceRetrieveFixedIncomeMarksRequest struct {
	// The correspondent id.
	CorrespondentID                       string                                           `pathParam:"style=simple,explode=false,name=correspondent_id"`
	RetrieveFixedIncomeMarksRequestCreate components.RetrieveFixedIncomeMarksRequestCreate `request:"mediaType=application/json"`
}

func (o *OrderPriceServiceRetrieveFixedIncomeMarksRequest) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *OrderPriceServiceRetrieveFixedIncomeMarksRequest) GetRetrieveFixedIncomeMarksRequestCreate() components.RetrieveFixedIncomeMarksRequestCreate {
	if o == nil {
		return components.RetrieveFixedIncomeMarksRequestCreate{}
	}
	return o.RetrieveFixedIncomeMarksRequestCreate
}

type OrderPriceServiceRetrieveFixedIncomeMarksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	RetrieveFixedIncomeMarksResponse *components.RetrieveFixedIncomeMarksResponse
	// INVALID_ARGUMENT: There was an issue with one or more fields in the request.  The message field will contain details about which field failed validation and why.
	Status *components.Status
}

func (o *OrderPriceServiceRetrieveFixedIncomeMarksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrderPriceServiceRetrieveFixedIncomeMarksResponse) GetRetrieveFixedIncomeMarksResponse() *components.RetrieveFixedIncomeMarksResponse {
	if o == nil {
		return nil
	}
	return o.RetrieveFixedIncomeMarksResponse
}

func (o *OrderPriceServiceRetrieveFixedIncomeMarksResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
