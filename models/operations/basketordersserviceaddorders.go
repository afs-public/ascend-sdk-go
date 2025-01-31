// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type BasketOrdersServiceAddOrdersRequest struct {
	// The correspondent id.
	CorrespondentID string `pathParam:"style=simple,explode=false,name=correspondent_id"`
	// The basket id.
	BasketID               string                            `pathParam:"style=simple,explode=false,name=basket_id"`
	AddOrdersRequestCreate components.AddOrdersRequestCreate `request:"mediaType=application/json"`
}

func (o *BasketOrdersServiceAddOrdersRequest) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *BasketOrdersServiceAddOrdersRequest) GetBasketID() string {
	if o == nil {
		return ""
	}
	return o.BasketID
}

func (o *BasketOrdersServiceAddOrdersRequest) GetAddOrdersRequestCreate() components.AddOrdersRequestCreate {
	if o == nil {
		return components.AddOrdersRequestCreate{}
	}
	return o.AddOrdersRequestCreate
}

type BasketOrdersServiceAddOrdersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Basket *components.Basket
	// INVALID_ARGUMENT: There was an issue with one or more fields in the request.  The message field will contain details about which field failed validation and why.
	// FAILED_PRECONDITION: The requested basket has already been submitted.
	Status *components.Status
}

func (o *BasketOrdersServiceAddOrdersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BasketOrdersServiceAddOrdersResponse) GetBasket() *components.Basket {
	if o == nil {
		return nil
	}
	return o.Basket
}

func (o *BasketOrdersServiceAddOrdersResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
