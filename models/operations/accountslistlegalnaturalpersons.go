// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type AccountsListLegalNaturalPersonsRequest struct {
	// The maximum number of legal natural persons to return. The service may return fewer than this value. If unspecified, at most 25 legal natural persons will be returned. The maximum value is 100; values above 100 will be coerced to 100.
	PageSize *int `queryParam:"style=form,explode=true,name=page_size"`
	// A page token, received from a previous `ListLegalNaturalPersons` call. Provide this to retrieve the subsequent page.
	//
	//  When paginating, all other parameters provided to `ListLegalNaturalPersons` must match the call that provided the page token.
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
	// The order in which legal natural persons are listed.
	OrderBy *string `queryParam:"style=form,explode=true,name=order_by"`
	// A CEL string to filter results; Use `upperAscii()` for case-insensitive searches; E.g. `given_name.upperAscii()=="rUsTy".upperAscii()`; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:
	//  `legal_natural_person_id`
	//  `correspondent_id`
	//  `given_name`
	//  `family_name`
	//  `tax_id`
	//  `tax_id_type`
	//  `investigation_id`
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *AccountsListLegalNaturalPersonsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *AccountsListLegalNaturalPersonsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *AccountsListLegalNaturalPersonsRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *AccountsListLegalNaturalPersonsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type AccountsListLegalNaturalPersonsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	ListLegalNaturalPersonsResponse *components.ListLegalNaturalPersonsResponse
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *AccountsListLegalNaturalPersonsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountsListLegalNaturalPersonsResponse) GetListLegalNaturalPersonsResponse() *components.ListLegalNaturalPersonsResponse {
	if o == nil {
		return nil
	}
	return o.ListLegalNaturalPersonsResponse
}

func (o *AccountsListLegalNaturalPersonsResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
