// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type CashJournalsCheckPartyTypeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CheckPartyTypeResponse *components.CheckPartyTypeResponse
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *CashJournalsCheckPartyTypeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CashJournalsCheckPartyTypeResponse) GetCheckPartyTypeResponse() *components.CheckPartyTypeResponse {
	if o == nil {
		return nil
	}
	return o.CheckPartyTypeResponse
}

func (o *CashJournalsCheckPartyTypeResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
