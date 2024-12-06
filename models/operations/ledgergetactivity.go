// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type LedgerGetActivityRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The activity id.
	ActivityID string `pathParam:"style=simple,explode=false,name=activity_id"`
}

func (o *LedgerGetActivityRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *LedgerGetActivityRequest) GetActivityID() string {
	if o == nil {
		return ""
	}
	return o.ActivityID
}

type LedgerGetActivityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Activity *components.Activity
	// INVALID_ARGUMENT: The request is not valid, additional information may be present in the BadRequest details.
	Status *components.Status
}

func (o *LedgerGetActivityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *LedgerGetActivityResponse) GetActivity() *components.Activity {
	if o == nil {
		return nil
	}
	return o.Activity
}

func (o *LedgerGetActivityResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}