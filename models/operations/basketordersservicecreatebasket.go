// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type BasketOrdersServiceCreateBasketRequest struct {
	// The correspondent id.
	CorrespondentID string                  `pathParam:"style=simple,explode=false,name=correspondent_id"`
	BasketCreate    components.BasketCreate `request:"mediaType=application/json"`
}

func (o *BasketOrdersServiceCreateBasketRequest) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *BasketOrdersServiceCreateBasketRequest) GetBasketCreate() components.BasketCreate {
	if o == nil {
		return components.BasketCreate{}
	}
	return o.BasketCreate
}

type BasketOrdersServiceCreateBasketResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Basket *components.Basket
	// INVALID_ARGUMENT: Either the correspondent doesn't have a valid average price account, or there was an issue with one or more fields in the request.  In the latter case, the message field will contain details about which field failed validation and why.
	Status *components.Status
}

func (o *BasketOrdersServiceCreateBasketResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *BasketOrdersServiceCreateBasketResponse) GetBasket() *components.Basket {
	if o == nil {
		return nil
	}
	return o.Basket
}

func (o *BasketOrdersServiceCreateBasketResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
