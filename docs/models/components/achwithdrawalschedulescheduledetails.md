# AchWithdrawalScheduleScheduleDetails

The transfer schedule details


## Fields

| Field                                                                                                                     | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `Amount`                                                                                                                  | [*components.AchWithdrawalScheduleAmount](../../models/components/achwithdrawalscheduleamount.md)                         | :heavy_minus_sign:                                                                                                        | A cash amount in the format of decimal value (mutually exclusive with 'full_disbursement')                                | {<br/>"value": "100.00"<br/>}                                                                                             |
| `ClientScheduleID`                                                                                                        | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | External identifier supplied by the API caller. Each request must have a unique pairing of client_schedule_id and account | ABC-123                                                                                                                   |
| `FullDisbursement`                                                                                                        | **bool*                                                                                                                   | :heavy_minus_sign:                                                                                                        | Flag to indicate a full disbursement transfer (mutually exclusive with 'amount')                                          | false                                                                                                                     |
| `ScheduleProperties`                                                                                                      | [*components.AchWithdrawalScheduleScheduleProperties](../../models/components/achwithdrawalschedulescheduleproperties.md) | :heavy_minus_sign:                                                                                                        | Common schedule properties                                                                                                |                                                                                                                           |