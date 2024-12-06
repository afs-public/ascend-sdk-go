# AchWithdrawalSchedule

A withdrawal transfer schedule using the ACH mechanism


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         | Example                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `BankRelationship`                                                                                                  | **string*                                                                                                           | :heavy_minus_sign:                                                                                                  | The name of the bank relationship to be used in the ACH transaction                                                 | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e                                      |
| `IraDistribution`                                                                                                   | [*components.IraDistribution](../../models/components/iradistribution.md)                                           | :heavy_minus_sign:                                                                                                  | The ira distribution info for an IRA account                                                                        |                                                                                                                     |
| `Name`                                                                                                              | **string*                                                                                                           | :heavy_minus_sign:                                                                                                  | The name of the ACH Withdrawal transfer schedule                                                                    | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawalSchedules/40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                     |
| `ScheduleDetails`                                                                                                   | [*components.AchWithdrawalScheduleScheduleDetails](../../models/components/achwithdrawalschedulescheduledetails.md) | :heavy_minus_sign:                                                                                                  | The transfer schedule details                                                                                       |                                                                                                                     |