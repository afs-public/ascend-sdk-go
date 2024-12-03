// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
)

type OrderServiceCancelOrderRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The order id.
	OrderID                  string                              `pathParam:"style=simple,explode=false,name=order_id"`
	CancelOrderRequestCreate components.CancelOrderRequestCreate `request:"mediaType=application/json"`
}

func (o *OrderServiceCancelOrderRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *OrderServiceCancelOrderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *OrderServiceCancelOrderRequest) GetCancelOrderRequestCreate() components.CancelOrderRequestCreate {
	if o == nil {
		return components.CancelOrderRequestCreate{}
	}
	return o.CancelOrderRequestCreate
}

type OrderServiceCancelOrderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Order *components.Order
	// INVALID_ARGUMENT: The account_id or the order_id could not be determined for the request.
	// FAILED_PRECONDITION: The order is not in a cancelable state
	Status *components.Status
}

func (o *OrderServiceCancelOrderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrderServiceCancelOrderResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *OrderServiceCancelOrderResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
