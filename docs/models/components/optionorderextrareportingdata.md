# OptionOrderExtraReportingData

Post-cancel reporting data provided via the SetOptionExtraReportingData endpoint.


## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `CancelConfirmedTime`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                                         | :heavy_minus_sign:                                                                                 | Correspondent-side timestamp when the cancellation was confirmed. Maps to CAT CANCEL_CONFIRM_TIME. | 2025-12-13T15:28:17.262732Z                                                                        |