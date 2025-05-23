// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type BookingCancelTradeRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The trade id.
	TradeID                  string                              `pathParam:"style=simple,explode=false,name=trade_id"`
	CancelTradeRequestCreate components.CancelTradeRequestCreate `request:"mediaType=application/json"`
}

func (o *BookingCancelTradeRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *BookingCancelTradeRequest) GetTradeID() string {
	if o == nil {
		return ""
	}
	return o.TradeID
}

func (o *BookingCancelTradeRequest) GetCancelTradeRequestCreate() components.CancelTradeRequestCreate {
	if o == nil {
		return components.CancelTradeRequestCreate{}
	}
	return o.CancelTradeRequestCreate
}

type BookingCancelTradeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CancelTradeResponse *components.CancelTradeResponse
	// INVALID_ARGUMENT: The request is not valid.
	// FAILED_PRECONDITION: The operation was rejected because the system is not in a state required for the operation's processing.
	Status *components.Status
}

func (o *BookingCancelTradeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BookingCancelTradeResponse) GetCancelTradeResponse() *components.CancelTradeResponse {
	if o == nil {
		return nil
	}
	return o.CancelTradeResponse
}

func (o *BookingCancelTradeResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
