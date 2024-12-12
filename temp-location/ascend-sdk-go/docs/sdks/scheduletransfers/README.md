# ScheduleTransfers
(*ScheduleTransfers*)

## Overview

### Available Operations

* [ListScheduleSummaries](#listschedulesummaries) - List Schedule Summaries
* [CreateAchDepositSchedule](#createachdepositschedule) - Create ACH Deposit Schedule
* [ListAchDepositSchedules](#listachdepositschedules) - List ACH Deposit Schedules
* [GetAchDepositSchedule](#getachdepositschedule) - Get ACH Deposit Schedule
* [UpdateAchDepositSchedule](#updateachdepositschedule) - Update ACH Deposit Schedules
* [CancelAchDepositSchedule](#cancelachdepositschedule) - Cancel ACH Deposit Schedule
* [CreateAchWithdrawalSchedule](#createachwithdrawalschedule) - Create ACH Withdrawal Schedule
* [ListAchWithdrawalSchedules](#listachwithdrawalschedules) - List ACH Withdrawal Schedules
* [GetAchWithdrawalSchedule](#getachwithdrawalschedule) - Get ACH Withdrawal Schedule
* [UpdateAchWithdrawalSchedule](#updateachwithdrawalschedule) - Update ACH Withdrawal Schedule
* [CancelAchWithdrawalSchedule](#cancelachwithdrawalschedule) - Cancel ACH Withdrawal Schedule
* [CreateWireWithdrawalSchedule](#createwirewithdrawalschedule) - Create Wire Withdrawal Schedule
* [ListWireWithdrawalSchedules](#listwirewithdrawalschedules) - List Wire Withdrawal Schedules
* [GetWireWithdrawalSchedule](#getwirewithdrawalschedule) - Get Wire Withdrawal Schedule
* [UpdateWireWithdrawalSchedule](#updatewirewithdrawalschedule) - Update Wire Withdrawal Schedule
* [CancelWireWithdrawalSchedule](#cancelwirewithdrawalschedule) - Cancel Wire Withdrawal Schedule

## ListScheduleSummaries

Lists transfer schedule summaries that match the filter in the request

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.ListScheduleSummaries(ctx, ascendsdkgo.String("mechanism == 'ACH' && direction == DEPOSIT && state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'"), ascendsdkgo.Int(100), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListScheduleSummariesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `account`<br/> `mechanism`<br/> `direction`<br/> `state`<br/> `start_date`<br/> `end_date` | mechanism == 'ACH' && direction == DEPOSIT && state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'                                                                                                                                              |
| `pageSize`                                                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                   | The maximum number of schedules to return. The service may return fewer than this value. If unspecified, at most 25 schedules will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.                                                | 100                                                                                                                                                                                                                                                                  |
| `pageToken`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | The page token to request                                                                                                                                                                                                                                            | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                                                          |
| `opts`                                                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                      |

### Response

**[*operations.TransferScheduleSummariesListScheduleSummariesResponse](../../models/operations/transferschedulesummarieslistschedulesummariesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAchDepositSchedule

Creates an ACH deposit transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CreateAchDepositSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.AchDepositScheduleCreate{
        BankRelationship: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
        IraContribution: &components.IraContribution{
            TaxYear: 2024,
            Type: components.AchDepositScheduleCreateTypeRegular,
        },
        RetirementContribution: &components.ScheduledRetirementContributionCreate{
            TaxYear: ascendsdkgo.Int(2024),
            TemporalTaxYear: components.TemporalTaxYearCurrentCalendarYear.ToPointer(),
            Type: components.ScheduledRetirementContributionCreateTypeRegular,
        },
        ScheduleDetails: components.DepositScheduleDetailsCreate{
            Amount: components.DecimalCreate{},
            ClientScheduleID: "ABC-123",
            ScheduleProperties: components.SchedulePropertiesCreate{
                Occurrences: ascendsdkgo.Int(12),
                StartDate: components.DateCreate{},
                TimeUnit: components.TimeUnitMonth,
                UnitMultiplier: 1,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDepositSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | The account id.                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                 |
| `achDepositScheduleCreate`                                                                 | [components.AchDepositScheduleCreate](../../models/components/achdepositschedulecreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.AchDepositSchedulesCreateAchDepositScheduleResponse](../../models/operations/achdepositschedulescreateachdepositscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAchDepositSchedules

Return a list of ACH deposit schedules for the specified account and filter params

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.ListAchDepositSchedules(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", ascendsdkgo.String("state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'"), ascendsdkgo.Int(100), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAchDepositSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                             |                                                                                                                                                                                                                                 |
| `accountID`                                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                              | The account id.                                                                                                                                                                                                                 | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `state`<br/> `start_date`<br/> `end_date` | state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'                                                                                                                                                       |
| `pageSize`                                                                                                                                                                                                                      | **int*                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                              | The maximum number of schedules to return. The service may return fewer than this value. If unspecified, at most 25 schedules will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.           | 100                                                                                                                                                                                                                             |
| `pageToken`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | The page token to request                                                                                                                                                                                                       | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                   |                                                                                                                                                                                                                                 |

### Response

**[*operations.AchDepositSchedulesListAchDepositSchedulesResponse](../../models/operations/achdepositscheduleslistachdepositschedulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAchDepositSchedule

Gets an ACH deposit transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.GetAchDepositSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1")
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDepositSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `achDepositScheduleID`                                   | *string*                                                 | :heavy_check_mark:                                       | The achDepositSchedule id.                               | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AchDepositSchedulesGetAchDepositScheduleResponse](../../models/operations/achdepositschedulesgetachdepositscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAchDepositSchedule

Updates the amount of an ACH deposit transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.UpdateAchDepositSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.AchDepositScheduleUpdate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDepositSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `accountID`                                                                                                               | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The account id.                                                                                                           | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                |
| `achDepositScheduleID`                                                                                                    | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The achDepositSchedule id.                                                                                                | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                      |
| `achDepositScheduleUpdate`                                                                                                | [components.AchDepositScheduleUpdate](../../models/components/achdepositscheduleupdate.md)                                | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |                                                                                                                           |
| `updateMask`                                                                                                              | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | A field mask representing the update. Note: only the 'schedule_details.amount' field of a schedule is currently updatable |                                                                                                                           |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.AchDepositSchedulesUpdateAchDepositScheduleResponse](../../models/operations/achdepositschedulesupdateachdepositscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelAchDepositSchedule

Cancels an ACH deposit transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CancelAchDepositSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.CancelAchDepositScheduleRequestCreate{
        Comment: ascendsdkgo.String("User Request"),
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDepositSchedules/40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDepositSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `accountID`                                                                                                          | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The account id.                                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                           |
| `achDepositScheduleID`                                                                                               | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The achDepositSchedule id.                                                                                           | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                 |
| `cancelAchDepositScheduleRequestCreate`                                                                              | [components.CancelAchDepositScheduleRequestCreate](../../models/components/cancelachdepositschedulerequestcreate.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.AchDepositSchedulesCancelAchDepositScheduleResponse](../../models/operations/achdepositschedulescancelachdepositscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAchWithdrawalSchedule

Creates an ACH withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CreateAchWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.AchWithdrawalScheduleCreate{
        BankRelationship: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/bankRelationships/651ef9de0dee00240813e60e",
        IraDistribution: &components.RetirementDistributionCreate{
            Type: components.RetirementDistributionCreateTypeNormal,
        },
        ScheduleDetails: components.WithdrawalScheduleDetailsCreate{
            ClientScheduleID: "ABC-123",
            FullDisbursement: ascendsdkgo.Bool(false),
            ScheduleProperties: components.SchedulePropertiesCreate{
                Occurrences: ascendsdkgo.Int(12),
                StartDate: components.DateCreate{},
                TimeUnit: components.TimeUnitMonth,
                UnitMultiplier: 1,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |                                                                                                  |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | The account id.                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                       |
| `achWithdrawalScheduleCreate`                                                                    | [components.AchWithdrawalScheduleCreate](../../models/components/achwithdrawalschedulecreate.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |                                                                                                  |

### Response

**[*operations.AchWithdrawalSchedulesCreateAchWithdrawalScheduleResponse](../../models/operations/achwithdrawalschedulescreateachwithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAchWithdrawalSchedules

Return a list of ACH withdrawal schedules for the specified account and filter params

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.ListAchWithdrawalSchedules(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", ascendsdkgo.String("state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'"), ascendsdkgo.Int(100), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAchWithdrawalSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                             |                                                                                                                                                                                                                                 |
| `accountID`                                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                              | The account id.                                                                                                                                                                                                                 | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `state`<br/> `start_date`<br/> `end_date` | state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'                                                                                                                                                       |
| `pageSize`                                                                                                                                                                                                                      | **int*                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                              | The maximum number of schedules to return. The service may return fewer than this value. If unspecified, at most 25 schedules will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.           | 100                                                                                                                                                                                                                             |
| `pageToken`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | The page token to request                                                                                                                                                                                                       | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                   |                                                                                                                                                                                                                                 |

### Response

**[*operations.AchWithdrawalSchedulesListAchWithdrawalSchedulesResponse](../../models/operations/achwithdrawalscheduleslistachwithdrawalschedulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAchWithdrawalSchedule

Gets an ACH withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.GetAchWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1")
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `achWithdrawalScheduleID`                                | *string*                                                 | :heavy_check_mark:                                       | The achWithdrawalSchedule id.                            | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.AchWithdrawalSchedulesGetAchWithdrawalScheduleResponse](../../models/operations/achwithdrawalschedulesgetachwithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAchWithdrawalSchedule

Updates the amount of an ACH withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.UpdateAchWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.AchWithdrawalScheduleUpdate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `accountID`                                                                                                               | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The account id.                                                                                                           | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                |
| `achWithdrawalScheduleID`                                                                                                 | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The achWithdrawalSchedule id.                                                                                             | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                      |
| `achWithdrawalScheduleUpdate`                                                                                             | [components.AchWithdrawalScheduleUpdate](../../models/components/achwithdrawalscheduleupdate.md)                          | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |                                                                                                                           |
| `updateMask`                                                                                                              | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | A field mask representing the update. Note: only the 'schedule_details.amount' field of a schedule is currently updatable |                                                                                                                           |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.AchWithdrawalSchedulesUpdateAchWithdrawalScheduleResponse](../../models/operations/achwithdrawalschedulesupdateachwithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelAchWithdrawalSchedule

Cancels an ACH withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CancelAchWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.CancelAchWithdrawalScheduleRequestCreate{
        Comment: ascendsdkgo.String("User Request"),
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawalSchedules/40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                | Example                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |                                                                                                                            |
| `accountID`                                                                                                                | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The account id.                                                                                                            | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                 |
| `achWithdrawalScheduleID`                                                                                                  | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The achWithdrawalSchedule id.                                                                                              | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                       |
| `cancelAchWithdrawalScheduleRequestCreate`                                                                                 | [components.CancelAchWithdrawalScheduleRequestCreate](../../models/components/cancelachwithdrawalschedulerequestcreate.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |                                                                                                                            |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |                                                                                                                            |

### Response

**[*operations.AchWithdrawalSchedulesCancelAchWithdrawalScheduleResponse](../../models/operations/achwithdrawalschedulescancelachwithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateWireWithdrawalSchedule

Creates a Wire withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CreateWireWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.WireWithdrawalScheduleCreate{
        Beneficiary: components.WireWithdrawalBeneficiaryCreate{
            Account: "73849218650987",
            AccountTitle: ascendsdkgo.String("Jane Dough"),
            ThirdParty: ascendsdkgo.Bool(false),
        },
        Intermediary: &components.WireWithdrawalIntermediaryCreate{
            Account: "NL02ABNA0123456789",
            AccountTitle: "Jane Dough",
            Address: components.AddressCreate{},
        },
        RecipientBank: components.WireWithdrawalRecipientBankCreate{
            BankID: components.RecipientBankBankIDCreate{
                ID: "ABNANL2AXXX",
                Type: components.RecipientBankBankIDCreateTypeBic,
            },
            InternationalBankDetails: &components.RecipientBankBankDetailsCreate{
                AdditionalInfo: ascendsdkgo.String("Jane Dough transfer through intermediary account"),
                Address: components.AddressCreate{},
                BankName: "ABN AMRO BANK N.V.",
            },
        },
        RetirementDistribution: &components.RetirementDistributionCreate{
            Type: components.RetirementDistributionCreateTypeNormal,
        },
        ScheduleDetails: components.WithdrawalScheduleDetailsCreate{
            ClientScheduleID: "ABC-123",
            FullDisbursement: ascendsdkgo.Bool(false),
            ScheduleProperties: components.SchedulePropertiesCreate{
                Occurrences: ascendsdkgo.Int(12),
                StartDate: components.DateCreate{},
                TimeUnit: components.TimeUnitMonth,
                UnitMultiplier: 1,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `accountID`                                                                                        | *string*                                                                                           | :heavy_check_mark:                                                                                 | The account id.                                                                                    | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                         |
| `wireWithdrawalScheduleCreate`                                                                     | [components.WireWithdrawalScheduleCreate](../../models/components/wirewithdrawalschedulecreate.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.WireWithdrawalSchedulesCreateWireWithdrawalScheduleResponse](../../models/operations/wirewithdrawalschedulescreatewirewithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 409      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListWireWithdrawalSchedules

Return a list of Wire withdrawal schedules for the specified account and filter params

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.ListWireWithdrawalSchedules(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", ascendsdkgo.String("state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'"), ascendsdkgo.Int(100), ascendsdkgo.String("4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWireWithdrawalSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                             |                                                                                                                                                                                                                                 |
| `accountID`                                                                                                                                                                                                                     | *string*                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                              | The account id.                                                                                                                                                                                                                 | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; Filter options include:<br/> `state`<br/> `start_date`<br/> `end_date` | state == 'ACTIVE' && start_date > '2024-04-05' && end_date < '2024-08-10'                                                                                                                                                       |
| `pageSize`                                                                                                                                                                                                                      | **int*                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                              | The maximum number of schedules to return. The service may return fewer than this value. If unspecified, at most 25 schedules will be returned. The maximum value is 1000; values above 1000 will be coerced to 1000.           | 100                                                                                                                                                                                                                             |
| `pageToken`                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                              | The page token to request                                                                                                                                                                                                       | 4ZHd3wAaMD1IQ0ZKS2BKV0FSRVdLW4VLWkY1R1B3MU4                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                   |                                                                                                                                                                                                                                 |

### Response

**[*operations.WireWithdrawalSchedulesListWireWithdrawalSchedulesResponse](../../models/operations/wirewithdrawalscheduleslistwirewithdrawalschedulesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetWireWithdrawalSchedule

Gets a Wire withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.GetWireWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1")
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `wireWithdrawalScheduleID`                               | *string*                                                 | :heavy_check_mark:                                       | The wireWithdrawalSchedule id.                           | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.WireWithdrawalSchedulesGetWireWithdrawalScheduleResponse](../../models/operations/wirewithdrawalschedulesgetwirewithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateWireWithdrawalSchedule

Updates the amount of a Wire withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.UpdateWireWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.WireWithdrawalScheduleUpdate{}, ascendsdkgo.String("{\"update_mask\":\"schedule_details.amount\"}"))
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               | Example                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |                                                                                                                           |
| `accountID`                                                                                                               | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The account id.                                                                                                           | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                |
| `wireWithdrawalScheduleID`                                                                                                | *string*                                                                                                                  | :heavy_check_mark:                                                                                                        | The wireWithdrawalSchedule id.                                                                                            | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                      |
| `wireWithdrawalScheduleUpdate`                                                                                            | [components.WireWithdrawalScheduleUpdate](../../models/components/wirewithdrawalscheduleupdate.md)                        | :heavy_check_mark:                                                                                                        | N/A                                                                                                                       |                                                                                                                           |
| `updateMask`                                                                                                              | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | A field mask representing the update. Note: only the 'schedule_details.amount' field of a schedule is currently updatable | {<br/>"update_mask": "schedule_details.amount"<br/>}                                                                      |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |                                                                                                                           |

### Response

**[*operations.WireWithdrawalSchedulesUpdateWireWithdrawalScheduleResponse](../../models/operations/wirewithdrawalschedulesupdatewirewithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelWireWithdrawalSchedule

Cancels a Wire withdrawal transfer schedule

### Example Usage

```go
package main

import(
	"github.com/afs-public/ascend-sdk-go/models/components"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"context"
	"log"
)

func main() {
    s := ascendsdkgo.New(
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey: "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
                Name: "FinFirm",
                Organization: "correspondents/00000000-0000-0000-0000-000000000000",
                Type: "serviceAccount",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ScheduleTransfers.CancelWireWithdrawalSchedule(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1", components.CancelWireWithdrawalScheduleRequestCreate{
        Comment: ascendsdkgo.String("User Request"),
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawalSchedules/40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawalSchedule != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `accountID`                                                                                                                  | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The account id.                                                                                                              | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                   |
| `wireWithdrawalScheduleID`                                                                                                   | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The wireWithdrawalSchedule id.                                                                                               | 40eb6b6f-76ff-4dc9-b8a0-b65a7658f8b1                                                                                         |
| `cancelWireWithdrawalScheduleRequestCreate`                                                                                  | [components.CancelWireWithdrawalScheduleRequestCreate](../../models/components/cancelwirewithdrawalschedulerequestcreate.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |                                                                                                                              |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*operations.WireWithdrawalSchedulesCancelWireWithdrawalScheduleResponse](../../models/operations/wirewithdrawalschedulescancelwirewithdrawalscheduleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |