# SetOptionExtraReportingDataRequestCreate

Request message for SetOptionExtraReportingData.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `CancelConfirmedTime`                                                                 | [time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_check_mark:                                                                    | The time the correspondent confirmed receipt of the cancellation.                     | 2025-12-13T15:28:17.262732Z                                                           |
| `Name`                                                                                | *string*                                                                              | :heavy_check_mark:                                                                    | Format: accounts/{account_id}/optionOrders/{option_order_id}                          | accounts/01HBRQ5BW6ZAY4BNWP4GWRD80X/optionOrders/ebb0c9b5-2c74-45c9-a4ab-40596b778706 |