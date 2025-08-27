# TestSimulation
(*TestSimulation*)

## Overview

### Available Operations

* [SimulateCreateCheckDeposit](#simulatecreatecheckdeposit) - Simulate Check Deposit Creation
* [ForceApproveAchDeposit](#forceapproveachdeposit) - ACH Deposit Approval
* [ForceNocAchDeposit](#forcenocachdeposit) - NOC for a Deposit
* [ForceRejectAchDeposit](#forcerejectachdeposit) - ACH Deposit Rejection
* [ForceReturnAchDeposit](#forcereturnachdeposit) - ACH Deposit Return
* [ForceApproveAchWithdrawal](#forceapproveachwithdrawal) - ACH Withdrawal Approval
* [ForceNocAchWithdrawal](#forcenocachwithdrawal) - ACH Withdrawal NOC
* [ForceRejectAchWithdrawal](#forcerejectachwithdrawal) - ACH Withdrawal Rejection
* [ForceReturnAchWithdrawal](#forcereturnachwithdrawal) - ACH Withdrawal Return
* [GetMicroDepositAmounts](#getmicrodepositamounts) - Get Relationship Micro Deposit Verification
* [ForceApproveIctDeposit](#forceapproveictdeposit) - Force Approve ICT Deposit
* [ForceRejectIctDeposit](#forcerejectictdeposit) - Force Reject ICT Deposit
* [ForceApproveIctWithdrawal](#forceapproveictwithdrawal) - Force Approve ICT Withdrawal
* [ForceRejectIctWithdrawal](#forcerejectictwithdrawal) - Force Reject ICT Withdrawal
* [ForceApproveWireWithdrawal](#forceapprovewirewithdrawal) - Force Approve Wire Withdrawal
* [ForceRejectWireWithdrawal](#forcerejectwirewithdrawal) - Force Reject Wire Withdrawal
* [ForceApproveCashJournal](#forceapprovecashjournal) - Force Approve Cash Journal
* [ForceRejectCashJournal](#forcerejectcashjournal) - Force Reject Cash Journal

## SimulateCreateCheckDeposit

Creates a check deposit for a specific account FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="CheckDeposits_SimulateCreateCheckDeposit" method="post" path="/transfers/v1/accounts/{account_id}/checkDeposits:simulate" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.SimulateCreateCheckDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", components.SimulateCreateCheckDepositRequestCreate{
        Amount: components.DecimalCreate{},
        Parent: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |                                                                                                                          |
| `accountID`                                                                                                              | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The account id.                                                                                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                               |
| `simulateCreateCheckDepositRequestCreate`                                                                                | [components.SimulateCreateCheckDepositRequestCreate](../../models/components/simulatecreatecheckdepositrequestcreate.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |                                                                                                                          |

### Response

**[*operations.CheckDepositsSimulateCreateCheckDepositResponse](../../models/operations/checkdepositssimulatecreatecheckdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveAchDeposit

Forces approval of an existing ACH deposit that is pending review. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_ForceApproveAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceApproveAchDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `accountID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The account id.                                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                       |
| `achDepositID`                                                                                                   | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The achDeposit id.                                                                                               | 20230817000319                                                                                                   |
| `forceApproveAchDepositRequestCreate`                                                                            | [components.ForceApproveAchDepositRequestCreate](../../models/components/forceapproveachdepositrequestcreate.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.AchDepositsForceApproveAchDepositResponse](../../models/operations/achdepositsforceapproveachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceNocAchDeposit

Forces a Nacha notice of change (NOC) on a completed ACH deposit. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_ForceNocAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}:forceNoc" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceNocAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceNocAchDepositRequestCreate{
        NachaNoc: components.NachaNocCreate{
            Code: components.CodeC01,
        },
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              | Example                                                                                                  |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |                                                                                                          |
| `accountID`                                                                                              | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The account id.                                                                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                               |
| `achDepositID`                                                                                           | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The achDeposit id.                                                                                       | 20230817000319                                                                                           |
| `forceNocAchDepositRequestCreate`                                                                        | [components.ForceNocAchDepositRequestCreate](../../models/components/forcenocachdepositrequestcreate.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |                                                                                                          |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |                                                                                                          |

### Response

**[*operations.AchDepositsForceNocAchDepositResponse](../../models/operations/achdepositsforcenocachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectAchDeposit

Forces rejection of an existing ACH deposit that is pending review. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_ForceRejectAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceRejectAchDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                     |
| `achDepositID`                                                                                                 | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The achDeposit id.                                                                                             | 20230817000319                                                                                                 |
| `forceRejectAchDepositRequestCreate`                                                                           | [components.ForceRejectAchDepositRequestCreate](../../models/components/forcerejectachdepositrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.AchDepositsForceRejectAchDepositResponse](../../models/operations/achdepositsforcerejectachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceReturnAchDeposit

Forces a Nacha return on a completed ACH deposit. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchDeposits_ForceReturnAchDeposit" method="post" path="/transfers/v1/accounts/{account_id}/achDeposits/{achDeposit_id}:forceReturn" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceReturnAchDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceReturnAchDepositRequestCreate{
        NachaReturn: components.NachaReturnCreate{
            Code: components.NachaReturnCreateCodeR13,
        },
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achDeposits/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                     |
| `achDepositID`                                                                                                 | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The achDeposit id.                                                                                             | 20230817000319                                                                                                 |
| `forceReturnAchDepositRequestCreate`                                                                           | [components.ForceReturnAchDepositRequestCreate](../../models/components/forcereturnachdepositrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.AchDepositsForceReturnAchDepositResponse](../../models/operations/achdepositsforcereturnachdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveAchWithdrawal

Forces approval of an existing ACH withdrawal that is pending review. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_ForceApproveAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726", components.ForceApproveAchWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawals/20230620500726",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `accountID`                                                                                                            | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The account id.                                                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                             |
| `achWithdrawalID`                                                                                                      | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The achWithdrawal id.                                                                                                  | 20230620500726                                                                                                         |
| `forceApproveAchWithdrawalRequestCreate`                                                                               | [components.ForceApproveAchWithdrawalRequestCreate](../../models/components/forceapproveachwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.AchWithdrawalsForceApproveAchWithdrawalResponse](../../models/operations/achwithdrawalsforceapproveachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceNocAchWithdrawal

Forces a Nacha notice of change (NOC) on a completed ACH withdrawal. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_ForceNocAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}:forceNoc" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceNocAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726", components.ForceNocAchWithdrawalRequestCreate{
        NachaNoc: components.NachaNocCreate{
            Code: components.CodeC01,
        },
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawals/20230620500726",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                     |
| `achWithdrawalID`                                                                                              | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The achWithdrawal id.                                                                                          | 20230620500726                                                                                                 |
| `forceNocAchWithdrawalRequestCreate`                                                                           | [components.ForceNocAchWithdrawalRequestCreate](../../models/components/forcenocachwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.AchWithdrawalsForceNocAchWithdrawalResponse](../../models/operations/achwithdrawalsforcenocachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectAchWithdrawal

Forces rejection of an existing ACH withdrawal that is pending review. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_ForceRejectAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726", components.ForceRejectAchWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawals/20230620500726",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `accountID`                                                                                                          | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The account id.                                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                           |
| `achWithdrawalID`                                                                                                    | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The achWithdrawal id.                                                                                                | 20230620500726                                                                                                       |
| `forceRejectAchWithdrawalRequestCreate`                                                                              | [components.ForceRejectAchWithdrawalRequestCreate](../../models/components/forcerejectachwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.AchWithdrawalsForceRejectAchWithdrawalResponse](../../models/operations/achwithdrawalsforcerejectachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceReturnAchWithdrawal

Forces a Nacha return on a completed ACH withdrawal. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="AchWithdrawals_ForceReturnAchWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/achWithdrawals/{achWithdrawal_id}:forceReturn" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceReturnAchWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230620500726", components.ForceReturnAchWithdrawalRequestCreate{
        NachaReturn: components.NachaReturnCreate{
            Code: components.NachaReturnCreateCodeR15,
        },
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/achWithdrawals/20230620500726",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AchWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `accountID`                                                                                                          | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The account id.                                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                           |
| `achWithdrawalID`                                                                                                    | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The achWithdrawal id.                                                                                                | 20230620500726                                                                                                       |
| `forceReturnAchWithdrawalRequestCreate`                                                                              | [components.ForceReturnAchWithdrawalRequestCreate](../../models/components/forcereturnachwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.AchWithdrawalsForceReturnAchWithdrawalResponse](../../models/operations/achwithdrawalsforcereturnachwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMicroDepositAmounts

Gets micro deposit amounts for bank relationships with the `MICRO_DEPOSIT` verification method. FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="BankRelationships_GetMicroDepositAmounts" method="get" path="/transfers/v1/accounts/{account_id}/bankRelationships/{bankRelationship_id}:getMicroDepositAmounts" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.GetMicroDepositAmounts(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "651ef9de0dee00240813e60e")
    if err != nil {
        log.Fatal(err)
    }
    if res.MicroDepositAmounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                               |
| `bankRelationshipID`                                     | *string*                                                 | :heavy_check_mark:                                       | The bankRelationship id.                                 | 651ef9de0dee00240813e60e                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.BankRelationshipsGetMicroDepositAmountsResponse](../../models/operations/bankrelationshipsgetmicrodepositamountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveIctDeposit

Forces an approval on an existing ICT deposit pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="IctDeposits_ForceApproveIctDeposit" method="post" path="/transfers/v1/accounts/{account_id}/ictDeposits/{ictDeposit_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveIctDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.ForceApproveIctDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictDeposits/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `accountID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The account id.                                                                                                  | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                       |
| `ictDepositID`                                                                                                   | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The ictDeposit id.                                                                                               | 20240321000472                                                                                                   |
| `forceApproveIctDepositRequestCreate`                                                                            | [components.ForceApproveIctDepositRequestCreate](../../models/components/forceapproveictdepositrequestcreate.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.IctDepositsForceApproveIctDepositResponse](../../models/operations/ictdepositsforceapproveictdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectIctDeposit

Forces a rejection on an existing ICT deposit pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="IctDeposits_ForceRejectIctDeposit" method="post" path="/transfers/v1/accounts/{account_id}/ictDeposits/{ictDeposit_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectIctDeposit(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.ForceRejectIctDepositRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictDeposits/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctDeposit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    | Example                                                                                                        |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |                                                                                                                |
| `accountID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The account id.                                                                                                | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                     |
| `ictDepositID`                                                                                                 | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The ictDeposit id.                                                                                             | 20240321000472                                                                                                 |
| `forceRejectIctDepositRequestCreate`                                                                           | [components.ForceRejectIctDepositRequestCreate](../../models/components/forcerejectictdepositrequestcreate.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |                                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |                                                                                                                |

### Response

**[*operations.IctDepositsForceRejectIctDepositResponse](../../models/operations/ictdepositsforcerejectictdepositresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveIctWithdrawal

Forces an approval on an existing ICT withdrawal pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="IctWithdrawals_ForceApproveIctWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/ictWithdrawals/{ictWithdrawal_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveIctWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.ForceApproveIctWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictWithdrawals/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `accountID`                                                                                                            | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The account id.                                                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                             |
| `ictWithdrawalID`                                                                                                      | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The ictWithdrawal id.                                                                                                  | 20240321000472                                                                                                         |
| `forceApproveIctWithdrawalRequestCreate`                                                                               | [components.ForceApproveIctWithdrawalRequestCreate](../../models/components/forceapproveictwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.IctWithdrawalsForceApproveIctWithdrawalResponse](../../models/operations/ictwithdrawalsforceapproveictwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectIctWithdrawal

Forces a rejection on an existing ICT withdrawal pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="IctWithdrawals_ForceRejectIctWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/ictWithdrawals/{ictWithdrawal_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectIctWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20240321000472", components.ForceRejectIctWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/ictWithdrawals/20240321000472",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IctWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `accountID`                                                                                                          | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The account id.                                                                                                      | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                           |
| `ictWithdrawalID`                                                                                                    | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | The ictWithdrawal id.                                                                                                | 20240321000472                                                                                                       |
| `forceRejectIctWithdrawalRequestCreate`                                                                              | [components.ForceRejectIctWithdrawalRequestCreate](../../models/components/forcerejectictwithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.IctWithdrawalsForceRejectIctWithdrawalResponse](../../models/operations/ictwithdrawalsforcerejectictwithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveWireWithdrawal

Forces an approval on an existing wire withdrawal pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="WireWithdrawals_ForceApproveWireWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/wireWithdrawals/{wireWithdrawal_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveWireWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceApproveWireWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawals/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |                                                                                                                          |
| `accountID`                                                                                                              | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The account id.                                                                                                          | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                               |
| `wireWithdrawalID`                                                                                                       | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The wireWithdrawal id.                                                                                                   | 20230817000319                                                                                                           |
| `forceApproveWireWithdrawalRequestCreate`                                                                                | [components.ForceApproveWireWithdrawalRequestCreate](../../models/components/forceapprovewirewithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |                                                                                                                          |

### Response

**[*operations.WireWithdrawalsForceApproveWireWithdrawalResponse](../../models/operations/wirewithdrawalsforceapprovewirewithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectWireWithdrawal

Forces a rejection on an existing wire withdrawal pending review - FOR TESTING

### Example Usage

<!-- UsageSnippet language="go" operationID="WireWithdrawals_ForceRejectWireWithdrawal" method="post" path="/transfers/v1/accounts/{account_id}/wireWithdrawals/{wireWithdrawal_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectWireWithdrawal(ctx, "01H8FB90ZRRFWXB4XC2JPJ1D4Y", "20230817000319", components.ForceRejectWireWithdrawalRequestCreate{
        Name: "accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y/wireWithdrawals/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WireWithdrawal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            | Example                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |                                                                                                                        |
| `accountID`                                                                                                            | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The account id.                                                                                                        | 01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                             |
| `wireWithdrawalID`                                                                                                     | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The wireWithdrawal id.                                                                                                 | 20230817000319                                                                                                         |
| `forceRejectWireWithdrawalRequestCreate`                                                                               | [components.ForceRejectWireWithdrawalRequestCreate](../../models/components/forcerejectwirewithdrawalrequestcreate.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |                                                                                                                        |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |                                                                                                                        |

### Response

**[*operations.WireWithdrawalsForceRejectWireWithdrawalResponse](../../models/operations/wirewithdrawalsforcerejectwirewithdrawalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403, 404      | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceApproveCashJournal

Forces approval of an existing cash journal that is pending review FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="CashJournals_ForceApproveCashJournal" method="post" path="/transfers/v1/cashJournals/{cashJournal_id}:forceApprove" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceApproveCashJournal(ctx, "20230817000319", components.ForceApproveCashJournalRequestCreate{
        Name: "cashJournals/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `cashJournalID`                                                                                                    | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The cashJournal id.                                                                                                | 20230817000319                                                                                                     |
| `forceApproveCashJournalRequestCreate`                                                                             | [components.ForceApproveCashJournalRequestCreate](../../models/components/forceapprovecashjournalrequestcreate.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.CashJournalsForceApproveCashJournalResponse](../../models/operations/cashjournalsforceapprovecashjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ForceRejectCashJournal

Forces rejection of an existing cash journal that is pending review FOR TESTING ONLY!

### Example Usage

<!-- UsageSnippet language="go" operationID="CashJournals_ForceRejectCashJournal" method="post" path="/transfers/v1/cashJournals/{cashJournal_id}:forceReject" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

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

    res, err := s.TestSimulation.ForceRejectCashJournal(ctx, "20230817000319", components.ForceRejectCashJournalRequestCreate{
        Name: "cashJournals/20230817000319",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CashJournal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `cashJournalID`                                                                                                  | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The cashJournal id.                                                                                              | 20230817000319                                                                                                   |
| `forceRejectCashJournalRequestCreate`                                                                            | [components.ForceRejectCashJournalRequestCreate](../../models/components/forcerejectcashjournalrequestcreate.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.CashJournalsForceRejectCashJournalResponse](../../models/operations/cashjournalsforcerejectcashjournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 403           | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |