# OptionInstructions
(*OptionInstructions*)

## Overview

### Available Operations

* [CreateOptionInstruction](#createoptioninstruction) - Create Option Instruction
* [ListOptionInstructions](#listoptioninstructions) - List Option Instructions
* [GetOptionInstruction](#getoptioninstruction) - Get Option Instruction
* [CancelOptionInstruction](#canceloptioninstruction) - Cancel Option Instruction

## CreateOptionInstruction

CreateOptionInstruction creates a new option instruction for trading actions

### Example Usage

<!-- UsageSnippet language="go" operationID="ExerciseService_CreateOptionInstruction" method="post" path="/options/v1/accounts/{account_id}/assets/{asset_id}/instructions" -->
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

    res, err := s.OptionInstructions.CreateOptionInstruction(ctx, "ACC123456", "12345", components.OptionInstructionCreate{
        AccountID: "01JTNGZM8PWRR6MHBCC6TEG9HE",
        Identifier: "AAPL250620P00090000",
        IdentifierType: components.OptionInstructionCreateIdentifierTypeOsi,
        Quantity: components.DecimalCreate{},
        Type: components.OptionInstructionCreateTypeDoNotExercise,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionInstruction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `accountID`                                                                              | *string*                                                                                 | :heavy_check_mark:                                                                       | The account id.                                                                          | ACC123456                                                                                |
| `assetID`                                                                                | *string*                                                                                 | :heavy_check_mark:                                                                       | The asset id.                                                                            | 12345                                                                                    |
| `optionInstructionCreate`                                                                | [components.OptionInstructionCreate](../../models/components/optioninstructioncreate.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.ExerciseServiceCreateOptionInstructionResponse](../../models/operations/exerciseservicecreateoptioninstructionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListOptionInstructions

ListOptionInstructions lists option instructions with optional filtering and pagination

### Example Usage

<!-- UsageSnippet language="go" operationID="ExerciseService_ListOptionInstructions" method="get" path="/options/v1/accounts/{account_id}/assets/{asset_id}/instructions" -->
```go
package main

import(
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
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

    res, err := s.OptionInstructions.ListOptionInstructions(ctx, operations.ExerciseServiceListOptionInstructionsRequest{
        AccountID: "ACC123456",
        AssetID: "12345",
        PageSize: ascendsdkgo.Int(50),
        PageToken: ascendsdkgo.String("eyJvZmZzZXQiOjUwfQ=="),
        Filter: ascendsdkgo.String("type == 'DO_NOT_EXERCISE' && account_id == '12345'"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListOptionInstructionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ExerciseServiceListOptionInstructionsRequest](../../models/operations/exerciseservicelistoptioninstructionsrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.ExerciseServiceListOptionInstructionsResponse](../../models/operations/exerciseservicelistoptioninstructionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetOptionInstruction

GetOptionInstruction retrieves an existing instruction by name

### Example Usage

<!-- UsageSnippet language="go" operationID="ExerciseService_GetOptionInstruction" method="get" path="/options/v1/accounts/{account_id}/assets/{asset_id}/instructions/{instruction_id}" -->
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

    res, err := s.OptionInstructions.GetOptionInstruction(ctx, "ACC123456", "12345", "abc12345")
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionInstruction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | The account id.                                          | ACC123456                                                |
| `assetID`                                                | *string*                                                 | :heavy_check_mark:                                       | The asset id.                                            | 12345                                                    |
| `instructionID`                                          | *string*                                                 | :heavy_check_mark:                                       | The instruction id.                                      | abc12345                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ExerciseServiceGetOptionInstructionResponse](../../models/operations/exerciseservicegetoptioninstructionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CancelOptionInstruction

CancelOptionInstruction cancels an existing instruction by name

### Example Usage

<!-- UsageSnippet language="go" operationID="ExerciseService_CancelOptionInstruction" method="post" path="/options/v1/accounts/{account_id}/assets/{asset_id}/instructions/{instruction_id}:cancel" -->
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

    res, err := s.OptionInstructions.CancelOptionInstruction(ctx, "ACC123456", "12345", "abc12345", components.CancelOptionInstructionRequestCreate{
        Name: "accounts/ACC123456/assets/12345/instructions/abc12345",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OptionInstruction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `accountID`                                                                                                        | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The account id.                                                                                                    | ACC123456                                                                                                          |
| `assetID`                                                                                                          | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The asset id.                                                                                                      | 12345                                                                                                              |
| `instructionID`                                                                                                    | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The instruction id.                                                                                                | abc12345                                                                                                           |
| `cancelOptionInstructionRequestCreate`                                                                             | [components.CancelOptionInstructionRequestCreate](../../models/components/canceloptioninstructionrequestcreate.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.ExerciseServiceCancelOptionInstructionResponse](../../models/operations/exerciseservicecanceloptioninstructionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |