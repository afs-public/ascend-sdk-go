// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/afs-public/ascend-sdk-go/models/components"
)

type CashJournalsCancelCashJournalRequest struct {
	// The cashJournal id.
	CashJournalID                  string                                    `pathParam:"style=simple,explode=false,name=cashJournal_id"`
	CancelCashJournalRequestCreate components.CancelCashJournalRequestCreate `request:"mediaType=application/json"`
}

func (o *CashJournalsCancelCashJournalRequest) GetCashJournalID() string {
	if o == nil {
		return ""
	}
	return o.CashJournalID
}

func (o *CashJournalsCancelCashJournalRequest) GetCancelCashJournalRequestCreate() components.CancelCashJournalRequestCreate {
	if o == nil {
		return components.CancelCashJournalRequestCreate{}
	}
	return o.CancelCashJournalRequestCreate
}

type CashJournalsCancelCashJournalResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CashJournal *components.CashJournal
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *CashJournalsCancelCashJournalResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CashJournalsCancelCashJournalResponse) GetCashJournal() *components.CashJournal {
	if o == nil {
		return nil
	}
	return o.CashJournal
}

func (o *CashJournalsCancelCashJournalResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
