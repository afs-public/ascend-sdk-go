// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EndReason - The end reason of the LTID.
type EndReason string

const (
	EndReasonReportableAccountEventReasonUnspecified EndReason = "REPORTABLE_ACCOUNT_EVENT_REASON_UNSPECIFIED"
	EndReasonEventReasonCreated                      EndReason = "EVENT_REASON_CREATED"
	EndReasonEventReasonCorrection                   EndReason = "EVENT_REASON_CORRECTION"
	EndReasonEventReasonEnded                        EndReason = "EVENT_REASON_ENDED"
	EndReasonEventReasonReplaced                     EndReason = "EVENT_REASON_REPLACED"
	EndReasonEventReasonTransfer                     EndReason = "EVENT_REASON_TRANSFER"
	EndReasonEventReasonOther                        EndReason = "EVENT_REASON_OTHER"
)

func (e EndReason) ToPointer() *EndReason {
	return &e
}

// EndLargeTraderRequestCreate - The request to end a Large Trader on a Legal Natural Person/Legal Entity.
type EndLargeTraderRequestCreate struct {
	// The end reason of the LTID.
	EndReason EndReason `json:"end_reason"`
}

func (o *EndLargeTraderRequestCreate) GetEndReason() EndReason {
	if o == nil {
		return EndReason("")
	}
	return o.EndReason
}