// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountTransfersListTransfersRequest struct {
	// The correspondent id.
	CorrespondentID string `pathParam:"style=simple,explode=false,name=correspondent_id"`
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The maximum number of transfers to return. The service may return fewer than this value. If unspecified, at most 100 transfers will be returned. The maximum value is 200; values above 200 will be coerced to 200.
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListTransfers` call. Provide this to retrieve the subsequent page.
	//
	//  When paginating, all other parameters provided to `ListTransfers` must match the call that provided the page token.
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;
	//
	//  Currently supported CEL filters: -------------------------------- * acat_control_number        (add to queries for better performance) * state     == State.* * nscc_status  == NsccStatus.* * direction   == Direction.* * transfer_type == TransferType.* * reject_code  == RejectCode.* * create_time * last_nscc_status_updated_time * receiver.account_id * receiver.account_number * receiver.participant_number * receiver.correspondent_code * receiver.correspondent_id * deliverer.account_id * deliverer.account_number * deliverer.participant_number * deliverer.correspondent_code * deliverer.correspondent_id
	//
	//  - Empty filters are allowed, which return the most recent page_size worth of transfers, in practice this is not performant  and should be avoided if possible
	//
	//  - Macros are NOT allowed, using them will result in an INVALID_ARGUMENT being returned
	//
	//  - The following CEL operators are NOT allowed, using them will result in an INVALID_ARGUMENT being returned:    string.matches(substr)    +    -    /    *
	//
	//  - Queries using acat_control_number will result in increased performance
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *AccountTransfersListTransfersRequest) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *AccountTransfersListTransfersRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountTransfersListTransfersRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *AccountTransfersListTransfersRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *AccountTransfersListTransfersRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type AccountTransfersListTransfersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListTransfersResponse *components.ListTransfersResponse
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *AccountTransfersListTransfersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountTransfersListTransfersResponse) GetListTransfersResponse() *components.ListTransfersResponse {
	if o == nil {
		return nil
	}
	return o.ListTransfersResponse
}

func (o *AccountTransfersListTransfersResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
