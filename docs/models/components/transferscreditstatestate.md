# TransfersCreditStateState

The high level state of a transfer, one of:
- `PROCESSING` - The transfer is being processed and will be posted if successful.
- `PENDING_REVIEW` - The transfer is pending review and will continue processing if approved.
- `POSTED` - The transfer has been posted to the ledger and will be completed at the end of the processing window if not canceled first.
- `COMPLETED` - The transfer has been batched and completed.
- `REJECTED` - The transfer was rejected.
- `CANCELED` - The transfer was canceled.
- `RETURNED` - The transfer was returned.
- `POSTPONED` - The transfer is postponed and will resume processing during the next processing window.


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `TransfersCreditStateStateStateUnspecified` | STATE_UNSPECIFIED                           |
| `TransfersCreditStateStateProcessing`       | PROCESSING                                  |
| `TransfersCreditStateStatePendingReview`    | PENDING_REVIEW                              |
| `TransfersCreditStateStatePosted`           | POSTED                                      |
| `TransfersCreditStateStateCompleted`        | COMPLETED                                   |
| `TransfersCreditStateStateRejected`         | REJECTED                                    |
| `TransfersCreditStateStateCanceled`         | CANCELED                                    |
| `TransfersCreditStateStateReturned`         | RETURNED                                    |
| `TransfersCreditStateStatePostponed`        | POSTPONED                                   |