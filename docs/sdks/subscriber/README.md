# Subscriber
(*Subscriber*)

## Overview

### Available Operations

* [CreatePushSubscription](#createpushsubscription) - Create Push Subscription
* [ListPushSubscriptions](#listpushsubscriptions) - List Push Subscriptions
* [GetPushSubscription](#getpushsubscription) - Get Push Subscription
* [UpdatePushSubscription](#updatepushsubscription) - Update Subscription
* [DeletePushSubscription](#deletepushsubscription) - Delete Subscription
* [GetPushSubscriptionDelivery](#getpushsubscriptiondelivery) - Get Subscription Event Delivery
* [ListPushSubscriptionDeliveries](#listpushsubscriptiondeliveries) - List Push Subscription Event Deliveries

## CreatePushSubscription

Creates a new push subscription for event notifications.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_CreatePushSubscription" method="post" path="/events/v1/subscriptions" -->
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

    res, err := s.Subscriber.CreatePushSubscription(ctx, components.PushSubscriptionCreate{
        DisplayName: "This is an example HTTP configuration.",
        EventTypes: []string{
            "position.v1.updated",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.PushSubscriptionCreate](../../models/components/pushsubscriptioncreate.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.SubscriberCreatePushSubscriptionResponse](../../models/operations/subscribercreatepushsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 409 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPushSubscriptions

Gets a list of push subscriptions.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_ListPushSubscriptions" method="get" path="/events/v1/subscriptions" -->
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

    res, err := s.Subscriber.ListPushSubscriptions(ctx, ascendsdkgo.String("correspondent_id==\"01H8MCDXH4HYJJAV921BDKCC83\""), ascendsdkgo.Int(50), ascendsdkgo.String("ZXhhbXBsZQo"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPushSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `filter`                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; If empty, all subscriptions the user has permission to view will be returned; Filter options include:<br/> `name`<br/> `subscription_id`<br/> `client_id`<br/> `correspondent_id`<br/> `display_name`<br/> `event_types`<br/> `state`<br/> `http_callback.url`<br/> `http_callback.timeout_seconds` | correspondent_id=="01H8MCDXH4HYJJAV921BDKCC83"                                                                                                                                                                                                                                                                                                                                                                                       |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | The number of entries to return in a single page; Default = 100; Maximum = 1000                                                                                                                                                                                                                                                                                                                                                      | 50                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `pageToken`                                                                                                                                                                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | Page token used for pagination; Supplying a page token returns the next page of results                                                                                                                                                                                                                                                                                                                                              | ZXhhbXBsZQo                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                      |

### Response

**[*operations.SubscriberListPushSubscriptionsResponse](../../models/operations/subscriberlistpushsubscriptionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403      | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPushSubscription

Gets the details of a specific push subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_GetPushSubscription" method="get" path="/events/v1/subscriptions/{subscription_id}" -->
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

    res, err := s.Subscriber.GetPushSubscription(ctx, "01H8MCDXH4JVH7KVNB2YY42907")
    if err != nil {
        log.Fatal(err)
    }
    if res.PushSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | *string*                                                 | :heavy_check_mark:                                       | The subscription id.                                     | 01H8MCDXH4JVH7KVNB2YY42907                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SubscriberGetPushSubscriptionResponse](../../models/operations/subscribergetpushsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePushSubscription

Updates the details of a push subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_UpdatePushSubscription" method="patch" path="/events/v1/subscriptions/{subscription_id}" -->
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

    res, err := s.Subscriber.UpdatePushSubscription(ctx, "01H8MCDXH4JVH7KVNB2YY42907", components.PushSubscriptionUpdate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PushSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `subscriptionID`                                                                       | *string*                                                                               | :heavy_check_mark:                                                                     | The subscription id.                                                                   | 01H8MCDXH4JVH7KVNB2YY42907                                                             |
| `pushSubscriptionUpdate`                                                               | [components.PushSubscriptionUpdate](../../models/components/pushsubscriptionupdate.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `updateMask`                                                                           | **string*                                                                              | :heavy_minus_sign:                                                                     | The fields to update in subscription                                                   |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.SubscriberUpdatePushSubscriptionResponse](../../models/operations/subscriberupdatepushsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePushSubscription

Stops receiving events from a push subscription, and then deletes it.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_DeletePushSubscription" method="delete" path="/events/v1/subscriptions/{subscription_id}" -->
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

    res, err := s.Subscriber.DeletePushSubscription(ctx, "01H8MCDXH4JVH7KVNB2YY42907")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | *string*                                                 | :heavy_check_mark:                                       | The subscription id.                                     | 01H8MCDXH4JVH7KVNB2YY42907                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SubscriberDeletePushSubscriptionResponse](../../models/operations/subscriberdeletepushsubscriptionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPushSubscriptionDelivery

Gets the details of a specific push subscription delivery.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_GetPushSubscriptionDelivery" method="get" path="/events/v1/subscriptions/{subscription_id}/deliveries/{delivery_id}" -->
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

    res, err := s.Subscriber.GetPushSubscriptionDelivery(ctx, "01H8MCDXH4JVH7KVNB2YY42907", "01H8MCDXH415BJ962YDN4B02JK")
    if err != nil {
        log.Fatal(err)
    }
    if res.PushSubscriptionDelivery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `subscriptionID`                                         | *string*                                                 | :heavy_check_mark:                                       | The subscription id.                                     | 01H8MCDXH4JVH7KVNB2YY42907                               |
| `deliveryID`                                             | *string*                                                 | :heavy_check_mark:                                       | The delivery id.                                         | 01H8MCDXH415BJ962YDN4B02JK                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SubscriberGetPushSubscriptionDeliveryResponse](../../models/operations/subscribergetpushsubscriptiondeliveryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPushSubscriptionDeliveries

Gets a list of a push subscription's event deliveries.

### Example Usage

<!-- UsageSnippet language="go" operationID="Subscriber_ListPushSubscriptionDeliveries" method="get" path="/events/v1/subscriptions/{subscription_id}/deliveries" -->
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

    res, err := s.Subscriber.ListPushSubscriptionDeliveries(ctx, "01H8MCDXH4JVH7KVNB2YY42907", ascendsdkgo.String("event_publish_time==timestamp(\"2023-06-13T23:48:58.343Z\")"), ascendsdkgo.Int(50), ascendsdkgo.String("ZXhhbXBsZQo"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPushSubscriptionDeliveriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                         |
| `subscriptionID`                                                                                                                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                      | The subscription id.                                                                                                                                                                                                                                                                                                                                                                    | 01H8MCDXH4JVH7KVNB2YY42907                                                                                                                                                                                                                                                                                                                                                              |
| `filter`                                                                                                                                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                      | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; If left empty, all deliveries the user has permission to view are returned; Filter options include:<br/> `name`<br/> `delivery_id`<br/> `event`<br/> `event_publish_time`<br/> `result`<br/> `last_response`<br/> `last_send_time`<br/> `duration` | event_publish_time==timestamp("2023-06-13T23:48:58.343Z")                                                                                                                                                                                                                                                                                                                               |
| `pageSize`                                                                                                                                                                                                                                                                                                                                                                              | **int*                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                      | The number of entries to return in a single page; Default = 100; Maximum = 1000                                                                                                                                                                                                                                                                                                         | 50                                                                                                                                                                                                                                                                                                                                                                                      |
| `pageToken`                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                      | Page token used for pagination; Supplying a page token returns the next page of results                                                                                                                                                                                                                                                                                                 | ZXhhbXBsZQo                                                                                                                                                                                                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                      | The options for this request.                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                         |

### Response

**[*operations.SubscriberListPushSubscriptionDeliveriesResponse](../../models/operations/subscriberlistpushsubscriptiondeliveriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Status   | 400, 401, 403, 404 | application/json   |
| sdkerrors.Status   | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |