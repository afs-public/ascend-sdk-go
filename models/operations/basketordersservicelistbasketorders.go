// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type BasketOrdersServiceListBasketOrdersRequest struct {
	// The correspondent id.
	CorrespondentID string `pathParam:"style=simple,explode=false,name=correspondent_id"`
	// The basket id.
	BasketID string `pathParam:"style=simple,explode=false,name=basket_id"`
	// The maximum number of basket orders to return. The service may return fewer than this value. If unspecified, at most 1000 basket orders will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListBasketOrders` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListBasketOrders` must match the call that provided the page token.
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *BasketOrdersServiceListBasketOrdersRequest) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *BasketOrdersServiceListBasketOrdersRequest) GetBasketID() string {
	if o == nil {
		return ""
	}
	return o.BasketID
}

func (o *BasketOrdersServiceListBasketOrdersRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *BasketOrdersServiceListBasketOrdersRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type BasketOrdersServiceListBasketOrdersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListBasketOrdersResponse *components.ListBasketOrdersResponse
	// INVALID_ARGUMENT: The correspondent_id or the basket_id could not be determined for the request.
	Status *components.Status
}

func (o *BasketOrdersServiceListBasketOrdersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BasketOrdersServiceListBasketOrdersResponse) GetListBasketOrdersResponse() *components.ListBasketOrdersResponse {
	if o == nil {
		return nil
	}
	return o.ListBasketOrdersResponse
}

func (o *BasketOrdersServiceListBasketOrdersResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
